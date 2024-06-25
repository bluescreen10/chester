package main

type Piece uint8

const (
	Empty Piece = iota
	WhitePawn
	WhiteKnight
	WhiteBishop
	WhiteRook
	WhiteQueen
	WhiteKing

	BlackPawn = iota + 2
	BlackKnight
	BlackBishop
	BlackRook
	BlackQueen
	BlackKing

	White Piece = 0
	Black Piece = 8
)

func (p Piece) String() string {
	switch p {
	case Empty:
		return "."
	case WhitePawn:
		return "P"
	case WhiteKnight:
		return "N"
	case WhiteBishop:
		return "B"
	case WhiteRook:
		return "R"
	case WhiteQueen:
		return "Q"
	case WhiteKing:
		return "K"
	case BlackPawn:
		return "p"
	case BlackKnight:
		return "n"
	case BlackBishop:
		return "b"
	case BlackRook:
		return "r"
	case BlackQueen:
		return "q"
	case BlackKing:
		return "k"
	default:
		return "?"
	}
}

func (p Piece) Color() Piece {
	return p & colorMask
}

func IsSlidingPiece(p Piece) bool {
	return p == WhiteBishop || p == BlackBishop || p == WhiteRook || p == BlackRook || p == WhiteQueen || p == BlackQueen
}
