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

func LegalMoves(moves []Move, pos *Position) ([]Move, bool) {
	us, them := pos.SideToMove()
	king := pos.Pieces[us][King]
	kingSq, _ := king.PopLSB()

	checkers, pinnedStraight, pinnedDiagonal := checkersAndPinned(pos, us, them, kingSq, king)
	pinned := pinnedStraight | pinnedDiagonal

	inCheck := checkers != 0

	mask := FullBoard
	enemiesOrEmpty := ^pos.AllPieces[us]

	switch checkers.OnesCount() {
	case 1:
		captureMask := checkers
		pushMask := EmptyBoard

		if (pos.Pieces[them][Bishop]|pos.Pieces[them][Rook]|pos.Pieces[them][Queen])&checkers != 0 {
			checkerSq, _ := checkers.PopLSB()
			pushMask = squaresBetween[checkerSq][kingSq]
		}
		mask = captureMask | pushMask
		fallthrough
	case 0:
		// FIX ME: currently not working correctly because genBishopMoves and genRookMoves
		// assume those are only bishop and rooks, and in the move they mark queens as
		// either rooks or bishops
		//bishopsAndQueens := (pos.Pieces[us][Bishop] | pos.Pieces[us][Queen]) &^ pinnedStraight
		//rookAndQueens := (pos.Pieces[us][Rook] | pos.Pieces[us][Queen]) &^ pinnedDiagonal
		bishopsAndQueens := (pos.Pieces[us][Bishop]) &^ pinnedStraight
		rookAndQueens := (pos.Pieces[us][Rook]) &^ pinnedDiagonal

		moves = genPawnMoves(moves, pos, us, them, mask, pinnedStraight, pinnedDiagonal, kingSq)
		moves = genKnightMoves(moves, pos.Pieces[us][Knight]&^pinned, enemiesOrEmpty&mask)
		moves = genBishopMoves(moves, bishopsAndQueens, pos.Occupied, enemiesOrEmpty&mask, pinnedDiagonal)
		moves = genRookMoves(moves, rookAndQueens, pos.Occupied, enemiesOrEmpty&mask, pinnedStraight)
		moves = genQueenMoves(moves, pos, us, mask, pinnedStraight, pinnedDiagonal)
		moves = genKingMoves(moves, king, pos.Occupied, enemiesOrEmpty, pos, us, them, inCheck)
	default:
		moves = genKingMoves(moves, king, pos.Occupied, enemiesOrEmpty, pos, us, them, false)
	}
	return moves, inCheck
}

func checkersAndPinned(p *Position, us, them Color, kingSq Square, king BitBoard) (BitBoard, BitBoard, BitBoard) {
	// king := p.Pieces[us][King]
	// kingSq, _ := king.PopLSB()

	checkers := knightMoves[kingSq] & p.Pieces[them][Knight]
	config := pawnConfig[us]
	checkers |= (king & File_Not_A).RotateLeft(config.leftAttacks) & p.Pieces[them][Pawn]
	checkers |= (king & File_Not_H).RotateLeft(config.rightAttacks) & p.Pieces[them][Pawn]

	kingDiagonalRays := rays[NorthWest][kingSq] | rays[NorthEast][kingSq] | rays[SouthWest][kingSq] | rays[SouthEast][kingSq]
	diagonalAttackers := (p.Pieces[them][Queen] | p.Pieces[them][Bishop])

	occupied := p.Occupied

	var pinnedStraight, pinnedDiagonal BitBoard

	for potentialCheckers := diagonalAttackers & kingDiagonalRays; potentialCheckers != 0; {
		var sq Square
		sq, potentialCheckers = potentialCheckers.PopLSB()

		path := squaresBetween[sq][kingSq]
		potentialyPinned := path & occupied
		switch potentialyPinned.OnesCount() {
		case 0:
			checkers |= 1 << sq
		case 1:

			pinnedDiagonal |= lineFromTo[kingSq][sq]
		}
	}

	kingStraightRays := rays[North][kingSq] | rays[South][kingSq] | rays[East][kingSq] | rays[West][kingSq]
	straightAttackers := (p.Pieces[them][Queen] | p.Pieces[them][Rook])

	for potentialCheckers := straightAttackers & kingStraightRays; potentialCheckers != 0; {
		var sq Square
		sq, potentialCheckers = potentialCheckers.PopLSB()

		path := squaresBetween[sq][kingSq]
		potentialyPinned := path & occupied
		switch potentialyPinned.OnesCount() {
		case 0:
			checkers |= 1 << sq
		case 1:
			pinnedStraight |= lineFromTo[kingSq][sq]
		}
	}

	return checkers, pinnedStraight, pinnedDiagonal
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

func genPawnMoves(moves []Move, p *Position, us, them Color, moveMask, pinnedStraight, pinnedDiagonal BitBoard, kingSq Square) []Move {
	pawns := p.Pieces[us][Pawn]
	occupied := p.Occupied
	enemies := p.AllPieces[them]
	//kingSq, _ := p.Pieces[us][King].PopLSB()
	cfg := pawnConfig[us]

	pawnsNotPinnedStraight := pawns &^ pinnedStraight
	pawnsNotPinnedDiagonal := pawns &^ pinnedDiagonal

	moves = genForwardMoves(moves, cfg, pawnsNotPinnedDiagonal, occupied, moveMask, pinnedStraight)
	moves = genLeftAttackMoves(moves, cfg, pawnsNotPinnedStraight, enemies, moveMask, pinnedDiagonal)
	moves = genRightAttackMoves(moves, cfg, pawnsNotPinnedStraight, enemies, moveMask, pinnedDiagonal)
	if p.EnPassantTarget != 0 {
		enemyQueenOrRooks := p.Pieces[them][Queen] | p.Pieces[them][Rook]
		moves = genEnPassantMoves(moves, cfg, pawns&cfg.enPassantRank, occupied, p.EnPassantTarget&moveMask, enemyQueenOrRooks, pinnedStraight|pinnedDiagonal, kingSq)
	}

	return moves
}

func genForwardMoves(moves []Move, cfg config, pawns, occupied BitBoard, moveMask, pinned BitBoard) []Move {
	pinnedPawns := pawns & pinned.RotateLeft(-cfg.singlePushes)
	unPinnedPawns := pawns &^ pinned
	pawns = pinnedPawns | unPinnedPawns

	singlePushes := pawns.RotateLeft(cfg.singlePushes) &^ occupied
	for pushes := singlePushes & moveMask &^ cfg.promotionRank; pushes != 0; {
		var to Square
		to, pushes = pushes.PopLSB()
		from := to - Square(cfg.singlePushes)
		moves = append(moves, NewMove(from, to, Pawn))
	}

	for pushes := singlePushes & moveMask & cfg.promotionRank; pushes != 0; {
		var to Square
		to, pushes = pushes.PopLSB()
		from := to - Square(cfg.singlePushes)
		moves = append(moves,
			NewPromotionMove(from, to, Queen),
			NewPromotionMove(from, to, Rook),
			NewPromotionMove(from, to, Bishop),
			NewPromotionMove(from, to, Knight),
		)
	}

	doublePushes := (singlePushes & cfg.startPlusOneRank).RotateLeft(cfg.singlePushes) &^ occupied & moveMask
	for doublePushes != 0 {
		var to Square
		to, doublePushes = doublePushes.PopLSB()
		from := to - Square(2*cfg.singlePushes)
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

func genKnightMoves(moves []Move, knights, enemiesOrEmpty BitBoard) []Move {
	for knights != 0 {
		var from Square
		from, knights = knights.PopLSB()
		targets := knightMoves[from] & enemiesOrEmpty

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to, Knight))
		}

	}

	return moves
}

func genBishopMoves(moves []Move, bishops, occupied, enemiesOrEmpty, pinned BitBoard) []Move {
	for b := bishops & pinned; b != 0; {
		var from Square
		from, b = b.PopLSB()
		targets := genBishopAttacks(from, occupied) & enemiesOrEmpty & pinned

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to, Bishop))
		}
	}

	for b := bishops &^ pinned; b != 0; {
		var from Square
		from, b = b.PopLSB()
		targets := genBishopAttacks(from, occupied) & enemiesOrEmpty

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to, Bishop))
		}
	}

	return moves
}

func genRookMoves(moves []Move, rooks, occupied, enemiesOrEmpty, pinned BitBoard) []Move {
	//rooks := p.Pieces[us][Rook] &^ pinnedDiagonal
	//enemiesOrEmpty := ^p.AllPieces[us] & moveMask
	//occupied := p.Occupied

	for r := rooks & pinned; r != 0; {
		var from Square
		from, r = r.PopLSB()
		targets := genRookAttacks(from, occupied) & enemiesOrEmpty & pinned

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to, Rook))
		}
	}

	for r := rooks &^ pinned; r != 0; {
		var from Square
		from, r = r.PopLSB()
		targets := genRookAttacks(from, occupied) & enemiesOrEmpty

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to, Rook))
		}
	}

	return moves
}

func genQueenMoves(moves []Move, p *Position, us Color, moveMask, pinnedStraight, pinnedDiagonal BitBoard) []Move {
	queens := p.Pieces[us][Queen]
	enemiesOrEmpty := ^p.AllPieces[us] & moveMask
	occupied := p.Occupied
	pinned := pinnedStraight | pinnedDiagonal

	for q := queens & pinnedDiagonal; q != 0; {
		var from Square
		from, q = q.PopLSB()
		targets := genBishopAttacks(from, occupied) & enemiesOrEmpty & pinnedDiagonal

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to, Queen))
		}
	}

	for q := queens & pinnedStraight; q != 0; {
		var from Square
		from, q = q.PopLSB()
		targets := genRookAttacks(from, occupied) & enemiesOrEmpty & pinnedStraight

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to, Queen))
		}
	}

	for q := queens &^ pinned; q != 0; {
		var from Square
		from, q = q.PopLSB()
		targets := (genRookAttacks(from, occupied) | genBishopAttacks(from, occupied)) & enemiesOrEmpty

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			moves = append(moves, NewMove(from, to, Queen))
		}
	}

	return moves
}

func genKingMoves(moves []Move, king, occupied, enemiesOrEmpty BitBoard, p *Position, us, them Color, inCheck bool) []Move {
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
			if WhiteKingSideCastleFree&occupied == 0 && WhiteKingSideCastleNotAttacked&attacked == 0 {
				moves = append(moves, NewCastleKingSideMove(SQ_E1, SQ_G1))
			}
		}

		if p.CanWhiteCastleQueenSide() {

			if WhiteQueenSideCastleFree&occupied == 0 && WhiteQueenSideCastleNotAttacked&attacked == 0 {
				moves = append(moves, NewCastleQueenSideMove(SQ_E1, SQ_C1))
			}
		}
	} else {
		if p.CanBlackCastleKingSide() {
			if BlackKingSideCastleFree&occupied == 0 && BlackKingSideCastleNotAttacked&attacked == 0 {
				moves = append(moves, NewCastleKingSideMove(SQ_E8, SQ_G8))
			}
		}

		if p.CanBlackCastleQueenSide() {
			if BlackQueenSideCastleFree&occupied == 0 && BlackQueenSideCastleNotAttacked&attacked == 0 {
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
