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
	checkers     bitboard
	diagonalPins bitboard
	straightPins bitboard
	allPins      bitboard
	moveMask     bitboard
}

func LegalMoves(moves []Move, p *Position) ([]Move, bool) {
	cpm := checkersAndPinned(p)
	inCheck := true

	switch cpm.checkers.OnesCount() {
	case 0:
		cpm.moveMask = ^p.AllPieces[p.Active()]
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

func checkersAndPinned(p *Position) checkersPinsAndMask {
	us := p.Active()
	them := p.Inactive()
	king := p.Pieces[King] & p.AllPieces[us]
	kingSq, _ := king.PopLSB()

	cpm := checkersPinsAndMask{}

	cpm.checkers |= knightMoves[kingSq] & p.Pieces[Knight] & p.AllPieces[them]

	leftAttacks := int(16*us - 9)
	rightAttacks := int(16*us - 7)
	pawns := p.AllPieces[them] & p.Pieces[Pawn]
	cpm.checkers |= (king & File_Not_A).RotateLeft(leftAttacks) & pawns
	cpm.checkers |= (king & File_Not_H).RotateLeft(rightAttacks) & pawns

	kingDiagonalRays := rays[NorthWest][kingSq] | rays[NorthEast][kingSq] | rays[SouthWest][kingSq] | rays[SouthEast][kingSq]
	diagonalAttackers := (p.Pieces[Queen] | p.Pieces[Bishop]) & p.AllPieces[them]

	var sq Square

	for potentialCheckers := diagonalAttackers & kingDiagonalRays; potentialCheckers != 0; {
		sq, potentialCheckers = potentialCheckers.PopLSB()

		path := lineFromTo[kingSq][sq]
		potentialyPinned := path & p.Occupied()
		switch potentialyPinned.OnesCount() {
		case 1:
			cpm.checkers |= 1 << sq
			cpm.moveMask |= path
		case 2:
			cpm.diagonalPins |= path
		}
	}

	kingStraightRays := rays[North][kingSq] | rays[South][kingSq] | rays[East][kingSq] | rays[West][kingSq]
	straightAttackers := (p.Pieces[Queen] | p.Pieces[Rook]) & p.AllPieces[them]

	for potentialCheckers := straightAttackers & kingStraightRays; potentialCheckers != 0; {
		sq, potentialCheckers = potentialCheckers.PopLSB()

		path := lineFromTo[kingSq][sq]
		potentialyPinned := path & p.Occupied()
		switch potentialyPinned.OnesCount() {
		case 1:
			cpm.checkers |= 1 << sq
			cpm.moveMask |= path
		case 2:
			cpm.straightPins |= path
		}
	}

	cpm.moveMask |= cpm.checkers
	cpm.allPins = cpm.diagonalPins | cpm.straightPins
	return cpm
}

func genPawnsAttacks(p *Position, us Color) bitboard {
	pawns := p.Pieces[Pawn] & p.AllPieces[us]
	//config := pawnConfig[us]
	leftAttacks := 16*int(us) - 9
	rightAttacks := 16*int(us) - 7
	left := (pawns & File_Not_A).RotateLeft(leftAttacks)
	right := (pawns & File_Not_H).RotateLeft(rightAttacks)
	return left | right
}

func genKnightsAttacks(p *Position, us Color) bitboard {
	var attacks bitboard

	knights := p.Pieces[Knight] & p.AllPieces[us]

	var sq Square

	for knights != 0 {
		sq, knights = knights.PopLSB()
		attacks |= knightMoves[sq]
	}

	return attacks
}

func genDiagonalAttacks(p *Position, us, them Color) bitboard {
	var attacks bitboard

	attacker := (p.Pieces[Bishop] | p.Pieces[Queen]) & p.AllPieces[us]
	occupied := p.Occupied() &^ (p.Pieces[King] & p.AllPieces[them])

	var sq Square

	for attacker != 0 {
		sq, attacker = attacker.PopLSB()
		attacks |= genBishopAttacks(sq, occupied)
	}

	return attacks
}

func genBishopAttacks(sq Square, occupied bitboard) bitboard {

	occupied &= BishopMagic[sq].Mask
	occupied *= BishopMagic[sq].Magic
	occupied >>= BishopMagic[sq].Shift
	return BishopMagic[sq].Attacks[occupied]
}

func genStraightAttacks(p *Position, us, them Color) bitboard {
	var attacks bitboard

	attackers := (p.Pieces[Rook] | p.Pieces[Queen]) & p.AllPieces[us]
	occupied := p.Occupied() &^ (p.Pieces[King] & p.AllPieces[them])

	var sq Square

	for attackers != 0 {
		sq, attackers = attackers.PopLSB()
		attacks |= genRookAttacks(sq, occupied)
	}

	return attacks
}

func genRookAttacks(sq Square, occupied bitboard) bitboard {
	//m := RookMagic[sq]
	occupied &= RookMagic[sq].Mask
	occupied *= RookMagic[sq].Magic
	occupied >>= RookMagic[sq].Shift
	return RookMagic[sq].Attacks[occupied]
}

func genKingAttacks(p *Position, us Color) bitboard {
	king := p.Pieces[King] & p.AllPieces[us]
	sq, _ := king.PopLSB()
	return kingMoves[sq]
}

func genPawnForwardMoves(moves []Move, p *Position, cpm *checkersPinsAndMask) []Move {
	us := p.Active()
	singlePushes := -8 + 16*int(us)
	startPlusOneRank := (Rank_3 * (1 - bitboard(us))) | (Rank_6 * bitboard(us))

	pawns := (p.Pieces[Pawn] & p.AllPieces[us]) &^ cpm.diagonalPins
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
	them := p.Inactive()
	leftAttacks := 16*int(us) - 9
	pawns := (p.Pieces[Pawn] & p.AllPieces[us]) &^ cpm.straightPins & File_Not_A
	pinnedPawns := pawns & (cpm.diagonalPins & File_Not_H).RotateLeft(-leftAttacks)
	unPinnedPawns := pawns &^ cpm.diagonalPins
	pawns = pinnedPawns | unPinnedPawns

	attacks := pawns.RotateLeft(leftAttacks) & p.AllPieces[them] & cpm.moveMask
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
	them := p.Inactive()
	rightAttacks := 16*int(us) - 7

	pawns := (p.Pieces[Pawn] & p.AllPieces[us]) &^ cpm.straightPins & File_Not_H
	pinnedPawns := pawns & (cpm.diagonalPins & File_Not_A).RotateLeft(-rightAttacks)
	unPinnedPawns := pawns &^ cpm.diagonalPins
	pawns = pinnedPawns | unPinnedPawns

	attacks := pawns.RotateLeft(rightAttacks) & p.AllPieces[them] & cpm.moveMask
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
	them := p.Inactive()

	kingSq, _ := (p.Pieces[King] & p.AllPieces[us]).PopLSB()
	enemyQueenOrRooks := (p.Pieces[Queen] | p.Pieces[Rook]) & p.AllPieces[them]

	pawnsOnRank := (p.Pieces[Pawn] & p.AllPieces[us]) &^ cpm.allPins

	left := pawnsOnRank & (File_Not_A & p.EnPassantTarget >> 1)
	if left != 0 {
		occupiedWithoutPawns := p.Occupied() &^ (left | p.EnPassantTarget)
		path := genRookAttacks(kingSq, occupiedWithoutPawns) & enPassantRanks

		if enemyQueenOrRooks&path == 0 {
			from, _ := left.PopLSB()
			to := from + 16*Square(us) - 7
			moves = append(moves, NewEnPassantMove(from, to))
		}
	}

	right := pawnsOnRank & (File_Not_H & p.EnPassantTarget << 1)
	if right != 0 {
		occupiedWithoutPawns := p.Occupied() &^ (right | p.EnPassantTarget)
		path := genRookAttacks(kingSq, occupiedWithoutPawns) & enPassantRanks

		if enemyQueenOrRooks&path == 0 {
			from, _ := right.PopLSB()
			to := from + 16*Square(us) - 9
			moves = append(moves, NewEnPassantMove(from, to))
		}
	}

	return moves
}

func genKnightMoves(moves []Move, p *Position, cpm *checkersPinsAndMask) []Move {
	knights := (p.AllPieces[p.Active()] & p.Pieces[Knight]) &^ cpm.allPins
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
	bishops := (p.AllPieces[p.Active()] & p.Pieces[Bishop]) &^ cpm.straightPins

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
	rooks := (p.AllPieces[p.Active()] & p.Pieces[Rook]) &^ cpm.diagonalPins

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
	queens := (p.AllPieces[p.Active()] & p.Pieces[Queen])
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
	king := (p.AllPieces[us] & p.Pieces[King])
	enemiesOrEmpty := ^p.AllPieces[us]
	from, _ := king.PopLSB()

	potentialTargets := kingMoves[from] & enemiesOrEmpty

	if potentialTargets == 0 {
		return moves
	}

	them := p.Inactive()
	enemyKing := (p.AllPieces[them] & p.Pieces[King])
	attacked := attacks(p, them, us)

	for targets := potentialTargets &^ (enemyKing | attacked); targets != 0; {
		var to Square
		to, targets = targets.PopLSB()
		moves = append(moves, NewMove(from, to, King))
	}

	// castling
	if us == White &&
		p.CanWhiteCastleKingSide() &&
		WhiteKingSideCastleFree&p.Occupied() == 0 &&
		WhiteKingSideCastleNotAttacked&attacked == 0 {
		moves = append(moves, NewCastleKingSideMove(SQ_E1, SQ_G1))
	}

	if us == White &&
		p.CanWhiteCastleQueenSide() &&
		WhiteQueenSideCastleFree&p.Occupied() == 0 &&
		WhiteQueenSideCastleNotAttacked&attacked == 0 {
		moves = append(moves, NewCastleQueenSideMove(SQ_E1, SQ_C1))
	}

	if us == Black &&
		p.CanBlackCastleKingSide() &&
		BlackKingSideCastleFree&p.Occupied() == 0 &&
		BlackKingSideCastleNotAttacked&attacked == 0 {
		moves = append(moves, NewCastleKingSideMove(SQ_E8, SQ_G8))
	}

	if us == Black &&
		p.CanBlackCastleQueenSide() &&
		BlackQueenSideCastleFree&p.Occupied() == 0 &&
		BlackQueenSideCastleNotAttacked&attacked == 0 {
		moves = append(moves, NewCastleQueenSideMove(SQ_E8, SQ_C8))
	}

	return moves
}

func attacks(p *Position, us, them Color) bitboard {
	return genPawnsAttacks(p, us) |
		genKnightsAttacks(p, us) |
		genDiagonalAttacks(p, us, them) |
		genStraightAttacks(p, us, them) |
		genKingAttacks(p, us)
}
