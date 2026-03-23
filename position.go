package chester

import (
	"fmt"
	"strconv"
	"strings"
)

const DefaultFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

type Color uint8

const (
	White Color = iota
	Black
)

type Piece uint8

const (
	Pawn = Piece(iota)
	Knight
	Bishop
	Rook
	Queen
	King
	Empty
)

type castlingRights uint8

const (
	whiteKingSideCastle castlingRights = 1 << iota
	whiteQueenSideCastle
	blackKingSideCastle
	blackQueenSideCastle
)

type Position struct {
	pieces    [Piece(6)]Bitboard
	allPieces [Color(2)]Bitboard

	hash uint64

	mailbox [64]Piece

	enPassantTarget Square

	castlingRights   castlingRights
	active, inactive Color

	halfMoves uint8
	fullMoves uint16
}

func ParseFEN(fen string) (Position, error) {
	var pos Position

	parts := strings.Split(strings.TrimSpace(fen), " ")
	if len(parts) < 6 {
		return pos, fmt.Errorf("invalid fen: %s", fen)
	}

	for i := range 64 {
		pos.mailbox[i] = Empty
	}

	bit := Bitboard(1)
	sq := Square(0)

	for _, row := range strings.Split(parts[0], "/") {
		for _, char := range row {
			switch char {
			case 'P':
				pos.mailbox[sq] = Pawn
				pos.pieces[Pawn] |= bit
				pos.allPieces[White] |= bit
			case 'N':
				pos.mailbox[sq] = Knight
				pos.pieces[Knight] |= bit
				pos.allPieces[White] |= bit
			case 'B':
				pos.mailbox[sq] = Bishop
				pos.pieces[Bishop] |= bit
				pos.allPieces[White] |= bit
			case 'R':
				pos.mailbox[sq] = Rook
				pos.pieces[Rook] |= bit
				pos.allPieces[White] |= bit
			case 'Q':
				pos.mailbox[sq] = Queen
				pos.pieces[Queen] |= bit
				pos.allPieces[White] |= bit
			case 'K':
				pos.mailbox[sq] = King
				pos.pieces[King] |= bit
				pos.allPieces[White] |= bit
			case 'p':
				pos.mailbox[sq] = Pawn
				pos.pieces[Pawn] |= bit
				pos.allPieces[Black] |= bit
			case 'n':
				pos.mailbox[sq] = Knight
				pos.pieces[Knight] |= bit
				pos.allPieces[Black] |= bit
			case 'b':
				pos.mailbox[sq] = Bishop
				pos.pieces[Bishop] |= bit
				pos.allPieces[Black] |= bit
			case 'r':
				pos.mailbox[sq] = Rook
				pos.pieces[Rook] |= bit
				pos.allPieces[Black] |= bit
			case 'q':
				pos.mailbox[sq] = Queen
				pos.pieces[Queen] |= bit
				pos.allPieces[Black] |= bit
			case 'k':
				pos.mailbox[sq] = King
				pos.pieces[King] |= bit
				pos.allPieces[Black] |= bit
			case '1', '2', '3', '4', '5', '6', '7', '8':
				bit <<= uint(char - '1')
				sq += Square(char - '1')
			default:
				return Position{}, fmt.Errorf("invalid piece: %c", char)
			}
			bit <<= 1
			sq++
		}
	}

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
			pos.castlingRights |= whiteKingSideCastle
		case 'Q':
			pos.castlingRights |= whiteQueenSideCastle
		case 'k':
			pos.castlingRights |= blackKingSideCastle
		case 'q':
			pos.castlingRights |= blackQueenSideCastle
		}
	}

	pos.enPassantTarget = SquareFromString(parts[3])

	halfMoves, err := strconv.Atoi(parts[4])
	if err != nil {
		return pos, fmt.Errorf("invalid half moves: %s", fen)
	}
	pos.halfMoves = uint8(halfMoves)

	fullMoves, err := strconv.Atoi(parts[5])
	if err != nil {
		return pos, fmt.Errorf("invalid full moves: %s", fen)
	}
	pos.fullMoves = uint16(fullMoves)

	pos.hash = computeHash(&pos)
	return pos, nil
}

func (p Position) FEN() string {
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
	fen.WriteString(p.enPassantTarget.String())
	fen.WriteString(fmt.Sprintf(" %d %d", p.halfMoves, p.fullMoves))

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
	return p.castlingRights&whiteKingSideCastle != 0
}

func (p *Position) CanWhiteCastleQueenSide() bool {
	return p.castlingRights&whiteQueenSideCastle != 0
}

func (p *Position) CanBlackCastleKingSide() bool {
	return p.castlingRights&blackKingSideCastle != 0
}

func (p *Position) CanBlackCastleQueenSide() bool {
	return p.castlingRights&blackQueenSideCastle != 0
}

func (p *Position) Occupied() Bitboard {
	return p.allPieces[White] | p.allPieces[Black]
}

func (p *Position) Do(m Move) {
	from := m.From()
	to := m.To()

	enPassantTarget := p.enPassantTarget
	p.enPassantTarget = SQ_NULL
	p.halfMoves++

	piece := p.mailbox[from]
	captured := p.mailbox[to]
	isCapture := captured != Empty
	diff := int(to) - int(from)

	if enPassantTarget != SQ_NULL {
		var pawnSq Square
		if p.active == White {
			pawnSq = enPassantTarget + 8 // black pawn moved to higher index
		} else {
			pawnSq = enPassantTarget - 8 // white pawn moved to lower index
		}
		if p.adjacentPawns(p.active, pawnSq) {
			p.hash ^= polyglotTable.EnPassant[enPassantTarget.File()]
		}
	}

	if isCapture {
		p.halfMoves = 0
		p.remove(captured, p.inactive, to)
		if captured == Rook {
			p.updateCastlingRights(to)
		}
	}

	// if enPassantTarget != SQ_NULL {
	// 	p.hash ^= polyglotTable.EnPassant[enPassantTarget.File()]
	// }

	switch piece {
	case Pawn:
		p.halfMoves = 0
		switch {
		case m.IsPromotion():
			p.remove(Pawn, p.active, from)
			p.put(m.PromoPiece(), p.active, to)
		case to == enPassantTarget:
			p.move(Pawn, p.active, from, to)
			if p.active == White {
				p.remove(Pawn, p.inactive, to+8)
			} else {
				p.remove(Pawn, p.inactive, to-8)
			}
		case diff == 16 || diff == -16:
			p.move(Pawn, p.active, from, to)
			p.enPassantTarget = (to + from) >> 1
			if p.adjacentPawns(p.inactive, to) {
				p.hash ^= polyglotTable.EnPassant[p.enPassantTarget.File()]
			}
		default:
			p.move(Pawn, p.active, from, to)
		}
	case King:
		p.move(King, p.active, from, to)
		p.updateCastlingRights(from)

		switch diff {
		case 2:
			if p.active == White {
				p.move(Rook, White, SQ_H1, SQ_F1)
			} else {
				p.move(Rook, Black, SQ_H8, SQ_F8)
			}
		case -2:
			if p.active == White {
				p.move(Rook, White, SQ_A1, SQ_D1)
			} else {
				p.move(Rook, Black, SQ_A8, SQ_D8)
			}
		}

	default:
		p.move(piece, p.active, from, to)
		if piece == Rook {
			p.updateCastlingRights(from)
		}
	}

	p.hash ^= polyglotTable.WhiteToMove
	p.fullMoves += uint16(p.active)
	p.active, p.inactive = p.inactive, p.active
}

func (p *Position) adjacentPawns(color Color, sq Square) bool {
	pawns := p.pieces[Pawn] & p.allPieces[color]
	bb := NewBitboardFromSquare(sq)
	return bb&((pawns&File_Not_H)<<1|(pawns&File_Not_A)>>1) != 0
}

func (p *Position) updateCastlingRights(sq Square) {
	fromTo := NewBitboardFromSquare(sq)
	p.hash ^= polyglotTable.Castling[p.castlingRights]
	p.castlingRights &^= whiteKingSideCastle * castlingRights((fromTo&BB_SQ_H1)>>SQ_H1)
	p.castlingRights &^= (whiteKingSideCastle | whiteQueenSideCastle) * castlingRights((fromTo&BB_SQ_E1)>>SQ_E1)
	p.castlingRights &^= whiteQueenSideCastle * castlingRights((fromTo&BB_SQ_A1)>>SQ_A1)

	p.castlingRights &^= blackKingSideCastle * castlingRights((fromTo&BB_SQ_H8)>>SQ_H8)
	p.castlingRights &^= (blackKingSideCastle | blackQueenSideCastle) * castlingRights((fromTo&BB_SQ_E8)>>SQ_E8)
	p.castlingRights &^= blackQueenSideCastle * castlingRights((fromTo&BB_SQ_A8)>>SQ_A8)
	p.hash ^= polyglotTable.Castling[p.castlingRights]
}

func (p *Position) move(piece Piece, color Color, from, to Square) {
	p.mailbox[from] = Empty
	p.mailbox[to] = piece
	fromAndTo := NewBitboardFromSquare(from) | NewBitboardFromSquare(to)
	p.allPieces[color] ^= fromAndTo
	p.pieces[piece] ^= fromAndTo
	p.hash ^= polyglotTable.Pieces[color][piece][from]
	p.hash ^= polyglotTable.Pieces[color][piece][to]
}

func (p *Position) put(piece Piece, color Color, sq Square) {
	p.mailbox[sq] = piece

	bb := NewBitboardFromSquare(sq)
	p.allPieces[color] |= bb
	p.pieces[piece] |= bb
	p.hash ^= polyglotTable.Pieces[color][piece][sq]
}

func (p *Position) remove(piece Piece, color Color, sq Square) {
	p.mailbox[sq] = Empty

	bb := NewBitboardFromSquare(sq)
	p.allPieces[color] &^= bb
	p.pieces[piece] &^= bb
	p.hash ^= polyglotTable.Pieces[color][piece][sq]
}

func (p *Position) Get(sq Square) Piece {
	return p.mailbox[sq]
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
	return p.pieces[Knight] & p.allPieces[Black]
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

func (p *Position) FullMoves() uint16 {
	return p.fullMoves
}

func (p *Position) HalfMoves() uint8 {
	return p.halfMoves
}

func (p *Position) EnPassantTarget() Square {
	return p.enPassantTarget
}

func (p *Position) Hash() uint64 {
	return p.hash
}

func computeHash(p *Position) uint64 {
	var hash uint64

	for color := range Color(2) {
		for piece, bb := range p.pieces {
			bb = bb & p.allPieces[color]
			var sq Square
			for bb != 0 {
				sq, bb = bb.PopLSB()
				hash ^= polyglotTable.Pieces[color][piece][sq]
			}
		}
	}

	hash ^= polyglotTable.Castling[p.castlingRights]

	if p.enPassantTarget != SQ_NULL && p.adjacentPawns(p.active, p.enPassantTarget) {
		file := p.enPassantTarget.File()
		hash ^= polyglotTable.EnPassant[file]
	}

	if p.Active() == White {
		hash ^= polyglotTable.WhiteToMove
	}

	return hash
}
