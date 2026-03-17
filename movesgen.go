package main

const (
	WhiteQueenSideCastleFree        = BB_SQ_B1 | BB_SQ_C1 | BB_SQ_D1
	WhiteQueenSideCastleNotAttacked = BB_SQ_C1 | BB_SQ_D1 | BB_SQ_E1
	WhiteKingSideCastleFree         = BB_SQ_F1 | BB_SQ_G1
	WhiteKingSideCastleNotAttacked  = BB_SQ_E1 | BB_SQ_F1 | BB_SQ_G1

	BlackQueenSideCastleFree        = BB_SQ_B8 | BB_SQ_C8 | BB_SQ_D8
	BlackQueenSideCastleNotAttacked = BB_SQ_C8 | BB_SQ_D8 | BB_SQ_E8
	BlackKingSideCastleFree         = BB_SQ_F8 | BB_SQ_G8
	BlackKingSideCastleNotAttacked  = BB_SQ_E8 | BB_SQ_F8 | BB_SQ_G8
)

type checkersPinsAndMask struct {
	checkers     Bitboard
	diagonalPins Bitboard
	straightPins Bitboard
	allPins      Bitboard
	moveMask     Bitboard
}

func LegalMoves(moves []Move, p *Position) ([]Move, bool) {
	cpm := checkersPinsAndMask{}
	checkersAndPinned(p, &cpm)
	inCheck := true

	switch cpm.checkers.OnesCount() {
	case 0:
		cpm.moveMask = p.EnemiesOrEmpty()
		inCheck = false
		fallthrough
	case 1:
		moves = genPawnForwardMoves(moves, p, &cpm)
		moves = genPawnLeftAttackMoves(moves, p, &cpm)
		moves = genPawnRightAttackMoves(moves, p, &cpm)

		if p.EnPassantTarget != 0 {
			moves = genPawnEnPassantMoves(moves, p, &cpm)
		}
		moves = genKnightMoves(moves, p, &cpm)
		moves = genBishopMoves(moves, p, &cpm)
		moves = genRookMoves(moves, p, &cpm)
		moves = genQueenMoves(moves, p, &cpm)
		fallthrough
	default:
		moves = genKingMoves(moves, p)
	}
	return moves, inCheck
}

func checkersAndPinned(p *Position, cpm *checkersPinsAndMask) {
	us := p.Active()
	king := p.King()
	kingSq, _ := king.PopLSB()

	//cpm := checkersPinsAndMask{}

	cpm.checkers |= knightMoves[kingSq] & p.EnemyKnights()

	leftAttacks := int(16*us - 9)
	rightAttacks := int(16*us - 7)
	pawns := p.EnemyPawns()
	cpm.checkers |= (king & File_Not_A).RotateLeft(leftAttacks) & pawns
	cpm.checkers |= (king & File_Not_H).RotateLeft(rightAttacks) & pawns

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
				cpm.checkers |= 1 << sq
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
				cpm.checkers |= 1 << sq
				cpm.moveMask |= path
			case 2:
				cpm.straightPins |= path
			}
		}
	}

	cpm.moveMask |= cpm.checkers
	cpm.allPins = cpm.diagonalPins | cpm.straightPins
}

func genPawnsAttacks(p *Position) Bitboard {
	pawns := p.EnemyPawns()
	//config := pawnConfig[us]
	color := p.Inactive()
	leftAttacks := 16*int(color) - 9
	rightAttacks := 16*int(color) - 7
	left := (pawns & File_Not_A).RotateLeft(leftAttacks)
	right := (pawns & File_Not_H).RotateLeft(rightAttacks)
	return left | right
}

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

func genBishopAttacks(sq Square, occupied Bitboard) Bitboard {

	occupied &= BishopMagic[sq].Mask
	occupied *= BishopMagic[sq].Magic
	occupied >>= BishopMagic[sq].Shift
	return BishopMagic[sq].Attacks[occupied]
}

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

func genRookAttacks(sq Square, occupied Bitboard) Bitboard {
	//m := RookMagic[sq]
	occupied &= RookMagic[sq].Mask
	occupied *= RookMagic[sq].Magic
	occupied >>= RookMagic[sq].Shift
	return RookMagic[sq].Attacks[occupied]
}

func genKingAttacks(p *Position) Bitboard {
	king := p.EnemyKing()
	sq, _ := king.PopLSB()
	return kingMoves[sq]
}

func genPawnForwardMoves(moves []Move, p *Position, cpm *checkersPinsAndMask) []Move {
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
			moves = append(moves, NewMove(from, to, Pawn))
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
		moves = append(moves, NewDoublePushMove(from, to))
	}
	return moves
}

func genPawnLeftAttackMoves(moves []Move, p *Position, cpm *checkersPinsAndMask) []Move {
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
			moves = append(moves, NewMove(from, to, Pawn))
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

func genPawnRightAttackMoves(moves []Move, p *Position, cpm *checkersPinsAndMask) []Move {
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
			moves = append(moves, NewMove(from, to, Pawn))
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

func genPawnEnPassantMoves(moves []Move, p *Position, cpm *checkersPinsAndMask) []Move {
	const enPassantRanks = Rank_5 | Rank_4
	us := p.Active()

	kingSq, _ := p.King().PopLSB()
	enemyQueensOrRooks := p.EnemyQueensOrRooks()

	pawnsOnRank := p.Pawns() &^ cpm.allPins

	left := pawnsOnRank & (File_Not_A & p.EnPassantTarget >> 1)
	if left != 0 {
		occupiedWithoutPawns := p.Occupied() &^ (left | p.EnPassantTarget)
		path := genRookAttacks(kingSq, occupiedWithoutPawns) & enPassantRanks

		if enemyQueensOrRooks&path == 0 {
			from, _ := left.PopLSB()
			to := from + 16*Square(us) - 7
			moves = append(moves, NewEnPassantMove(from, to))
		}
	}

	right := pawnsOnRank & (File_Not_H & p.EnPassantTarget << 1)
	if right != 0 {
		occupiedWithoutPawns := p.Occupied() &^ (right | p.EnPassantTarget)
		path := genRookAttacks(kingSq, occupiedWithoutPawns) & enPassantRanks

		if enemyQueensOrRooks&path == 0 {
			from, _ := right.PopLSB()
			to := from + 16*Square(us) - 9
			moves = append(moves, NewEnPassantMove(from, to))
		}
	}

	return moves
}

func genKnightMoves(moves []Move, p *Position, cpm *checkersPinsAndMask) []Move {
	knights := p.Knights() &^ cpm.allPins
	var from, to Square
	for knights != 0 {
		from, knights = knights.PopLSB()
		targets := knightMoves[from] & cpm.moveMask

		for targets != 0 {
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to, Knight))
		}

	}

	return moves
}

func genBishopMoves(moves []Move, p *Position, cpm *checkersPinsAndMask) []Move {
	bishops := p.Bishops() &^ cpm.straightPins

	var from, to Square
	for b := bishops & cpm.diagonalPins; b != 0; {
		from, b = b.PopLSB()
		targets := genBishopAttacks(from, p.Occupied()) & cpm.moveMask & cpm.diagonalPins

		for targets != 0 {
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to, Bishop))
		}
	}

	for b := bishops &^ cpm.diagonalPins; b != 0; {
		from, b = b.PopLSB()
		targets := genBishopAttacks(from, p.Occupied()) & cpm.moveMask

		for targets != 0 {
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to, Bishop))
		}
	}

	return moves
}

func genRookMoves(moves []Move, p *Position, cpm *checkersPinsAndMask) []Move {
	rooks := p.Rooks() &^ cpm.diagonalPins

	var from, to Square
	for r := rooks & cpm.straightPins; r != 0; {
		from, r = r.PopLSB()
		targets := genRookAttacks(from, p.Occupied()) & cpm.moveMask & cpm.straightPins

		for targets != 0 {
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to, Rook))
		}
	}

	for r := rooks &^ cpm.straightPins; r != 0; {
		from, r = r.PopLSB()
		targets := genRookAttacks(from, p.Occupied()) & cpm.moveMask

		for targets != 0 {
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to, Rook))
		}
	}

	return moves
}

func genQueenMoves(moves []Move, p *Position, cpm *checkersPinsAndMask) []Move {
	queens := p.Queens()
	//enemiesOrEmpty := ^p.AllPieces[us] & cpm.moveMask
	occupied := p.Occupied()

	var from, to Square
	for q := queens & cpm.diagonalPins; q != 0; {
		from, q = q.PopLSB()
		targets := genBishopAttacks(from, occupied) & cpm.moveMask & cpm.diagonalPins

		for targets != 0 {
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to, Queen))
		}
	}

	for q := queens & cpm.straightPins; q != 0; {
		from, q = q.PopLSB()
		targets := genRookAttacks(from, occupied) & cpm.moveMask & cpm.straightPins

		for targets != 0 {
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to, Queen))
		}
	}

	for q := queens &^ cpm.allPins; q != 0; {
		from, q = q.PopLSB()
		targets := (genRookAttacks(from, occupied) | genBishopAttacks(from, occupied)) & cpm.moveMask

		for targets != 0 {
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to, Queen))
		}
	}

	return moves
}

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
		moves = append(moves, NewMove(from, to, King))
	}

	// castling
	if us == White {
		if p.CanWhiteCastleKingSide() &&
			WhiteKingSideCastleFree&p.Occupied() == 0 &&
			WhiteKingSideCastleNotAttacked&attacked == 0 {
			moves = append(moves, NewCastleKingSideMove(SQ_E1, SQ_G1))
		}

		if p.CanWhiteCastleQueenSide() &&
			WhiteQueenSideCastleFree&p.Occupied() == 0 &&
			WhiteQueenSideCastleNotAttacked&attacked == 0 {
			moves = append(moves, NewCastleQueenSideMove(SQ_E1, SQ_C1))
		}
	}

	if us == Black {
		if p.CanBlackCastleKingSide() &&
			BlackKingSideCastleFree&p.Occupied() == 0 &&
			BlackKingSideCastleNotAttacked&attacked == 0 {
			moves = append(moves, NewCastleKingSideMove(SQ_E8, SQ_G8))
		}

		if p.CanBlackCastleQueenSide() &&
			BlackQueenSideCastleFree&p.Occupied() == 0 &&
			BlackQueenSideCastleNotAttacked&attacked == 0 {
			moves = append(moves, NewCastleQueenSideMove(SQ_E8, SQ_C8))
		}
	}

	return moves
}

func attacks(p *Position) Bitboard {
	return genPawnsAttacks(p) |
		genKnightsAttacks(p) |
		genDiagonalAttacks(p) |
		genStraightAttacks(p) |
		genKingAttacks(p)
}
