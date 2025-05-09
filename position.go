package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Color uint8

const (
	White = iota
	Black
)

type Piece uint8

const (
	Knight = 0
	Bishop = 2
	Rook   = 4
	Queen  = 6
	Pawn   = 8
	King   = 10
	Empty  = 12
)

const (
	AllPieces = 12
	Occupied  = 14
	Metadata  = 15
)

// Position
// Index 0: White Knights
// Index 1: Black Knights
// Index 2: White Bishops
// Index 3: Black Bishops
// Index 4: White Rook
// Index 5: Black Rook
// Index 6: White Queen
// Index 7: Black Queen
// Index 8: White Pawn
// Index 9: Black Pawn
// Index 10: White King
// Index 11: Black King
// Index 12: White Pieces
// Index 13: Black Pieces
// Index 14: Occupied
// Index 15: Metadata
// |- 0  Active Player (White / Black)
// |- 1-4 Castling Rights
// |- 11-23 En Passant Target BitBoard
// |- 31-47 Full Moves
// |- 48-55 Half Moves

type Position [16]BitBoard

const (
	ActivePlayerMask     = 0x01
	WhiteKingSideCastle  = 0x02
	WhiteQueenSideCastle = 0x04
	BlackKingSideCastle  = 0x08
	BlackQueenSideCastle = 0x10
	EnPassantTargetMask  = 0x000000ffff000000
	HalfMoveMask         = 0x0000ff0000000000
	FullMoveMask         = 0xffff000000000000

	HalfMoveShift = 40
	FullMoveShift = 48
)

// type Position struct {
// 	Pieces    [Color(2)][Piece(6)]BitBoard
// 	AllPieces [Color(2)]BitBoard
// 	Occupied  BitBoard

// 	EnPassantTarget BitBoard

// 	FullMoves      uint16
// 	HalfMoves      uint8
// 	CastlingRights CastlingRights
// 	active         Color
// }

func Parse(fen string) (Position, error) {
	var p Position

	parts := strings.Split(strings.TrimSpace(fen), " ")
	if len(parts) < 6 {
		return p, fmt.Errorf("invalid fen: %s", fen)
	}

	bit := BitBoard(1)

	for _, row := range strings.Split(parts[0], "/") {
		for _, char := range row {
			switch char {
			case 'P':
				p[Pawn+White] |= bit
				p[AllPieces+White] |= bit
				p[Occupied] |= bit
			case 'N':
				p[Knight+White] |= bit
				p[AllPieces+White] |= bit
				p[Occupied] |= bit
			case 'B':
				p[Bishop+White] |= bit
				p[AllPieces+White] |= bit
				p[Occupied] |= bit
			case 'R':
				p[Rook+White] |= bit
				p[AllPieces+White] |= bit
				p[Occupied] |= bit
			case 'Q':
				p[Queen+White] |= bit
				p[AllPieces+White] |= bit
				p[Occupied] |= bit
			case 'K':
				p[King+White] |= bit
				p[AllPieces+White] |= bit
				p[Occupied] |= bit
			case 'p':
				p[Pawn+Black] |= bit
				p[AllPieces+Black] |= bit
				p[Occupied] |= bit
			case 'n':
				p[Knight+Black] |= bit
				p[AllPieces+Black] |= bit
				p[Occupied] |= bit
			case 'b':
				p[Bishop+Black] |= bit
				p[AllPieces+Black] |= bit
				p[Occupied] |= bit
			case 'r':
				p[Rook+Black] |= bit
				p[AllPieces+Black] |= bit
				p[Occupied] |= bit
			case 'q':
				p[Queen+Black] |= bit
				p[AllPieces+Black] |= bit
				p[Occupied] |= bit
			case 'k':
				p[King+Black] |= bit
				p[AllPieces+Black] |= bit
				p[Occupied] |= bit
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

	if parts[1] == "b" || parts[1] == "B" {
		p[Metadata] |= ActivePlayerMask
	}

	for _, c := range parts[2] {
		switch c {
		case 'K':
			p[Metadata] |= WhiteKingSideCastle
		case 'Q':
			p[Metadata] |= WhiteQueenSideCastle
		case 'k':
			p[Metadata] |= BlackKingSideCastle
		case 'q':
			p[Metadata] |= BlackQueenSideCastle
		}
	}

	if sq := SquareFromString(parts[3]); sq != SQ_NULL {
		if p.Active() == White {
			p[Metadata] |= NewBitBoardFromSquare(sq + 8)
		} else {
			p[Metadata] |= NewBitBoardFromSquare(sq - 8)
		}
	}

	halfMoves, err := strconv.Atoi(parts[4])
	if err != nil {
		return p, fmt.Errorf("invalid half moves: %s", fen)
	}
	p[Metadata] |= BitBoard(halfMoves) << HalfMoveShift

	fullMoves, err := strconv.Atoi(parts[5])
	if err != nil {
		return p, fmt.Errorf("invalid full moves: %s", fen)
	}
	p[Metadata] |= BitBoard(fullMoves) << FullMoveShift

	return p, nil
}

func (p Position) Fen() string {
	fen := strings.Builder{}

	for bit := BitBoard(1); bit != 0; bit <<= 1 {
		if bit&File_A != 0 && bit > 1 {
			fen.WriteByte('/')
		}

		if p[Pawn+White]&bit != 0 {
			fen.WriteByte('P')
		} else if p[Pawn+Black]&bit != 0 {
			fen.WriteByte('p')
		} else if p[Knight+White]&bit != 0 {
			fen.WriteByte('N')
		} else if p[Knight+Black]&bit != 0 {
			fen.WriteByte('n')
		} else if p[Bishop+White]&bit != 0 {
			fen.WriteByte('B')
		} else if p[Bishop+Black]&bit != 0 {
			fen.WriteByte('b')
		} else if p[Rook+White]&bit != 0 {
			fen.WriteByte('R')
		} else if p[Rook+Black]&bit != 0 {
			fen.WriteByte('r')
		} else if p[Queen+White]&bit != 0 {
			fen.WriteByte('Q')
		} else if p[Queen+Black]&bit != 0 {
			fen.WriteByte('q')
		} else if p[King+White]&bit != 0 {
			fen.WriteByte('K')
		} else if p[King+Black]&bit != 0 {
			fen.WriteByte('k')
		} else {
			empty := 1
			bit <<= 1
			for ; bit != 0; bit, empty = bit<<1, empty+1 {
				if p[Occupied]&bit != 0 || bit&File_A != 0 {
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
		fen.WriteString((p.EnPassantTarget() >> 8).Square().String())
	} else {
		fen.WriteString((p.EnPassantTarget() << 8).Square().String())
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
			if p[Pawn+White]&bit != 0 {
				builder.WriteString(" P |")
			} else if p[Pawn+Black]&bit != 0 {
				builder.WriteString(" p |")
			} else if p[Knight+White]&bit != 0 {
				builder.WriteString(" N |")
			} else if p[Knight+Black]&bit != 0 {
				builder.WriteString(" n |")
			} else if p[Bishop+White]&bit != 0 {
				builder.WriteString(" B |")
			} else if p[Bishop+Black]&bit != 0 {
				builder.WriteString(" b |")
			} else if p[Rook+White]&bit != 0 {
				builder.WriteString(" R |")
			} else if p[Rook+Black]&bit != 0 {
				builder.WriteString(" r |")
			} else if p[Queen+White]&bit != 0 {
				builder.WriteString(" Q |")
			} else if p[Queen+Black]&bit != 0 {
				builder.WriteString(" q |")
			} else if p[King+White]&bit != 0 {
				builder.WriteString(" K |")
			} else if p[King+Black]&bit != 0 {
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
	return Color(p[Metadata] & ActivePlayerMask)
}

func (p *Position) Inactive() Color {
	return Color(^p[Metadata] & ActivePlayerMask)
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
	return p[Metadata]&WhiteKingSideCastle != 0
}

func (p *Position) CanWhiteCastleQueenSide() bool {
	return p[Metadata]&WhiteQueenSideCastle != 0
}

func (p *Position) CanBlackCastleKingSide() bool {
	return p[Metadata]&BlackKingSideCastle != 0
}

func (p *Position) CanBlackCastleQueenSide() bool {
	return p[Metadata]&BlackQueenSideCastle != 0
}

func Do(p *Position, m Move) {

	us, them := p.Active(), p.Inactive()
	from := BitBoard(1) << m.From()
	to := BitBoard(1) << m.To()
	isCapture := p[AllPieces+them]&to != 0
	piece := m.Piece()
	enPassantTarget := p.EnPassantTarget()
	p[Metadata] &^= EnPassantTargetMask

	switch m.Type() {
	case Promotion:
		if isCapture {
			p.removeAll(them, to)
			p.updateCastlingRights(from | to)
		}
		p.remove(Pawn, us, from)
		p.put(m.PromoPiece(), us, to)

	case EnPassant:
		p.remove(Pawn, them, enPassantTarget)
		p.move(Pawn, us, from, to)
	case CastleKingSide:
		p.move(King, us, from, to)
		p.move(Rook, us, from>>3, from>>1)
		p.updateCastlingRights(from)

	case CastleQueenSide:
		p.move(King, us, from, to)
		p.move(Rook, us, from<<4, from<<2)
		p.updateCastlingRights(from)
	case DoublePush:
		p.move(Pawn, us, from, to)
		p[Metadata] |= to // set en passant target
	default:
		if isCapture {
			p.removeAll(them, to)
		}
		p.move(piece, us, from, to)
		p.updateCastlingRights(from | to)
	}

	if !isCapture && piece != Pawn {
		p[Metadata] += 0x10000000000
	}

	if us == Black {
		p[Metadata] += 0x1000000000000
	}

	p[Metadata] ^= ActivePlayerMask
}

func (p *Position) updateCastlingRights(fromTo BitBoard) {
	p[Metadata] &^= WhiteKingSideCastle * ((fromTo & BB_SQ_H1) >> SQ_H1)
	p[Metadata] &^= (WhiteKingSideCastle | WhiteQueenSideCastle) * ((fromTo & BB_SQ_E1) >> SQ_E1)
	p[Metadata] &^= WhiteQueenSideCastle * (fromTo & BB_SQ_A1) >> SQ_A1

	p[Metadata] &^= BlackKingSideCastle * ((fromTo & BB_SQ_H8) >> SQ_H8)
	p[Metadata] &^= (BlackKingSideCastle | BlackQueenSideCastle) * ((fromTo & BB_SQ_E8) >> SQ_E8)
	p[Metadata] &^= BlackQueenSideCastle * ((fromTo & BB_SQ_A8) >> SQ_A8)
}

//func (p *Position) Undo () {}

func (p *Position) move(piece Piece, color Color, from, to BitBoard) {
	fromAndTo := from | to
	p[Occupied] ^= fromAndTo
	p[AllPieces+color] ^= fromAndTo
	p[int8(piece)+int8(color)] ^= fromAndTo
}

func (p *Position) put(piece Piece, color Color, sq BitBoard) {
	p[Occupied] |= sq
	p[AllPieces+color] |= sq
	p[int8(piece)+int8(color)] |= sq
}

func (p *Position) remove(piece Piece, color Color, sq BitBoard) {
	p[Occupied] &^= sq
	p[AllPieces+color] &^= sq
	p[int8(piece)+int8(color)] &^= sq
}

func (p *Position) removeAll(color Color, sq BitBoard) {
	p[Occupied] &^= sq
	p[AllPieces+color] &^= sq
	p[Pawn+color] &^= sq
	p[Knight+color] &^= sq
	p[Bishop+color] &^= sq
	p[Rook+color] &^= sq
	p[Queen+color] &^= sq
}

func (p *Position) EnPassantTarget() BitBoard {
	return p[Metadata] & EnPassantTargetMask
}

func (p *Position) HalfMoves() uint8 {
	return uint8((p[Metadata] & HalfMoveMask) >> HalfMoveShift)
}

func (p *Position) FullMoves() uint8 {
	return uint8((p[Metadata] & FullMoveMask) >> FullMoveShift)
}

func (p *Position) Pawn(color Color) BitBoard {
	return p[Pawn+color]
}
func (p *Position) Bishop(color Color) BitBoard {
	return p[Bishop+color]
}
func (p *Position) Knight(color Color) BitBoard {
	return p[Knight+color]
}

func (p *Position) Rook(color Color) BitBoard {
	return p[Rook+color]
}

func (p *Position) Queen(color Color) BitBoard {
	return p[Queen+color]
}

func (p *Position) King(color Color) BitBoard {
	return p[King+color]
}

func (p *Position) AllPieces(color Color) BitBoard {
	return p[AllPieces+color]
}

func (p *Position) Occupied() BitBoard {
	return p[Occupied]
}

func (p Position) Get(sq Square) Piece {
	bit := BitBoard(1) << sq
	color := int(p[Metadata] & ActivePlayerMask)

	if p[Pawn+color]&bit != 0 {
		return Pawn
	} else if p[Knight+color]&bit != 0 {
		return Knight
	} else if p[Bishop+color]&bit != 0 {
		return Bishop
	} else if p[Rook+color]&bit != 0 {
		return Rook
	} else if p[Queen+color]&bit != 0 {
		return Queen
	} else if p[King+color]&bit != 0 {
		return King
	}
	return Empty
}
