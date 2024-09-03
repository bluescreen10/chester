package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

const (
	colorMask OldPiece = 0x08
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

type Color uint8

const (
	White Color = iota
	Black
)

type Piece uint8

const (
	Pawn Piece = iota
	Knight
	Bishop
	Rook
	Queen
	King
	Empty
)

type Position struct {
	Pieces    [Color(2)][Piece(6)]BitBoard
	AllPieces [Color(2)]BitBoard
	Occupied  BitBoard
	State     *State
}

func Parse(fen string) (Position, error) {
	var pos Position

	parts := strings.Split(strings.TrimSpace(fen), " ")
	if len(parts) < 6 {
		return pos, fmt.Errorf("invalid fen: %s", fen)
	}

	bit := BitBoard(1)

	for _, row := range strings.Split(parts[0], "/") {
		for _, char := range row {
			switch char {
			case 'P':
				pos.Pieces[White][Pawn] |= bit
				pos.AllPieces[White] |= bit
				pos.Occupied |= bit
			case 'N':
				pos.Pieces[White][Knight] |= bit
				pos.AllPieces[White] |= bit
				pos.Occupied |= bit
			case 'B':
				pos.Pieces[White][Bishop] |= bit
				pos.AllPieces[White] |= bit
				pos.Occupied |= bit
			case 'R':
				pos.Pieces[White][Rook] |= bit
				pos.AllPieces[White] |= bit
				pos.Occupied |= bit
			case 'Q':
				pos.Pieces[White][Queen] |= bit
				pos.AllPieces[White] |= bit
				pos.Occupied |= bit
			case 'K':
				pos.Pieces[White][King] |= bit
				pos.AllPieces[White] |= bit
				pos.Occupied |= bit
			case 'p':
				pos.Pieces[Black][Pawn] |= bit
				pos.AllPieces[Black] |= bit
				pos.Occupied |= bit
			case 'n':
				pos.Pieces[Black][Knight] |= bit
				pos.AllPieces[Black] |= bit
				pos.Occupied |= bit
			case 'b':
				pos.Pieces[Black][Bishop] |= bit
				pos.AllPieces[Black] |= bit
				pos.Occupied |= bit
			case 'r':
				pos.Pieces[Black][Rook] |= bit
				pos.AllPieces[Black] |= bit
				pos.Occupied |= bit
			case 'q':
				pos.Pieces[Black][Queen] |= bit
				pos.AllPieces[Black] |= bit
				pos.Occupied |= bit
			case 'k':
				pos.Pieces[Black][King] |= bit
				pos.AllPieces[Black] |= bit
				pos.Occupied |= bit
			case '1', '2', '3', '4', '5', '6', '7', '8':
				bit <<= uint(char - '1')
			default:
				return Position{}, fmt.Errorf("invalid piece: %c", char)
			}
			bit <<= 1
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

// func (p *Position) Set(rank, file int, piece OldPiece) {
// 	index := (7-rank)*8 + file
// 	bit := BitBoard(1) << uint(index)

// 	if piece == OldEmpty {
// 		p.Occupied &^= bit
// 	} else {
// 		p.Occupied |= bit
// 	}

// 	p.Board[index] = piece
// }

// func (p Position) Get(rank, file int) OldPiece {
// 	index := (7-rank)*8 + file
// 	return p.Board[index]
// }

func (p Position) Fen() string {
	fen := strings.Builder{}

	for bit := BitBoard(1); bit != 0; bit <<= 1 {
		if bit&File_A != 0 && bit > 1 {
			fen.WriteByte('/')
		}

		if p.Pieces[White][Pawn]&bit != 0 {
			fen.WriteByte('P')
		} else if p.Pieces[Black][Pawn]&bit != 0 {
			fen.WriteByte('p')
		} else if p.Pieces[White][Knight]&bit != 0 {
			fen.WriteByte('N')
		} else if p.Pieces[Black][Knight]&bit != 0 {
			fen.WriteByte('n')
		} else if p.Pieces[White][Bishop]&bit != 0 {
			fen.WriteByte('B')
		} else if p.Pieces[Black][Bishop]&bit != 0 {
			fen.WriteByte('b')
		} else if p.Pieces[White][Rook]&bit != 0 {
			fen.WriteByte('R')
		} else if p.Pieces[Black][Rook]&bit != 0 {
			fen.WriteByte('r')
		} else if p.Pieces[White][Queen]&bit != 0 {
			fen.WriteByte('Q')
		} else if p.Pieces[Black][Queen]&bit != 0 {
			fen.WriteByte('q')
		} else if p.Pieces[White][King]&bit != 0 {
			fen.WriteByte('K')
		} else if p.Pieces[Black][King]&bit != 0 {
			fen.WriteByte('k')
		} else {
			empty := 1
			bit <<= 1
			for ; bit != 0; bit, empty = bit<<1, empty+1 {
				if p.Occupied&bit != 0 || bit&File_A != 0 {
					bit >>= 1
					break
				}
			}

			fen.WriteByte('0' + byte(empty))
		}
	}

	fen.WriteByte(' ')

	stm, _ := p.SideToMove()
	if stm == White {
		fen.WriteByte('w')
	} else {
		fen.WriteByte('b')
	}

	fen.WriteByte(' ')

	if p.CanWhiteCastleKingSide() || p.CanWhiteCastleQueenSide() || p.CanBlackCastleKingSide() || p.CanBlackCastleQueenSide() {
		if p.CanWhiteCastleKingSide() {
			fen.WriteByte('K')
		}

		if p.CanWhiteCastleQueenSide() {
			fen.WriteByte('Q')
		}

		if p.CanBlackCastleKingSide() {
			fen.WriteByte('k')
		}

		if p.CanBlackCastleQueenSide() {
			fen.WriteByte('q')
		}
	} else {
		fen.WriteString("-")
	}

	fen.WriteByte(' ')

	if p.IsEnPassant() {
		file := p.State.EnPassantFile - 1

		if stm == White {
			fen.WriteString(fmt.Sprintf("%c6", 'a'+file))
		} else {
			fen.WriteString(fmt.Sprintf("%c3", 'a'+file))
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
	bit := BitBoard(1)
	for rank := 7; rank >= 0; rank-- {
		builder.WriteByte('|')
		for file := 0; file < 8; file++ {
			if p.Pieces[White][Pawn]&bit != 0 {
				builder.WriteString(" P |")
			} else if p.Pieces[Black][Pawn]&bit != 0 {
				builder.WriteString(" p |")
			} else if p.Pieces[White][Knight]&bit != 0 {
				builder.WriteString(" N |")
			} else if p.Pieces[Black][Knight]&bit != 0 {
				builder.WriteString(" n |")
			} else if p.Pieces[White][Bishop]&bit != 0 {
				builder.WriteString(" B |")
			} else if p.Pieces[Black][Bishop]&bit != 0 {
				builder.WriteString(" b |")
			} else if p.Pieces[White][Rook]&bit != 0 {
				builder.WriteString(" R |")
			} else if p.Pieces[Black][Rook]&bit != 0 {
				builder.WriteString(" r |")
			} else if p.Pieces[White][Queen]&bit != 0 {
				builder.WriteString(" Q |")
			} else if p.Pieces[Black][Queen]&bit != 0 {
				builder.WriteString(" q |")
			} else if p.Pieces[White][King]&bit != 0 {
				builder.WriteString(" K |")
			} else if p.Pieces[Black][King]&bit != 0 {
				builder.WriteString(" k |")
			} else {
				builder.WriteString("   |")
			}
			bit <<= 1
		}
		builder.WriteString(fmt.Sprintf(" %d\n", rank+1))
		builder.WriteString("+---+---+---+---+---+---+---+---+\n")
	}
	builder.WriteString("  a   b   c   d   e   f   g   h\n")
	return builder.String()
}

func (p Position) SideToMove() (Color, Color) {
	if p.State.WhiteToMove {
		return White, Black
	}
	return Black, White
}

func (p Position) IsEnPassant() bool {
	return p.State.EnPassantFile != 0
}

func (p Position) EnPassantFile() BitBoard {

	switch p.State.EnPassantFile {
	case 1:
		return File_A
	case 2:
		return File_B
	case 3:
		return File_C
	case 4:
		return File_D
	case 5:
		return File_E
	case 6:
		return File_F
	case 7:
		return File_G
	case 8:
		return File_H
	default:
		return BitBoard(0)
	}
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

	us, them := p.SideToMove()

	capture := p.get(them, m.To)
	if capture != Empty {
		if capture == King {
			fmt.Println(p)
			panic("invalid move")
		}
		p.remove(capture, them, m.To)
	}
	newState.Capture = capture

	switch m.Piece {
	case Bishop, Knight, Queen:
		p.put(m.Piece, us, m.To)
		p.remove(m.Piece, us, m.From)
	case Rook:
		if us == White {
			if m.From == SQ_A1 {
				newState.WhiteQueenSideCastle = false
			}

			if m.From == SQ_H1 {
				newState.WhiteKingSideCastle = false
			}
		} else {
			if m.From == SQ_A8 {
				newState.BlackQueenSideCastle = false
			}

			if m.From == SQ_H8 {
				newState.BlackKingSideCastle = false
			}
		}
		p.put(m.Piece, us, m.To)
		p.remove(m.Piece, us, m.From)

	case King:
		if us == White {
			newState.WhiteKingSideCastle = false
			newState.WhiteQueenSideCastle = false
		} else {
			newState.BlackKingSideCastle = false
			newState.BlackQueenSideCastle = false
		}

		if m.Type == Castle {
			if us == White {
				p.put(m.Piece, us, m.To)
				p.remove(m.Piece, us, m.From)
				if m.To == SQ_G1 {
					p.put(Rook, us, SQ_F1)
					p.remove(Rook, us, SQ_H1)
				} else {
					p.put(Rook, us, SQ_D1)
					p.remove(Rook, us, SQ_A1)
				}
			} else {
				p.put(m.Piece, us, m.To)
				p.remove(m.Piece, us, m.From)
				if m.To == SQ_G8 {
					p.put(Rook, us, SQ_F8)
					p.remove(Rook, us, SQ_H8)
				} else {
					p.put(Rook, us, SQ_D8)
					p.remove(Rook, us, SQ_A8)
				}
			}
		} else {
			p.put(m.Piece, us, m.To)
			p.remove(m.Piece, us, m.From)

		}

	case Pawn:
		piece := m.Piece
		switch m.Type {
		case PromotionToKnight:
			piece = Knight
		case PromotionToBishop:
			piece = Bishop
		case PromotionToRook:
			piece = Rook
		case PromotionToQueen:
			piece = Queen
		case EnPassant:
			if us == White {
				p.remove(Pawn, them, m.To+8)
			} else {
				p.remove(Pawn, them, m.To-8)
			}
		}

		if m.From-m.To == 16 || m.To-m.From == 16 {
			newState.EnPassantFile = uint8(m.To%8) + 1
		}

		p.put(piece, us, m.To)
		p.remove(piece, us, m.From)
	}

	if us == Black {
		newState.WhiteToMove = true
		newState.FullMoves++
	} else {
		newState.WhiteToMove = false
	}

	if us == Black && m.To == SQ_A1 {
		newState.WhiteQueenSideCastle = false
	}

	if us == Black && m.To == SQ_H1 {
		newState.WhiteKingSideCastle = false
	}

	if us == White && m.To == SQ_A8 {
		newState.BlackQueenSideCastle = false
	}

	if us == White && m.To == SQ_H8 {
		newState.BlackKingSideCastle = false
	}

	p.State = newState
}

func (p *Position) Undo() {
	m := p.State.LastMove
	if m == nil {
		return
	}

	capture := p.State.Capture
	us, them := p.SideToMove()

	switch m.Piece {
	case Knight, Bishop, Rook, Queen:
		p.put(m.Piece, them, m.From)
		if capture != Empty {
			p.put(capture, us, m.To)
		} else {
			p.remove(m.Piece, them, m.To)
		}

	case Pawn:
		switch m.Type {
		case PromotionToKnight:
			p.remove(Knight, them, m.To)

		case PromotionToBishop:
			p.remove(Bishop, them, m.To)

		case PromotionToRook:
			p.remove(Rook, them, m.To)

		case PromotionToQueen:
			p.remove(Queen, them, m.To)

		case EnPassant:
			if us == White {
				p.put(Pawn, us, m.To-8)
			} else {
				p.put(Pawn, us, m.To+8)
			}

		}

		if capture != Empty {
			p.put(capture, us, m.To)
		} else {
			p.remove(Pawn, them, m.To)
		}

		p.put(m.Piece, them, m.From)
	case King:
		if m.Type != Castle {
			if them == White {
				if m.To == SQ_G1 {
					p.put(King, them, m.From)
					p.remove(King, them, m.To)
					p.put(Rook, them, SQ_H1)
					p.remove(Rook, them, SQ_F1)
				} else {
					p.put(King, them, m.From)
					p.remove(King, them, m.To)
					p.put(Rook, them, SQ_A1)
					p.remove(Rook, them, SQ_D1)
				}
			} else {
				if m.To == SQ_G8 {
					p.put(King, them, m.From)
					p.remove(King, them, m.To)
					p.put(Rook, them, SQ_H8)
					p.remove(Rook, them, SQ_F8)
				} else {
					p.put(King, them, m.From)
					p.remove(King, them, m.To)
					p.put(Rook, them, SQ_A8)
					p.remove(Rook, them, SQ_D8)
				}
			}
		} else {
			p.put(King, them, m.From)
			if capture != Empty {
				p.put(capture, us, m.To)
			} else {
				p.remove(King, them, m.To)
			}
		}
	}

	defer statePool.Put(p.State)
	p.State = p.State.Previous
}

func (p *Position) put(piece Piece, color Color, sq Square) {
	bit := BitBoard(1) << sq
	p.Occupied |= bit
	p.AllPieces[color] |= bit
	p.Pieces[color][piece] |= bit
}

func (p *Position) get(color Color, sq Square) Piece {
	bit := BitBoard(1) << sq
	if p.Pieces[color][Pawn]&bit != 0 {
		return Pawn
	} else if p.Pieces[color][Knight]&bit != 0 {
		return Knight
	} else if p.Pieces[color][Bishop]&bit != 0 {
		return Bishop
	} else if p.Pieces[color][Rook]&bit != 0 {
		return Rook
	} else if p.Pieces[color][Queen]&bit != 0 {
		return Queen
	} else if p.Pieces[color][King]&bit != 0 {
		return King
	} else {
		return Empty
	}
}

func (p *Position) remove(piece Piece, color Color, sq Square) {
	bit := BitBoard(1) << sq
	p.Occupied &^= bit
	p.AllPieces[color] &^= bit
	p.Pieces[color][piece] &^= bit
}
