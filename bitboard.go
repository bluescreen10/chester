package main

import (
	"fmt"
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
	builder.WriteString("  a   b   c   d   e   f   g  h\n")
	return builder.String()
}
