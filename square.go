package chester

import "fmt"

// Square identifies a board square by its index. The encoding places a8 at
// index 0 and h1 at index 63, matching the LSB-first bitboard layout:
//
//	a8=0  b8=1  ... h8=7
//	a7=8  b7=9  ... h7=15
//	...
//	a1=56 b1=57 ... h1=63
//
// SQ_NULL is the sentinel value for "no square".
type Square int8

// RankAndFile returns the rank (0=rank1 .. 7=rank8) and file (0=a .. 7=h)
// of the square.
func (s Square) RankAndFile() (int8, int8) {
	return 7 - int8(s/8), int8(s % 8)
}

// File returns the file of the square (0=a .. 7=h).
func (s Square) File() int8 {
	return int8(s % 8)
}

// Rank returns the rank of the square (0=rank1 .. 7=rank8).
func (s Square) Rank() int8 {
	return 7 - int8(s/8)
}

// SquareFromRankAndFile constructs a Square from a rank (0=rank1 .. 7=rank8)
// and file (0=a .. 7=h).
func SquareFromRankAndFile(rank, file int8) Square {
	return Square((7-rank)*8 + file)
}

// String returns the algebraic name of the square (e.g. "e4"), or "-" for
// SQ_NULL. This is the format used in FEN strings.
func (s Square) String() string {
	if s == SQ_NULL {
		return "-"
	}
	rank, file := s.RankAndFile()
	return fmt.Sprintf("%c%d", file+'a', rank+1)
}

// ParseSquare parses an algebraic square name (e.g., "e4") and returns the Square.
func ParseSquare(s string) (Square, error) {
	if len(s) != 2 {
		return SQ_NULL, fmt.Errorf("invalid Square: %s", s)
	}
	rank := int8(s[1] - '1')
	file := int8(s[0] - 'a')
	if rank < 0 || rank > 7 || file < 0 || file > 7 {
		return SQ_NULL, fmt.Errorf("invalid Square: %s", s)
	}
	return SquareFromRankAndFile(rank, file), nil
}
