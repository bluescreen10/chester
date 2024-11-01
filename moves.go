package main

import (
	"fmt"
	"io"
	"math/bits"
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
	PromotionToQueen
	PromotionToRook
	PromotionToBishop
	PromotionToKnight
)

// Eventualy we want to use uint16
// First 3 bytes move type
// Next 6 bytes from square
// Last 6 bytes to square
type Move struct {
	Piece    Piece
	Type     MoveType
	From, To Square
}

type config struct {
	singlePushes     int
	leftAttacks      int
	rightAttacks     int
	startPlusOneRank BitBoard
	enPassantRank    BitBoard
	promotionRank    BitBoard
}

var pawnConfig [2]config = [2]config{
	{-8, -9, -7, Rank_3, Rank_5, Rank_8},
	{8, 7, 9, Rank_6, Rank_4, Rank_1},
}

func (m Move) String() string {
	fromRank, fromFile := m.From.RankAndFile()
	toRank, toFile := m.To.RankAndFile()

	suffix := ""

	switch m.Type {
	case PromotionToBishop:
		suffix = "b"
	case PromotionToKnight:
		suffix = "n"
	case PromotionToRook:
		suffix = "r"
	case PromotionToQueen:
		suffix = "q"
	default:
	}

	return fmt.Sprintf("%c%d%c%d%s", 'a'+fromFile, fromRank+1, 'a'+toFile, toRank+1, suffix)
}

var movesPool = sync.Pool{
	New: func() any {
		s := make([]Move, 0, maxMoves)
		return &s
	},
}

func Perft(p Position, depth int, print bool, output io.Writer) int {
	var nodes int
	moves := movesPool.Get().(*[]Move)
	defer movesPool.Put(moves)
	*moves = (*moves)[:0]
	LegalMoves(moves, &p)

	if depth == 1 {
		return len(*moves)
	}

	for _, m := range *moves {
		newPos := p.Do(m)
		newNodes := Perft(newPos, depth-1, false, nil)
		if print {
			fmt.Fprintf(output, "%s: %d\n", m, newNodes)
		}
		nodes += newNodes
		//p.Undo()
	}
	//movesPool.Put(moves)
	return nodes
}

func LegalMoves(moves *[]Move, pos *Position) {
	us, them := pos.SideToMove()

	checkers, pinned := checkersAndPinned(pos, us, them)

	switch c := checkers.OnesCount(); {
	case c >= 2:
		genKingMoves(moves, pos, us, them, false)
	case c == 1:
		captureMask := checkers
		pushMask := EmptyBoard
		kingSq, _ := pos.Pieces[us][King].PopLSB()

		if (pos.Pieces[them][Bishop]|pos.Pieces[them][Rook]|pos.Pieces[them][Queen])&checkers != 0 {
			checkerSq, _ := checkers.PopLSB()
			pushMask = squaresBetween[checkerSq][kingSq]
		}

		genPawnMoves(moves, pos, us, them, captureMask|pushMask, pinned)
		genKnightMoves(moves, pos, us, captureMask|pushMask, pinned)
		genBishopMoves(moves, pos, us, captureMask|pushMask, pinned)
		genRookMoves(moves, pos, us, captureMask|pushMask, pinned)
		genQueenMoves(moves, pos, us, captureMask|pushMask, pinned)
		genKingMoves(moves, pos, us, them, false)
	default:
		genPawnMoves(moves, pos, us, them, FullBoard, pinned)
		genKnightMoves(moves, pos, us, FullBoard, pinned)
		genBishopMoves(moves, pos, us, FullBoard, pinned)
		genRookMoves(moves, pos, us, FullBoard, pinned)
		genQueenMoves(moves, pos, us, FullBoard, pinned)
		genKingMoves(moves, pos, us, them, true)
	}
}

func checkersAndPinned(p *Position, us, them Color) (BitBoard, BitBoard) {
	king := p.Pieces[us][King]
	kingSq, _ := king.PopLSB()

	checkers := knightMoves[kingSq] & p.Pieces[them][Knight]
	config := pawnConfig[us]
	checkers |= (king & File_Not_A).RotateLeft(config.leftAttacks) & p.Pieces[them][Pawn]
	checkers |= (king & File_Not_H).RotateLeft(config.rightAttacks) & p.Pieces[them][Pawn]

	kingDiagonalRays := rays[NorthWest][kingSq] | rays[NorthEast][kingSq] | rays[SouthWest][kingSq] | rays[SouthEast][kingSq]
	diagonalAttackers := (p.Pieces[them][Queen] | p.Pieces[them][Bishop])

	kingStraightRays := rays[North][kingSq] | rays[South][kingSq] | rays[East][kingSq] | rays[West][kingSq]
	straightAttackers := (p.Pieces[them][Queen] | p.Pieces[them][Rook])

	potentialCheckers := (diagonalAttackers & kingDiagonalRays) | (straightAttackers & kingStraightRays)
	occupied := p.Occupied

	var pinned BitBoard

	for potentialCheckers != 0 {
		var sq Square
		sq, potentialCheckers = potentialCheckers.PopLSB()

		path := squaresBetween[sq][kingSq]
		potentialyPinned := path & occupied
		switch potentialyPinned.OnesCount() {
		case 0:
			checkers |= 1 << sq
		case 1:
			pinned |= potentialyPinned
		}
	}

	return checkers, pinned
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
	attacks := rays[NorthWest][sq]
	if intersect := rays[NorthWest][sq] & occupied; intersect != 0 {
		//fmt.Println(intersect)
		index := bits.LeadingZeros64(uint64(intersect))
		attacks &^= rays[NorthWest][63-index]
	}
	// fmt.Println(attacks)
	// fmt.Println(rays[NorthEast][sq])
	attacks |= rays[NorthEast][sq]
	if intersect := rays[NorthEast][sq] & occupied; intersect != 0 {
		// fmt.Println(intersect)
		index := bits.LeadingZeros64(uint64(intersect))
		attacks &^= rays[NorthEast][63-index]
	}
	// fmt.Println(attacks)
	// fmt.Println(rays[SouthWest][sq])
	attacks |= rays[SouthWest][sq]
	if intersect := rays[SouthWest][sq] & occupied; intersect != 0 {
		// fmt.Println(intersect)
		index := bits.TrailingZeros64(uint64(intersect))
		attacks &^= rays[SouthWest][index]
	}
	// fmt.Println(attacks)
	// fmt.Println(rays[SouthEast][sq])
	attacks |= rays[SouthEast][sq]
	if intersect := rays[SouthEast][sq] & occupied; intersect != 0 {
		// fmt.Println(intersect)
		index := bits.TrailingZeros64(uint64(intersect))
		attacks &^= rays[SouthEast][index]
	}
	// fmt.Println(attacks)
	return attacks
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
	// fmt.Println(rays[North][sq])
	attacks := rays[North][sq]
	if intersect := rays[North][sq] & occupied; intersect != 0 {
		// fmt.Println(intersect)
		index := bits.LeadingZeros64(uint64(intersect))
		attacks &^= rays[North][63-index]
	}
	// fmt.Println(attacks)
	// fmt.Println(rays[South][sq])

	attacks |= rays[South][sq]
	if intersect := rays[South][sq] & occupied; intersect != 0 {
		// fmt.Println(intersect)
		index := bits.TrailingZeros64(uint64(intersect))
		attacks &^= rays[South][index]
	}
	// fmt.Println(attacks)
	// fmt.Println(rays[East][sq])
	attacks |= rays[East][sq]
	if intersect := rays[East][sq] & occupied; intersect != 0 {
		// fmt.Println(intersect)
		index := bits.TrailingZeros64(uint64(intersect))
		attacks &^= rays[East][index]
	}
	// fmt.Println(attacks)
	// fmt.Println(rays[West][sq])
	attacks |= rays[West][sq]
	if intersect := rays[West][sq] & occupied; intersect != 0 {
		// fmt.Println(intersect)
		index := bits.LeadingZeros64(uint64(intersect))
		attacks &^= rays[West][63-index]
	}
	// fmt.Println(attacks)
	return attacks
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

func genPawnMoves(moves *[]Move, p *Position, us, them Color, moveMask, pinned BitBoard) {
	pawns := p.Pieces[us][Pawn]
	occupied := p.Occupied
	enemies := p.AllPieces[them]
	kingSq, _ := p.Pieces[us][King].PopLSB()
	enemyQueenOrRooks := p.Pieces[them][Queen] | p.Pieces[them][Rook]

	genForwardMoves(moves, pawns, occupied, us, moveMask, pinned, kingSq)
	genAttackMoves(moves, pawns, enemies, us, moveMask, pinned, kingSq)
	if p.IsEnPassant() {
		genEnPassantMoves(moves, pawns, occupied, p.EnPassantFile(), us, enemyQueenOrRooks, moveMask, pinned, kingSq)
	}
}

func genForwardMoves(moves *[]Move, pawns, occupied BitBoard, us Color, moveMask, pinned BitBoard, kingSq Square) {
	config := pawnConfig[us]
	singlePushes := (pawns & pinned).RotateLeft(config.singlePushes) &^ occupied

	for pushes := singlePushes & moveMask; pushes != 0; {
		var to Square
		to, pushes = pushes.PopLSB()
		from := to - Square(config.singlePushes)
		if lineFromTo[kingSq][from]&lineFromTo[kingSq][to] != 0 {
			*moves = append(*moves, Move{Piece: Pawn, From: from, To: to})
		}
	}

	doublePushes := (singlePushes & config.startPlusOneRank).RotateLeft(config.singlePushes) &^ occupied & moveMask
	for doublePushes != 0 {
		var to Square
		to, doublePushes = doublePushes.PopLSB()
		from := to - Square(2*config.singlePushes)
		if lineFromTo[kingSq][from]&lineFromTo[kingSq][to] != 0 {
			*moves = append(*moves, Move{Piece: Pawn, From: from, To: to})
		}
	}

	singlePushes = (pawns &^ pinned).RotateLeft(config.singlePushes) &^ occupied
	for pushes := singlePushes &^ pinned & moveMask; pushes != 0; {
		var to Square
		to, pushes = pushes.PopLSB()
		from := to - Square(config.singlePushes)

		if config.promotionRank&(1<<to) == 0 {
			*moves = append(*moves, Move{Piece: Pawn, From: from, To: to})
		} else {
			*moves = append(*moves,
				Move{Piece: Pawn, From: from, To: to, Type: PromotionToQueen},
				Move{Piece: Pawn, From: from, To: to, Type: PromotionToRook},
				Move{Piece: Pawn, From: from, To: to, Type: PromotionToBishop},
				Move{Piece: Pawn, From: from, To: to, Type: PromotionToKnight},
			)
		}
	}

	doublePushes = (singlePushes & config.startPlusOneRank).RotateLeft(config.singlePushes) &^ occupied & moveMask
	for doublePushes != 0 {
		var to Square
		to, doublePushes = doublePushes.PopLSB()
		*moves = append(*moves, Move{Piece: Pawn, From: to - Square(2*config.singlePushes), To: to})
	}
}

func genAttackMoves(moves *[]Move, pawns, enemies BitBoard, us Color, moveMask, pinned BitBoard, kingSq Square) {
	config := pawnConfig[us]

	// left attacks
	leftAttacks := (pawns & pinned &^ File_A).RotateLeft(config.leftAttacks) & enemies
	for attacks := leftAttacks & moveMask; attacks != 0; {
		var to Square
		to, attacks = attacks.PopLSB()
		from := to - Square(config.leftAttacks)
		if lineFromTo[kingSq][from]&lineFromTo[kingSq][to] == 0 {
			continue
		}

		if config.promotionRank&(1<<to) == 0 {
			*moves = append(*moves, Move{Piece: Pawn, From: from, To: to})
		} else {
			*moves = append(*moves,
				Move{Piece: Pawn, From: from, To: to, Type: PromotionToQueen},
				Move{Piece: Pawn, From: from, To: to, Type: PromotionToRook},
				Move{Piece: Pawn, From: from, To: to, Type: PromotionToBishop},
				Move{Piece: Pawn, From: from, To: to, Type: PromotionToKnight},
			)
		}
	}

	leftAttacks = (pawns &^ pinned &^ File_A).RotateLeft(config.leftAttacks) & enemies
	for attacks := leftAttacks & moveMask; attacks != 0; {
		var to Square
		to, attacks = attacks.PopLSB()
		from := to - Square(config.leftAttacks)

		if config.promotionRank&(1<<to) == 0 {
			*moves = append(*moves, Move{Piece: Pawn, From: from, To: to})
		} else {
			*moves = append(*moves,
				Move{Piece: Pawn, From: from, To: to, Type: PromotionToQueen},
				Move{Piece: Pawn, From: from, To: to, Type: PromotionToRook},
				Move{Piece: Pawn, From: from, To: to, Type: PromotionToBishop},
				Move{Piece: Pawn, From: from, To: to, Type: PromotionToKnight},
			)
		}
	}

	// right attacks
	rightAttacks := (pawns & pinned &^ File_H).RotateLeft(config.rightAttacks) & enemies
	for attacks := rightAttacks & moveMask; attacks != 0; {
		var to Square
		to, attacks = attacks.PopLSB()
		from := to - Square(config.rightAttacks)
		if lineFromTo[kingSq][from]&lineFromTo[kingSq][to] == 0 {
			continue
		}

		if config.promotionRank&(1<<to) == 0 {
			*moves = append(*moves, Move{Piece: Pawn, From: from, To: to})
		} else {
			*moves = append(*moves,
				Move{Piece: Pawn, From: from, To: to, Type: PromotionToQueen},
				Move{Piece: Pawn, From: from, To: to, Type: PromotionToRook},
				Move{Piece: Pawn, From: from, To: to, Type: PromotionToBishop},
				Move{Piece: Pawn, From: from, To: to, Type: PromotionToKnight},
			)
		}
	}

	rightAttacks = (pawns &^ pinned &^ File_H).RotateLeft(config.rightAttacks) & enemies
	for attacks := rightAttacks & moveMask; attacks != 0; {
		var to Square
		to, attacks = attacks.PopLSB()
		from := to - Square(config.rightAttacks)
		if config.promotionRank&(1<<to) == 0 {
			*moves = append(*moves, Move{Piece: Pawn, From: from, To: to})
		} else {
			*moves = append(*moves,
				Move{Piece: Pawn, From: from, To: to, Type: PromotionToQueen},
				Move{Piece: Pawn, From: from, To: to, Type: PromotionToRook},
				Move{Piece: Pawn, From: from, To: to, Type: PromotionToBishop},
				Move{Piece: Pawn, From: from, To: to, Type: PromotionToKnight},
			)
		}
	}
}

func genEnPassantMoves(moves *[]Move, pawns, occupied, enPassantFile BitBoard, us Color, enemyQueenOrRooks, moveMask, pinned BitBoard, kingSq Square) {
	config := pawnConfig[us]
	pawnsOnRank := pawns & config.enPassantRank
	left := (pawnsOnRank &^ File_A).RotateLeft(config.leftAttacks) & enPassantFile & moveMask
	if left&^pinned != 0 {
		sq, _ := left.PopLSB()
		path := lineFromTo[kingSq][sq-Square(config.leftAttacks)]

		occupiedWithoutPawns := occupied &^ (1 << (sq - Square(config.leftAttacks))) &^ (1 << sq)
		inCheck := false
		for potentialCheckers := enemyQueenOrRooks & path; potentialCheckers != 0; {
			var checkerSq Square
			checkerSq, potentialCheckers = potentialCheckers.PopLSB()
			if (squaresBetween[checkerSq][kingSq] & occupiedWithoutPawns).OnesCount() == 0 {
				inCheck = true
				break
			}
		}
		if !inCheck {
			*moves = append(*moves, Move{Piece: Pawn, From: sq - Square(config.leftAttacks), To: sq, Type: EnPassant})
		}
	}

	right := (pawnsOnRank &^ File_H).RotateLeft(config.rightAttacks) & enPassantFile & moveMask
	if right&^pinned != 0 {
		sq, _ := right.PopLSB()
		path := lineFromTo[kingSq][sq-Square(config.rightAttacks)]
		occupiedWithoutPawns := occupied &^ (1 << (sq - Square(config.rightAttacks))) &^ (1 << sq)
		inCheck := false
		for potentialCheckers := enemyQueenOrRooks & path; potentialCheckers != 0; {
			var checkerSq Square
			checkerSq, potentialCheckers = potentialCheckers.PopLSB()
			if (squaresBetween[checkerSq][kingSq] & occupiedWithoutPawns).OnesCount() == 0 {
				inCheck = true
				break
			}
		}
		if !inCheck {
			*moves = append(*moves, Move{Piece: Pawn, From: sq - Square(config.rightAttacks), To: sq, Type: EnPassant})
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

func genBishopMoves(moves *[]Move, p *Position, us Color, moveMask, pinned BitBoard) {
	bishops := p.Pieces[us][Bishop]
	enemiesOrEmpty := ^p.AllPieces[us]
	occupied := p.Occupied

	for b := bishops & pinned; b != 0; {
		var from Square
		from, b = b.PopLSB()
		kingSq, _ := p.Pieces[us][King].PopLSB()
		rayMask := lineFromTo[kingSq][from]
		targets := genBishopAttacks(from, occupied) & enemiesOrEmpty & moveMask & rayMask

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			*moves = append(*moves, Move{Piece: Bishop, From: from, To: to})
		}
	}

	for b := bishops &^ pinned; b != 0; {
		var from Square
		from, b = b.PopLSB()
		targets := genBishopAttacks(from, occupied) & enemiesOrEmpty & moveMask

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			*moves = append(*moves, Move{Piece: Bishop, From: from, To: to})
		}
	}
}

func genRookMoves(moves *[]Move, p *Position, us Color, moveMask, pinned BitBoard) {
	rooks := p.Pieces[us][Rook]
	enemiesOrEmpty := ^p.AllPieces[us]
	occupied := p.Occupied

	for r := rooks & pinned; r != 0; {
		var from Square
		from, r = r.PopLSB()
		kingSq, _ := p.Pieces[us][King].PopLSB()
		rayMask := lineFromTo[kingSq][from]
		targets := genRookAttacks(from, occupied) & enemiesOrEmpty & moveMask & rayMask

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			*moves = append(*moves, Move{Piece: Rook, From: from, To: to})
		}
	}

	for r := rooks &^ pinned; r != 0; {
		var from Square
		from, r = r.PopLSB()
		targets := genRookAttacks(from, occupied) & enemiesOrEmpty & moveMask

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			*moves = append(*moves, Move{Piece: Rook, From: from, To: to})
		}
	}
}

func genQueenMoves(moves *[]Move, p *Position, us Color, moveMask, pinned BitBoard) {
	queens := p.Pieces[us][Queen]
	enemiesOrEmpty := ^p.AllPieces[us]
	occupied := p.Occupied

	for q := queens & pinned; q != 0; {
		var from Square
		from, q = q.PopLSB()
		kingSq, _ := p.Pieces[us][King].PopLSB()
		raysMask := lineFromTo[kingSq][from]
		targets := (genRookAttacks(from, occupied) | genBishopAttacks(from, occupied)) & enemiesOrEmpty & moveMask & raysMask

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			*moves = append(*moves, Move{Piece: Queen, From: from, To: to})
		}
	}

	for q := queens &^ pinned; q != 0; {
		var from Square
		from, q = q.PopLSB()
		targets := (genRookAttacks(from, occupied) | genBishopAttacks(from, occupied)) & enemiesOrEmpty & moveMask

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			*moves = append(*moves, Move{Piece: Queen, From: from, To: to})
		}
	}
}

func genKingMoves(moves *[]Move, p *Position, us, them Color, castling bool) {
	king := p.Pieces[us][King]
	enemyKing := p.Pieces[them][King]
	enemiesOrEmpty := ^p.AllPieces[us]
	occupied := p.Occupied
	attacked := attacks(p, them, us)

	// normal moves
	from, _ := king.PopLSB()
	targets := kingMoves[from] & enemiesOrEmpty &^ (enemyKing | attacked)

	for targets != 0 {
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
