package main

type OldPiece uint8

const (
	OldEmpty OldPiece = iota
	WhitePawn
	WhiteKnight
	WhiteBishop
	WhiteRook
	WhiteQueen
	WhiteKing

	BlackPawn OldPiece = iota + 2
	BlackKnight
	BlackBishop
	BlackRook
	BlackQueen
	BlackKing

	OldWhite OldPiece = 0
	OldBlack OldPiece = 8
)

func (p OldPiece) String() string {
	switch p {
	case OldEmpty:
		return " "
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

func (p OldPiece) Color() Color {
	if p&colorMask != 0 {
		return Black
	} else {
		return White

	}
}

func IsSlidingOldPiece(p OldPiece) bool {
	return p == WhiteBishop || p == BlackBishop || p == WhiteRook || p == BlackRook || p == WhiteQueen || p == BlackQueen
}
