package chester

import (
	"fmt"
	"math/bits"
	"strings"
)

type Bitboard uint64

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

func (b Bitboard) PopLSB() (Square, Bitboard) {
	s := Square(bits.TrailingZeros64(uint64(b)))
	b &= b - 1
	return s, b
}

func (b Bitboard) OnesCount() int {
	return bits.OnesCount64(uint64(b))
}

func (b Bitboard) RotateLeft(offset int) Bitboard {
	return Bitboard(bits.RotateLeft64(uint64(b), offset))
}

func (b Bitboard) Square() Square {
	if b == 0 {
		return SQ_NULL
	}

	sq := bits.TrailingZeros64(uint64(b))
	return Square(sq)
}

func NewBitboardFromSquare(sq Square) Bitboard {
	return 1 << sq
}
