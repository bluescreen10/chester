package main

import "fmt"

const (
	maxMoves = 218
)

type MoveType uint8
type Square uint8

const (
	Default MoveType = iota
	CastleKingSide
	CastleQueenSide
	Capture
	EnPassant
	Promotion
)

type Move struct {
	Type         MoveType
	From, To     Square
	Piece, Enemy Piece
}

// func (m Move) String() string {
// 	target := fmt.Sprintf("%c%d", 'a'+int(m.To%8), int(m.To/8)+1)

//		switch m.Piece {
//		case WhitePawn, BlackPawn:
//			return target
//		case WhiteKnight, BlackKnight:
//			return "N" + target
//		case WhiteBishop, BlackBishop:
//			return "B" + target
//		case WhiteRook, BlackRook:
//			return "R" + target
//		case WhiteQueen, BlackQueen:
//			return "Q" + target
//		case WhiteKing, BlackKing:
//			return "K" + target
//		default:
//			return "??"
//		}
//	}
func (m Move) String() string {
	return fmt.Sprintf("%c%d%c%d", 'a'+int(m.From%8), int(m.From/8)+1, 'a'+int(m.To%8), int(m.To/8)+1)
}

func LegalMoves(position Position) []Move {
	moves := make([]Move, 0, maxMoves)

	stm := position.SideToMove()

	inCheck := IsInCheck(position, stm)

	for sq, piece := range position.Board {
		if piece == Empty || piece.Color() != stm {
			continue
		}

		switch piece {
		case WhitePawn, BlackPawn:
			genPawnMoves(&moves, position, sq, piece, stm)
		case WhiteKnight, BlackKnight:
			genKnightMoves(&moves, position, sq, piece, stm)
		case WhiteBishop, BlackBishop:
			genBishopMoves(&moves, position, sq, piece, stm)
		case WhiteRook, BlackRook:
			genRookMoves(&moves, position, sq, piece, stm)
		case WhiteQueen, BlackQueen:
			genQueenMoves(&moves, position, sq, piece, stm)
		case WhiteKing, BlackKing:
			genKingMoves(&moves, position, sq, piece, stm, inCheck)
		}
	}

	for i := 0; i < len(moves); i++ {
		move := moves[i]
		position.DoMove(move)
		if IsInCheck(position, stm) {
			moves = append(moves[:i], moves[i+1:]...)
			i--
		}
		position.UndoMove(move)
	}

	return moves
}

func genPawnMoves(moves *[]Move, position Position, sq int, piece, stm Piece) {
	rank, file := sq/8, sq%8

	dir := 1
	startRank := 1
	enPassantRank := 4
	promoRank := 6
	promoPieces := []Piece{WhiteQueen, WhiteRook, WhiteBishop, WhiteKnight}

	if stm == Black {
		dir = -1
		startRank = 6
		enPassantRank = 3
		promoRank = 1
		promoPieces = []Piece{BlackQueen, BlackRook, BlackBishop, BlackKnight}
	}

	targetSq := sq + dir*8
	target := position.Board[targetSq]

	if target == Empty && rank != promoRank {
		*moves = append(*moves, Move{From: Square(sq), To: Square(targetSq), Piece: piece})

		if rank == startRank {
			targetSq += dir * 8
			target = position.Board[targetSq]

			if target == Empty {
				*moves = append(*moves, Move{From: Square(sq), To: Square(targetSq), Piece: piece})
			}
		}
	}

	if target == Empty && rank == promoRank {
		for _, p := range promoPieces {
			*moves = append(*moves, Move{Type: Promotion, From: Square(sq), To: Square(targetSq), Piece: p})
		}
	}

	if file > 0 {
		targetSq = sq + dir*8 - 1
		target = position.Board[targetSq]

		if target != Empty && target.Color() != stm {
			if rank == promoRank {
				for _, p := range promoPieces {
					*moves = append(*moves, Move{Type: Promotion, From: Square(sq), To: Square(targetSq), Piece: p, Enemy: target})
				}
			} else {
				*moves = append(*moves, Move{Type: Capture, From: Square(sq), To: Square(targetSq), Piece: piece, Enemy: target})
			}
		}
	}

	if file < 7 {
		targetSq = sq + dir*8 + 1
		target = position.Board[targetSq]

		if target != Empty && target.Color() != stm {
			if rank == promoRank {
				for _, p := range promoPieces {
					*moves = append(*moves, Move{Type: Promotion, From: Square(sq), To: Square(targetSq), Piece: p, Enemy: target})
				}
			} else {
				*moves = append(*moves, Move{Type: Capture, From: Square(sq), To: Square(targetSq), Piece: piece, Enemy: target})
			}
		}
	}

	if rank == enPassantRank && position.IsEnPassant() {
		enPassantFile := position.EnPassantFile()
		if Abs(enPassantFile-file) == 1 {
			targetSq = sq + dir*8 + enPassantFile - file
			target = position.Board[targetSq]
			*moves = append(*moves, Move{Type: EnPassant, From: Square(sq), To: Square(targetSq), Piece: piece, Enemy: target})
		}
	}
}

func genKnightMoves(moves *[]Move, position Position, sq int, piece, stm Piece) {
	rank, file := sq/8, sq%8

	knightMoves := [][]int{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}}

	for _, offset := range knightMoves {

		targetRank, targetFile := rank+offset[0], file+offset[1]
		if targetRank < 0 || targetRank > 7 || targetFile < 0 || targetFile > 7 {
			continue
		}

		targetSq := targetRank*8 + targetFile
		target := position.Board[targetSq]

		if target != Empty && target.Color() == stm {
			continue
		}

		if target == Empty || target.Color() != stm {
			t := Default
			if target != Empty {
				t = Capture
			}
			*moves = append(*moves, Move{Type: t, From: Square(sq), To: Square(targetSq), Piece: piece, Enemy: target})
		}
	}
}

func genBishopMoves(moves *[]Move, position Position, sq int, piece, stm Piece) {
	rank, file := sq/8, sq%8

	bishopMoves := [][]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	for _, offset := range bishopMoves {
		incX, incY := offset[0], offset[1]

		for targetRank, targetFile := rank+incX, file+incY; targetRank >= 0 && targetRank <= 7 && targetFile >= 0 && targetFile <= 7; targetRank, targetFile = targetRank+incX, targetFile+incY {
			targetSq := targetRank*8 + targetFile
			target := position.Board[targetSq]

			if target != Empty {
				if target.Color() != stm {
					*moves = append(*moves, Move{Type: Capture, From: Square(sq), To: Square(targetSq), Piece: piece, Enemy: target})
				}
				break
			}

			if target == Empty {
				*moves = append(*moves, Move{Type: Default, From: Square(sq), To: Square(targetSq), Piece: piece, Enemy: target})
			}
		}
	}
}

func genRookMoves(moves *[]Move, position Position, sq int, piece, stm Piece) {
	rank, file := sq/8, sq%8

	rookMoves := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for _, offset := range rookMoves {
		incX, incY := offset[0], offset[1]

		for targetRank, targetFile := rank+incX, file+incY; targetRank >= 0 && targetRank <= 7 && targetFile >= 0 && targetFile <= 7; targetRank, targetFile = targetRank+incX, targetFile+incY {
			targetSq := targetRank*8 + targetFile
			target := position.Board[targetSq]

			if target != Empty {
				if target.Color() != stm {
					*moves = append(*moves, Move{Type: Capture, From: Square(sq), To: Square(targetSq), Piece: piece, Enemy: target})
				}
				break
			}

			if target == Empty {
				*moves = append(*moves, Move{Type: Default, From: Square(sq), To: Square(targetSq), Piece: piece, Enemy: target})
			}
		}
	}
}

func genQueenMoves(moves *[]Move, position Position, sq int, piece, stm Piece) {
	genBishopMoves(moves, position, sq, piece, stm)
	genRookMoves(moves, position, sq, piece, stm)
}

func genKingMoves(moves *[]Move, position Position, sq int, piece, stm Piece, inCheck bool) {
	rank, file := sq/8, sq%8

	kingMoves := [][]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}

	for _, offset := range kingMoves {
		targetRank, targetFile := rank+offset[0], file+offset[1]
		if targetRank < 0 || targetRank > 7 || targetFile < 0 || targetFile > 7 {
			continue
		}

		targetSq := targetRank*8 + targetFile
		target := position.Board[targetSq]

		if target != Empty && target.Color() == stm && target != WhiteKing && target != BlackKing {
			continue
		}

		if target == Empty || target.Color() != stm {
			t := Default
			if target != Empty {
				t = Capture
			}
			*moves = append(*moves, Move{Type: t, From: Square(sq), To: Square(targetSq), Piece: piece, Enemy: target})
		}
	}

	if position.CanCastleKingSide(stm) && !inCheck {
		if position.Board[sq+1] == Empty && position.Board[sq+2] == Empty && !IsAttacked(position, sq+1, stm) && !IsAttacked(position, sq+2, stm) {
			*moves = append(*moves, Move{Type: CastleKingSide})
		}
	}

	if position.CanCastleQueenSide(stm) && !inCheck {
		if position.Board[sq-1] == Empty && position.Board[sq-2] == Empty && position.Board[sq-3] == Empty && !IsAttacked(position, sq-1, stm) && !IsAttacked(position, sq-2, stm) && !IsAttacked(position, sq-3, stm) {
			*moves = append(*moves, Move{Type: CastleQueenSide})
		}
	}
}

func IsAttacked(position Position, sq int, stm Piece) bool {
	for i, piece := range position.Board {
		if piece == Empty || piece.Color() == stm {
			continue
		}

		switch piece {
		case WhitePawn, BlackPawn:
			if isPawnAttacking(i, sq, piece.Color()) {
				return true
			}
		case WhiteKnight, BlackKnight:
			if isKnightAttacking(i, sq) {
				return true
			}
		case WhiteBishop, BlackBishop:
			if isBishopAttacking(position, i, sq) {
				return true
			}
		case WhiteRook, BlackRook:
			if isRookAttacking(position, i, sq) {
				return true
			}
		case WhiteQueen, BlackQueen:
			if isQueenAttacking(position, i, sq) {
				return true
			}
		}
	}

	return false

}

func IsInCheck(position Position, stm Piece) bool {

	kingSq := 0

	for sq, piece := range position.Board {
		if piece == Empty || piece.Color() != stm {
			continue
		}

		if piece == WhiteKing || piece == BlackKing {
			kingSq = sq
		}
	}

	return IsAttacked(position, kingSq, stm)
}

func isPawnAttacking(pawnSq, targetSq int, stm Piece) bool {
	rank, file := pawnSq/8, pawnSq%8
	targetRank, targetFile := targetSq/8, targetSq%8

	if stm == White {
		if targetRank == rank+1 && Abs(targetFile-file) == 1 {
			return true
		}
	} else {
		if targetRank == rank-1 && Abs(targetFile-file) == 1 {
			return true
		}
	}

	return false
}

func isKnightAttacking(knightSq, targetSq int) bool {
	rank, file := knightSq/8, knightSq%8
	targetRank, targetFile := targetSq/8, targetSq%8

	knightMoves := [][]int{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}}

	for _, offset := range knightMoves {
		if targetRank == rank+offset[0] && targetFile == file+offset[1] {
			return true
		}
	}

	return false
}

func isBishopAttacking(position Position, sq, targetSq int) bool {
	rank, file := sq/8, sq%8
	targetRank, targetFile := targetSq/8, targetSq%8

	bishopMoves := [][]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	for _, offset := range bishopMoves {
		incX, incY := offset[0], offset[1]

		for r, f := rank+incX, file+incY; r >= 0 && r <= 7 && f >= 0 && f <= 7; r, f = r+incX, f+incY {
			if r == targetRank && f == targetFile {
				return true
			}

			if position.Board[r*8+f] != Empty {
				break
			}
		}
	}

	return false
}

func isRookAttacking(position Position, sq, targetSq int) bool {
	rank, file := sq/8, sq%8
	targetRank, targetFile := targetSq/8, targetSq%8

	rookMoves := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for _, offset := range rookMoves {
		incX, incY := offset[0], offset[1]

		for r, f := rank+incX, file+incY; r >= 0 && r <= 7 && f >= 0 && f <= 7; r, f = r+incX, f+incY {
			if r == targetRank && f == targetFile {
				return true
			}

			if position.Board[r*8+f] != Empty {
				break
			}
		}
	}

	return false
}

func isQueenAttacking(position Position, sq int, targetSq int) bool {
	return isBishopAttacking(position, sq, targetSq) || isRookAttacking(position, sq, targetSq)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
