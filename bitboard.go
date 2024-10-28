package main

import (
	"fmt"
	"math/bits"
	"strings"
)

type BitBoard uint64

const (
	File_A BitBoard = 0x0101010101010101
	File_B BitBoard = 0x0202020202020202
	File_C BitBoard = 0x0404040404040404
	File_D BitBoard = 0x0808080808080808
	File_E BitBoard = 0x1010101010101010
	File_F BitBoard = 0x2020202020202020
	File_G BitBoard = 0x4040404040404040
	File_H BitBoard = 0x8080808080808080

	Rank_1 BitBoard = 0xFF00000000000000
	Rank_2 BitBoard = 0x00FF000000000000
	Rank_3 BitBoard = 0x0000FF0000000000
	Rank_4 BitBoard = 0x000000FF00000000
	Rank_5 BitBoard = 0x00000000FF000000
	Rank_6 BitBoard = 0x0000000000FF0000
	Rank_7 BitBoard = 0x000000000000FF00
	Rank_8 BitBoard = 0x00000000000000FF

	File_Not_A BitBoard = ^File_A
	File_Not_H BitBoard = ^File_H

	EmptyBoard BitBoard = 0
	FullBoard  BitBoard = 0xFFFFFFFFFFFFFFFF
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

func (b BitBoard) OnesCount() int8 {
	return int8(bits.OnesCount64(uint64(b)))
}

func (b BitBoard) RotateLeft(offset int) BitBoard {
	return BitBoard(bits.RotateLeft64(uint64(b), offset))
}
