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

func (s Square) String() string {
	if s == SQ_NULL {
		return "-"
	}
	rank, file := s.RankAndFile()
	return fmt.Sprintf("%c%d", file+'a', rank+1)
}
