package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type Square int8

func (s Square) RankAndFile() (int8, int8) {
	return 7 - int8(s/8), int8(s % 8)
}

func SquareFromRankAndFile(rank, file int8) Square {
	return Square((7-rank)*8 + file)
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
)

const (
	colorMask Piece = 0x08
)

var statePool = sync.Pool{
	New: func() any {
		return &State{}
	},
}

type State struct {
	WhiteToMove bool

	// Castling rights
	WhiteQueenSideCastle bool
	WhiteKingSideCastle  bool
	BlackQueenSideCastle bool
	BlackKingSideCastle  bool

	// En passant file
	EnPassantFile uint8

	// Moves
	HalfMoves uint8
	FullMoves uint16

	// Previous state
	LastMove *Move
	Capture  Piece
	Previous *State
}

func (s *State) Reset() {
	s.WhiteToMove = false
	s.WhiteKingSideCastle = false
	s.WhiteQueenSideCastle = false
	s.BlackKingSideCastle = false
	s.BlackQueenSideCastle = false
	s.EnPassantFile = 0
	s.HalfMoves = 0
	s.FullMoves = 0
	s.Previous = nil
}

type Position struct {
	Board    [Square(64)]Piece
	State    *State
	Occupied BitBoard
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

	pos.State = statePool.Get().(*State)
	pos.State.Reset()

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
	index := (7-rank)*8 + file
	bit := BitBoard(1) << uint(index)

	if piece == Empty {
		p.Occupied &^= bit
	} else {
		p.Occupied |= bit
	}

	p.Board[index] = piece
}

func (p Position) Get(rank, file int) Piece {
	index := (7-rank)*8 + file
	return p.Board[index]
}

func (p Position) Fen() string {
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

	if p.CanWhiteCastleKingSide() {
		fen.WriteByte('K')
	} else {
		fen.WriteByte('-')
	}

	if p.CanWhiteCastleQueenSide() {
		fen.WriteByte('Q')
	} else {
		fen.WriteByte('-')
	}

	if p.CanBlackCastleKingSide() {
		fen.WriteByte('k')
	} else {
		fen.WriteByte('-')
	}

	if p.CanBlackCastleQueenSide() {
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

func (p Position) String() string {
	builder := strings.Builder{}

	builder.WriteString("+---+---+---+---+---+---+---+---+\n")
	for rank := 7; rank >= 0; rank-- {
		builder.WriteByte('|')
		for file := 0; file < 8; file++ {
			builder.WriteString(fmt.Sprintf(" %s |", p.Get(rank, file)))
		}
		builder.WriteString(fmt.Sprintf(" %d\n", rank+1))
		builder.WriteString("+---+---+---+---+---+---+---+---+\n")
	}
	builder.WriteString("  a   b   c   d   e   f   g   h\n")
	return builder.String()
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

func (p Position) EnPassantFile() int8 {
	return int8(p.State.EnPassantFile - 1)
}

func (p *Position) CanWhiteCastleKingSide() bool {
	return p.State.WhiteKingSideCastle
}

func (p *Position) CanWhiteCastleQueenSide() bool {
	return p.State.WhiteQueenSideCastle
}

func (p *Position) CanBlackCastleKingSide() bool {
	return p.State.BlackKingSideCastle
}

func (p *Position) CanBlackCastleQueenSide() bool {
	return p.State.BlackQueenSideCastle
}

func (p *Position) HalfMoves() uint8 {
	return p.State.HalfMoves
}

func (p *Position) FullMoves() uint16 {
	return p.State.FullMoves
}

func (p *Position) Do(m Move) {

	newState := statePool.Get().(*State)

	newState.Previous = p.State
	newState.HalfMoves = p.State.HalfMoves + 1
	newState.FullMoves = p.State.FullMoves
	newState.WhiteKingSideCastle = p.State.WhiteKingSideCastle
	newState.WhiteQueenSideCastle = p.State.WhiteQueenSideCastle
	newState.BlackKingSideCastle = p.State.BlackKingSideCastle
	newState.BlackQueenSideCastle = p.State.BlackQueenSideCastle
	newState.EnPassantFile = 0
	newState.LastMove = &m

	from, to := m.From, m.To
	bitFrom, bitTo := BitBoard(1)<<uint(m.From), BitBoard(1)<<uint(m.To)
	piece := p.Board[from]
	enemy := p.Board[to]
	newState.Capture = enemy

	switch piece {
	case WhiteBishop, BlackBishop, WhiteKnight, BlackKnight, WhiteQueen, BlackQueen:
		p.Board[to] = piece
		p.Board[from] = Empty
		p.Occupied |= bitTo
		p.Occupied &^= bitFrom
	case WhiteRook:
		if from == SQ_A1 {
			newState.WhiteQueenSideCastle = false
		}

		if from == SQ_H1 {
			newState.WhiteKingSideCastle = false
		}

		p.Board[to] = piece
		p.Board[from] = Empty
		p.Occupied |= bitTo
		p.Occupied &^= bitFrom
	case BlackRook:
		if from == SQ_A8 {
			newState.BlackQueenSideCastle = false
		}

		if from == SQ_H8 {
			newState.BlackKingSideCastle = false
		}

		p.Board[to] = piece
		p.Board[from] = Empty
		p.Occupied |= bitTo
		p.Occupied &^= bitFrom

	case WhiteKing:
		newState.WhiteKingSideCastle = false
		newState.WhiteQueenSideCastle = false
		if m.Type == Castle {
			p.Board[to] = piece
			if to == SQ_G1 {
				p.Board[from] = Empty
				p.Board[SQ_F1] = WhiteRook
				p.Board[SQ_H1] = Empty
				p.Board[to] = piece
				p.Occupied |= BitBoard(3) << SQ_F1
				p.Occupied &^= BitBoard(9) << SQ_E1
			} else {
				p.Board[from] = Empty
				p.Board[SQ_D1] = WhiteRook
				p.Board[to] = piece
				p.Occupied |= bitTo
				p.Occupied |= BitBoard(3) << SQ_C1
				p.Occupied &^= BitBoard(17) << SQ_A1
			}
		} else {
			p.Board[to] = piece
			p.Board[from] = Empty
			p.Occupied |= bitTo
			p.Occupied &^= bitFrom
		}
	case BlackKing:
		newState.BlackKingSideCastle = false
		newState.BlackQueenSideCastle = false
		if m.Type == Castle {
			p.Board[to] = piece
			if to == SQ_G8 {
				p.Board[from] = BlackRook
				p.Board[SQ_H8] = Empty
				p.Board[to] = piece
				p.Occupied |= BitBoard(3) << SQ_F8
				p.Occupied &^= BitBoard(9) << SQ_E8
			} else {
				p.Board[from] = Empty
				p.Board[SQ_D8] = BlackRook
				p.Board[to] = piece
				p.Occupied |= bitTo
				p.Occupied |= BitBoard(3) << SQ_C8
				p.Occupied &^= BitBoard(17) << SQ_A8
			}
		} else {
			p.Board[to] = piece
			p.Board[from] = Empty
			p.Occupied |= bitTo
			p.Occupied &^= bitFrom
		}
	case WhitePawn:
		switch m.Type {
		case PromotionToKnight:
			piece = WhiteKnight
		case PromotionToBishop:
			piece = WhiteBishop
		case PromotionToRook:
			piece = WhiteRook
		case PromotionToQueen:
			piece = WhiteQueen
		case EnPassant:
			newState.Capture = BlackPawn
			p.Board[to+8] = Empty
			p.Occupied &^= bitTo << 8
		}

		if from-to == 16 {
			newState.EnPassantFile = uint8(to%8) + 1
		}

		p.Board[to] = piece
		p.Board[from] = Empty
		p.Occupied |= bitTo
		p.Occupied &^= bitFrom
	case BlackPawn:
		switch m.Type {
		case PromotionToKnight:
			piece = BlackKnight
		case PromotionToBishop:
			piece = BlackBishop
		case PromotionToRook:
			piece = BlackRook
		case PromotionToQueen:
			piece = BlackQueen
		case EnPassant:
			newState.Capture = WhitePawn
			p.Board[to-8] = Empty
			p.Occupied &^= bitTo >> 8
		}

		if to-from == 16 {
			newState.EnPassantFile = uint8(to%8) + 1
		}

		p.Board[to] = piece
		p.Board[from] = Empty
		p.Occupied |= bitTo
		p.Occupied &^= bitFrom
	}

	if p.SideToMove() == Black {
		newState.WhiteToMove = true
		newState.FullMoves++
	} else {
		newState.WhiteToMove = false
	}

	p.State = newState
}

func (p *Position) Undo() {
	m := p.State.LastMove
	if m == nil {
		return
	}

	from, to := m.From, m.To
	bitFrom, bitTo := BitBoard(1)<<uint8(from), BitBoard(1)<<uint(to)
	piece := p.Board[to]
	enemy := p.State.Capture

	switch piece {
	case WhiteKnight, WhiteBishop, WhiteRook, WhiteQueen:
		if m.Type == PromotionToKnight || m.Type == PromotionToBishop || m.Type == PromotionToRook || m.Type == PromotionToQueen {
			piece = WhitePawn
		}
		p.Board[to] = enemy
		p.Board[from] = piece
		p.Occupied |= bitFrom
		if enemy == Empty {
			p.Occupied &^= bitTo
		}
	case BlackKnight, BlackBishop, BlackRook, BlackQueen:
		if m.Type == PromotionToKnight || m.Type == PromotionToBishop || m.Type == PromotionToRook || m.Type == PromotionToQueen {
			piece = BlackPawn
		}
		p.Board[to] = enemy
		p.Board[from] = piece
		p.Occupied |= bitFrom
		if enemy == Empty {
			p.Occupied &^= bitTo
		}
	case WhitePawn:
		if m.Type != EnPassant {
			p.Board[to] = enemy
			p.Board[from] = piece
			p.Occupied |= bitFrom
			if enemy == Empty {
				p.Occupied &^= bitTo
			}
		} else {
			p.Board[to] = Empty
			p.Board[to+8] = enemy
			p.Board[from] = piece
			p.Occupied |= bitFrom
			p.Occupied |= bitTo << uint(8)
			p.Occupied &^= bitTo
		}
	case BlackPawn:
		if m.Type != EnPassant {
			p.Board[to] = enemy
			p.Board[from] = piece
			p.Occupied |= bitFrom
			if enemy == Empty {
				p.Occupied &^= bitTo
			}
		} else {
			p.Board[to] = Empty
			p.Board[to-8] = enemy
			p.Board[from] = piece
			p.Occupied |= bitFrom
			p.Occupied |= bitTo >> uint(8)
			p.Occupied &^= bitTo
		}
	case WhiteKing:
		if m.Type != Castle {
			p.Board[to] = enemy
			p.Board[from] = piece
			p.Occupied |= bitFrom
			if enemy == Empty {
				p.Occupied &^= bitTo
			}
		} else {
			if to == SQ_G1 {
				p.Board[from] = piece
				p.Board[to] = Empty
				p.Board[SQ_F1] = Empty
				p.Board[SQ_H1] = WhiteRook
				p.Occupied |= BitBoard(9) << SQ_E1
				p.Occupied &^= BitBoard(3) << SQ_F1
			} else {
				p.Board[from] = piece
				p.Board[to] = Empty
				p.Board[SQ_D1] = Empty
				p.Board[SQ_A1] = WhiteRook
				p.Occupied |= BitBoard(17) << SQ_A1
				p.Occupied &^= BitBoard(3) << SQ_C1
			}
		}
	case BlackKing:
		if m.Type != Castle {
			p.Board[to] = enemy
			p.Board[from] = piece
			p.Occupied |= bitFrom
			if enemy == Empty {
				p.Occupied &^= bitTo
			}
		} else {
			if to == SQ_G8 {
				p.Board[from] = piece
				p.Board[to] = Empty
				p.Board[SQ_F8] = Empty
				p.Board[SQ_H8] = BlackRook
				p.Occupied |= BitBoard(9) << SQ_E8
				p.Occupied &^= BitBoard(3) << SQ_F8
			} else {
				p.Board[from] = piece
				p.Board[to] = Empty
				p.Board[SQ_D8] = Empty
				p.Board[SQ_A8] = BlackRook
				p.Occupied |= BitBoard(17) << SQ_A8
				p.Occupied &^= BitBoard(3) << SQ_C8
			}
		}
	}

	defer statePool.Put(p.State)
	p.State = p.State.Previous
}
