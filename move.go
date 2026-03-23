package chester

import "fmt"

// Move encodes a chess move in 16 bits.
//
//	15 14 13 12 11 10  9  8  7  6  5  4  3  2  1  0
//
// +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
// |           |                 |                 |
// | promotion |      from       |        to       |
// |           |                 |                 |
// +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//
//	4 bits       6 bits          6 bits
//
// promotion
// 0000 -> no promotion
// 0001 -> knight
// 0010 -> bishop
// 0011 -> rook
// 0100 -> queen
type Move uint16

// NewMove encodes a move from from to to with no promotion.
func NewMove(from, to Square) Move {
	return Move(from)<<6 | Move(to)
}

// NewPromotionMove encodes a pawn promotion move from from to to, promoting
// to the given piece type.
func NewPromotionMove(from, to Square, promotion Piece) Move {
	return Move(promotion)<<12 | Move(from)<<6 | Move(to)
}

// From returns the origin square of the move.
func (m Move) From() Square {
	return Square(m >> 6 & 0x3f)
}

// To returns the destination square of the move.
func (m Move) To() Square {
	return Square(m & 0x3f)
}

// PromoPiece returns the promotion piece encoded in the move, or a zero value
// (Knight) if the move is not a promotion. Always check IsPromotion first.
func (m Move) PromoPiece() Piece {
	return Piece(m >> 12)
}

// IsPromotion reports whether the move encodes a pawn promotion.
func (m Move) IsPromotion() bool {
	return m&0xf000 != 0
}

// ParseMove parses a move string in pure algebraic coordinate notation
// (e.g. "e2e4", "e7e8q") relative to position p. An optional fifth character
// specifies the promotion piece: 'n', 'b', 'r', or 'q'.
// Returns an error if the string is malformed or the promotion character is
// unrecognised.
func ParseMove(m string, p *Position) (Move, error) {
	from, err := ParseSquare(m[:2])
	if err != nil {
		return Move(0), err
	}

	to, err := ParseSquare(m[2:4])
	if err != nil {
		return Move(0), err
	}

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

// String returns the move in pure algebraic coordinate notation, e.g. "e2e4"
// or "e7e8q" for a promotion to queen.
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
