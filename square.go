package main

import "fmt"

type Square int8

func (s Square) RankAndFile() (int8, int8) {
	return 7 - int8(s/8), int8(s % 8)
}

func SquareFromRankAndFile(rank, file int8) Square {
	return Square((7-rank)*8 + file)
}

func SquareFromString(s string) Square {
	if len(s) != 2 {
		return SQ_NULL
	}
	return SquareFromRankAndFile(int8(s[1]-'1'), int8(s[0]-'a'))
}

const (
	SQ_A8 Square = iota
	SQ_B8
	SQ_C8
	SQ_D8
	SQ_E8
	SQ_F8
	SQ_G8
	SQ_H8
	SQ_A7
	SQ_B7
	SQ_C7
	SQ_D7
	SQ_E7
	SQ_F7
	SQ_G7
	SQ_H7
	SQ_A6
	SQ_B6
	SQ_C6
	SQ_D6
	SQ_E6
	SQ_F6
	SQ_G6
	SQ_H6
	SQ_A5
	SQ_B5
	SQ_C5
	SQ_D5
	SQ_E5
	SQ_F5
	SQ_G5
	SQ_H5
	SQ_A4
	SQ_B4
	SQ_C4
	SQ_D4
	SQ_E4
	SQ_F4
	SQ_G4
	SQ_H4
	SQ_A3
	SQ_B3
	SQ_C3
	SQ_D3
	SQ_E3
	SQ_F3
	SQ_G3
	SQ_H3
	SQ_A2
	SQ_B2
	SQ_C2
	SQ_D2
	SQ_E2
	SQ_F2
	SQ_G2
	SQ_H2
	SQ_A1
	SQ_B1
	SQ_C1
	SQ_D1
	SQ_E1
	SQ_F1
	SQ_G1
	SQ_H1
	SQ_NULL
)

func (s Square) String() string {
	if s == SQ_NULL {
		return "-"
	}
	rank, file := s.RankAndFile()
	return fmt.Sprintf("%c%d", file+'a', rank+1)
}
