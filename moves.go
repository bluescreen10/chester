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
	Type     MoveType
	From, To Square
}

type direction uint8

const (
	North direction = iota
	East
	South
	West

	NorthEast
	SouthEast
	SouthWest
	NorthWest
)

var rays [direction(8)][Square(64)]BitBoard

var knightMoves [64]BitBoard
var kingMoves [64]BitBoard

func init() {
	// pre-compute rays
	for i := 0; i < 64; i++ {
		r, f := 7-i/8, i%8

		for j := 1; j < 8; j++ {
			if r-j >= 0 && f-j >= 0 {
				rays[SouthWest][i] |= 1 << uint8((7-r+j)*8+f-j)
			}
			if r-j >= 0 && f+j < 8 {
				rays[SouthEast][i] |= 1 << uint8((7-r+j)*8+f+j)
			}
			if r+j < 8 && f-j >= 0 {
				rays[NorthWest][i] |= 1 << uint8((7-r-j)*8+f-j)
			}
			if r+j < 8 && f+j < 8 {
				rays[NorthEast][i] |= 1 << uint8((7-r-j)*8+f+j)
			}
		}
	}

	for i := 0; i < 64; i++ {
		r, f := 7-i/8, i%8

		for j := 1; j < 8; j++ {
			if r-j >= 0 {
				rays[South][i] |= 1 << uint8((7-r+j)*8+f)
			}
			if r+j < 8 {
				rays[North][i] |= 1 << uint8((7-r-j)*8+f)
			}
			if f-j >= 0 {
				rays[West][i] |= 1 << uint8((7-r)*8+f-j)
			}
			if f+j < 8 {
				rays[East][i] |= 1 << uint8((7-r)*8+f+j)
			}
		}
	}

	// pre-compute knight moves
	for i := int8(0); i < 64; i++ {
		rank, file := 7-i/8, i%8
		moves := [][]int8{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}}

		for _, offsets := range moves {
			targetRank, targetFile := rank+offsets[0], file+offsets[1]
			if targetRank < 0 || targetRank > 7 || targetFile < 0 || targetFile > 7 {
				continue
			}

			knightMoves[i] |= 1 << uint8((7-targetRank)*8+targetFile)
		}
	}

	// pre-compute king moves
	for i := int8(0); i < 64; i++ {
		rank, file := 7-i/8, i%8
		moves := [][]int8{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}

		for _, offsets := range moves {
			targetRank, targetFile := rank+offsets[0], file+offsets[1]
			if targetRank < 0 || targetRank > 7 || targetFile < 0 || targetFile > 7 {
				continue
			}

			kingMoves[i] |= 1 << uint8((7-targetRank)*8+targetFile)
		}
	}
}

// func (m Move) String() string {
// 	target := fmt.Sprintf("%c%d", 'a'+int(m.To%8), int(m.To/8)+1)

//		switch m.Piece {
//		case WhitePawn, BlackPawn:
//			return target
//		case WhiteKnight, BlackKnight:
//			return "N" + target
//		case WhiteBishop, BlackBishop:
//			return "B" + target
//		case WhiteRook, BlackRook:
//			return "R" + target
//		case WhiteQueen, BlackQueen:
//			return "Q" + target
//		case WhiteKing, BlackKing:
//			return "K" + target
//		default:
//			return "??"
//		}
//	}
func (m Move) String() string {
	fromRank, fromFile := m.From.RankAndFile()
	toRank, toFile := m.To.RankAndFile()
	return fmt.Sprintf("%c%d%c%d", 'a'+fromFile, fromRank+1, 'a'+toFile, toRank+1)
}

var movesPool = sync.Pool{
	New: func() any {
		s := make([]Move, 0, maxMoves)
		return &s
	},
}

func Perft(p Position, depth int, print bool, output io.Writer) int {
	if depth == 0 {
		return 1
	}
	var nodes int
	moves := movesPool.Get().(*[]Move)
	*moves = (*moves)[:0]
	LegalMoves(moves, p)
	for _, m := range *moves {
		p.Do(m)
		newNodes := Perft(p, depth-1, false, output)
		if print {
			fmt.Fprintf(output, "%s: %d\n", m, newNodes)
		}
		nodes += newNodes
		p.Undo()
	}
	movesPool.Put(moves)
	return nodes
}

func LegalMoves(moves *[]Move, position Position) {
	*moves = (*moves)[:0]

	us := position.SideToMove()

	var attacks BitBoard
	occupied := position.Occupied

	if us == White {
		for sq := Square(0); sq < 64; sq++ {

			switch position.Board[sq] {
			case BlackPawn:
				attacks |= genBlackPawnAttacks(sq)
			case BlackKnight:
				attacks |= genKnightAttacks(sq)
			case BlackBishop:
				attacks |= genBishopAttacks(sq, occupied)
			case BlackRook:
				attacks |= genRookAttacks(sq, occupied)
			case BlackQueen:
				attacks |= genBishopAttacks(sq, occupied)
				attacks |= genRookAttacks(sq, occupied)
			case BlackKing:
				attacks |= genKingAttacks(sq)
				//case WhiteKing:
				//	king = BitBoard(1 << uint(sq))
			}
		}
	} else {
		for sq := Square(0); sq < 64; sq++ {
			switch position.Board[sq] {
			case WhitePawn:
				attacks |= genWhitePawnAttacks(sq)
			case WhiteKnight:
				attacks |= genKnightAttacks(sq)
			case WhiteBishop:
				attacks |= genBishopAttacks(sq, occupied)
			case WhiteRook:
				attacks |= genRookAttacks(sq, occupied)
			case WhiteQueen:
				attacks |= genBishopAttacks(sq, occupied)
				attacks |= genRookAttacks(sq, occupied)
			case WhiteKing:
				attacks |= genKingAttacks(sq)
				//case BlackKing:
				//	king = BitBoard(1 << uint(sq))
			}
		}
	}

	// fmt.Println(king)
	// fmt.Println(attacks)
	//inCheck := king&attacks != 0

	for sq := Square(0); sq < 64; sq++ {
		piece := position.Board[sq]

		if piece == Empty || piece.Color() != us {
			continue
		}

		switch piece {
		case WhitePawn:
			genWhitePawnMoves(moves, position, sq)
		case BlackPawn:
			genBlackPawnMoves(moves, position, sq)
		case WhiteKnight, BlackKnight:
			genKnightMoves(moves, position, sq, us)
		case WhiteBishop, BlackBishop:
			genBishopMoves(moves, position, sq, us)
		case WhiteRook, BlackRook:
			genRookMoves(moves, position, sq, us)
		case WhiteQueen, BlackQueen:
			genQueenMoves(moves, position, sq, us)
		case WhiteKing:
			genWhiteKingMoves(moves, position, sq, attacks)
		case BlackKing:
			genBlackKingMoves(moves, position, sq, attacks)
		}
	}

	for i := 0; i < len(*moves); i++ {
		move := (*moves)[i]
		//fmt.Println(move)
		position.Do(move)
		var king, attacks BitBoard

		occupied = position.Occupied
		if us == White {
			for sq, bit := Square(0), BitBoard(1); sq < 64; sq, bit = sq+1, bit<<1 {
				switch position.Board[sq] {
				case BlackPawn:
					attacks |= genBlackPawnAttacks(sq)
				case BlackKnight:
					attacks |= genKnightAttacks(sq)
				case BlackBishop:
					attacks |= genBishopAttacks(sq, occupied)
				case BlackRook:
					attacks |= genRookAttacks(sq, occupied)
				case BlackQueen:
					attacks |= genBishopAttacks(sq, occupied) | genRookAttacks(sq, occupied)
				case BlackKing:
					attacks |= genKingAttacks(sq)
				case WhiteKing:
					king = bit
				}
			}
		} else {
			for sq, bit := Square(0), BitBoard(1); sq < 64; sq, bit = sq+1, bit<<1 {
				switch position.Board[sq] {
				case WhitePawn:
					attacks |= genWhitePawnAttacks(sq)
				case WhiteKnight:
					attacks |= genKnightAttacks(sq)
				case WhiteBishop:
					attacks |= genBishopAttacks(sq, occupied)
				case WhiteRook:
					attacks |= genRookAttacks(sq, occupied)
				case WhiteQueen:
					attacks |= genBishopAttacks(sq, occupied) | genRookAttacks(sq, occupied)
				case BlackKing:
					king = bit
				}
			}
		}

		inCheck := king&attacks != 0

		if inCheck {
			*moves = append((*moves)[:i], (*moves)[i+1:]...)
			i--
		}
		position.Undo()
	}
}

func genWhitePawnAttacks(sq Square) BitBoard {
	attacks := BitBoard(0)

	file := sq % 8
	if file > 0 {
		attacks |= 1 << uint(sq-9)
	}

	if file < 7 {
		attacks |= 1 << uint(sq-7)
	}

	return attacks
}

func genBlackPawnAttacks(sq Square) BitBoard {
	var attacks BitBoard

	file := sq % 8
	if file > 0 {
		attacks |= 1 << uint(sq+7)
	}

	if file < 7 {
		attacks |= 1 << uint(sq+9)
	}
	return attacks
}

func genKnightAttacks(sq Square) BitBoard {
	return knightMoves[sq]
}

func genBishopAttacks(sq Square, occupied BitBoard) BitBoard {
	attacks := rays[NorthWest][sq]
	//fmt.Println(rays[NorthWest][sq])
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

func genKingAttacks(sq Square) BitBoard {
	return kingMoves[sq]
}

func genWhitePawnMoves(moves *[]Move, position Position, sq Square) {
	rank, file := sq.RankAndFile()

	// forward moves
	if to := sq - 8; position.Board[to] == Empty {

		if rank < 6 {
			*moves = append(*moves, Move{From: sq, To: to})
		} else {
			*moves = append(*moves,
				Move{From: sq, To: to, Type: PromotionToKnight},
				Move{From: sq, To: to, Type: PromotionToBishop},
				Move{From: sq, To: to, Type: PromotionToRook},
				Move{From: sq, To: to, Type: PromotionToQueen},
			)
		}

		if to := sq - 16; rank == 1 && position.Board[to] == Empty {
			*moves = append(*moves, Move{From: sq, To: to})
		}
	}

	// captures
	if file > 0 {
		if to := sq - 9; position.Board[to] != Empty && position.Board[to].Color() == Black {
			if rank < 6 {
				*moves = append(*moves, Move{From: sq, To: to})
			} else {
				*moves = append(*moves,
					Move{From: sq, To: to, Type: PromotionToKnight},
					Move{From: sq, To: to, Type: PromotionToBishop},
					Move{From: sq, To: to, Type: PromotionToRook},
					Move{From: sq, To: to, Type: PromotionToQueen},
				)
			}
		}
	}

	if file < 7 {
		if to := sq - 7; position.Board[to] != Empty && position.Board[to].Color() == Black {
			if rank < 6 {
				*moves = append(*moves, Move{From: sq, To: to})
			} else {
				*moves = append(*moves,
					Move{From: sq, To: to, Type: PromotionToKnight},
					Move{From: sq, To: to, Type: PromotionToBishop},
					Move{From: sq, To: to, Type: PromotionToRook},
					Move{From: sq, To: to, Type: PromotionToQueen},
				)
			}
		}
	}

	// en passant
	if rank == 4 && position.IsEnPassant() {
		enPassantFile := position.EnPassantFile()

		if file-enPassantFile == 1 {
			to := sq - 9
			*moves = append(*moves, Move{From: sq, To: to, Type: EnPassant})
		}

		if file-enPassantFile == -1 {
			to := sq - 7
			*moves = append(*moves, Move{From: sq, To: to, Type: EnPassant})
		}

	}
}

func genBlackPawnMoves(moves *[]Move, position Position, sq Square) {
	rank, file := sq.RankAndFile()

	// forward moves
	if to := sq + 8; position.Board[to] == Empty {

		if rank > 1 {
			*moves = append(*moves, Move{From: sq, To: to})
		} else {
			*moves = append(*moves,
				Move{From: sq, To: to, Type: PromotionToKnight},
				Move{From: sq, To: to, Type: PromotionToBishop},
				Move{From: sq, To: to, Type: PromotionToRook},
				Move{From: sq, To: to, Type: PromotionToQueen},
			)
		}

		if to := sq + 16; rank == 6 && position.Board[to] == Empty {
			*moves = append(*moves, Move{From: sq, To: to})
		}
	}

	// captures
	if file > 0 {
		if to := sq + 7; position.Board[to] != Empty && position.Board[to].Color() == White {
			if rank > 1 {
				*moves = append(*moves, Move{From: sq, To: to})
			} else {
				*moves = append(*moves,
					Move{From: sq, To: to, Type: PromotionToKnight},
					Move{From: sq, To: to, Type: PromotionToBishop},
					Move{From: sq, To: to, Type: PromotionToRook},
					Move{From: sq, To: to, Type: PromotionToQueen},
				)
			}
		}
	}

	if file < 7 {
		if to := sq + 9; position.Board[to] != Empty && position.Board[to].Color() == White {
			if rank > 1 {
				*moves = append(*moves, Move{From: sq, To: to})
			} else {
				*moves = append(*moves,
					Move{From: sq, To: to, Type: PromotionToKnight},
					Move{From: sq, To: to, Type: PromotionToBishop},
					Move{From: sq, To: to, Type: PromotionToRook},
					Move{From: sq, To: to, Type: PromotionToQueen},
				)
			}
		}
	}

	// en passant
	if rank == 3 && position.IsEnPassant() {
		enPassantFile := position.EnPassantFile()

		if file-enPassantFile == 1 {
			to := sq + 7
			*moves = append(*moves, Move{From: sq, To: to, Type: EnPassant})
		}

		if file-enPassantFile == -1 {
			to := sq + 9
			*moves = append(*moves, Move{From: sq, To: to, Type: EnPassant})
		}

	}
}

func genKnightMoves(moves *[]Move, position Position, sq Square, us Piece) {
	rank, file := sq.RankAndFile()

	knightMoves := [][]int8{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}}

	for _, offset := range knightMoves {
		toRank, toFile := rank+offset[0], file+offset[1]
		if toRank < 0 || toRank > 7 || toFile < 0 || toFile > 7 {
			continue
		}

		to := SquareFromRankAndFile(toRank, toFile)
		enemy := position.Board[to]

		if enemy == Empty || enemy.Color() != us {
			*moves = append(*moves, Move{From: sq, To: to})
		}
	}
}

func genBishopMoves(moves *[]Move, position Position, sq Square, us Piece) {
	rank, file := sq.RankAndFile()

	bishopMoves := [][]int8{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	for _, offset := range bishopMoves {
		incX, incY := offset[0], offset[1]

		for toRank, toFile := rank+incY, file+incX; toRank >= 0 && toRank <= 7 && toFile >= 0 && toFile <= 7; toRank, toFile = toRank+incY, toFile+incX {
			to := SquareFromRankAndFile(toRank, toFile)
			enemy := position.Board[to]

			if enemy != Empty {
				if enemy.Color() != us {
					*moves = append(*moves, Move{From: sq, To: to})
				}
				break
			}

			if enemy == Empty {
				*moves = append(*moves, Move{From: sq, To: to})
			}
		}
	}
}

func genRookMoves(moves *[]Move, position Position, sq Square, us Piece) {
	rank, file := sq.RankAndFile()

	rookMoves := [][]int8{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for _, offset := range rookMoves {
		incX, incY := offset[0], offset[1]

		for toRank, toFile := rank+incY, file+incX; toRank >= 0 && toRank <= 7 && toFile >= 0 && toFile <= 7; toRank, toFile = toRank+incY, toFile+incX {
			to := SquareFromRankAndFile(toRank, toFile)
			enemy := position.Board[to]

			if enemy != Empty {
				if enemy.Color() != us {
					*moves = append(*moves, Move{From: sq, To: to})
				}
				break
			}

			if enemy == Empty {
				*moves = append(*moves, Move{From: sq, To: to})
			}
		}
	}
}

func genQueenMoves(moves *[]Move, position Position, sq Square, us Piece) {
	genBishopMoves(moves, position, sq, us)
	genRookMoves(moves, position, sq, us)
}

func genWhiteKingMoves(moves *[]Move, position Position, sq Square, attacks BitBoard) {
	rank, file := sq.RankAndFile()

	occupied := position.Occupied

	kingMoves := [][]int8{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}

	for _, offset := range kingMoves {
		toRank, toFile := rank+offset[0], file+offset[1]
		if toRank < 0 || toRank > 7 || toFile < 0 || toFile > 7 {
			continue
		}

		to := SquareFromRankAndFile(toRank, toFile)
		enemy := position.Board[to]
		toBit := BitBoard(1) << uint8(to)

		if isAttacked := attacks&toBit != 0; isAttacked {
			continue
		}

		if enemy == Empty || (enemy.Color() == Black && enemy != BlackKing) {
			*moves = append(*moves, Move{From: sq, To: to})
		}
	}

	if position.CanWhiteCastleKingSide() {
		emptySquares := BitBoard(3) << uint8(SQ_F1)
		mustNotBeAttacked := BitBoard(7) << uint8(SQ_E1)

		if emptySquares&occupied == 0 && mustNotBeAttacked&attacks == 0 {
			*moves = append(*moves, Move{Type: Castle, From: SQ_E1, To: SQ_G1})
		}
	}

	if position.CanWhiteCastleQueenSide() {
		emptySquares := BitBoard(7) << uint8(SQ_B1)
		mustNotBeAttacked := BitBoard(7) << uint8(SQ_C1)

		if emptySquares&occupied == 0 && mustNotBeAttacked&attacks == 0 {
			*moves = append(*moves, Move{Type: Castle, From: SQ_E1, To: SQ_C1})
		}
	}
}

func genBlackKingMoves(moves *[]Move, position Position, sq Square, attacks BitBoard) {
	rank, file := sq.RankAndFile()

	//inCheck := (BitBoard(1)<<uint8(sq))&attacks != 0
	occupied := position.Occupied

	kingMoves := [][]int8{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}

	for _, offset := range kingMoves {
		toRank, toFile := rank+offset[0], file+offset[1]
		if toRank < 0 || toRank > 7 || toFile < 0 || toFile > 7 {
			continue
		}

		to := SquareFromRankAndFile(toRank, toFile)
		enemy := position.Board[to]
		toBit := BitBoard(1) << uint8(to)

		if isAttacked := attacks&toBit != 0; isAttacked {
			continue
		}

		if enemy == Empty || (enemy.Color() == White && enemy != WhiteKing) {
			*moves = append(*moves, Move{From: sq, To: to})
		}
	}

	if position.CanBlackCastleKingSide() {
		emptySquares := BitBoard(3) << uint8(SQ_F8)
		mustNotBeAttacked := BitBoard(7) << uint8(SQ_E8)

		if emptySquares&occupied == 0 && mustNotBeAttacked&attacks == 0 {
			*moves = append(*moves, Move{Type: Castle, From: SQ_E8, To: SQ_G8})
		}
	}

	if position.CanBlackCastleQueenSide() {
		emptySquares := BitBoard(7) << uint8(SQ_B8)
		mustNotBeAttacked := BitBoard(7) << uint8(SQ_C8)

		if emptySquares&occupied == 0 && mustNotBeAttacked&attacks == 0 {
			*moves = append(*moves, Move{Type: Castle, From: SQ_E8, To: SQ_C8})
		}
	}
}
