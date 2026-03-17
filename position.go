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
	pieces [Piece(6)]Bitboard

	allPieces [Color(2)]Bitboard

	EnPassantTarget Bitboard

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

	bit := Bitboard(1)

	for _, row := range strings.Split(parts[0], "/") {
		for _, char := range row {
			switch char {
			case 'P':
				pos.pieces[Pawn] |= bit
				pos.allPieces[White] |= bit
			case 'N':
				pos.pieces[Knight] |= bit
				pos.allPieces[White] |= bit
			case 'B':
				pos.pieces[Bishop] |= bit
				pos.allPieces[White] |= bit
			case 'R':
				pos.pieces[Rook] |= bit
				pos.allPieces[White] |= bit
			case 'Q':
				pos.pieces[Queen] |= bit
				pos.allPieces[White] |= bit
			case 'K':
				pos.pieces[King] |= bit
				pos.allPieces[White] |= bit
			case 'p':
				pos.pieces[Pawn] |= bit
				pos.allPieces[Black] |= bit
			case 'n':
				pos.pieces[Knight] |= bit
				pos.allPieces[Black] |= bit
			case 'b':
				pos.pieces[Bishop] |= bit
				pos.allPieces[Black] |= bit
			case 'r':
				pos.pieces[Rook] |= bit
				pos.allPieces[Black] |= bit
			case 'q':
				pos.pieces[Queen] |= bit
				pos.allPieces[Black] |= bit
			case 'k':
				pos.pieces[King] |= bit
				pos.allPieces[Black] |= bit
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
			pos.EnPassantTarget = NewBitboardFromSquare(sq + 8)
		} else {
			pos.EnPassantTarget = NewBitboardFromSquare(sq - 8)
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

	for bit := Bitboard(1); bit != 0; bit <<= 1 {
		if bit&File_A != 0 && bit > 1 {
			fen.WriteByte('/')
		}

		if p.allPieces[White]&p.pieces[Pawn]&bit != 0 {
			fen.WriteByte('P')
		} else if p.allPieces[Black]&p.pieces[Pawn]&bit != 0 {
			fen.WriteByte('p')
		} else if p.allPieces[White]&p.pieces[Knight]&bit != 0 {
			fen.WriteByte('N')
		} else if p.allPieces[Black]&p.pieces[Knight]&bit != 0 {
			fen.WriteByte('n')
		} else if p.allPieces[White]&p.pieces[Bishop]&bit != 0 {
			fen.WriteByte('B')
		} else if p.allPieces[Black]&p.pieces[Bishop]&bit != 0 {
			fen.WriteByte('b')
		} else if p.allPieces[White]&p.pieces[Rook]&bit != 0 {
			fen.WriteByte('R')
		} else if p.allPieces[Black]&p.pieces[Rook]&bit != 0 {
			fen.WriteByte('r')
		} else if p.allPieces[White]&p.pieces[Queen]&bit != 0 {
			fen.WriteByte('Q')
		} else if p.allPieces[Black]&p.pieces[Queen]&bit != 0 {
			fen.WriteByte('q')
		} else if p.allPieces[White]&p.pieces[King]&bit != 0 {
			fen.WriteByte('K')
		} else if p.allPieces[Black]&p.pieces[King]&bit != 0 {
			fen.WriteByte('k')
		} else {
			empty := 1
			bit <<= 1
			for ; bit != 0; bit, empty = bit<<1, empty+1 {
				if (p.allPieces[White]|p.allPieces[Black])&bit != 0 || bit&File_A != 0 {
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
	bit := Bitboard(1)
	for rank := 7; rank >= 0; rank-- {
		builder.WriteByte('|')
		for file := 0; file < 8; file++ {
			if p.allPieces[White]&p.pieces[Pawn]&bit != 0 {
				builder.WriteString(" P |")
			} else if p.allPieces[Black]&p.pieces[Pawn]&bit != 0 {
				builder.WriteString(" p |")
			} else if p.allPieces[White]&p.pieces[Knight]&bit != 0 {
				builder.WriteString(" N |")
			} else if p.allPieces[Black]&p.pieces[Knight]&bit != 0 {
				builder.WriteString(" n |")
			} else if p.allPieces[White]&p.pieces[Bishop]&bit != 0 {
				builder.WriteString(" B |")
			} else if p.allPieces[Black]&p.pieces[Bishop]&bit != 0 {
				builder.WriteString(" b |")
			} else if p.allPieces[White]&p.pieces[Rook]&bit != 0 {
				builder.WriteString(" R |")
			} else if p.allPieces[Black]&p.pieces[Rook]&bit != 0 {
				builder.WriteString(" r |")
			} else if p.allPieces[White]&p.pieces[Queen]&bit != 0 {
				builder.WriteString(" Q |")
			} else if p.allPieces[Black]&p.pieces[Queen]&bit != 0 {
				builder.WriteString(" q |")
			} else if p.allPieces[White]&p.pieces[King]&bit != 0 {
				builder.WriteString(" K |")
			} else if p.allPieces[Black]&p.pieces[King]&bit != 0 {
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

func (p *Position) Occupied() Bitboard {
	return p.allPieces[White] | p.allPieces[Black]
}

func Do(p *Position, m Move) {

	from := Bitboard(1) << m.From()
	to := Bitboard(1) << m.To()
	isCapture := p.Occupied()&to != 0
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
	} else {
		p.HalfMoves = 0
	}

	if p.active == Black {
		p.FullMoves++
	}

	p.active, p.inactive = p.inactive, p.active
}

func (p *Position) updateCastlingRights(fromTo Bitboard) {
	p.CastlingRights &^= WhiteKingSideCastle * CastlingRights((fromTo&BB_SQ_H1)>>SQ_H1)
	p.CastlingRights &^= (WhiteKingSideCastle | WhiteQueenSideCastle) * CastlingRights((fromTo&BB_SQ_E1)>>SQ_E1)
	p.CastlingRights &^= WhiteQueenSideCastle * CastlingRights((fromTo&BB_SQ_A1)>>SQ_A1)

	p.CastlingRights &^= BlackKingSideCastle * CastlingRights((fromTo&BB_SQ_H8)>>SQ_H8)
	p.CastlingRights &^= (BlackKingSideCastle | BlackQueenSideCastle) * CastlingRights((fromTo&BB_SQ_E8)>>SQ_E8)
	p.CastlingRights &^= BlackQueenSideCastle * CastlingRights((fromTo&BB_SQ_A8)>>SQ_A8)
}

//func (p *Position) Undo () {}

func (p *Position) move(piece Piece, color Color, from, to Bitboard) {
	fromAndTo := from | to
	p.allPieces[color] ^= fromAndTo
	p.pieces[piece] ^= fromAndTo
}

func (p *Position) put(piece Piece, color Color, sq Bitboard) {
	p.allPieces[color] |= sq
	p.pieces[piece] |= sq
}

func (p *Position) remove(piece Piece, color Color, sq Bitboard) {
	p.allPieces[color] &^= sq
	p.pieces[piece] &^= sq
}

func (p *Position) removeAll(color Color, sq Bitboard) {
	p.allPieces[color] &^= sq
	p.pieces[Pawn] &^= sq
	p.pieces[Knight] &^= sq
	p.pieces[Bishop] &^= sq
	p.pieces[Rook] &^= sq
	p.pieces[Queen] &^= sq
}

func (p *Position) Get(sq Square) Piece {
	bit := Bitboard(1) << sq

	if p.allPieces[p.active]&p.pieces[Pawn]&bit != 0 {
		return Pawn
	} else if p.allPieces[p.active]&p.pieces[Knight]&bit != 0 {
		return Knight
	} else if p.allPieces[p.active]&p.pieces[Bishop]&bit != 0 {
		return Bishop
	} else if p.allPieces[p.active]&p.pieces[Rook]&bit != 0 {
		return Rook
	} else if p.allPieces[p.active]&p.pieces[Queen]&bit != 0 {
		return Queen
	} else if p.allPieces[p.active]&p.pieces[King]&bit != 0 {
		return King
	}
	return Empty
}

func (p *Position) Pawns() Bitboard {
	return p.pieces[Pawn] & p.allPieces[p.active]
}

func (p *Position) WhitePawns() Bitboard {
	return p.pieces[Pawn] & p.allPieces[White]
}

func (p *Position) BlackPawns() Bitboard {
	return p.pieces[Pawn] & p.allPieces[Black]
}

func (p *Position) EnemyPawns() Bitboard {
	return p.pieces[Pawn] & p.allPieces[p.inactive]
}

func (p *Position) Knights() Bitboard {
	return p.pieces[Knight] & p.allPieces[p.active]
}

func (p *Position) WhiteKnights() Bitboard {
	return p.pieces[Knight] & p.allPieces[White]
}

func (p *Position) BlackKnights() Bitboard {
	return p.pieces[Pawn] & p.allPieces[Black]
}

func (p *Position) EnemyKnights() Bitboard {
	return p.pieces[Knight] & p.allPieces[p.inactive]
}

func (p *Position) Bishops() Bitboard {
	return p.pieces[Bishop] & p.allPieces[p.active]
}

func (p *Position) WhiteBishops() Bitboard {
	return p.pieces[Bishop] & p.allPieces[White]
}

func (p *Position) BlackBishops() Bitboard {
	return p.pieces[Bishop] & p.allPieces[Black]
}

func (p *Position) EnemyBishops() Bitboard {
	return p.pieces[Bishop] & p.allPieces[p.inactive]
}

func (p *Position) Rooks() Bitboard {
	return p.pieces[Rook] & p.allPieces[p.active]
}

func (p *Position) WhiteRooks() Bitboard {
	return p.pieces[Rook] & p.allPieces[White]
}

func (p *Position) BlackRooks() Bitboard {
	return p.pieces[Rook] & p.allPieces[Black]
}

func (p *Position) EnemyRooks() Bitboard {
	return p.pieces[Rook] & p.allPieces[p.inactive]
}

func (p *Position) Queens() Bitboard {
	return p.pieces[Queen] & p.allPieces[p.active]
}

func (p *Position) WhiteQueens() Bitboard {
	return p.pieces[Queen] & p.allPieces[White]
}

func (p *Position) BlackQueens() Bitboard {
	return p.pieces[Queen] & p.allPieces[Black]
}

func (p *Position) EnemyQueens() Bitboard {
	return p.pieces[Queen] & p.allPieces[p.inactive]
}

func (p *Position) King() Bitboard {
	return p.pieces[King] & p.allPieces[p.active]
}

func (p *Position) WhiteKing() Bitboard {
	return p.pieces[King] & p.allPieces[White]
}

func (p *Position) BlackKing() Bitboard {
	return p.pieces[King] & p.allPieces[Black]
}

func (p *Position) EnemyKing() Bitboard {
	return p.pieces[King] & p.allPieces[p.inactive]
}

func (p *Position) EnemyQueensOrBishops() Bitboard {
	return (p.pieces[Queen] | p.pieces[Bishop]) & p.allPieces[p.inactive]
}

func (p *Position) EnemyQueensOrRooks() Bitboard {
	return (p.pieces[Queen] | p.pieces[Rook]) & p.allPieces[p.inactive]
}

func (p *Position) Enemies() Bitboard {
	return p.allPieces[p.inactive]
}

func (p *Position) EnemiesOrEmpty() Bitboard {
	return ^p.allPieces[p.active]
}
