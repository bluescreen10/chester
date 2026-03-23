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

// LegalMoves appends all legal moves for the active color in position p to
// moves and returns the updated slice. The second return value reports
// whether the active king is currently in check.
//
// Double check restricts generation to king moves only. A single check
// restricts all other pieces via moveMask. When not in check, moveMask
// is set to EnemiesOrEmpty and all piece generators run.
func LegalMoves(moves []Move, p *Position) ([]Move, bool) {
	cpm := checkersPinsAndMask{}
	numCheckers := checkersAndPinned(p, &cpm)
	inCheck := true

	switch numCheckers {
	case 0:
		cpm.moveMask = p.EnemiesOrEmpty()
		inCheck = false
		fallthrough
	case 1:
		moves = genPawnForwardMoves(moves, p, cpm)
		moves = genPawnLeftAttackMoves(moves, p, cpm)
		moves = genPawnRightAttackMoves(moves, p, cpm)

		if p.EnPassantTarget() != SQ_NULL {
			moves = genPawnEnPassantMoves(moves, p, cpm)
		}
		moves = genKnightMoves(moves, p, cpm)
		moves = genBishopMoves(moves, p, cpm)
		moves = genRookMoves(moves, p, cpm)
		moves = genQueenMoves(moves, p, cpm)
		fallthrough
	default:
		moves = genKingMoves(moves, p)
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

	for potentialCheckers := diagonalAttackers & kingDiagonalRays; potentialCheckers != 0; {
		sq, potentialCheckers = potentialCheckers.PopLSB()

		path := lineFromTo[kingSq][sq]
		potentialyPinned := path & p.Occupied()
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
		potentialyPinned := path & p.Occupied()
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
// expanded into all four promotion piece types.
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
	for pushes := singlePush & cpm.moveMask; pushes != 0; {
		to, pushes = pushes.PopLSB()
		from = to - sp

		if to < SQ_A1 && to > SQ_H8 {
			moves = append(moves, NewMove(from, to))
		} else {
			moves = append(moves,
				NewPromotionMove(from, to, Queen),
				NewPromotionMove(from, to, Rook),
				NewPromotionMove(from, to, Bishop),
				NewPromotionMove(from, to, Knight),
			)
		}
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

// genPawnLeftAttackMoves appends all legal pawn left-diagonal capture moves
// for the active color. "Left" is toward the a-file for White, toward the
// h-file for Black. Straight-pinned pawns cannot capture. Diagonally pinned
// pawns may only capture along their pin ray. Captures on the back rank are
// expanded into all four promotion piece types.
func genPawnLeftAttackMoves(moves []Move, p *Position, cpm checkersPinsAndMask) []Move {
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
			moves = append(moves,
				NewPromotionMove(from, to, Queen),
				NewPromotionMove(from, to, Rook),
				NewPromotionMove(from, to, Bishop),
				NewPromotionMove(from, to, Knight),
			)
		}
	}

	return moves
}

// genPawnRightAttackMoves appends all legal pawn right-diagonal capture moves
// for the active color. "Right" is toward the h-file for White, toward the
// a-file for Black. Straight-pinned pawns cannot capture. Diagonally pinned
// pawns may only capture along their pin ray. Captures on the back rank are
// expanded into all four promotion piece types.
func genPawnRightAttackMoves(moves []Move, p *Position, cpm checkersPinsAndMask) []Move {
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
			moves = append(moves,
				NewPromotionMove(from, to, Queen),
				NewPromotionMove(from, to, Rook),
				NewPromotionMove(from, to, Bishop),
				NewPromotionMove(from, to, Knight),
			)
		}
	}

	return moves
}

// genPawnEnPassantMoves appends any legal en passant capture moves. Pinned
// pawns are excluded. The horizontal-pin edge case is handled explicitly: after
// removing both the capturing and the captured pawn from the occupancy, the
// king's rank is re-checked for rook or queen attacks to ensure the capture
// does not expose the king.
func genPawnEnPassantMoves(moves []Move, p *Position, cpm checkersPinsAndMask) []Move {
	const enPassantRanks = Rank_5 | Rank_4
	us := p.Active()

	kingSq, _ := p.King().PopLSB()
	enemyQueensOrRooks := p.EnemyQueensOrRooks()

	pawns := p.Pawns() &^ (cpm.diagonalPins | cpm.straightPins)
	leftAttacks := int(16*us - 9)
	rightAttacks := int(16*us - 7)
	enPassantTarget := NewBitboardFromSquare(p.EnPassantTarget())

	left := (pawns & File_Not_A).RotateLeft(leftAttacks) & enPassantTarget
	if left != 0 {
		occupiedWithoutPawns := p.Occupied() &^ (left | enPassantTarget)
		path := genRookAttacks(kingSq, occupiedWithoutPawns) & enPassantRanks

		if enemyQueensOrRooks&path == 0 {
			to, _ := left.PopLSB()
			from := to - Square(leftAttacks)
			moves = append(moves, NewMove(from, to))
		}
	}

	right := (pawns & File_Not_H).RotateLeft(rightAttacks) & enPassantTarget
	if right != 0 {
		occupiedWithoutPawns := p.Occupied() &^ (right | enPassantTarget)
		path := genRookAttacks(kingSq, occupiedWithoutPawns) & enPassantRanks

		if enemyQueensOrRooks&path == 0 {
			to, _ := right.PopLSB()
			from := to - Square(rightAttacks)
			moves = append(moves, NewMove(from, to))
		}
	}

	return moves
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

// genKingMoves appends all legal king moves including castling for the active
// color. The full enemy attack map is computed and subtracted from candidate
// targets. Castling is only added when the rights flag is set, the path is
// unoccupied, and no square the king crosses is under attack.
func genKingMoves(moves []Move, p *Position) []Move {
	us := p.Active()
	king := p.King()
	enemiesOrEmpty := p.EnemiesOrEmpty()
	from, _ := king.PopLSB()

	potentialTargets := kingMoves[from] & enemiesOrEmpty

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
