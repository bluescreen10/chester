package main

import "fmt"

type MoveType uint8

const (
	Default MoveType = iota
	CastleKingSide
	CastleQueenSide
	EnPassant
	DoublePush
	Promotion
)

//  15 14 13 12 11 10  9  8  7  6  5  4  3  2  1  0
// +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
// |           |                 |                 |
// | promotion |      from       |        to       |
// |           |                 |                 |
// +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    4 bits       6 bits          6 bits
//
// promotion
// 0000 -> no promotion
// 0001 -> knight
// 0010 -> bishop
// 0011 -> rook
// 0100 -> queen

type Move uint16

func NewMove(from, to Square) Move {
	return Move(from)<<6 | Move(to)
}

func NewPromotionMove(from, to Square, promotion Piece) Move {
	return Move(promotion)<<12 | Move(from)<<6 | Move(to)
}

func (m Move) From() Square {
	return Square(m >> 6 & 0x3f)
}

func (m Move) To() Square {
	return Square(m & 0x3f)
}

func (m Move) PromoPiece() Piece {
	return Piece(m >> 12)
}

func (m Move) IsPromotion() bool {
	return m&0xf000 != 0
}

func ParseMove(m string, p Position) (Move, error) {
	from := SquareFromString(m[:2])
	to := SquareFromString(m[2:4])
	if len(m) == 5 {
		switch m[4] {
		case 'b':
			return NewPromotionMove(from, to, Bishop), nil
		case 'n':
			return NewPromotionMove(from, to, Knight), nil
		case 'r':
			return NewPromotionMove(from, to, Rook), nil
		case 'q':
			return NewPromotionMove(from, to, Queen), nil
		default:
			return Move(0), fmt.Errorf("invalid move suffix")
		}
	} else {
		return NewMove(from, to), nil
	}
}

func (m Move) String() string {
	fromRank, fromFile := m.From().RankAndFile()
	toRank, toFile := m.To().RankAndFile()

	suffix := ""

	switch m.PromoPiece() {
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

	return fmt.Sprintf("%c%d%c%d%s", 'a'+fromFile, fromRank+1, 'a'+toFile, toRank+1, suffix)
}
