package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	colorMask Piece = 0x08
)

type State struct {
	WhiteToMove          bool
	WhiteQueenSideCastle bool
	WhiteKingSideCastle  bool
	BlackQueenSideCastle bool
	BlackKingSideCastle  bool
	EnPassantFile        uint8
	HalfMoves            uint8
	FullMoves            uint16
	Previous             *State
}

type Position struct {
	Board [64]Piece
	State *State
}

func Parse(fen string) (Position, error) {
	var pos Position

	parts := strings.Split(strings.TrimSpace(fen), " ")
	if len(parts) < 6 {
		return pos, fmt.Errorf("invalid fen: %s", fen)
	}

	rank, file := 7, 0

	for _, c := range parts[0] {
		switch c {
		case '/':
			if file != 8 {
				return pos, fmt.Errorf("invalid row: %s", fen)
			}
			rank--
			file = 0
		case '1', '2', '3', '4', '5', '6', '7', '8':
			n := int(c - '0')
			if file+n > 8 {
				return pos, fmt.Errorf("invalid rank: %s", fen)
			}
			file += n
		default:
			piece, err := parsePiece(byte(c))
			if err != nil {
				return pos, err
			}
			pos.Set(rank, file, piece)
			file++
		}
	}

	pos.State = &State{}

	if parts[1] == "w" || parts[1] == "W" {
		pos.State.WhiteToMove = true
	}

	for _, c := range parts[2] {
		switch c {
		case 'K':
			pos.State.WhiteKingSideCastle = true
		case 'Q':
			pos.State.WhiteQueenSideCastle = true
		case 'k':
			pos.State.BlackKingSideCastle = true

		case 'q':
			pos.State.BlackQueenSideCastle = true
		}
	}

	if parts[3] != "-" {
		file := uint8(parts[3][0] - 'a' + 1)
		pos.State.EnPassantFile = file
	}

	halfMoves, err := strconv.Atoi(parts[4])
	if err != nil {
		return pos, fmt.Errorf("invalid half moves: %s", fen)
	}
	pos.State.HalfMoves = uint8(halfMoves)

	fullMoves, err := strconv.Atoi(parts[5])
	if err != nil {
		return pos, fmt.Errorf("invalid full moves: %s", fen)
	}
	pos.State.FullMoves = uint16(fullMoves)

	return pos, nil
}

func parsePiece(c byte) (Piece, error) {
	switch c {
	case 'P':
		return WhitePawn, nil
	case 'N':
		return WhiteKnight, nil
	case 'B':
		return WhiteBishop, nil
	case 'R':
		return WhiteRook, nil
	case 'Q':
		return WhiteQueen, nil
	case 'K':
		return WhiteKing, nil
	case 'p':
		return BlackPawn, nil
	case 'n':
		return BlackKnight, nil
	case 'b':
		return BlackBishop, nil
	case 'r':
		return BlackRook, nil
	case 'q':
		return BlackQueen, nil
	case 'k':
		return BlackKing, nil
	default:
		return Empty, fmt.Errorf("invalid piece: %c", c)
	}
}

func (p *Position) Set(rank, file int, piece Piece) {
	p.Board[rank*8+file] = piece
}

func (p Position) Get(rank, file int) Piece {
	return p.Board[rank*8+file]
}

func (p Position) String() string {
	fen := strings.Builder{}

	for rank := 7; rank >= 0; rank-- {
		for file := 0; file < 8; file++ {
			piece := p.Get(rank, file)
			emptySpaces := 0
			for piece == Empty && file < 8 {
				emptySpaces++
				file++
				piece = p.Get(rank, file)
			}
			if emptySpaces > 0 {
				fen.WriteString(fmt.Sprintf("%d", emptySpaces))
			} else {
				fen.WriteString(piece.String())
			}
		}
		if rank > 0 {
			fen.WriteByte('/')
		}
	}

	fen.WriteByte(' ')

	if p.SideToMove() == White {
		fen.WriteByte('w')
	} else {
		fen.WriteByte('b')
	}

	fen.WriteByte(' ')

	if p.CanCastleKingSide(White) {
		fen.WriteByte('K')
	} else {
		fen.WriteByte('-')
	}

	if p.CanCastleQueenSide(White) {
		fen.WriteByte('Q')
	} else {
		fen.WriteByte('-')
	}

	if p.CanCastleKingSide(Black) {
		fen.WriteByte('k')
	} else {
		fen.WriteByte('-')
	}

	if p.CanCastleQueenSide(Black) {
		fen.WriteByte('q')
	} else {
		fen.WriteByte('-')
	}

	fen.WriteByte(' ')

	if p.IsEnPassant() {
		file := p.EnPassantFile()

		if p.SideToMove() == White {
			fen.WriteString(fmt.Sprintf("%c5", 'a'+file))
		} else {
			fen.WriteString(fmt.Sprintf("%c4", 'a'+file))
		}
	} else {
		fen.WriteByte('-')
	}

	fen.WriteString(fmt.Sprintf(" %d %d", p.HalfMoves(), p.FullMoves()))

	return fen.String()
}

func (p Position) PrettyPrint() string {
	board := strings.Builder{}

	for rank := 7; rank >= 0; rank-- {
		for file := 0; file < 8; file++ {
			board.WriteString(p.Get(rank, file).String())
		}
		board.WriteByte('\n')
	}

	return board.String()
}

func (p Position) SideToMove() Piece {
	if p.State.WhiteToMove {
		return White
	}
	return Black
}

func (p Position) IsEnPassant() bool {
	return p.State.EnPassantFile != 0
}

func (p Position) EnPassantFile() int {
	return int(p.State.EnPassantFile - 1)
}

func (p *Position) CanCastleKingSide(stm Piece) bool {
	if p.SideToMove() == White {
		return p.State.WhiteKingSideCastle
	} else {
		return p.State.BlackKingSideCastle
	}
}

func (p *Position) CanCastleQueenSide(stm Piece) bool {
	if p.SideToMove() == White {
		return p.State.WhiteQueenSideCastle
	} else {
		return p.State.BlackKingSideCastle
	}
}

func (p *Position) HalfMoves() uint8 {
	return p.State.HalfMoves
}

func (p *Position) FullMoves() uint16 {
	return p.State.FullMoves
}

func (p *Position) DoMove(m Move) {

	newState := *p.State

	from, to := int(m.From), int(m.To)
	newState.EnPassantFile = 0

	switch m.Type {
	case Default, Capture:
		p.Board[to] = p.Board[from]
		p.Board[from] = Empty
		if (m.Piece == WhitePawn || m.Piece == BlackPawn) && (to-from == 16 || to-from == -16) {
			newState.EnPassantFile = uint8(to%8 + 1)
		}

		if m.Piece == WhiteKing {
			newState.WhiteKingSideCastle = false
			newState.WhiteQueenSideCastle = false
		}

		if m.Piece == BlackKing {
			newState.BlackKingSideCastle = false
			newState.BlackQueenSideCastle = false
		}

		if m.Piece == WhiteRook && from == 0 {
			newState.WhiteKingSideCastle = false
		}

		if m.Piece == WhiteRook && from == 7 {
			newState.WhiteQueenSideCastle = false
		}

		if m.Piece == BlackRook && from == 56 {
			newState.BlackKingSideCastle = false
		}

		if m.Piece == BlackRook && from == 63 {
			newState.BlackQueenSideCastle = false
		}

	case EnPassant:
		p.Board[to] = p.Board[from]
		p.Board[from] = Empty
		if p.SideToMove() == White {
			p.Board[to-8] = Empty
		} else {
			p.Board[to+8] = Empty
		}
	case CastleKingSide:
		if p.SideToMove() == White {
			newState.WhiteKingSideCastle = false
			newState.WhiteQueenSideCastle = false
			p.Board[5] = WhiteRook
			p.Board[7] = Empty
			p.Board[4] = WhiteKing
		} else {
			newState.BlackKingSideCastle = false
			newState.BlackQueenSideCastle = false
			p.Board[61] = BlackRook
			p.Board[63] = Empty
			p.Board[60] = BlackKing
		}
	case CastleQueenSide:
		if p.SideToMove() == White {
			newState.WhiteKingSideCastle = false
			newState.WhiteQueenSideCastle = false
			p.Board[3] = WhiteRook
			p.Board[0] = Empty
			p.Board[4] = WhiteKing
		} else {
			newState.BlackKingSideCastle = false
			newState.BlackQueenSideCastle = false
			p.Board[59] = BlackRook
			p.Board[56] = Empty
			p.Board[60] = BlackKing
		}
	case Promotion:
		p.Board[to] = m.Piece
		p.Board[from] = Empty
	}

	newState.HalfMoves++

	if p.SideToMove() == Black {
		newState.WhiteToMove = true
		newState.FullMoves++
	} else {
		newState.WhiteToMove = false
	}

	newState.Previous = p.State
	p.State = &newState
}

func (p *Position) UndoMove(m Move) {
	from, to := int(m.From), int(m.To)

	switch m.Type {
	case Default:
		p.Board[from] = p.Board[to]
		p.Board[to] = Empty
	case Capture:
		p.Board[from] = p.Board[to]
		p.Board[to] = m.Enemy
	case EnPassant:
		p.Board[from] = p.Board[to]
		p.Board[to] = Empty
		if p.SideToMove() == White {
			p.Board[to+8] = WhitePawn
		} else {
			p.Board[to-8] = BlackPawn
		}
	case CastleKingSide:
		if p.SideToMove() == White {
			p.Board[7] = WhiteRook
			p.Board[5] = Empty
		} else {
			p.Board[63] = BlackRook
			p.Board[61] = Empty
		}
	case CastleQueenSide:
		if p.SideToMove() == White {
			p.Board[0] = WhiteRook
			p.Board[3] = Empty
		} else {
			p.Board[56] = BlackRook
			p.Board[59] = Empty
		}
	case Promotion:
		p.Board[from] = WhitePawn
		p.Board[to] = m.Enemy
	}

	if p.State.Previous != nil {
		p.State = p.State.Previous
	}
}
