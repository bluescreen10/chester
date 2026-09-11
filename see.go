package chester

// Static exchange evaluation.
//
// A capture's worth is not decided by the two pieces named in the move. Taking
// a knight with a rook looks like winning a knight until you notice a pawn
// defends it, at which point it is losing two pawns' worth of material. Move
// ordering by most-valuable-victim cannot see the defender at all; this can,
// and it does so without searching a single node.
//
// The method is to play out the whole exchange on one square, each side always
// recapturing with its least valuable attacker, and let either side stop when
// continuing would cost it material.

// seeValue is the material each piece is worth to an exchange. The king has no
// value because it is never traded: an exchange it takes part in ends there,
// handled as its own case below.
var seeValue = [7]int{
	Pawn:   100,
	Knight: 320,
	Bishop: 330,
	Rook:   500,
	Queen:  900,
	King:   0,
	Empty:  0,
}

// pawnAttackersTo returns the squares a pawn of the given colour would have to
// stand on to attack sq. It is the inverse of the usual question, which is why
// the shifts run opposite to the direction pawns move.
func pawnAttackersTo(sq Square, color Color) Bitboard {
	bb := NewBitboardFromSquare(sq)

	// Plain shifts rather than rotates. A rotate carries a square off one edge
	// of the board and back on at the other, inventing attackers that are not
	// there. The engine's own generators get away with rotating because they
	// always intersect the result with real pawns, and the wrapped bits happen
	// to land on the back ranks where no pawn can stand -- but a function that
	// returns a set of squares should return only squares that exist.
	//
	// The file masks stop the diagonal step from crossing into the next rank.
	if color == White {
		// White advances towards lower indices, so an attacker of sq sits at a
		// higher index: one rank behind it, one file to either side.
		return ((bb & File_Not_A) << 7) | ((bb & File_Not_H) << 9)
	}

	return ((bb & File_Not_A) >> 9) | ((bb & File_Not_H) >> 7)
}

// attackersTo returns every piece of either colour that attacks sq under the
// given occupancy.
//
// The occupancy is a parameter rather than taken from the position because the
// exchange removes pieces as it proceeds, and removing one can uncover another
// behind it -- a rook behind a rook, a bishop behind a pawn. Recomputing the
// sliding attacks against the updated occupancy is what makes those x-rays
// appear, and the magic tables make it almost free.
func (p *Position) attackersTo(sq Square, occupied Bitboard) Bitboard {
	attackers := knightMoves[sq] & p.pieces[Knight]
	attackers |= kingMoves[sq] & p.pieces[King]
	attackers |= genBishopAttacks(sq, occupied) & (p.pieces[Bishop] | p.pieces[Queen])
	attackers |= genRookAttacks(sq, occupied) & (p.pieces[Rook] | p.pieces[Queen])
	attackers |= pawnAttackersTo(sq, White) & p.pieces[Pawn] & p.allPieces[White]
	attackers |= pawnAttackersTo(sq, Black) & p.pieces[Pawn] & p.allPieces[Black]

	return attackers & occupied
}

// leastValuableAttacker returns the cheapest piece among the given attackers,
// and the bitboard of that piece type's attackers. Recapturing with anything
// dearer than necessary can only make the exchange worse.
func (p *Position) leastValuableAttacker(attackers Bitboard) (Piece, Bitboard) {
	for piece := Pawn; piece <= King; piece++ {
		if bb := attackers & p.pieces[piece]; bb != 0 {
			return piece, bb
		}
	}
	return Empty, 0
}

// seeGE reports whether the static exchange evaluation of m is at least
// threshold centipawns.
//
// It answers a yes/no question rather than computing the exact value, which
// lets it stop as soon as the outcome is settled. Callers almost always want
// to know whether a capture loses material, not by how much.
func seeGE(p *Position, m Move, threshold int) bool {
	from, to := m.From(), m.To()

	// A promotion adds a queen part way through the sequence, and the pawn an
	// en passant capture removes is not on the destination square. Neither
	// fits the swap, both are rare, and mismodelling them is worse than
	// declining to judge them, so they are reported as neutral.
	if m.IsPromotion() || (p.mailbox[from] == Pawn && to == p.enPassantTarget) {
		return threshold <= 0
	}

	// What the capture wins outright. Every later step in an exchange only
	// takes material away, so falling short here can never be recovered.
	swap := seeValue[p.mailbox[to]] - threshold
	if swap < 0 {
		return false
	}

	// What it costs if the opponent simply recaptures. Clearing the threshold
	// even after losing the moving piece makes the rest of the sequence moot.
	swap = seeValue[p.mailbox[from]] - swap
	if swap <= 0 {
		return true
	}

	occupied := p.Occupied() ^ NewBitboardFromSquare(from) ^ NewBitboardFromSquare(to)
	attackers := p.attackersTo(to, occupied)

	side := p.active
	result := 1

	for {
		side = 1 - side
		attackers &= occupied

		mine := attackers & p.allPieces[side]
		if mine == 0 {
			// This side has run out of attackers, so it is the one that has
			// to accept the balance as it stands.
			break
		}

		result ^= 1

		piece, bb := p.leastValuableAttacker(mine)

		if piece == King {
			// A king may only capture on an undefended square. If the other
			// side still has an attacker, this recapture is not available and
			// the exchange ended one step earlier than it appeared to.
			if attackers&p.allPieces[1-side] != 0 {
				result ^= 1
			}
			break
		}

		// Would continuing still clear the threshold? The comparison is
		// against result rather than zero because the side that is about to
		// run out of material is the one holding the balance.
		swap = seeValue[piece] - swap
		if swap < result {
			break
		}

		sq, _ := bb.PopLSB()
		occupied ^= NewBitboardFromSquare(sq)

		// Taking a piece off the square it attacked from can uncover a slider
		// behind it, which now joins the exchange.
		switch piece {
		case Pawn, Bishop, Queen:
			attackers |= genBishopAttacks(to, occupied) & (p.pieces[Bishop] | p.pieces[Queen])
		}
		switch piece {
		case Rook, Queen:
			attackers |= genRookAttacks(to, occupied) & (p.pieces[Rook] | p.pieces[Queen])
		}
	}

	return result != 0
}
