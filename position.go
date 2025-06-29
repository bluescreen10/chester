package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Color uint8

const (
	White Color = iota
	Black
)

type Piece uint8

const (
	Knight Piece = iota
	Bishop
	Rook
	Queen
	Pawn
	King
	Empty
)

type CastlingRights uint8

const (
	WhiteKingSideCastle CastlingRights = 1 << iota
	WhiteQueenSideCastle
	BlackKingSideCastle
	BlackQueenSideCastle
)

type Position struct {
	Pieces    [Color(2)][Piece(6)]BitBoard
	AllPieces [Color(2)]BitBoard
	Occupied  BitBoard

	EnPassantTarget BitBoard

	FullMoves        uint16
	HalfMoves        uint8
	CastlingRights   CastlingRights
	active, inactive Color
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

	// pos.State = statePool.Get().(*State)
	// pos.State.Reset()

	if parts[1] == "w" || parts[1] == "W" {
		pos.active = White
		pos.inactive = Black
	} else {
		pos.active = Black
		pos.inactive = White
	}

	for _, c := range parts[2] {
		switch c {
		case 'K':
			pos.CastlingRights |= WhiteKingSideCastle
		case 'Q':
			pos.CastlingRights |= WhiteQueenSideCastle
		case 'k':
			pos.CastlingRights |= BlackKingSideCastle
		case 'q':
			pos.CastlingRights |= BlackQueenSideCastle
		}
	}

	if sq := SquareFromString(parts[3]); sq != SQ_NULL {
		if pos.Active() == White {
			pos.EnPassantTarget = NewBitBoardFromSquare(sq + 8)
		} else {
			pos.EnPassantTarget = NewBitBoardFromSquare(sq - 8)
		}
	}

	halfMoves, err := strconv.Atoi(parts[4])
	if err != nil {
		return pos, fmt.Errorf("invalid half moves: %s", fen)
	}
	pos.HalfMoves = uint8(halfMoves)

	fullMoves, err := strconv.Atoi(parts[5])
	if err != nil {
		return pos, fmt.Errorf("invalid full moves: %s", fen)
	}
	pos.FullMoves = uint16(fullMoves)

	return pos, nil
}

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

	if p.Active() == White {
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
	if p.Active() == White {
		fen.WriteString((p.EnPassantTarget >> 8).Square().String())
	} else {
		fen.WriteString((p.EnPassantTarget << 8).Square().String())
	}
	fen.WriteString(fmt.Sprintf(" %d %d", p.HalfMoves, p.FullMoves))

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

func (p *Position) Active() Color {
	return p.active
}

func (p *Position) Inactive() Color {
	return p.inactive
}

// func (p Position) EnPassantFile() BitBoard {
// 	_, file := p.EnPassantSquare.RankAndFile()
// 	switch file {
// 	case 0:
// 		return File_A
// 	case 1:
// 		return File_B
// 	case 2:
// 		return File_C
// 	case 3:
// 		return File_D
// 	case 4:
// 		return File_E
// 	case 5:
// 		return File_F
// 	case 6:
// 		return File_G
// 	case 7:
// 		return File_H
// 	default:
// 		return BitBoard(0)
// 	}
// }

func (p *Position) CanWhiteCastleKingSide() bool {
	return p.CastlingRights&WhiteKingSideCastle != 0
}

func (p *Position) CanWhiteCastleQueenSide() bool {
	return p.CastlingRights&WhiteQueenSideCastle != 0
}

func (p *Position) CanBlackCastleKingSide() bool {
	return p.CastlingRights&BlackKingSideCastle != 0
}

func (p *Position) CanBlackCastleQueenSide() bool {
	return p.CastlingRights&BlackQueenSideCastle != 0
}

func Do(p *Position, m Move) {

	from := BitBoard(1) << m.From()
	to := BitBoard(1) << m.To()
	isCapture := p.Occupied&to != 0
	piece := m.Piece()
	enPassantTarget := p.EnPassantTarget
	p.EnPassantTarget = 0

	switch m.Type() {
	case Promotion:
		if isCapture {
			p.removeAll(p.inactive, to)
			p.updateCastlingRights(from | to)
		}
		p.remove(Pawn, p.active, from)
		p.put(m.PromoPiece(), p.active, to)

	case EnPassant:
		p.remove(Pawn, p.inactive, enPassantTarget)
		p.move(Pawn, p.active, from, to)
	case CastleKingSide:
		p.move(King, p.active, from, to)
		p.move(Rook, p.active, from>>3, from>>1)
		p.updateCastlingRights(from)

	case CastleQueenSide:
		p.move(King, p.active, from, to)
		p.move(Rook, p.active, from<<4, from<<2)
		p.updateCastlingRights(from)
	case DoublePush:
		p.move(Pawn, p.active, from, to)
		p.EnPassantTarget = to
	default:
		if isCapture {
			p.removeAll(p.inactive, to)
		}
		p.move(piece, p.active, from, to)
		p.updateCastlingRights(from | to)
	}

	if !isCapture && piece != Pawn {
		p.HalfMoves++
	}

	if p.active == Black {
		p.FullMoves++
	}

	p.active, p.inactive = p.inactive, p.active
}

func (p *Position) updateCastlingRights(fromTo BitBoard) {
	p.CastlingRights &^= WhiteKingSideCastle * CastlingRights((fromTo&BB_SQ_H1)>>SQ_H1)
	p.CastlingRights &^= (WhiteKingSideCastle | WhiteQueenSideCastle) * CastlingRights((fromTo&BB_SQ_E1)>>SQ_E1)
	p.CastlingRights &^= WhiteQueenSideCastle * CastlingRights((fromTo&BB_SQ_A1)>>SQ_A1)

	p.CastlingRights &^= BlackKingSideCastle * CastlingRights((fromTo&BB_SQ_H8)>>SQ_H8)
	p.CastlingRights &^= (BlackKingSideCastle | BlackQueenSideCastle) * CastlingRights((fromTo&BB_SQ_E8)>>SQ_E8)
	p.CastlingRights &^= BlackQueenSideCastle * CastlingRights((fromTo&BB_SQ_A8)>>SQ_A8)
}

//func (p *Position) Undo () {}

func (p *Position) move(piece Piece, color Color, from, to BitBoard) {
	fromAndTo := from | to
	p.Occupied ^= fromAndTo
	p.AllPieces[color] ^= fromAndTo
	p.Pieces[color][piece] ^= fromAndTo
}

func (p *Position) put(piece Piece, color Color, sq BitBoard) {
	p.Occupied |= sq
	p.AllPieces[color] |= sq
	p.Pieces[color][piece] |= sq
}

func (p *Position) remove(piece Piece, color Color, sq BitBoard) {
	p.Occupied &^= sq
	p.AllPieces[color] &^= sq
	p.Pieces[color][piece] &^= sq
}

func (p *Position) removeAll(color Color, sq BitBoard) {
	p.Occupied &^= sq
	p.AllPieces[color] &^= sq
	p.Pieces[color][Pawn] &^= sq
	p.Pieces[color][Knight] &^= sq
	p.Pieces[color][Bishop] &^= sq
	p.Pieces[color][Rook] &^= sq
	p.Pieces[color][Queen] &^= sq
}

func (p Position) Get(sq Square) Piece {
	bit := BitBoard(1) << sq

	if p.Pieces[p.active][Pawn]&bit != 0 {
		return Pawn
	} else if p.Pieces[p.active][Knight]&bit != 0 {
		return Knight
	} else if p.Pieces[p.active][Bishop]&bit != 0 {
		return Bishop
	} else if p.Pieces[p.active][Rook]&bit != 0 {
		return Rook
	} else if p.Pieces[p.active][Queen]&bit != 0 {
		return Queen
	} else if p.Pieces[p.active][King]&bit != 0 {
		return King
	}
	return Empty
}
