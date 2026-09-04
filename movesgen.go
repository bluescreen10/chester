package chester

// Castling path constants define the squares that must be unoccupied (Free)
// or not under attack (NotAttacked) for each castling option.
//
// Free squares are all squares between the king and rook (exclusive).
// NotAttacked squares are those the king passes through or lands on.
const (
	whiteQueenSideCastleFree        = BB_SQ_B1 | BB_SQ_C1 | BB_SQ_D1
	whiteQueenSideCastleNotAttacked = BB_SQ_C1 | BB_SQ_D1 | BB_SQ_E1
	whiteKingSideCastleFree         = BB_SQ_F1 | BB_SQ_G1
	whiteKingSideCastleNotAttacked  = BB_SQ_E1 | BB_SQ_F1 | BB_SQ_G1

	blackQueenSideCastleFree        = BB_SQ_B8 | BB_SQ_C8 | BB_SQ_D8
	blackQueenSideCastleNotAttacked = BB_SQ_C8 | BB_SQ_D8 | BB_SQ_E8
	blackKingSideCastleFree         = BB_SQ_F8 | BB_SQ_G8
	blackKingSideCastleNotAttacked  = BB_SQ_E8 | BB_SQ_F8 | BB_SQ_G8
)

// appendPromotions appends the promotion moves for a pawn going from from to
// to, and reports the moves slice.
//
// allPieces selects whether the dominated promotions are included. A queen is
// a rook plus a bishop, so a rook or bishop promotion is only ever played to
// avoid stalemate -- and the quiescence search stands pat rather than
// reasoning about stalemate, so there they are pure branching factor. A knight
// is not dominated by anything: it attacks squares a queen cannot, which is
// what makes promoting with check to fork the king and queen work, so it is
// always generated.
func appendPromotions(moves []Move, from, to Square, allPieces bool) []Move {
	moves = append(moves,
		NewPromotionMove(from, to, Queen),
		NewPromotionMove(from, to, Knight),
	)

	if allPieces {
		moves = append(moves,
			NewPromotionMove(from, to, Rook),
			NewPromotionMove(from, to, Bishop),
		)
	}

	return moves
}

// promotionRanks are the two ranks a pawn promotes on. Splitting pawn pushes
// by destination rank separates promotions from ordinary pushes without a
// per-move branch in either loop.
const promotionRanks = Rank_1 | Rank_8

// checkersPinsAndMask accumulates the check and pin state of the active
// king, computed once per position by checkersAndPinned before dispatching
// to the per-piece generators.
type checkersPinsAndMask struct {
	// diagonalPins is the union of rays along which an active-color piece is
	// pinned diagonally against its king by an enemy bishop or queen.
	// A piece on this mask may only move along the ray itself.
	diagonalPins Bitboard

	// straightPins is the union of rays along which an active-color piece is
	// pinned on a rank or file against its king by an enemy rook or queen.
	// A piece on this mask may only move along the ray itself.
	straightPins Bitboard

	// moveMask restricts the destination squares of all non-king pieces.
	// When not in check it equals EnemiesOrEmpty (all moves allowed).
	// When in check by one piece it is the union of the checker's square and
	// the ray between checker and king, so a legal response must either
	// capture the checker or interpose on the ray.
	moveMask Bitboard
}

// LegalMoves appends all legal moves for the active color to moves and returns
// the updated slice and whether the king is in check.
func LegalMoves(moves []Move, p *Position) ([]Move, bool) {
	return legalMoves(moves, p, false)
}

// NoisyMoves appends the legal noisy moves for the active color to moves and
// returns the updated slice and whether the king is in check.
//
// A noisy move is one that changes material: a capture, or a promotion whether
// or not it captures. It is the exact complement of isQuiet, and it is the
// move set the quiescence search runs on -- quiescence exists to search noisy
// moves until none are left, which is what makes the resulting position quiet.
//
// Promotions to rook and bishop are omitted. They are noisy by the definition
// above, but a queen is a rook plus a bishop, so they are only ever played to
// avoid stalemate, which quiescence never reasons about.
func NoisyMoves(moves []Move, p *Position) ([]Move, bool) {
	return legalMoves(moves, p, true)
}

// legalMoves is the core move generator that produces all legal moves for the
// current player. If noisyOnly is true it generates only the moves that change
// material -- captures and promotions -- which is the set quiescence searches.
// It returns the updated moves slice and a boolean indicating if the king is
// currently in check.
func legalMoves(moves []Move, p *Position, noisyOnly bool) ([]Move, bool) {
	cpm := checkersPinsAndMask{}
	numCheckers := checkersAndPinned(p, &cpm)
	inCheck := true

	switch numCheckers {
	case 0:
		if noisyOnly {
			cpm.moveMask = p.Enemies()
		} else {
			cpm.moveMask = p.EnemiesOrEmpty()
		}
		inCheck = false
		fallthrough
	case 1:
		if !noisyOnly {
			moves = genPawnForwardMoves(moves, p, cpm)
		}

		// Promotions are generated even when only noisy moves are wanted.
		// They are not captures, but they are not quiet either: leaving one
		// out of the quiescence search hides a queen appearing on the board
		// just past the horizon.
		moves = genPawnPromotions(moves, p, cpm, inCheck, noisyOnly)

		moves = genPawnLeftAttackMoves(moves, p, cpm, noisyOnly)
		moves = genPawnRightAttackMoves(moves, p, cpm, noisyOnly)

		if p.EnPassantTarget() != SQ_NULL {
			moves = genPawnEnPassantMoves(moves, p, cpm, inCheck)
		}
		moves = genKnightMoves(moves, p, cpm)
		moves = genBishopMoves(moves, p, cpm)
		moves = genRookMoves(moves, p, cpm)
		moves = genQueenMoves(moves, p, cpm)
		fallthrough
	default:
		moves = genKingMoves(moves, p, noisyOnly)
	}
	return moves, inCheck
}

// checkersAndPinned computes checkers, pinned pieces, and the move mask for
// the active king and stores the results in cpm. It returns the number of
// pieces currently giving check (0, 1, or 2).
//
// Knight and pawn checkers are found with direct attack-table lookups.
// Sliding checkers are found by tracing diagonal and straight rays outward
// from the king and intersecting with enemy sliders:
//   - A ray with no intervening friendly piece is a direct check; the
//     checker's square and the ray are added to moveMask.
//   - A ray with exactly one intervening friendly piece is a pin; the ray
//     is added to diagonalPins or straightPins accordingly.
func checkersAndPinned(p *Position, cpm *checkersPinsAndMask) int {
	us := p.Active()
	king := p.King()
	kingSq, _ := king.PopLSB()

	//cpm := checkersPinsAndMask{}

	checkers := knightMoves[kingSq] & p.EnemyKnights()

	leftAttacks := int(16*us - 9)
	rightAttacks := int(16*us - 7)
	pawns := p.EnemyPawns()
	checkers |= (king & File_Not_A).RotateLeft(leftAttacks) & pawns
	checkers |= (king & File_Not_H).RotateLeft(rightAttacks) & pawns

	kingDiagonalRays := diagonalRays[kingSq]
	diagonalAttackers := p.EnemyQueensOrBishops()

	var sq Square

	occupied := p.Occupied()

	for potentialCheckers := diagonalAttackers & kingDiagonalRays; potentialCheckers != 0; {
		sq, potentialCheckers = potentialCheckers.PopLSB()

		path := lineFromTo[kingSq][sq]
		potentialyPinned := path & occupied
		if potentialyPinned != 0 {
			switch potentialyPinned.OnesCount() {
			case 1:
				checkers |= 1 << sq
				cpm.moveMask |= path
			case 2:
				cpm.diagonalPins |= path
			}
		}
	}

	kingStraightRays := straightRays[kingSq]
	straightAttackers := p.EnemyQueensOrRooks()

	for potentialCheckers := straightAttackers & kingStraightRays; potentialCheckers != 0; {
		sq, potentialCheckers = potentialCheckers.PopLSB()

		path := lineFromTo[kingSq][sq]
		potentialyPinned := path & occupied
		if potentialyPinned != 0 {
			switch potentialyPinned.OnesCount() {
			case 1:
				checkers |= 1 << sq
				cpm.moveMask |= path
			case 2:
				cpm.straightPins |= path
			}
		}
	}

	cpm.moveMask |= checkers
	return checkers.OnesCount()
}

// genPawnsAttacks returns a Bitboard of all squares attacked by the inactive
// color's pawns. The active king must not step onto these squares.
func genPawnsAttacks(p *Position) Bitboard {
	pawns := p.EnemyPawns()
	color := p.Inactive()
	leftAttacks := 16*int(color) - 9
	rightAttacks := 16*int(color) - 7
	left := (pawns & File_Not_A).RotateLeft(leftAttacks)
	right := (pawns & File_Not_H).RotateLeft(rightAttacks)
	return left | right
}

// genKnightsAttacks returns a Bitboard of all squares attacked by the
// inactive color's knights. The active king must not step onto these squares.
func genKnightsAttacks(p *Position) Bitboard {
	var attacks Bitboard

	knights := p.EnemyKnights()

	var sq Square

	for knights != 0 {
		sq, knights = knights.PopLSB()
		attacks |= knightMoves[sq]
	}

	return attacks
}

// genDiagonalAttacks returns a Bitboard of all squares attacked diagonally
// by the inactive color's bishops and queens. The active king is removed from
// the occupancy so it cannot block its own escape squares.
func genDiagonalAttacks(p *Position) Bitboard {
	var attacks Bitboard

	attacker := p.EnemyQueensOrBishops()
	occupied := p.Occupied() &^ p.King()

	var sq Square

	for attacker != 0 {
		sq, attacker = attacker.PopLSB()
		attacks |= genBishopAttacks(sq, occupied)
	}

	return attacks
}

// genBishopAttacks returns the set of squares a bishop on sq attacks given
// the provided occupancy, using magic bitboard lookup.
func genBishopAttacks(sq Square, occupied Bitboard) Bitboard {

	occupied &= bishopMagic[sq].Mask
	occupied *= bishopMagic[sq].Magic
	occupied >>= bishopMagic[sq].Shift
	return bishopMagic[sq].Attacks[occupied]
}

// genStraightAttacks returns a Bitboard of all squares attacked along ranks
// and files by the inactive color's rooks and queens. The active king is
// removed from the occupancy so it cannot block its own escape squares.
func genStraightAttacks(p *Position) Bitboard {
	var attacks Bitboard

	attackers := p.EnemyQueensOrRooks()
	occupied := p.Occupied() &^ p.King()

	var sq Square

	for attackers != 0 {
		sq, attackers = attackers.PopLSB()
		attacks |= genRookAttacks(sq, occupied)
	}

	return attacks
}

// genRookAttacks returns the set of squares a rook on sq attacks given the
// provided occupancy, using magic bitboard lookup.
func genRookAttacks(sq Square, occupied Bitboard) Bitboard {
	occupied &= rookMagic[sq].Mask
	occupied *= rookMagic[sq].Magic
	occupied >>= rookMagic[sq].Shift
	return rookMagic[sq].Attacks[occupied]
}

// genKingAttacks returns a Bitboard of all squares attacked by the inactive
// color's king. Used to prevent the active king moving adjacent to the enemy
// king.
func genKingAttacks(p *Position) Bitboard {
	king := p.EnemyKing()
	sq, _ := king.PopLSB()
	return kingMoves[sq]
}

// genPawnForwardMoves appends all legal pawn push moves (single and double)
// for the active color. Diagonally pinned pawns cannot push. Straight-pinned
// pawns may only push along their pin ray. Pushes to the back rank are
// promotions are generated separately, by genPawnPromotions.
func genPawnForwardMoves(moves []Move, p *Position, cpm checkersPinsAndMask) []Move {
	us := p.Active()
	singlePushes := -8 + 16*int(us)
	startPlusOneRank := (Rank_3 * (1 - Bitboard(us))) | (Rank_6 * Bitboard(us))

	pawns := p.Pawns() &^ cpm.diagonalPins
	pinnedPawns := pawns & cpm.straightPins.RotateLeft(-singlePushes)
	unPinnedPawns := pawns &^ cpm.straightPins
	pawns = pinnedPawns | unPinnedPawns

	var from, to Square

	singlePush := pawns.RotateLeft(singlePushes) &^ p.Occupied()
	sp := Square(singlePushes)
	for pushes := singlePush & cpm.moveMask &^ promotionRanks; pushes != 0; {
		to, pushes = pushes.PopLSB()
		from = to - sp
		moves = append(moves, NewMove(from, to))
	}

	doublePushes := (singlePush & startPlusOneRank).RotateLeft(singlePushes) &^ p.Occupied() & cpm.moveMask
	dp := Square(2 * singlePushes)
	for doublePushes != 0 {
		to, doublePushes = doublePushes.PopLSB()
		from = to - dp
		moves = append(moves, NewMove(from, to))
	}
	return moves
}

// genPawnPromotions appends all legal pawn pushes onto the last rank for the
// active color, expanded by appendPromotions into one move per promotion
// piece. Capturing promotions are not included here; they come from the two
// attack generators.
//
// This is split out of genPawnForwardMoves so that the quiescence search can
// ask for promotions without asking for quiet pushes. A pawn reaching the last
// rank swings the evaluation by about eight hundred centipawns, which is
// exactly the kind of change quiescence exists to resolve, but a pawn stepping
// forward onto an empty square is the definition of a quiet move.
//
// inCheck selects which restriction applies to the destination. While in
// check, cpm.moveMask holds the squares that answer the check and must be
// obeyed. Otherwise no mask is needed: a push already lands on an empty
// square, and pinned pawns are excluded below.
func genPawnPromotions(moves []Move, p *Position, cpm checkersPinsAndMask, inCheck, noisyOnly bool) []Move {
	us := p.Active()

	// Only a pawn one rank from the end can promote by pushing. In nearly
	// every position there are none, and this check is what keeps the
	// function affordable for quiescence, which calls it at every node.
	aboutToPromote := (Rank_7 * (1 - Bitboard(us))) | (Rank_2 * Bitboard(us))
	if p.Pawns()&aboutToPromote == 0 {
		return moves
	}

	singlePushes := -8 + 16*int(us)

	pawns := p.Pawns() &^ cpm.diagonalPins
	pinnedPawns := pawns & cpm.straightPins.RotateLeft(-singlePushes)
	unPinnedPawns := pawns &^ cpm.straightPins
	pawns = pinnedPawns | unPinnedPawns

	pushes := pawns.RotateLeft(singlePushes) &^ p.Occupied() & promotionRanks
	if inCheck {
		pushes &= cpm.moveMask
	}

	sp := Square(singlePushes)
	var from, to Square

	for pushes != 0 {
		to, pushes = pushes.PopLSB()
		from = to - sp
		moves = appendPromotions(moves, from, to, !noisyOnly)
	}

	return moves
}

// genPawnLeftAttackMoves appends all legal pawn left-diagonal capture moves
// for the active color. "Left" is toward the a-file for White, toward the
// h-file for Black. Straight-pinned pawns cannot capture. Diagonally pinned
// pawns may only capture along their pin ray. Captures on the back rank are
// expanded by appendPromotions into one move per promotion piece.
func genPawnLeftAttackMoves(moves []Move, p *Position, cpm checkersPinsAndMask, noisyOnly bool) []Move {
	us := p.Active()
	leftAttacks := 16*int(us) - 9
	pawns := p.Pawns() &^ cpm.straightPins & File_Not_A
	pinnedPawns := pawns & (cpm.diagonalPins & File_Not_H).RotateLeft(-leftAttacks)
	unPinnedPawns := pawns &^ cpm.diagonalPins
	pawns = pinnedPawns | unPinnedPawns

	attacks := pawns.RotateLeft(leftAttacks) & p.Enemies() & cpm.moveMask
	var from, to Square

	for attacks != 0 {

		to, attacks = attacks.PopLSB()
		from = to - Square(leftAttacks)
		if to < SQ_A1 && to > SQ_H8 {
			moves = append(moves, NewMove(from, to))
		} else {
			moves = appendPromotions(moves, from, to, !noisyOnly)
		}
	}

	return moves
}

// genPawnRightAttackMoves appends all legal pawn right-diagonal capture moves
// for the active color. "Right" is toward the h-file for White, toward the
// a-file for Black. Straight-pinned pawns cannot capture. Diagonally pinned
// pawns may only capture along their pin ray. Captures on the back rank are
// expanded by appendPromotions into one move per promotion piece.
func genPawnRightAttackMoves(moves []Move, p *Position, cpm checkersPinsAndMask, noisyOnly bool) []Move {
	us := p.Active()
	rightAttacks := 16*int(us) - 7

	pawns := p.Pawns() &^ cpm.straightPins & File_Not_H
	pinnedPawns := pawns & (cpm.diagonalPins & File_Not_A).RotateLeft(-rightAttacks)
	unPinnedPawns := pawns &^ cpm.diagonalPins
	pawns = pinnedPawns | unPinnedPawns

	attacks := pawns.RotateLeft(rightAttacks) & p.Enemies() & cpm.moveMask
	var from, to Square

	for attacks != 0 {
		to, attacks = attacks.PopLSB()
		from = to - Square(rightAttacks)
		if to < SQ_A1 && to > SQ_H8 {
			moves = append(moves, NewMove(from, to))
		} else {
			moves = appendPromotions(moves, from, to, !noisyOnly)
		}
	}

	return moves
}

// genPawnEnPassantMoves appends any legal en passant capture moves. Pinned
// pawns are excluded. The horizontal-pin edge case is handled explicitly: after
// removing both the capturing and the captured pawn from the occupancy, the
// king's rank is re-checked for rook or queen attacks to ensure the capture
// does not expose the king.
func genPawnEnPassantMoves(moves []Move, p *Position, cpm checkersPinsAndMask, inCheck bool) []Move {
	us := p.Active()

	kingSq, _ := p.King().PopLSB()
	enemyQueensOrRooks := p.EnemyQueensOrRooks()

	pawns := p.Pawns() &^ (cpm.diagonalPins | cpm.straightPins)
	leftAttacks := int(16*us - 9)
	rightAttacks := int(16*us - 7)
	enPassantTarget := NewBitboardFromSquare(p.EnPassantTarget())

	if left := (pawns & File_Not_A).RotateLeft(leftAttacks) & enPassantTarget; left != 0 {
		to, _ := left.PopLSB()
		from := to - Square(leftAttacks)

		if enPassantIsLegal(p, cpm, kingSq, enemyQueensOrRooks, from, to, us, inCheck) {
			moves = append(moves, NewMove(from, to))
		}
	}

	if right := (pawns & File_Not_H).RotateLeft(rightAttacks) & enPassantTarget; right != 0 {
		to, _ := right.PopLSB()
		from := to - Square(rightAttacks)

		if enPassantIsLegal(p, cpm, kingSq, enemyQueensOrRooks, from, to, us, inCheck) {
			moves = append(moves, NewMove(from, to))
		}
	}

	return moves
}

// enPassantIsLegal reports whether capturing en passan from and to is
// legal. It covers the two cases the shared check and pin masks cannot,
// both of which exist because en passant is the only move where the captured
// piece does not stand on the destination square.
func enPassantIsLegal(p *Position, cpm checkersPinsAndMask, kingSq Square, enemyQueensOrRooks Bitboard, from, to Square, us Color, inCheck bool) bool {
	enemySq := to + 8 - Square(16*us)
	enemy := NewBitboardFromSquare(enemySq)

	// While in check, en passant only helps if it captures the checking pawn
	// or interposes on the checking ray. moveMask describes both, but it has
	// to be tested against the captured pawn's square as well as the
	// destination, since a double-pushed pawn giving check is not standing
	// on the square that captures it.
	if inCheck && cpm.moveMask&(NewBitboardFromSquare(to)|enemy) == 0 {
		return false
	}

	// En passant is also the only move that clears two squares of one rank at
	// once: the capturing pawn leaves from, and the captured pawn is removed
	// from beside it. Ordinary pin detection looks for a single blocker
	// between the king and an enemy slider, so a rank whose only blockers are
	// these two pawns never registers as a pin. Recheck it on the occupancy
	// the move actually produces.
	//
	// Only a rank can hide this. A file or diagonal through the king loses at
	// most one of the two pawns, which the regular pin masks already cover.
	rank := Bitboard(0xff) << (from & 56)
	if NewBitboardFromSquare(kingSq)&rank == 0 {
		return true
	}

	occupied := p.Occupied() &^ (NewBitboardFromSquare(from) | enemy)
	return genRookAttacks(kingSq, occupied)&rank&enemyQueensOrRooks == 0
}

// genKnightMoves appends all legal knight moves for the active color. Knights
// that are pinned (diagonally or straight) cannot move and are excluded
// entirely.
func genKnightMoves(moves []Move, p *Position, cpm checkersPinsAndMask) []Move {
	knights := p.Knights() &^ (cpm.diagonalPins | cpm.straightPins)
	var from, to Square
	for knights != 0 {
		from, knights = knights.PopLSB()
		targets := knightMoves[from] & cpm.moveMask

		for targets != 0 {
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to))
		}

	}

	return moves
}

// genBishopMoves appends all legal bishop moves for the active color.
// Straight-pinned bishops cannot move. Diagonally pinned bishops may only
// move along their pin ray.
func genBishopMoves(moves []Move, p *Position, cpm checkersPinsAndMask) []Move {
	bishops := p.Bishops() &^ cpm.straightPins

	var from, to Square
	for b := bishops & cpm.diagonalPins; b != 0; {
		from, b = b.PopLSB()
		targets := genBishopAttacks(from, p.Occupied()) & cpm.moveMask & cpm.diagonalPins

		for targets != 0 {
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to))
		}
	}

	for b := bishops &^ cpm.diagonalPins; b != 0; {
		from, b = b.PopLSB()
		targets := genBishopAttacks(from, p.Occupied()) & cpm.moveMask

		for targets != 0 {
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to))
		}
	}

	return moves
}

// genRookMoves appends all legal rook moves for the active color.
// Diagonally pinned rooks cannot move. Straight-pinned rooks may only move
// along their pin ray.
func genRookMoves(moves []Move, p *Position, cpm checkersPinsAndMask) []Move {
	rooks := p.Rooks() &^ cpm.diagonalPins

	var from, to Square
	for r := rooks & cpm.straightPins; r != 0; {
		from, r = r.PopLSB()
		targets := genRookAttacks(from, p.Occupied()) & cpm.moveMask & cpm.straightPins

		for targets != 0 {
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to))
		}
	}

	for r := rooks &^ cpm.straightPins; r != 0; {
		from, r = r.PopLSB()
		targets := genRookAttacks(from, p.Occupied()) & cpm.moveMask

		for targets != 0 {
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to))
		}
	}

	return moves
}

// genQueenMoves appends all legal queen moves for the active color. Pinned
// queens are restricted to their respective pin ray (diagonal or straight).
// Unpinned queens combine both bishop and rook attack sets.
func genQueenMoves(moves []Move, p *Position, cpm checkersPinsAndMask) []Move {
	queens := p.Queens()
	occupied := p.Occupied()

	var from, to Square
	for q := queens & cpm.diagonalPins; q != 0; {
		from, q = q.PopLSB()
		targets := genBishopAttacks(from, occupied) & cpm.moveMask & cpm.diagonalPins

		for targets != 0 {
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to))
		}
	}

	for q := queens & cpm.straightPins; q != 0; {
		from, q = q.PopLSB()
		targets := genRookAttacks(from, occupied) & cpm.moveMask & cpm.straightPins

		for targets != 0 {
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to))
		}
	}

	for q := queens &^ (cpm.diagonalPins | cpm.straightPins); q != 0; {
		from, q = q.PopLSB()
		targets := (genRookAttacks(from, occupied) | genBishopAttacks(from, occupied)) & cpm.moveMask

		for targets != 0 {
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to))
		}
	}

	return moves
}

// genKingMoves appends all legal king moves for the active color. The full
// enemy attack map is computed and subtracted from candidate targets.
//
// When noisyOnly is false, castling is also generated, but only when the rights
// flag is set, the path is unoccupied, and no square the king crosses is under
// attack. A castle moves no material, so it is never noisy.
func genKingMoves(moves []Move, p *Position, noisyOnly bool) []Move {
	us := p.Active()
	king := p.King()

	var mask Bitboard
	if noisyOnly {
		mask = p.Enemies()
	} else {
		mask = p.EnemiesOrEmpty()
	}

	from, _ := king.PopLSB()

	potentialTargets := kingMoves[from] & mask

	if potentialTargets == 0 {
		return moves
	}

	enemyKing := p.EnemyKing()
	attacked := attacks(p)

	for targets := potentialTargets &^ (enemyKing | attacked); targets != 0; {
		var to Square
		to, targets = targets.PopLSB()
		moves = append(moves, NewMove(from, to))
	}

	// castling
	if us == White {
		if p.CanWhiteCastleKingSide() &&
			whiteKingSideCastleFree&p.Occupied() == 0 &&
			whiteKingSideCastleNotAttacked&attacked == 0 {
			moves = append(moves, NewMove(SQ_E1, SQ_G1))
		}

		if p.CanWhiteCastleQueenSide() &&
			whiteQueenSideCastleFree&p.Occupied() == 0 &&
			whiteQueenSideCastleNotAttacked&attacked == 0 {
			moves = append(moves, NewMove(SQ_E1, SQ_C1))
		}
	}

	if us == Black {
		if p.CanBlackCastleKingSide() &&
			blackKingSideCastleFree&p.Occupied() == 0 &&
			blackKingSideCastleNotAttacked&attacked == 0 {
			moves = append(moves, NewMove(SQ_E8, SQ_G8))
		}

		if p.CanBlackCastleQueenSide() &&
			blackQueenSideCastleFree&p.Occupied() == 0 &&
			blackQueenSideCastleNotAttacked&attacked == 0 {
			moves = append(moves, NewMove(SQ_E8, SQ_C8))
		}
	}

	return moves
}

// attacks returns a Bitboard of every square attacked by at least one piece
// of the inactive color. Used by genKingMoves to determine safe king
// destinations.
func attacks(p *Position) Bitboard {
	return genPawnsAttacks(p) |
		genKnightsAttacks(p) |
		genDiagonalAttacks(p) |
		genStraightAttacks(p) |
		genKingAttacks(p)
}
