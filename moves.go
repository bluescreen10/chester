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

var knightMoves [Square(64)]BitBoard
var kingMoves [Square(64)]BitBoard
var squaresBetween [Square(64)][Square(64)]BitBoard

type config struct {
	singlePushes        int
	leftAttacks         int
	rightAttacks        int
	startingPlusOneRank BitBoard
	enPassantRank       BitBoard
	promotionRank       BitBoard
}

var pawnConfig [2]config

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

	// pre-compute squares between
	for i := Square(0); i < Square(64); i++ {
		for j := Square(0); j < Square(64); j++ {
			if i == j {
				continue
			}

			r1, f1 := i.RankAndFile()
			r2, f2 := j.RankAndFile()

			if r1 == r2 || f1 == f2 {
				squaresBetween[i][j] =
					(rays[North][i] & rays[South][j]) |
						(rays[South][i] & rays[North][j]) |
						(rays[East][i] & rays[West][j]) |
						(rays[West][i] & rays[East][j])
			} else if abs(r1-r2) == abs(f1-f2) {
				squaresBetween[i][j] =
					(rays[NorthEast][i] & rays[SouthWest][j]) |
						(rays[SouthWest][i] & rays[NorthEast][j]) |
						(rays[NorthWest][i] & rays[SouthEast][j]) |
						(rays[SouthEast][i] & rays[NorthWest][j])
			}
		}
	}

	//pawn config
	pawnConfig[White] = config{-8, -9, -7, Rank_3, Rank_5, Rank_8}
	pawnConfig[Black] = config{8, 7, 9, Rank_6, Rank_4, Rank_1}
}

func abs(num int8) int8 {
	if num < 0 {
		return -num
	}
	return num
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
		newPos := p.Do(m)
		newNodes := Perft(newPos, depth-1, false, output)
		if print {
			fmt.Fprintf(output, "%s: %d\n", m, newNodes)
		}
		nodes += newNodes
		//p.Undo()
	}
	movesPool.Put(moves)
	return nodes
}

func LegalMoves(moves *[]Move, pos *Position) {
	*moves = (*moves)[:0]

	us, them := pos.SideToMove()

	checkers, _, _ := checkersPinnedAndPinners(pos, us, them)

	if checkers.OnesCount() > 1 {
		genKingMoves(moves, pos, us, them, false)
		return
	} else {
		genPawnMoves(moves, pos, us, them)
		genKnightMoves(moves, pos, us)
		genBishopMoves(moves, pos, us)
		genRookMoves(moves, pos, us)
		genQueenMoves(moves, pos, us)
		genKingMoves(moves, pos, us, them, true)
	}

	for i := 0; i < len(*moves); i++ {
		move := (*moves)[i]
		newPos := pos.Do(move)

		if checkers, _, _ := checkersPinnedAndPinners(&newPos, us, them); checkers.OnesCount() > 0 {
			*moves = append((*moves)[:i], (*moves)[i+1:]...)
			i--
		}
	}
}

func checkersPinnedAndPinners(p *Position, us, them Color) (BitBoard, BitBoard, BitBoard) {
	//FIXME: implement pinned and pinners
	king := p.Pieces[us][King]
	kingSq, _ := king.PopLSB()

	checkers := knightMoves[kingSq] & p.Pieces[them][Knight]
	if us == White {
		checkers |= ((king & File_Not_A) >> 9) & p.Pieces[them][Pawn]
		checkers |= ((king & File_Not_H) >> 7) & p.Pieces[them][Pawn]
	} else {
		checkers |= ((king & File_Not_A) << 7) & p.Pieces[them][Pawn]
		checkers |= ((king & File_Not_H) << 9) & p.Pieces[them][Pawn]
	}

	kingDiagonalRays := rays[NorthWest][kingSq] | rays[NorthEast][kingSq] | rays[SouthWest][kingSq] | rays[SouthEast][kingSq]
	diagonalAttackers := (p.Pieces[them][Queen] | p.Pieces[them][Bishop])

	kingStraightRays := rays[North][kingSq] | rays[South][kingSq] | rays[East][kingSq] | rays[West][kingSq]
	straightAttackers := (p.Pieces[them][Queen] | p.Pieces[them][Rook])

	potentialCheckers := (diagonalAttackers & kingDiagonalRays) | (straightAttackers & kingStraightRays)
	occupied := p.Occupied

	for potentialCheckers != 0 {
		var sq Square
		sq, potentialCheckers = potentialCheckers.PopLSB()

		path := squaresBetween[sq][kingSq]
		if path&occupied == 0 {
			checkers |= 1 << sq
		}
	}

	return checkers, 0, 0
}

func genPawnsAttacks(p *Position, us Color) BitBoard {
	pawns := p.Pieces[us][Pawn]
	config := pawnConfig[us]

	left := (pawns & File_Not_A).RotateLeft(config.leftAttacks)
	right := (pawns & File_Not_H).RotateLeft(config.rightAttacks)
	return left | right
}

func genKnightsAttacks(p *Position, us Color) BitBoard {
	var attacks BitBoard

	knights := p.Pieces[us][Knight]

	for knights != 0 {
		var sq Square
		sq, knights = knights.PopLSB()
		attacks |= knightMoves[sq]
	}

	return attacks
}

func genBishopsAttacks(p *Position, us Color) BitBoard {
	var attacks BitBoard

	bishops := p.Pieces[us][Bishop]
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

func genRooksAttacks(p *Position, us Color) BitBoard {
	var attacks BitBoard

	rooks := p.Pieces[us][Rook]
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

func genQueensAttacks(p *Position, us Color) BitBoard {
	var attacks BitBoard

	queens := p.Pieces[us][Queen]
	occupied := p.Occupied

	for queens != 0 {
		var sq Square
		sq, queens = queens.PopLSB()
		attacks |= genBishopAttacks(sq, occupied) | genRookAttacks(sq, occupied)
	}

	return attacks
}

func genKingAttacks(p *Position, us Color) BitBoard {
	king := p.Pieces[us][King]
	sq, _ := king.PopLSB()
	return kingMoves[sq]
}

func genPawnMoves(moves *[]Move, p *Position, us, them Color) {
	pawns := p.Pieces[us][Pawn]
	occupied := p.Occupied
	enemies := p.AllPieces[them]

	genForwardMoves(moves, pawns, occupied, us)
	genAttackMoves(moves, pawns, enemies, us)
	if p.IsEnPassant() {
		genEnPassantMoves(moves, pawns, p.EnPassantFile(), us)
	}
}

func genForwardMoves(moves *[]Move, pawns, occupied BitBoard, us Color) {
	config := pawnConfig[us]

	singlePushes := pawns.RotateLeft(config.singlePushes) &^ occupied

	// non-promotion moves
	for pushes := singlePushes &^ config.promotionRank; pushes != 0; {
		var to Square
		to, pushes = pushes.PopLSB()
		*moves = append(*moves, Move{Piece: Pawn, From: to - Square(config.singlePushes), To: to})
	}

	// promotion
	for pushes := singlePushes & config.promotionRank; pushes != 0; {
		var to Square
		to, pushes = pushes.PopLSB()
		from := to - Square(config.singlePushes)
		*moves = append(*moves,
			Move{Piece: Pawn, From: from, To: to, Type: PromotionToQueen},
			Move{Piece: Pawn, From: from, To: to, Type: PromotionToRook},
			Move{Piece: Pawn, From: from, To: to, Type: PromotionToBishop},
			Move{Piece: Pawn, From: from, To: to, Type: PromotionToKnight},
		)
	}

	doublePushes := (singlePushes & config.startingPlusOneRank).RotateLeft(config.singlePushes) &^ occupied
	for doublePushes != 0 {
		var to Square
		to, doublePushes = doublePushes.PopLSB()
		*moves = append(*moves, Move{Piece: Pawn, From: to - Square(2*config.singlePushes), To: to})
	}
}

func genAttackMoves(moves *[]Move, pawns, enemies BitBoard, us Color) {
	config := pawnConfig[us]
	leftAttacks := (pawns &^ File_A).RotateLeft(config.leftAttacks) & enemies

	// non-promotion captures
	for attacks := leftAttacks &^ config.promotionRank; attacks != 0; {
		var to Square
		to, attacks = attacks.PopLSB()
		*moves = append(*moves, Move{Piece: Pawn, From: to - Square(config.leftAttacks), To: to})
	}

	// promotion
	for attacks := leftAttacks & config.promotionRank; attacks != 0; {
		var to Square
		to, attacks = attacks.PopLSB()
		from := to - Square(config.leftAttacks)
		*moves = append(*moves,
			Move{Piece: Pawn, From: from, To: to, Type: PromotionToQueen},
			Move{Piece: Pawn, From: from, To: to, Type: PromotionToRook},
			Move{Piece: Pawn, From: from, To: to, Type: PromotionToBishop},
			Move{Piece: Pawn, From: from, To: to, Type: PromotionToKnight},
		)
	}

	rightAttacks := (pawns &^ File_H).RotateLeft(config.rightAttacks) & enemies
	for attacks := rightAttacks &^ config.promotionRank; attacks != 0; {
		var to Square
		to, attacks = attacks.PopLSB()
		*moves = append(*moves, Move{Piece: Pawn, From: to - Square(config.rightAttacks), To: to})
	}

	// promotion
	for attacks := rightAttacks & config.promotionRank; attacks != 0; {
		var to Square
		to, attacks = attacks.PopLSB()
		from := to - Square(config.rightAttacks)
		*moves = append(*moves,
			Move{Piece: Pawn, From: from, To: to, Type: PromotionToQueen},
			Move{Piece: Pawn, From: from, To: to, Type: PromotionToRook},
			Move{Piece: Pawn, From: from, To: to, Type: PromotionToBishop},
			Move{Piece: Pawn, From: from, To: to, Type: PromotionToKnight},
		)
	}
}

func genEnPassantMoves(moves *[]Move, pawns, enPassantFile BitBoard, us Color) {
	config := pawnConfig[us]
	pawnsOnRank := pawns & config.enPassantRank
	left := (pawnsOnRank &^ File_A).RotateLeft(config.leftAttacks) & enPassantFile
	if left != 0 {
		sq, _ := left.PopLSB()
		*moves = append(*moves, Move{Piece: Pawn, From: sq - Square(config.leftAttacks), To: sq, Type: EnPassant})
	}

	right := (pawnsOnRank &^ File_H).RotateLeft(config.rightAttacks) & enPassantFile
	if right != 0 {
		sq, _ := right.PopLSB()
		*moves = append(*moves, Move{Piece: Pawn, From: sq - Square(config.rightAttacks), To: sq, Type: EnPassant})
	}
}

func genKnightMoves(moves *[]Move, p *Position, us Color) {
	knights := p.Pieces[us][Knight]
	enemiesOrEmpty := ^p.AllPieces[us]

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

func genBishopMoves(moves *[]Move, p *Position, us Color) {
	bishops := p.Pieces[us][Bishop]
	enemiesOrEmpty := ^p.AllPieces[us]
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

func genRookMoves(moves *[]Move, p *Position, us Color) {
	rooks := p.Pieces[us][Rook]
	enemiesOrEmpty := ^p.AllPieces[us]
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

func genQueenMoves(moves *[]Move, p *Position, us Color) {
	queens := p.Pieces[us][Queen]
	enemiesOrEmpty := ^p.AllPieces[us]
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

func genKingMoves(moves *[]Move, p *Position, us, them Color, castling bool) {
	king := p.Pieces[us][King]
	enemyKing := p.Pieces[them][King]
	enemiesOrEmpty := ^p.AllPieces[us]
	occupied := p.Occupied
	attacked := attacks(p, them)

	// normal moves
	from, _ := king.PopLSB()
	targets := kingMoves[from] & enemiesOrEmpty &^ (enemyKing | attacked)

	for targets != 0 {
		var to Square
		to, targets = targets.PopLSB()
		*moves = append(*moves, Move{Piece: King, From: from, To: to})
	}

	if !castling {
		return
	}

	// castling
	if us == White {
		if p.CanWhiteCastleKingSide() {
			emptySquares := squaresBetween[SQ_E1][SQ_H1]
			mustNotBeAttacked := squaresBetween[SQ_D1][SQ_H1]

			if emptySquares&occupied == 0 && mustNotBeAttacked&attacked == 0 {
				*moves = append(*moves, Move{Piece: King, Type: Castle, From: SQ_E1, To: SQ_G1})
			}
		}

		if p.CanWhiteCastleQueenSide() {
			emptySquares := squaresBetween[SQ_A1][SQ_E1]
			mustNotBeAttacked := squaresBetween[SQ_B1][SQ_E1]

			attacked := attacks(p, them)

			if emptySquares&occupied == 0 && mustNotBeAttacked&attacked == 0 {
				*moves = append(*moves, Move{Piece: King, Type: Castle, From: SQ_E1, To: SQ_C1})
			}
		}
	} else {
		if p.CanBlackCastleKingSide() {
			emptySquares := squaresBetween[SQ_E8][SQ_H8]
			mustNotBeAttacked := squaresBetween[SQ_D8][SQ_H8]

			if emptySquares&occupied == 0 && mustNotBeAttacked&attacked == 0 {
				*moves = append(*moves, Move{Piece: King, Type: Castle, From: SQ_E8, To: SQ_G8})
			}
		}

		if p.CanBlackCastleQueenSide() {
			emptySquares := squaresBetween[SQ_A8][SQ_E8]
			mustNotBeAttacked := squaresBetween[SQ_B8][SQ_E8]

			attacked := attacks(p, them)

			if emptySquares&occupied == 0 && mustNotBeAttacked&attacked == 0 {
				*moves = append(*moves, Move{Piece: King, Type: Castle, From: SQ_E8, To: SQ_C8})
			}
		}
	}
}

func attacks(p *Position, us Color) BitBoard {
	return genPawnsAttacks(p, us) |
		genKnightsAttacks(p, us) |
		genBishopsAttacks(p, us) |
		genRooksAttacks(p, us) |
		genQueensAttacks(p, us) |
		genKingAttacks(p, us)
}
