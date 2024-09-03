package main

import (
	"fmt"
	"io"
	"math/bits"
	"sync"
)

const (
	maxMoves = 218
)

type MoveType uint8

const (
	Default MoveType = iota
	Castle
	EnPassant
	PromotionToQueen
	PromotionToRook
	PromotionToBishop
	PromotionToKnight
)

// Eventualy we want to use uint16
// First 3 bytes move type
// Next 6 bytes from square
// Last 6 bytes to square
type Move struct {
	Piece    Piece
	Type     MoveType
	From, To Square
}

type direction uint8

const (
	North direction = iota
	East
	South
	West

	NorthEast
	SouthEast
	SouthWest
	NorthWest
)

var rays [direction(8)][Square(64)]BitBoard

var knightMoves [64]BitBoard
var kingMoves [64]BitBoard

func init() {
	// pre-compute rays
	for i := 0; i < 64; i++ {
		r, f := 7-i/8, i%8

		for j := 1; j < 8; j++ {
			if r-j >= 0 && f-j >= 0 {
				rays[SouthWest][i] |= 1 << uint8((7-r+j)*8+f-j)
			}
			if r-j >= 0 && f+j < 8 {
				rays[SouthEast][i] |= 1 << uint8((7-r+j)*8+f+j)
			}
			if r+j < 8 && f-j >= 0 {
				rays[NorthWest][i] |= 1 << uint8((7-r-j)*8+f-j)
			}
			if r+j < 8 && f+j < 8 {
				rays[NorthEast][i] |= 1 << uint8((7-r-j)*8+f+j)
			}
		}
	}

	for i := 0; i < 64; i++ {
		r, f := 7-i/8, i%8

		for j := 1; j < 8; j++ {
			if r-j >= 0 {
				rays[South][i] |= 1 << uint8((7-r+j)*8+f)
			}
			if r+j < 8 {
				rays[North][i] |= 1 << uint8((7-r-j)*8+f)
			}
			if f-j >= 0 {
				rays[West][i] |= 1 << uint8((7-r)*8+f-j)
			}
			if f+j < 8 {
				rays[East][i] |= 1 << uint8((7-r)*8+f+j)
			}
		}
	}

	// pre-compute knight moves
	for i := int8(0); i < 64; i++ {
		rank, file := 7-i/8, i%8
		moves := [][]int8{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}}

		for _, offsets := range moves {
			targetRank, targetFile := rank+offsets[0], file+offsets[1]
			if targetRank < 0 || targetRank > 7 || targetFile < 0 || targetFile > 7 {
				continue
			}

			knightMoves[i] |= 1 << uint8((7-targetRank)*8+targetFile)
		}
	}

	// pre-compute king moves
	for i := int8(0); i < 64; i++ {
		rank, file := 7-i/8, i%8
		moves := [][]int8{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}

		for _, offsets := range moves {
			targetRank, targetFile := rank+offsets[0], file+offsets[1]
			if targetRank < 0 || targetRank > 7 || targetFile < 0 || targetFile > 7 {
				continue
			}

			kingMoves[i] |= 1 << uint8((7-targetRank)*8+targetFile)
		}
	}
}

func (m Move) String() string {
	fromRank, fromFile := m.From.RankAndFile()
	toRank, toFile := m.To.RankAndFile()

	suffix := ""

	switch m.Type {
	case PromotionToBishop:
		suffix = "b"
	case PromotionToKnight:
		suffix = "n"
	case PromotionToRook:
		suffix = "r"
	case PromotionToQueen:
		suffix = "q"
	default:
	}

	return fmt.Sprintf("%c%d%c%d%s", 'a'+fromFile, fromRank+1, 'a'+toFile, toRank+1, suffix)
}

var movesPool = sync.Pool{
	New: func() any {
		s := make([]Move, 0, maxMoves)
		return &s
	},
}

func Perft(p Position, depth int, print bool, output io.Writer) int {
	if depth == 0 {
		return 1
	}
	var nodes int
	moves := movesPool.Get().(*[]Move)
	*moves = (*moves)[:0]
	LegalMoves(moves, &p)
	for _, m := range *moves {
		p.Do(m)
		newNodes := Perft(p, depth-1, false, output)
		if print {
			fmt.Fprintf(output, "%s: %d\n", m, newNodes)
		}
		nodes += newNodes
		p.Undo()
	}
	movesPool.Put(moves)
	return nodes
}

func LegalMoves(moves *[]Move, pos *Position) {
	*moves = (*moves)[:0]

	us, _ := pos.SideToMove()

	if us == White {
		genWhitePawnMoves(moves, pos)
		genWhiteKnightMoves(moves, pos)
		genWhiteBishopMoves(moves, pos)
		genWhiteRookMoves(moves, pos)
		genWhiteQueenMoves(moves, pos)
		genWhiteKingMoves(moves, pos)
	} else {
		genBlackPawnMoves(moves, pos)
		genBlackKnightMoves(moves, pos)
		genBlackBishopMoves(moves, pos)
		genBlackRookMoves(moves, pos)
		genBlackQueenMoves(moves, pos)
		genBlackKingMoves(moves, pos)
	}

	for i := 0; i < len(*moves); i++ {
		move := (*moves)[i]
		//fmt.Println(move)
		pos.Do(move)

		king := pos.Pieces[us][King]

		var attacks BitBoard
		if us == White {
			attacks = blackAttacks(pos)
		} else {
			attacks = whiteAttacks(pos)
		}

		// if in check remove the move from the valid moves
		if king&attacks != 0 {
			*moves = append((*moves)[:i], (*moves)[i+1:]...)
			i--
		}
		pos.Undo()
	}
}

func genWhitePawnAttacks(p *Position) BitBoard {
	pawns := p.Pieces[White][Pawn]

	left := (pawns &^ File_A) >> 7
	right := (pawns &^ File_H) >> 9

	return left | right
}

func genBlackPawnAttacks(p *Position) BitBoard {
	pawns := p.Pieces[Black][Pawn]

	left := (pawns &^ File_A) << 7
	right := (pawns &^ File_H) << 9

	return left | right
}

func genWhiteKnightAttacks(p *Position) BitBoard {
	var attacks BitBoard

	knights := p.Pieces[White][Knight]

	for knights != 0 {
		var sq Square
		sq, knights = knights.PopLSB()
		attacks |= knightMoves[sq]
	}

	return attacks
}

func genBlackKnightAttacks(p *Position) BitBoard {
	var attacks BitBoard

	knights := p.Pieces[Black][Knight]

	for knights != 0 {
		var sq Square
		sq, knights = knights.PopLSB()
		attacks |= knightMoves[sq]
	}

	return attacks
}

func genWhiteBishopAttacks(p *Position) BitBoard {
	var attacks BitBoard

	bishops := p.Pieces[White][Bishop]
	occupied := p.Occupied

	for bishops != 0 {
		var sq Square
		sq, bishops = bishops.PopLSB()
		attacks |= genBishopAttacks(sq, occupied)
	}

	return attacks
}

func genBlackBishopAttacks(p *Position) BitBoard {
	var attacks BitBoard

	bishops := p.Pieces[Black][Bishop]
	occupied := p.Occupied

	for bishops != 0 {
		var sq Square
		sq, bishops = bishops.PopLSB()
		attacks |= genBishopAttacks(sq, occupied)
	}

	return attacks
}

func genBishopAttacks(sq Square, occupied BitBoard) BitBoard {
	attacks := rays[NorthWest][sq]
	if intersect := rays[NorthWest][sq] & occupied; intersect != 0 {
		//fmt.Println(intersect)
		index := bits.LeadingZeros64(uint64(intersect))
		attacks &^= rays[NorthWest][63-index]
	}
	// fmt.Println(attacks)
	// fmt.Println(rays[NorthEast][sq])
	attacks |= rays[NorthEast][sq]
	if intersect := rays[NorthEast][sq] & occupied; intersect != 0 {
		// fmt.Println(intersect)
		index := bits.LeadingZeros64(uint64(intersect))
		attacks &^= rays[NorthEast][63-index]
	}
	// fmt.Println(attacks)
	// fmt.Println(rays[SouthWest][sq])
	attacks |= rays[SouthWest][sq]
	if intersect := rays[SouthWest][sq] & occupied; intersect != 0 {
		// fmt.Println(intersect)
		index := bits.TrailingZeros64(uint64(intersect))
		attacks &^= rays[SouthWest][index]
	}
	// fmt.Println(attacks)
	// fmt.Println(rays[SouthEast][sq])
	attacks |= rays[SouthEast][sq]
	if intersect := rays[SouthEast][sq] & occupied; intersect != 0 {
		// fmt.Println(intersect)
		index := bits.TrailingZeros64(uint64(intersect))
		attacks &^= rays[SouthEast][index]
	}
	// fmt.Println(attacks)
	return attacks
}

func genWhiteRookAttacks(p *Position) BitBoard {
	var attacks BitBoard

	rooks := p.Pieces[White][Rook]
	occupied := p.Occupied

	for rooks != 0 {
		var sq Square
		sq, rooks = rooks.PopLSB()
		attacks |= genRookAttacks(sq, occupied)
	}

	return attacks
}

func genBlackRookAttacks(p *Position) BitBoard {
	var attacks BitBoard

	rooks := p.Pieces[Black][Rook]
	occupied := p.Occupied

	for rooks != 0 {
		var sq Square
		sq, rooks = rooks.PopLSB()
		attacks |= genRookAttacks(sq, occupied)
	}

	return attacks
}

func genRookAttacks(sq Square, occupied BitBoard) BitBoard {
	// fmt.Println(rays[North][sq])
	attacks := rays[North][sq]
	if intersect := rays[North][sq] & occupied; intersect != 0 {
		// fmt.Println(intersect)
		index := bits.LeadingZeros64(uint64(intersect))
		attacks &^= rays[North][63-index]
	}
	// fmt.Println(attacks)
	// fmt.Println(rays[South][sq])

	attacks |= rays[South][sq]
	if intersect := rays[South][sq] & occupied; intersect != 0 {
		// fmt.Println(intersect)
		index := bits.TrailingZeros64(uint64(intersect))
		attacks &^= rays[South][index]
	}
	// fmt.Println(attacks)
	// fmt.Println(rays[East][sq])
	attacks |= rays[East][sq]
	if intersect := rays[East][sq] & occupied; intersect != 0 {
		// fmt.Println(intersect)
		index := bits.TrailingZeros64(uint64(intersect))
		attacks &^= rays[East][index]
	}
	// fmt.Println(attacks)
	// fmt.Println(rays[West][sq])
	attacks |= rays[West][sq]
	if intersect := rays[West][sq] & occupied; intersect != 0 {
		// fmt.Println(intersect)
		index := bits.LeadingZeros64(uint64(intersect))
		attacks &^= rays[West][63-index]
	}
	// fmt.Println(attacks)
	return attacks
}

func genWhiteQueenAttacks(p *Position) BitBoard {
	var attacks BitBoard

	queens := p.Pieces[White][Queen]
	occupied := p.Occupied

	for queens != 0 {
		var sq Square
		sq, queens = queens.PopLSB()
		attacks |= genBishopAttacks(sq, occupied) | genRookAttacks(sq, occupied)
	}

	return attacks
}

func genBlackQueenAttacks(p *Position) BitBoard {
	var attacks BitBoard

	queens := p.Pieces[Black][Queen]
	occupied := p.Occupied

	for queens != 0 {
		var sq Square
		sq, queens = queens.PopLSB()
		attacks |= genBishopAttacks(sq, occupied) | genRookAttacks(sq, occupied)
	}

	return attacks
}

func genWhiteKingAttacks(p *Position) BitBoard {
	king := p.Pieces[White][King]
	sq, _ := king.PopLSB()
	return kingMoves[sq]
}

func genBlackKingAttacks(p *Position) BitBoard {
	king := p.Pieces[Black][King]
	sq, _ := king.PopLSB()
	return kingMoves[sq]
}

func genWhitePawnMoves(moves *[]Move, p *Position) {
	pawns := p.Pieces[White][Pawn]
	occupied := p.Occupied
	enemies := p.AllPieces[Black]

	genWhiteForwardMoves(moves, pawns, occupied)
	genWhiteAttackMoves(moves, pawns, enemies)
	if p.IsEnPassant() {
		genWhiteEnPassantMoves(moves, pawns, p.EnPassantFile())
	}
}

func genWhiteForwardMoves(moves *[]Move, pawns, occupied BitBoard) {
	singlePushes := (pawns >> 8) &^ occupied

	// non-promotion moves
	for pushes := singlePushes &^ Rank_8; pushes != 0; {
		var to Square
		to, pushes = pushes.PopLSB()
		*moves = append(*moves, Move{Piece: Pawn, From: to + 8, To: to})
	}

	// promotion
	for pushes := singlePushes & Rank_8; pushes != 0; {
		var to Square
		to, pushes = pushes.PopLSB()
		*moves = append(*moves,
			Move{Piece: Pawn, From: to + 8, To: to, Type: PromotionToQueen},
			Move{Piece: Pawn, From: to + 8, To: to, Type: PromotionToRook},
			Move{Piece: Pawn, From: to + 8, To: to, Type: PromotionToBishop},
			Move{Piece: Pawn, From: to + 8, To: to, Type: PromotionToKnight},
		)
	}

	doublePushes := ((singlePushes >> 8) &^ occupied) & Rank_4
	for doublePushes != 0 {
		var to Square
		to, doublePushes = doublePushes.PopLSB()
		*moves = append(*moves, Move{Piece: Pawn, From: to + 16, To: to})
	}
}

func genWhiteAttackMoves(moves *[]Move, pawns, enemies BitBoard) {
	leftAttacks := (pawns &^ File_A) >> 9 & enemies

	// non-promotion captures
	for attacks := leftAttacks &^ Rank_8; attacks != 0; {
		var to Square
		to, attacks = attacks.PopLSB()
		*moves = append(*moves, Move{Piece: Pawn, From: to + 9, To: to})
	}

	// promotion
	for attacks := leftAttacks & Rank_8; attacks != 0; {
		var to Square
		to, attacks = attacks.PopLSB()
		*moves = append(*moves,
			Move{Piece: Pawn, From: to + 9, To: to, Type: PromotionToQueen},
			Move{Piece: Pawn, From: to + 9, To: to, Type: PromotionToRook},
			Move{Piece: Pawn, From: to + 9, To: to, Type: PromotionToBishop},
			Move{Piece: Pawn, From: to + 9, To: to, Type: PromotionToKnight},
		)
	}

	rightAttacks := (pawns &^ File_H) >> 7 & enemies
	for attacks := rightAttacks &^ Rank_8; attacks != 0; {
		var to Square
		to, attacks = attacks.PopLSB()
		*moves = append(*moves, Move{Piece: Pawn, From: to + 7, To: to})
	}

	// promotion
	for attacks := rightAttacks & Rank_8; attacks != 0; {
		var to Square
		to, attacks = attacks.PopLSB()
		*moves = append(*moves,
			Move{Piece: Pawn, From: to + 7, To: to, Type: PromotionToQueen},
			Move{Piece: Pawn, From: to + 7, To: to, Type: PromotionToRook},
			Move{Piece: Pawn, From: to + 7, To: to, Type: PromotionToBishop},
			Move{Piece: Pawn, From: to + 7, To: to, Type: PromotionToKnight},
		)
	}
}

func genWhiteEnPassantMoves(moves *[]Move, pawns, enPassantFile BitBoard) {
	pawnsOnRank := pawns & Rank_5
	left := (pawnsOnRank &^ File_A) >> 7 & enPassantFile
	if left != 0 {
		sq, _ := left.PopLSB()
		*moves = append(*moves, Move{Piece: Pawn, From: sq + 7, To: sq, Type: EnPassant})
	}

	right := (pawnsOnRank &^ File_H) >> 9 & enPassantFile
	if right != 0 {
		sq, _ := left.PopLSB()
		*moves = append(*moves, Move{Piece: Pawn, From: sq + 9, To: sq, Type: EnPassant})
	}
}

func genBlackPawnMoves(moves *[]Move, p *Position) {
	pawns := p.Pieces[Black][Pawn]
	occupied := p.Occupied
	enemies := p.AllPieces[White]

	genBlackForwardMoves(moves, pawns, occupied)
	genBlackAttackMoves(moves, pawns, enemies)
	if p.IsEnPassant() {
		genBlackEnPassantMoves(moves, pawns, p.EnPassantFile())
	}
}

func genBlackForwardMoves(moves *[]Move, pawns, occupied BitBoard) {
	singlePushes := (pawns << 8) &^ occupied

	// non-promotion moves
	for pushes := singlePushes &^ Rank_1; pushes != 0; {
		var to Square
		to, pushes = pushes.PopLSB()
		*moves = append(*moves, Move{Piece: Pawn, From: to - 8, To: to})
	}

	// promotion
	for pushes := singlePushes & Rank_1; pushes != 0; {
		var to Square
		to, pushes = pushes.PopLSB()
		*moves = append(*moves,
			Move{Piece: Pawn, From: to - 8, To: to, Type: PromotionToQueen},
			Move{Piece: Pawn, From: to - 8, To: to, Type: PromotionToRook},
			Move{Piece: Pawn, From: to - 8, To: to, Type: PromotionToBishop},
			Move{Piece: Pawn, From: to - 8, To: to, Type: PromotionToKnight},
		)
	}

	doublePushes := ((singlePushes << 8) &^ occupied) & Rank_5
	for doublePushes != 0 {
		var to Square
		to, doublePushes = doublePushes.PopLSB()
		*moves = append(*moves, Move{Piece: Pawn, From: to - 16, To: to})
	}
}

func genBlackAttackMoves(moves *[]Move, pawns, enemies BitBoard) {
	leftAttacks := (pawns &^ File_A) << 7 & enemies

	// non-promotion captures
	for attacks := leftAttacks &^ Rank_1; attacks != 0; {
		var to Square
		to, attacks = attacks.PopLSB()
		*moves = append(*moves, Move{Piece: Pawn, From: to - 7, To: to})
	}

	// promotion
	for attacks := leftAttacks & Rank_1; attacks != 0; {
		var to Square
		to, attacks = attacks.PopLSB()
		*moves = append(*moves,
			Move{Piece: Pawn, From: to - 7, To: to, Type: PromotionToQueen},
			Move{Piece: Pawn, From: to - 7, To: to, Type: PromotionToRook},
			Move{Piece: Pawn, From: to - 7, To: to, Type: PromotionToBishop},
			Move{Piece: Pawn, From: to - 7, To: to, Type: PromotionToKnight},
		)
	}

	rightAttacks := (pawns &^ File_H) << 9 & enemies
	for attacks := rightAttacks &^ Rank_1; attacks != 0; {
		var to Square
		to, attacks = attacks.PopLSB()
		*moves = append(*moves, Move{Piece: Pawn, From: to - 9, To: to})
	}

	// promotion
	for attacks := rightAttacks & Rank_1; attacks != 0; {
		var to Square
		to, attacks = attacks.PopLSB()
		*moves = append(*moves,
			Move{Piece: Pawn, From: to - 9, To: to, Type: PromotionToQueen},
			Move{Piece: Pawn, From: to - 9, To: to, Type: PromotionToRook},
			Move{Piece: Pawn, From: to - 9, To: to, Type: PromotionToBishop},
			Move{Piece: Pawn, From: to - 9, To: to, Type: PromotionToKnight},
		)
	}
}

func genBlackEnPassantMoves(moves *[]Move, pawns, enPassantFile BitBoard) {
	pawnsOnRank := pawns & Rank_4
	left := (pawnsOnRank &^ File_A) << 7 & enPassantFile
	if left != 0 {
		sq, _ := left.PopLSB()
		*moves = append(*moves, Move{Piece: Pawn, From: sq - 7, To: sq, Type: EnPassant})
	}

	right := (pawnsOnRank &^ File_H) << 9 & enPassantFile
	if right != 0 {
		sq, _ := left.PopLSB()
		*moves = append(*moves, Move{Piece: Pawn, From: sq - 9, To: sq, Type: EnPassant})
	}
}

func genWhiteKnightMoves(moves *[]Move, p *Position) {
	knights := p.Pieces[White][Knight]
	enemiesOrEmpty := ^p.AllPieces[White]

	for knights != 0 {
		var from Square
		from, knights = knights.PopLSB()
		targets := knightMoves[from] & enemiesOrEmpty

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			*moves = append(*moves, Move{Piece: Knight, From: from, To: to})
		}

	}
}

func genBlackKnightMoves(moves *[]Move, p *Position) {
	knights := p.Pieces[Black][Knight]
	enemiesOrEmpty := ^p.AllPieces[Black]

	for knights != 0 {
		var from Square
		from, knights = knights.PopLSB()
		targets := knightMoves[from] & enemiesOrEmpty

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			*moves = append(*moves, Move{Piece: Knight, From: from, To: to})
		}

	}
}

func genWhiteBishopMoves(moves *[]Move, p *Position) {
	bishops := p.Pieces[White][Bishop]
	enemiesOrEmpty := ^p.AllPieces[White]
	occupied := p.Occupied

	for bishops != 0 {
		var from Square
		from, bishops = bishops.PopLSB()
		targets := genBishopAttacks(from, occupied) & enemiesOrEmpty

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			*moves = append(*moves, Move{Piece: Bishop, From: from, To: to})
		}
	}
}

func genBlackBishopMoves(moves *[]Move, p *Position) {
	bishops := p.Pieces[Black][Bishop]
	enemiesOrEmpty := ^p.AllPieces[Black]
	occupied := p.Occupied

	for bishops != 0 {
		var from Square
		from, bishops = bishops.PopLSB()
		targets := genBishopAttacks(from, occupied) & enemiesOrEmpty

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			*moves = append(*moves, Move{Piece: Bishop, From: from, To: to})
		}
	}
}

func genWhiteRookMoves(moves *[]Move, p *Position) {
	rooks := p.Pieces[White][Rook]
	enemiesOrEmpty := ^p.AllPieces[White]
	occupied := p.Occupied

	for rooks != 0 {
		var from Square
		from, rooks = rooks.PopLSB()
		targets := genRookAttacks(from, occupied) & enemiesOrEmpty

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			*moves = append(*moves, Move{Piece: Rook, From: from, To: to})
		}
	}
}

func genBlackRookMoves(moves *[]Move, p *Position) {
	rooks := p.Pieces[Black][Rook]
	enemiesOrEmpty := ^p.AllPieces[Black]
	occupied := p.Occupied

	for rooks != 0 {
		var from Square
		from, rooks = rooks.PopLSB()
		targets := genRookAttacks(from, occupied) & enemiesOrEmpty

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			*moves = append(*moves, Move{Piece: Rook, From: from, To: to})
		}
	}
}

func genWhiteQueenMoves(moves *[]Move, p *Position) {
	queens := p.Pieces[White][Queen]
	enemiesOrEmpty := ^p.AllPieces[White]
	occupied := p.Occupied

	for queens != 0 {
		var from Square
		from, queens = queens.PopLSB()
		targets := (genRookAttacks(from, occupied) | genBishopAttacks(from, occupied)) & enemiesOrEmpty

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			*moves = append(*moves, Move{Piece: Queen, From: from, To: to})
		}
	}
}

func genBlackQueenMoves(moves *[]Move, p *Position) {
	queens := p.Pieces[Black][Queen]
	enemiesOrEmpty := ^p.AllPieces[Black]
	occupied := p.Occupied

	for queens != 0 {
		var from Square
		from, queens = queens.PopLSB()
		targets := (genRookAttacks(from, occupied) | genBishopAttacks(from, occupied)) & enemiesOrEmpty

		for targets != 0 {
			var to Square
			to, targets = targets.PopLSB()
			*moves = append(*moves, Move{Piece: Queen, From: from, To: to})
		}
	}
}

func genWhiteKingMoves(moves *[]Move, p *Position) {
	king := p.Pieces[White][King]
	enemyKing := p.Pieces[Black][King]
	enemiesOrEmpty := ^p.AllPieces[White]
	occupied := p.Occupied
	attacked := blackAttacks(p)

	// normal moves
	from, _ := king.PopLSB()
	targets := kingMoves[from] & enemiesOrEmpty &^ (enemyKing | attacked)

	for targets != 0 {
		var to Square
		to, targets = targets.PopLSB()
		*moves = append(*moves, Move{Piece: King, From: from, To: to})
	}

	// castling
	if p.CanWhiteCastleKingSide() {
		emptySquares := BitBoard(3) << uint8(SQ_F1)
		mustNotBeAttacked := BitBoard(7) << uint8(SQ_E1)

		if emptySquares&occupied == 0 && mustNotBeAttacked&attacked == 0 {
			*moves = append(*moves, Move{Piece: King, Type: Castle, From: SQ_E1, To: SQ_G1})
		}
	}

	if p.CanWhiteCastleQueenSide() {
		emptySquares := BitBoard(7) << uint8(SQ_B1)
		mustNotBeAttacked := BitBoard(7) << uint8(SQ_C1)

		if emptySquares&occupied == 0 && mustNotBeAttacked&attacked == 0 {
			*moves = append(*moves, Move{Piece: King, Type: Castle, From: SQ_E1, To: SQ_C1})
		}
	}
}

func genBlackKingMoves(moves *[]Move, p *Position) {
	king := p.Pieces[Black][King]
	enemyKing := p.Pieces[White][King]
	enemiesOrEmpty := ^p.AllPieces[Black]
	occupied := p.Occupied
	attacked := whiteAttacks(p)

	// normal moves
	from, _ := king.PopLSB()
	targets := kingMoves[from] & enemiesOrEmpty &^ (enemyKing | attacked)

	for targets != 0 {
		var to Square
		to, targets = targets.PopLSB()
		*moves = append(*moves, Move{Piece: King, From: from, To: to})
	}

	// castling
	if p.CanBlackCastleKingSide() {
		emptySquares := BitBoard(3) << uint8(SQ_F8)
		mustNotBeAttacked := BitBoard(7) << uint8(SQ_E8)

		if emptySquares&occupied == 0 && mustNotBeAttacked&attacked == 0 {
			*moves = append(*moves, Move{Piece: King, Type: Castle, From: SQ_E8, To: SQ_G8})
		}
	}

	if p.CanBlackCastleQueenSide() {
		emptySquares := BitBoard(7) << uint8(SQ_B8)
		mustNotBeAttacked := BitBoard(7) << uint8(SQ_C8)

		if emptySquares&occupied == 0 && mustNotBeAttacked&attacked == 0 {
			*moves = append(*moves, Move{Piece: King, Type: Castle, From: SQ_E8, To: SQ_C8})
		}
	}
}

func blackAttacks(p *Position) BitBoard {
	return genBlackPawnAttacks(p) |
		genBlackKnightAttacks(p) |
		genBlackBishopAttacks(p) |
		genBlackRookAttacks(p) |
		genBlackQueenAttacks(p) |
		genBlackKingAttacks(p)
}

func whiteAttacks(p *Position) BitBoard {
	return genWhitePawnAttacks(p) |
		genWhiteKnightAttacks(p) |
		genWhiteBishopAttacks(p) |
		genWhiteRookAttacks(p) |
		genWhiteQueenAttacks(p) |
		genWhiteKingAttacks(p)
}
