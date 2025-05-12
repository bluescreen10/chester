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

type config struct {
	singlePushes     int
	leftAttacks      int
	rightAttacks     int
	promotionRank    BitBoard
	startPlusOneRank BitBoard
	enPassantRank    BitBoard
}

var pawnConfig [2]config = [2]config{
	{-8, -9, -7, Rank_8, Rank_3, Rank_5},
	{8, 7, 9, Rank_1, Rank_6, Rank_4},
}

type checkersPinsAndMask struct {
	checkers     BitBoard
	diagonalPins BitBoard
	straightPins BitBoard
	moveMask     BitBoard
}

func LegalMoves(moves []Move, p *Position) ([]Move, bool) {
	us := p.Active()
	them := p.Inactive()
	king := p.Pieces[us][King]
	kingSq, _ := king.PopLSB()

	cpm := checkersAndPinned(p)
	//pinned := cpm.diagonalPins | cpm.straightPins

	inCheck := cpm.checkers != 0

	enemiesOrEmpty := ^p.AllPieces[us]

	switch cpm.checkers.OnesCount() {
	case 0:
		cpm.moveMask = enemiesOrEmpty
		fallthrough
	case 1:
		moves = genForwardMoves(moves, p, cpm)
		moves = genPawnMoves(moves, p, us, them, cpm, kingSq)
		moves = genKnightMoves(moves, p, cpm)
		moves = genBishopMoves(moves, p, cpm)
		moves = genRookMoves(moves, p, cpm)
		moves = genQueenMoves(moves, p, cpm)
		moves = genKingMoves(moves, p, king, enemiesOrEmpty, us, them, inCheck)
	default:
		moves = genKingMoves(moves, p, king, enemiesOrEmpty, us, them, false)
	}
	return moves, inCheck
}

func checkersAndPinned(p *Position) checkersPinsAndMask {
	us := p.Active()
	them := p.Inactive()
	king := p.Pieces[us][King]
	kingSq, _ := king.PopLSB()

	cpm := checkersPinsAndMask{}

	cpm.checkers |= knightMoves[kingSq] & p.Pieces[them][Knight]

	config := pawnConfig[us]
	cpm.checkers |= (king & File_Not_A).RotateLeft(config.leftAttacks) & p.Pieces[them][Pawn]
	cpm.checkers |= (king & File_Not_H).RotateLeft(config.rightAttacks) & p.Pieces[them][Pawn]

	kingDiagonalRays := rays[NorthWest][kingSq] | rays[NorthEast][kingSq] | rays[SouthWest][kingSq] | rays[SouthEast][kingSq]
	diagonalAttackers := (p.Pieces[them][Queen] | p.Pieces[them][Bishop])

	for potentialCheckers := diagonalAttackers & kingDiagonalRays; potentialCheckers != 0; {
		var sq Square
		sq, potentialCheckers = potentialCheckers.PopLSB()

		path := squaresBetween[sq][kingSq]
		potentialyPinned := path & p.Occupied
		switch potentialyPinned.OnesCount() {
		case 0:
			cpm.checkers |= 1 << sq
			cpm.moveMask |= path
		case 1:
			cpm.diagonalPins |= lineFromTo[kingSq][sq]
		}
	}

	kingStraightRays := rays[North][kingSq] | rays[South][kingSq] | rays[East][kingSq] | rays[West][kingSq]
	straightAttackers := (p.Pieces[them][Queen] | p.Pieces[them][Rook])

	for potentialCheckers := straightAttackers & kingStraightRays; potentialCheckers != 0; {
		var sq Square
		sq, potentialCheckers = potentialCheckers.PopLSB()

		path := squaresBetween[sq][kingSq]
		potentialyPinned := path & p.Occupied
		switch potentialyPinned.OnesCount() {
		case 0:
			cpm.checkers |= 1 << sq
			cpm.moveMask |= path
		case 1:
			cpm.straightPins |= lineFromTo[kingSq][sq]
		}
	}

	cpm.moveMask |= cpm.checkers

	return cpm
}

func genPawnsAttacks(p *Position, us Color) BitBoard {
	pawns := p.Pieces[us][Pawn]
	config := pawnConfig[us]

	left := (pawns & File_Not_A).RotateLeft(config.leftAttacks)
	right := (pawns & File_Not_H).RotateLeft(config.rightAttacks)
	return left | right
}

func genKnightsAttacks(p *Position, us Color) BitBoard {
	var attacks BitBoard

	knights := p.Pieces[us][Knight]

	for knights != 0 {
		var sq Square
		sq, knights = knights.PopLSB()
		attacks |= knightMoves[sq]
	}

	return attacks
}

func genBishopsAttacks(p *Position, us, them Color) BitBoard {
	var attacks BitBoard

	bishops := p.Pieces[us][Bishop]
	occupied := p.Occupied &^ p.Pieces[them][King]

	for bishops != 0 {
		var sq Square
		sq, bishops = bishops.PopLSB()
		attacks |= genBishopAttacks(sq, occupied)
	}

	return attacks
}

func genBishopAttacks(sq Square, occupied BitBoard) BitBoard {

	occupied &= BishopMagic[sq].Mask
	occupied *= BishopMagic[sq].Magic
	occupied >>= BishopMagic[sq].Shift
	return BishopMagic[sq].Attacks[occupied]
}

func genRooksAttacks(p *Position, us, them Color) BitBoard {
	var attacks BitBoard

	rooks := p.Pieces[us][Rook]
	occupied := p.Occupied &^ p.Pieces[them][King]

	for rooks != 0 {
		var sq Square
		sq, rooks = rooks.PopLSB()
		attacks |= genRookAttacks(sq, occupied)
	}

	return attacks
}

func genRookAttacks(sq Square, occupied BitBoard) BitBoard {
	//m := RookMagic[sq]
	occupied &= RookMagic[sq].Mask
	occupied *= RookMagic[sq].Magic
	occupied >>= RookMagic[sq].Shift
	return RookMagic[sq].Attacks[occupied]
}

func genQueensAttacks(p *Position, us, them Color) BitBoard {
	var attacks BitBoard

	queens := p.Pieces[us][Queen]
	occupied := p.Occupied &^ p.Pieces[them][King]

	for queens != 0 {
		var sq Square
		sq, queens = queens.PopLSB()
		attacks |= genBishopAttacks(sq, occupied) | genRookAttacks(sq, occupied)
	}

	return attacks
}

func genKingAttacks(p *Position, us Color) BitBoard {
	king := p.Pieces[us][King]
	sq, _ := king.PopLSB()
	return kingMoves[sq]
}

func genPawnMoves(moves []Move, p *Position, us, them Color, cpm checkersPinsAndMask, kingSq Square) []Move {
	pawns := p.Pieces[us][Pawn]
	occupied := p.Occupied
	enemies := p.AllPieces[them]
	//kingSq, _ := p.Pieces[us][King].PopLSB()
	cfg := pawnConfig[us]

	pawnsNotPinnedStraight := pawns &^ cpm.straightPins
	//pawnsNotPinnedDiagonal := pawns &^ cpm.diagonalPins

	moves = genLeftAttackMoves(moves, cfg, pawnsNotPinnedStraight, enemies, cpm.moveMask, cpm.diagonalPins)
	moves = genRightAttackMoves(moves, cfg, pawnsNotPinnedStraight, enemies, cpm.moveMask, cpm.diagonalPins)
	if p.EnPassantTarget != 0 {
		enemyQueenOrRooks := p.Pieces[them][Queen] | p.Pieces[them][Rook]
		moves = genEnPassantMoves(moves, cfg, pawns&cfg.enPassantRank, occupied, p.EnPassantTarget&cpm.moveMask, enemyQueenOrRooks, cpm.straightPins|cpm.diagonalPins, kingSq)
	}

	return moves
}

func genForwardMoves(moves []Move, p *Position, cpm checkersPinsAndMask) []Move {
	us := p.Active()
	singlePushes := -8 + 16*int(us)
	promoRank := Rank_8 | Rank_1
	startPlusOneRank := Rank_3*(1-BitBoard(us)) + Rank_6*BitBoard(us)
	//cfg := pawnConfig[us]
	pawns := p.Pieces[us][Pawn] &^ cpm.diagonalPins
	pinnedPawns := pawns & cpm.straightPins.RotateLeft(-singlePushes)
	unPinnedPawns := pawns &^ cpm.straightPins
	pawns = pinnedPawns | unPinnedPawns

	var from, to Square

	singlePush := pawns.RotateLeft(singlePushes) &^ p.Occupied
	for pushes := singlePush & cpm.moveMask &^ promoRank; pushes != 0; {
		to, pushes = pushes.PopLSB()
		from = to - Square(singlePushes)
		moves = append(moves, NewMove(from, to, Pawn))
	}

	for pushes := singlePush & cpm.moveMask & promoRank; pushes != 0; {
		to, pushes = pushes.PopLSB()
		from = to - Square(singlePushes)
		moves = append(moves,
			NewPromotionMove(from, to, Queen),
			NewPromotionMove(from, to, Rook),
			NewPromotionMove(from, to, Bishop),
			NewPromotionMove(from, to, Knight),
		)
	}

	doublePushes := (singlePush & startPlusOneRank).RotateLeft(singlePushes) &^ p.Occupied & cpm.moveMask
	for doublePushes != 0 {
		to, doublePushes = doublePushes.PopLSB()
		from = to - Square(2*singlePushes)
		moves = append(moves, NewDoublePushMove(from, to))
	}
	return moves
}

func genLeftAttackMoves(moves []Move, cfg config, pawns, enemies BitBoard, moveMask, pinned BitBoard) []Move {
	pawns = pawns & File_Not_A
	pinnedPawns := pawns & (pinned & File_Not_H).RotateLeft(-cfg.leftAttacks)
	unPinnedPawns := pawns &^ pinned
	pawns = pinnedPawns | unPinnedPawns

	pawnAttacks := pawns.RotateLeft(cfg.leftAttacks) & enemies & moveMask
	for attacks := pawnAttacks &^ cfg.promotionRank; attacks != 0; {
		var to Square
		to, attacks = attacks.PopLSB()
		from := to - Square(cfg.leftAttacks)
		moves = append(moves, NewMove(from, to, Pawn))
	}

	for attacks := pawnAttacks & cfg.promotionRank; attacks != 0; {
		var to Square
		to, attacks = attacks.PopLSB()
		from := to - Square(cfg.leftAttacks)

		moves = append(moves,
			NewPromotionMove(from, to, Queen),
			NewPromotionMove(from, to, Rook),
			NewPromotionMove(from, to, Bishop),
			NewPromotionMove(from, to, Knight),
		)

	}

	return moves
}

func genRightAttackMoves(moves []Move, cfg config, pawns, enemies BitBoard, moveMask, pinned BitBoard) []Move {
	pawns = pawns & File_Not_H
	pinnedPawns := pawns & (pinned & File_Not_A).RotateLeft(-cfg.rightAttacks)
	unPinnedPawns := pawns &^ pinned
	pawns = pinnedPawns | unPinnedPawns

	pawnAttacks := pawns.RotateLeft(cfg.rightAttacks) & enemies & moveMask
	for attacks := pawnAttacks &^ cfg.promotionRank; attacks != 0; {
		var to Square
		to, attacks = attacks.PopLSB()
		from := to - Square(cfg.rightAttacks)
		moves = append(moves, NewMove(from, to, Pawn))
	}

	for attacks := pawnAttacks & cfg.promotionRank; attacks != 0; {
		var to Square
		to, attacks = attacks.PopLSB()
		from := to - Square(cfg.rightAttacks)

		moves = append(moves,
			NewPromotionMove(from, to, Queen),
			NewPromotionMove(from, to, Rook),
			NewPromotionMove(from, to, Bishop),
			NewPromotionMove(from, to, Knight),
		)

	}

	return moves
}

func genEnPassantMoves(moves []Move, cfg config, pawnsOnRank, occupied, enPassantTarget BitBoard, enemyQueenOrRooks, pinned BitBoard, kingSq Square) []Move {
	left := pawnsOnRank & (File_Not_A & enPassantTarget >> 1)
	if left&^pinned != 0 {
		occupiedWithoutPawns := occupied &^ (left | enPassantTarget)
		path := genRookAttacks(kingSq, occupiedWithoutPawns) & cfg.enPassantRank
		if enemyQueenOrRooks&path == 0 {
			from, _ := left.PopLSB()
			to := left.RotateLeft(cfg.rightAttacks).Square()
			moves = append(moves, NewEnPassantMove(from, to))
		}
	}

	right := pawnsOnRank & (File_Not_H & enPassantTarget << 1)
	if right&^pinned != 0 {
		occupiedWithoutPawns := occupied &^ (right | enPassantTarget)
		path := genRookAttacks(kingSq, occupiedWithoutPawns) & cfg.enPassantRank

		if enemyQueenOrRooks&path == 0 {
			from, _ := right.PopLSB()
			to := right.RotateLeft(cfg.leftAttacks).Square()
			moves = append(moves, NewEnPassantMove(from, to))
		}
	}

	return moves
}

func genKnightMoves(moves []Move, p *Position, cpm checkersPinsAndMask) []Move {
	knights := p.Pieces[p.Active()][Knight] &^ (cpm.diagonalPins | cpm.straightPins)
	for knights != 0 {
		var from Square
		from, knights = knights.PopLSB()
		targets := knightMoves[from] & cpm.moveMask

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to, Knight))
		}

	}

	return moves
}

func genBishopMoves(moves []Move, p *Position, cpm checkersPinsAndMask) []Move {
	bishops := p.Pieces[p.Active()][Bishop] &^ cpm.straightPins
	for b := bishops & cpm.diagonalPins; b != 0; {
		var from Square
		from, b = b.PopLSB()
		targets := genBishopAttacks(from, p.Occupied) & cpm.moveMask & cpm.diagonalPins

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to, Bishop))
		}
	}

	for b := bishops &^ cpm.diagonalPins; b != 0; {
		var from Square
		from, b = b.PopLSB()
		targets := genBishopAttacks(from, p.Occupied) & cpm.moveMask

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to, Bishop))
		}
	}

	return moves
}

func genRookMoves(moves []Move, p *Position, cpm checkersPinsAndMask) []Move {
	rooks := p.Pieces[p.Active()][Rook] &^ cpm.diagonalPins

	for r := rooks & cpm.straightPins; r != 0; {
		var from Square
		from, r = r.PopLSB()
		targets := genRookAttacks(from, p.Occupied) & cpm.moveMask & cpm.straightPins

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to, Rook))
		}
	}

	for r := rooks &^ cpm.straightPins; r != 0; {
		var from Square
		from, r = r.PopLSB()
		targets := genRookAttacks(from, p.Occupied) & cpm.moveMask

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to, Rook))
		}
	}

	return moves
}

func genQueenMoves(moves []Move, p *Position, cpm checkersPinsAndMask) []Move {
	queens := p.Pieces[p.Active()][Queen]
	//enemiesOrEmpty := ^p.AllPieces[us] & cpm.moveMask
	occupied := p.Occupied
	pinned := cpm.straightPins | cpm.diagonalPins

	for q := queens & cpm.diagonalPins; q != 0; {
		var from Square
		from, q = q.PopLSB()
		targets := genBishopAttacks(from, occupied) & cpm.moveMask & cpm.diagonalPins

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to, Queen))
		}
	}

	for q := queens & cpm.straightPins; q != 0; {
		var from Square
		from, q = q.PopLSB()
		targets := genRookAttacks(from, occupied) & cpm.moveMask & cpm.straightPins

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to, Queen))
		}
	}

	for q := queens &^ pinned; q != 0; {
		var from Square
		from, q = q.PopLSB()
		targets := (genRookAttacks(from, occupied) | genBishopAttacks(from, occupied)) & cpm.moveMask

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to, Queen))
		}
	}

	return moves
}

func genKingMoves(moves []Move, p *Position, king, enemiesOrEmpty BitBoard, us, them Color, inCheck bool) []Move {
	// king := p.Pieces[us][King]
	//enemiesOrEmpty = ^p.AllPieces[us]
	// occupied := p.Occupied

	// normal moves
	from, _ := king.PopLSB()
	potentialTargets := kingMoves[from] & enemiesOrEmpty

	if potentialTargets == 0 {
		return moves
	}

	enemyKing := p.Pieces[them][King]
	attacked := attacks(p, them, us)

	for targets := potentialTargets &^ (enemyKing | attacked); targets != 0; {
		var to Square
		to, targets = targets.PopLSB()
		moves = append(moves, NewMove(from, to, King))
	}

	if inCheck {
		return moves
	}

	// castling
	if us == White {
		if p.CanWhiteCastleKingSide() {
			if WhiteKingSideCastleFree&p.Occupied == 0 && WhiteKingSideCastleNotAttacked&attacked == 0 {
				moves = append(moves, NewCastleKingSideMove(SQ_E1, SQ_G1))
			}
		}

		if p.CanWhiteCastleQueenSide() {

			if WhiteQueenSideCastleFree&p.Occupied == 0 && WhiteQueenSideCastleNotAttacked&attacked == 0 {
				moves = append(moves, NewCastleQueenSideMove(SQ_E1, SQ_C1))
			}
		}
	} else {
		if p.CanBlackCastleKingSide() {
			if BlackKingSideCastleFree&p.Occupied == 0 && BlackKingSideCastleNotAttacked&attacked == 0 {
				moves = append(moves, NewCastleKingSideMove(SQ_E8, SQ_G8))
			}
		}

		if p.CanBlackCastleQueenSide() {
			if BlackQueenSideCastleFree&p.Occupied == 0 && BlackQueenSideCastleNotAttacked&attacked == 0 {
				moves = append(moves, NewCastleQueenSideMove(SQ_E8, SQ_C8))
			}
		}
	}

	return moves
}

func attacks(p *Position, us, them Color) BitBoard {
	return genPawnsAttacks(p, us) |
		genKnightsAttacks(p, us) |
		genBishopsAttacks(p, us, them) |
		genRooksAttacks(p, us, them) |
		genQueensAttacks(p, us, them) |
		genKingAttacks(p, us)
}
