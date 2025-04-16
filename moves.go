package main

import (
	"fmt"
	"sync"
)

const (
	maxMoves = 218
)

type MoveType uint8

const (
	Default MoveType = iota
	Castle
	EnPassant
	Promotion
)

// Eventualy we want to use uint16
// First 3 bits move type
// Next 6 bits from square
// Last 6 bits to square
type Move struct {
	Piece, PromotionPiece Piece
	Type                  MoveType
	From, To              Square
}

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

func ParseMove(m string, p Position) (Move, error) {
	from := SquareFromString(m[:2])
	to := SquareFromString(m[2:4])
	if len(m) == 5 {
		switch m[4] {
		case 'b':
			return Move{Piece: Pawn, From: from, To: to, Type: Promotion, PromotionPiece: Bishop}, nil
		case 'n':
			return Move{Piece: Pawn, From: from, To: to, Type: Promotion, PromotionPiece: Knight}, nil
		case 'r':
			return Move{Piece: Pawn, From: from, To: to, Type: Promotion, PromotionPiece: Rook}, nil
		case 'q':
			return Move{Piece: Pawn, From: from, To: to, Type: Promotion, PromotionPiece: Queen}, nil
		default:
			return Move{}, fmt.Errorf("invalid move suffix")
		}
	} else {
		piece := p.Get(from)
		if piece == Empty {
			return Move{}, fmt.Errorf("no piece at from square: %s", from)
		}
		return Move{Piece: piece, From: from, To: to}, nil
	}
}

func (m Move) String() string {
	fromRank, fromFile := m.From.RankAndFile()
	toRank, toFile := m.To.RankAndFile()

	suffix := ""

	if m.Type == Promotion {
		switch m.PromotionPiece {
		case Bishop:
			suffix = "b"
		case Knight:
			suffix = "n"
		case Rook:
			suffix = "r"
		case Queen:
			suffix = "q"
		default:
		}
	}

	return fmt.Sprintf("%c%d%c%d%s", 'a'+fromFile, fromRank+1, 'a'+toFile, toRank+1, suffix)
}

var movesPool = sync.Pool{
	New: func() any {
		var s []Move
		return &s
	},
}

func LegalMoves(moves *[]Move, pos *Position) bool {
	us, them := pos.SideToMove()
	king := pos.Pieces[us][King]
	kingSq, _ := king.PopLSB()

	checkers, pinnedStraight, pinnedDiagonal := checkersAndPinned(pos, us, them, kingSq, king)
	pinned := pinnedStraight | pinnedDiagonal

	switch c := checkers.OnesCount(); {
	case c >= 2:
		genKingMoves(moves, pos, us, them, false)
		return true
	case c == 1:
		captureMask := checkers
		pushMask := EmptyBoard

		if (pos.Pieces[them][Bishop]|pos.Pieces[them][Rook]|pos.Pieces[them][Queen])&checkers != 0 {
			checkerSq, _ := checkers.PopLSB()
			pushMask = squaresBetween[checkerSq][kingSq]
		}

		genPawnMoves(moves, pos, us, them, captureMask|pushMask, pinnedStraight, pinnedDiagonal, kingSq)
		genKnightMoves(moves, pos, us, captureMask|pushMask, pinned)
		genBishopMoves(moves, pos, us, captureMask|pushMask, pinnedStraight, pinnedDiagonal)
		genRookMoves(moves, pos, us, captureMask|pushMask, pinnedStraight, pinnedStraight)
		genQueenMoves(moves, pos, us, captureMask|pushMask, pinnedStraight, pinnedDiagonal)
		genKingMoves(moves, pos, us, them, false)
		return true
	default:
		genPawnMoves(moves, pos, us, them, FullBoard, pinnedStraight, pinnedDiagonal, kingSq)
		genKnightMoves(moves, pos, us, FullBoard, pinned)
		genBishopMoves(moves, pos, us, FullBoard, pinnedStraight, pinnedDiagonal)
		genRookMoves(moves, pos, us, FullBoard, pinnedStraight, pinnedDiagonal)
		genQueenMoves(moves, pos, us, FullBoard, pinnedStraight, pinnedDiagonal)
		genKingMoves(moves, pos, us, them, true)
		return false
	}
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
	friends := p.AllPieces[us]

	for potentialCheckers := diagonalAttackers & kingDiagonalRays; potentialCheckers != 0; {
		var sq Square
		sq, potentialCheckers = potentialCheckers.PopLSB()

		path := squaresBetween[sq][kingSq]
		potentialyPinned := path & occupied
		switch potentialyPinned.OnesCount() {
		case 0:
			checkers |= 1 << sq
		case 1:
			if path&friends != 0 {
				pinnedDiagonal |= lineFromTo[kingSq][sq]
			}
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
			if path&friends != 0 {
				pinnedStraight |= lineFromTo[kingSq][sq]
			}
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

func genPawnMoves(moves *[]Move, p *Position, us, them Color, moveMask, pinnedStraight, pinnedDiagonal BitBoard, kingSq Square) {
	pawns := p.Pieces[us][Pawn]
	occupied := p.Occupied
	enemies := p.AllPieces[them]
	//kingSq, _ := p.Pieces[us][King].PopLSB()
	enemyQueenOrRooks := p.Pieces[them][Queen] | p.Pieces[them][Rook]
	cfg := pawnConfig[us]

	pawnsNotPinnedStraight := pawns &^ pinnedStraight
	pawnsNotPinnedDiagonal := pawns &^ pinnedDiagonal

	genForwardMoves(moves, cfg, pawnsNotPinnedDiagonal, occupied, moveMask, pinnedStraight)
	genLeftAttackMoves(moves, cfg, pawnsNotPinnedStraight, enemies, moveMask, pinnedDiagonal)
	genRightAttackMoves(moves, cfg, pawnsNotPinnedStraight, enemies, moveMask, pinnedDiagonal)
	if p.IsEnPassant() {
		genEnPassantMoves(moves, cfg, pawns, occupied, p.EnPassantFile(), p.EnPassantSquare, enemyQueenOrRooks, moveMask, pinnedStraight|pinnedDiagonal, kingSq)
	}
}

func genForwardMoves(moves *[]Move, cfg config, pawns, occupied BitBoard, moveMask, pinned BitBoard) {
	pinnedPawns := pawns & pinned.RotateLeft(-cfg.singlePushes)
	unPinnedPawns := pawns &^ pinned
	pawns = pinnedPawns | unPinnedPawns

	singlePushes := pawns.RotateLeft(cfg.singlePushes) &^ occupied
	for pushes := singlePushes & moveMask &^ cfg.promotionRank; pushes != 0; {
		var to Square
		to, pushes = pushes.PopLSB()
		from := to - Square(cfg.singlePushes)
		*moves = append(*moves, Move{Piece: Pawn, From: from, To: to})
	}

	for pushes := singlePushes & moveMask & cfg.promotionRank; pushes != 0; {
		var to Square
		to, pushes = pushes.PopLSB()
		from := to - Square(cfg.singlePushes)
		*moves = append(*moves,
			Move{Piece: Pawn, From: from, To: to, Type: Promotion, PromotionPiece: Queen},
			Move{Piece: Pawn, From: from, To: to, Type: Promotion, PromotionPiece: Rook},
			Move{Piece: Pawn, From: from, To: to, Type: Promotion, PromotionPiece: Bishop},
			Move{Piece: Pawn, From: from, To: to, Type: Promotion, PromotionPiece: Knight},
		)
	}

	doublePushes := (singlePushes & cfg.startPlusOneRank).RotateLeft(cfg.singlePushes) &^ occupied & moveMask
	for doublePushes != 0 {
		var to Square
		to, doublePushes = doublePushes.PopLSB()
		from := to - Square(2*cfg.singlePushes)
		*moves = append(*moves, Move{Piece: Pawn, From: from, To: to})
	}
}

func genLeftAttackMoves(moves *[]Move, cfg config, pawns, enemies BitBoard, moveMask, pinned BitBoard) {
	pawns = pawns & File_Not_A
	pinnedPawns := pawns & (pinned & File_Not_H).RotateLeft(-cfg.leftAttacks)
	unPinnedPawns := pawns &^ pinned
	pawns = pinnedPawns | unPinnedPawns

	pawnAttacks := pawns.RotateLeft(cfg.leftAttacks) & enemies & moveMask
	for attacks := pawnAttacks &^ cfg.promotionRank; attacks != 0; {
		var to Square
		to, attacks = attacks.PopLSB()
		from := to - Square(cfg.leftAttacks)
		*moves = append(*moves, Move{Piece: Pawn, From: from, To: to})
	}

	for attacks := pawnAttacks & cfg.promotionRank; attacks != 0; {
		var to Square
		to, attacks = attacks.PopLSB()
		from := to - Square(cfg.leftAttacks)

		*moves = append(*moves,
			Move{Piece: Pawn, From: from, To: to, Type: Promotion, PromotionPiece: Queen},
			Move{Piece: Pawn, From: from, To: to, Type: Promotion, PromotionPiece: Rook},
			Move{Piece: Pawn, From: from, To: to, Type: Promotion, PromotionPiece: Bishop},
			Move{Piece: Pawn, From: from, To: to, Type: Promotion, PromotionPiece: Knight},
		)

	}
}

func genRightAttackMoves(moves *[]Move, cfg config, pawns, enemies BitBoard, moveMask, pinned BitBoard) {
	pawns = pawns & File_Not_H
	pinnedPawns := pawns & (pinned & File_Not_A).RotateLeft(-cfg.rightAttacks)
	unPinnedPawns := pawns &^ pinned
	pawns = pinnedPawns | unPinnedPawns

	pawnAttacks := pawns.RotateLeft(cfg.rightAttacks) & enemies & moveMask
	for attacks := pawnAttacks &^ cfg.promotionRank; attacks != 0; {
		var to Square
		to, attacks = attacks.PopLSB()
		from := to - Square(cfg.rightAttacks)
		*moves = append(*moves, Move{Piece: Pawn, From: from, To: to})
	}

	for attacks := pawnAttacks & cfg.promotionRank; attacks != 0; {
		var to Square
		to, attacks = attacks.PopLSB()
		from := to - Square(cfg.rightAttacks)

		*moves = append(*moves,
			Move{Piece: Pawn, From: from, To: to, Type: Promotion, PromotionPiece: Queen},
			Move{Piece: Pawn, From: from, To: to, Type: Promotion, PromotionPiece: Rook},
			Move{Piece: Pawn, From: from, To: to, Type: Promotion, PromotionPiece: Bishop},
			Move{Piece: Pawn, From: from, To: to, Type: Promotion, PromotionPiece: Knight},
		)

	}
}

func genEnPassantMoves(moves *[]Move, cfg config, pawns, occupied, enPassantFile BitBoard, enPassantSquare Square, enemyQueenOrRooks, moveMask, pinned BitBoard, kingSq Square) {
	pawnsOnRank := pawns & cfg.enPassantRank
	left := (pawnsOnRank &^ File_A).RotateLeft(cfg.leftAttacks) & enPassantFile & moveMask
	if left&^pinned != 0 {
		sq, _ := left.PopLSB()
		from := sq - Square(cfg.leftAttacks)
		occupiedWithoutPawns := occupied&^(1<<from|1<<(int(enPassantSquare)-cfg.singlePushes)) | 1<<sq
		path := genRookAttacks(kingSq, occupiedWithoutPawns)
		potentialCheckers := enemyQueenOrRooks & path
		if potentialCheckers == 0 {
			*moves = append(*moves, Move{Piece: Pawn, From: from, To: sq, Type: EnPassant})
		} else {
			for potentialCheckers != 0 {
				var checker Square
				checker, potentialCheckers = potentialCheckers.PopLSB()
				if squaresBetween[kingSq][checker]&occupiedWithoutPawns != 0 {
					*moves = append(*moves, Move{Piece: Pawn, From: from, To: sq, Type: EnPassant})
				}
			}
		}
	}

	right := (pawnsOnRank &^ File_H).RotateLeft(cfg.rightAttacks) & enPassantFile & moveMask
	if right&^pinned != 0 {
		sq, _ := right.PopLSB()
		from := sq - Square(cfg.rightAttacks)
		occupiedWithoutPawns := occupied&^(1<<from|1<<(int(enPassantSquare)-cfg.singlePushes)) | 1<<sq
		path := genRookAttacks(kingSq, occupiedWithoutPawns)
		potentialCheckers := enemyQueenOrRooks & path
		if potentialCheckers == 0 {
			*moves = append(*moves, Move{Piece: Pawn, From: from, To: sq, Type: EnPassant})
		} else {
			for potentialCheckers != 0 {
				var checker Square
				checker, potentialCheckers = potentialCheckers.PopLSB()
				if squaresBetween[kingSq][checker]&occupiedWithoutPawns != 0 {
					*moves = append(*moves, Move{Piece: Pawn, From: from, To: sq, Type: EnPassant})
				}
			}
		}
	}
}

func genKnightMoves(moves *[]Move, p *Position, us Color, moveMask, pinned BitBoard) {
	knights := p.Pieces[us][Knight] &^ pinned
	enemiesOrEmpty := ^p.AllPieces[us]

	for knights != 0 {
		var from Square
		from, knights = knights.PopLSB()
		targets := knightMoves[from] & enemiesOrEmpty & moveMask

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			*moves = append(*moves, Move{Piece: Knight, From: from, To: to})
		}

	}
}

func genBishopMoves(moves *[]Move, p *Position, us Color, moveMask, pinnedStraight, pinnedDiagonal BitBoard) {
	bishops := p.Pieces[us][Bishop] &^ pinnedStraight
	enemiesOrEmpty := ^p.AllPieces[us] & moveMask
	occupied := p.Occupied

	for b := bishops & pinnedDiagonal; b != 0; {
		var from Square
		from, b = b.PopLSB()
		targets := genBishopAttacks(from, occupied) & enemiesOrEmpty & pinnedDiagonal

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			*moves = append(*moves, Move{Piece: Bishop, From: from, To: to})
		}
	}

	for b := bishops &^ pinnedDiagonal; b != 0; {
		var from Square
		from, b = b.PopLSB()
		targets := genBishopAttacks(from, occupied) & enemiesOrEmpty

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			*moves = append(*moves, Move{Piece: Bishop, From: from, To: to})
		}
	}
}

func genRookMoves(moves *[]Move, p *Position, us Color, moveMask, pinnedStraight, pinnedDiagonal BitBoard) {
	rooks := p.Pieces[us][Rook] &^ pinnedDiagonal
	enemiesOrEmpty := ^p.AllPieces[us] & moveMask
	occupied := p.Occupied

	for r := rooks & pinnedStraight; r != 0; {
		var from Square
		from, r = r.PopLSB()
		targets := genRookAttacks(from, occupied) & enemiesOrEmpty & pinnedStraight

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			*moves = append(*moves, Move{Piece: Rook, From: from, To: to})
		}
	}

	for r := rooks &^ pinnedStraight; r != 0; {
		var from Square
		from, r = r.PopLSB()
		targets := genRookAttacks(from, occupied) & enemiesOrEmpty

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			*moves = append(*moves, Move{Piece: Rook, From: from, To: to})
		}
	}
}

func genQueenMoves(moves *[]Move, p *Position, us Color, moveMask, pinnedStraight, pinnedDiagonal BitBoard) {
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
			*moves = append(*moves, Move{Piece: Queen, From: from, To: to})
		}
	}

	for q := queens & pinnedStraight; q != 0; {
		var from Square
		from, q = q.PopLSB()
		targets := genRookAttacks(from, occupied) & enemiesOrEmpty & pinnedStraight

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			*moves = append(*moves, Move{Piece: Queen, From: from, To: to})
		}
	}

	for q := queens &^ pinned; q != 0; {
		var from Square
		from, q = q.PopLSB()
		targets := (genRookAttacks(from, occupied) | genBishopAttacks(from, occupied)) & enemiesOrEmpty

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			*moves = append(*moves, Move{Piece: Queen, From: from, To: to})
		}
	}
}

func genKingMoves(moves *[]Move, p *Position, us, them Color, castling bool) {
	king := p.Pieces[us][King]
	enemiesOrEmpty := ^p.AllPieces[us]
	occupied := p.Occupied

	// normal moves
	from, _ := king.PopLSB()
	potentialTargets := kingMoves[from] & enemiesOrEmpty
	if potentialTargets == 0 {
		return
	}

	enemyKing := p.Pieces[them][King]
	attacked := attacks(p, them, us)

	for targets := potentialTargets &^ (enemyKing | attacked); targets != 0; {
		var to Square
		to, targets = targets.PopLSB()
		*moves = append(*moves, Move{Piece: King, From: from, To: to})
	}

	if !castling {
		return
	}

	// castling
	if us == White {
		if p.CanWhiteCastleKingSide() {
			emptySquares := squaresBetween[SQ_E1][SQ_H1]
			mustNotBeAttacked := squaresBetween[SQ_D1][SQ_H1]

			if emptySquares&occupied == 0 && mustNotBeAttacked&attacked == 0 {
				*moves = append(*moves, Move{Piece: King, Type: Castle, From: SQ_E1, To: SQ_G1})
			}
		}

		if p.CanWhiteCastleQueenSide() {
			emptySquares := squaresBetween[SQ_A1][SQ_E1]
			mustNotBeAttacked := squaresBetween[SQ_B1][SQ_E1]

			if emptySquares&occupied == 0 && mustNotBeAttacked&attacked == 0 {
				*moves = append(*moves, Move{Piece: King, Type: Castle, From: SQ_E1, To: SQ_C1})
			}
		}
	} else {
		if p.CanBlackCastleKingSide() {
			emptySquares := squaresBetween[SQ_E8][SQ_H8]
			mustNotBeAttacked := squaresBetween[SQ_D8][SQ_H8]

			if emptySquares&occupied == 0 && mustNotBeAttacked&attacked == 0 {
				*moves = append(*moves, Move{Piece: King, Type: Castle, From: SQ_E8, To: SQ_G8})
			}
		}

		if p.CanBlackCastleQueenSide() {
			emptySquares := squaresBetween[SQ_A8][SQ_E8]
			mustNotBeAttacked := squaresBetween[SQ_B8][SQ_E8]

			if emptySquares&occupied == 0 && mustNotBeAttacked&attacked == 0 {
				*moves = append(*moves, Move{Piece: King, Type: Castle, From: SQ_E8, To: SQ_C8})
			}
		}
	}
}

func attacks(p *Position, us, them Color) BitBoard {
	return genPawnsAttacks(p, us) |
		genKnightsAttacks(p, us) |
		genBishopsAttacks(p, us, them) |
		genRooksAttacks(p, us, them) |
		genQueensAttacks(p, us, them) |
		genKingAttacks(p, us)
}
