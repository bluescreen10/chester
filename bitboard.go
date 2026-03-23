package chester

import (
	"fmt"
	"math/bits"
	"strings"
)

// Bitboard is a 64-bit mask representing a set of squares. Each bit
// corresponds to one square using the same index as Square: bit 0 = a8,
// bit 63 = h1. Multiple bits may be set to represent a set of squares,
// such as all squares occupied by white pawns.
type Bitboard uint64

// Single-square Bitboard constants for the first and eighth ranks.
// Named BB_SQ_<file><rank>, e.g. BB_SQ_E1 has only the e1 bit set.
// Used as masks in castling legality checks and rights updates.
const (
	BB_SQ_A1 Bitboard = 1 << SQ_A1
	BB_SQ_B1 Bitboard = 1 << SQ_B1
	BB_SQ_C1 Bitboard = 1 << SQ_C1
	BB_SQ_D1 Bitboard = 1 << SQ_D1
	BB_SQ_E1 Bitboard = 1 << SQ_E1
	BB_SQ_F1 Bitboard = 1 << SQ_F1
	BB_SQ_G1 Bitboard = 1 << SQ_G1
	BB_SQ_H1 Bitboard = 1 << SQ_H1

	BB_SQ_A8 Bitboard = 1 << SQ_A8
	BB_SQ_B8 Bitboard = 1 << SQ_B8
	BB_SQ_C8 Bitboard = 1 << SQ_C8
	BB_SQ_D8 Bitboard = 1 << SQ_D8
	BB_SQ_E8 Bitboard = 1 << SQ_E8
	BB_SQ_F8 Bitboard = 1 << SQ_F8
	BB_SQ_G8 Bitboard = 1 << SQ_G8
	BB_SQ_H8 Bitboard = 1 << SQ_H8
)

// String returns a human-readable ASCII grid of the bitboard with rank
// numbers and file letters. Set bits are shown as X, clear bits as spaces.
func (b Bitboard) String() string {
	builder := strings.Builder{}

	builder.WriteString("+---+---+---+---+---+---+---+---+\n")
	bit := Bitboard(1)
	for r := 7; r >= 0; r-- {
		builder.WriteString("|")
		for f := 0; f < 8; f++ {
			if b&bit != 0 {
				builder.WriteString(" X |")
			} else {
				builder.WriteString("   |")
			}
			bit <<= 1
		}
		builder.WriteString(fmt.Sprintf(" %d \n", r+1))
		builder.WriteString("+---+---+---+---+---+---+---+---+\n")
	}
	builder.WriteString("  a   b   c   d   e   f   g   h\n")
	return builder.String()
}

// PopLSB removes the least significant set bit and returns its Square index
// together with the modified Bitboard. Typical usage:
//
//	for bb != 0 {
//	    sq, bb = bb.PopLSB()
//	    // process sq
//	}
func (b Bitboard) PopLSB() (Square, Bitboard) {
	s := Square(bits.TrailingZeros64(uint64(b)))
	b &= b - 1
	return s, b
}

// OnesCount returns the number of set bits (population count).
func (b Bitboard) OnesCount() int {
	return bits.OnesCount64(uint64(b))
}

// RotateLeft rotates all 64 bits left by offset positions. Negative values
// rotate right. Used to shift pawn attack and push masks without branching
// on color: a positive offset for Black and a negative offset for White.
func (b Bitboard) RotateLeft(offset int) Bitboard {
	return Bitboard(bits.RotateLeft64(uint64(b), offset))
}

// NewBitboardFromSquare returns a Bitboard with only the bit for sq set.
func NewBitboardFromSquare(sq Square) Bitboard {
	return 1 << sq
}
