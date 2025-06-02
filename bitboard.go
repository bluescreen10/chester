package main

import (
	"fmt"
	"math/bits"
	"strings"
)

type BitBoard uint64

const (
	BB_SQ_A1 BitBoard = 1 << SQ_A1
	BB_SQ_B1 BitBoard = 1 << SQ_B1
	BB_SQ_C1 BitBoard = 1 << SQ_C1
	BB_SQ_D1 BitBoard = 1 << SQ_D1
	BB_SQ_E1 BitBoard = 1 << SQ_E1
	BB_SQ_F1 BitBoard = 1 << SQ_F1
	BB_SQ_G1 BitBoard = 1 << SQ_G1
	BB_SQ_H1 BitBoard = 1 << SQ_H1

	BB_SQ_A8 BitBoard = 1 << SQ_A8
	BB_SQ_B8 BitBoard = 1 << SQ_B8
	BB_SQ_C8 BitBoard = 1 << SQ_C8
	BB_SQ_D8 BitBoard = 1 << SQ_D8
	BB_SQ_E8 BitBoard = 1 << SQ_E8
	BB_SQ_F8 BitBoard = 1 << SQ_F8
	BB_SQ_G8 BitBoard = 1 << SQ_G8
	BB_SQ_H8 BitBoard = 1 << SQ_H8
)

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

func (b BitBoard) OnesCount() int {
	return bits.OnesCount64(uint64(b))
}

func (b BitBoard) RotateLeft(offset int) BitBoard {
	return BitBoard(bits.RotateLeft64(uint64(b), offset))
}

func (b BitBoard) Square() Square {
	if sq := bits.TrailingZeros64(uint64(b)); sq != 0 {
		return Square(sq)
	}

	return SQ_NULL
}

func NewBitBoardFromSquare(sq Square) BitBoard {
	return 1 << sq
}
