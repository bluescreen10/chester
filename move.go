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
// |   type    |      from       |        to       |
// |           |                 |                 |
// +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    4 bits       6 bits          6 bits
//
// type
// 0xxx -> move or captures xxx = piece
// 1000 -> castle king side
// 1001 -> castle queen side
// 1010 -> en passant
// 1011 -> double push
// 11xx -> promotion xx = piece

type Move uint16

func NewMove(from, to Square, piece Piece) Move {
	return Move(piece)<<12 | Move(from)<<6 | Move(to)
}

func NewDoublePushMove(from, to Square) Move {
	return Move(0x0b)<<12 | Move(from)<<6 | Move(to)
}

func NewEnPassantMove(from, to Square) Move {
	return Move(0x0a)<<12 | Move(from)<<6 | Move(to)
}

func NewCastleKingSideMove(from, to Square) Move {
	return Move(0x08)<<12 | Move(from)<<6 | Move(to)
}

func NewCastleQueenSideMove(from, to Square) Move {
	return Move(0x09)<<12 | Move(from)<<6 | Move(to)
}

func NewPromotionMove(from, to Square, promotion Piece) Move {
	return Move(0x0c|promotion)<<12 | Move(from)<<6 | Move(to)
}

func (m Move) From() Square {
	return Square(m >> 6 & 0x3f)
}

func (m Move) To() Square {
	return Square(m & 0x3f)
}

func (m Move) Type() MoveType {
	switch m & 0xf000 {
	case 0xc000, 0xd000, 0xe000, 0xf000:
		return Promotion
	case 0x8000:
		return CastleKingSide
	case 0x9000:
		return CastleQueenSide
	case 0xa000:
		return EnPassant
	case 0xb000:
		return DoublePush
	default:
		return Default
	}
}

func (m Move) Piece() Piece {
	switch p := m & 0xf000; p {
	case 0xc000, 0xd000, 0xe000, 0xf000:
		return Pawn
	case 0x8000, 0x9000:
		return King
	case 0xa000, 0xb000:
		return Pawn
	default:
		return Piece(p >> 12 & 0x07)
	}
}

func (m Move) PromoPiece() Piece {
	return Piece(m >> 12 & 0x03)
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
		piece := p.Get(from)
		if piece == Empty {
			return Move(0), fmt.Errorf("no piece at from square: %s", from)
		}
		return NewMove(from, to, piece), nil
	}
}

func (m Move) String() string {
	fromRank, fromFile := m.From().RankAndFile()
	toRank, toFile := m.To().RankAndFile()

	suffix := ""

	if m.Type() == Promotion {
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
	}

	return fmt.Sprintf("%c%d%c%d%s", 'a'+fromFile, fromRank+1, 'a'+toFile, toRank+1, suffix)
}
