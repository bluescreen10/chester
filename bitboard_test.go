package main_test

import (
	"fmt"
	"testing"

	pawned "github.com/bluescreen10/pawned"
)

func TestPopLSB(t *testing.T) {
	tests := []struct {
		in, out pawned.BitBoard
		sq      pawned.Square
	}{
		{0x01, 0x00, 0},
		{0x02, 0x00, 1},
		{0x03, 0x02, 0},
		{0x04, 0x00, 2},
		{0x06, 0x04, 1},
		{0x09, 0x08, 0},
		{0xf000_0000_0000_0000, 0xe000_0000_0000_0000, 60},
		{0xffff_ffff_ffff_fffe, 0xffff_ffff_ffff_fffc, 1},
		{0xffff_ffff_ffff_ffff, 0xffff_ffff_ffff_fffe, 0},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("PopLSB(%d)", i), func(t *testing.T) {
			b := pawned.BitBoard(test.in)
			sq, out := b.PopLSB()
			if sq != test.sq || out != test.out {
				t.Fatalf("PopLSB(%d) got sq(%d) bb(%b), want sq(%d) bb(%b)", i, sq, out, test.sq, test.out)
			}
		})
	}
}
