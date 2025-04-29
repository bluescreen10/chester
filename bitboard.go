package main

import (
	"fmt"
	"math/bits"
	"strings"
)

type BitBoard uint64

func (b BitBoard) String() string {
	builder := strings.Builder{}

	builder.WriteString("+---+---+---+---+---+---+---+---+\n")
	bit := BitBoard(1)
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

func (b BitBoard) PopLSB() (Square, BitBoard) {
	s := Square(bits.TrailingZeros64(uint64(b)))
	b &= b - 1
	return s, b
}

func (b BitBoard) OnesCount() int8 {
	return int8(bits.OnesCount64(uint64(b)))
}

func (b BitBoard) RotateLeft(offset int) BitBoard {
	return BitBoard(bits.RotateLeft64(uint64(b), offset))
}

func (b BitBoard) Square() Square {
	return Square(bits.TrailingZeros64(uint64(b)))
}

func NewBitBoardFromSquare(sq Square) BitBoard {
	return 1 << sq
}
