package main

import "fmt"

type Square int8

func (s Square) RankAndFile() (int8, int8) {
	return 7 - int8(s/8), int8(s % 8)
}

func (s Square) File() int8 {
	return int8(s % 8)
}

func (s Square) Rank() int8 {
	return 7 - int8(s/8)
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

func ParseSquare(s string) (Square, error) {
	if len(s) != 2 {
		return SQ_NULL, fmt.Errorf("invalid square: %s", s)
	}
	rank := int8(s[1] - '1')
	file := int8(s[0] - 'a')
	if rank < 0 || rank > 7 || file < 0 || file > 7 {
		return SQ_NULL, fmt.Errorf("invalid square: %s", s)
	}
	return SquareFromRankAndFile(rank, file), nil
}
