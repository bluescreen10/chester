package chester

import (
	"math/rand/v2"
	"testing"
)

// evalPestoBySquareScan is the previous implementation of EvalPesto, kept as
// a reference. It walks all sixty-four squares through the mailbox rather
// than walking the piece bitboards.
//
// It exists so that the faster implementation has something to be checked
// against. An evaluation that is quietly wrong does not crash or fail a
// perft; it just loses games, which is the most expensive way to find a bug.
func evalPestoBySquareScan(p *Position) int {
	var mg [2]int
	var eg [2]int
	gamePhase := 0

	whiteBB := p.WhitePieces()

	bb := Bitboard(1)
	for sq := range Square(64) {
		piece := p.mailbox[sq]
		if piece != Empty {
			if bb&whiteBB != 0 {
				mg[White] += mgTable[White][piece][sq]
				eg[White] += egTable[White][piece][sq]
			} else {
				mg[Black] += mgTable[Black][piece][sq]
				eg[Black] += egTable[Black][piece][sq]
			}
			gamePhase += gamephaseInc[piece]
		}
		bb <<= 1
	}

	mgScore := mg[p.active] - mg[p.inactive]
	egScore := eg[p.active] - eg[p.inactive]
	mgPhase := gamePhase
	if mgPhase > 24 {
		mgPhase = 24
	}

	egPhase := 24 - mgPhase
	return (mgScore*mgPhase + egScore*egPhase) / 24
}

var evalFENs = []string{
	DefaultFEN,
	"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1",
	"8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - - 0 1",
	"r2q1rk1/1b1nbppp/p2ppn2/1p6/3NPP2/1BN1B3/PPP3PP/R2Q1RK1 w - - 0 12",
	"8/8/4k3/8/2p5/8/B2P2K1/8 w - - 0 1",
	"rnbq1k1r/pp1Pbppp/2p5/8/2B5/8/PPP1NnPP/RNBQK2R w KQ - 1 8",
	"4k3/8/8/8/8/8/4P3/4K3 w - - 0 1",
}

func evalPositions(t testing.TB) []*Position {
	out := make([]*Position, 0, len(evalFENs))
	for _, fen := range evalFENs {
		p, err := ParseFEN(fen)
		if err != nil {
			t.Fatal(err)
		}
		out = append(out, p)
	}
	return out
}

// TestEvalPestoMatchesReference walks a tree of real positions and checks
// that the bitboard implementation scores every one of them exactly as the
// square scan did. Both sides of the board, every piece type, promotions and
// captures are reached by playing the moves rather than by listing positions.
func TestEvalPestoMatchesReference(t *testing.T) {
	r := rand.New(rand.NewPCG(7, 11))

	checked := 0
	for _, root := range evalPositions(t) {
		for game := 0; game < 60; game++ {
			p := *root

			for ply := 0; ply < 40; ply++ {
				if got, want := EvalPesto(&p), evalPestoBySquareScan(&p); got != want {
					t.Fatalf("%s: bitboard %d, square scan %d", p.FEN(), got, want)
				}
				checked++

				var moves []Move
				moves, _ = LegalMoves(moves, &p)
				if len(moves) == 0 {
					break
				}
				p.Do(moves[r.IntN(len(moves))])
			}
		}
	}

	t.Logf("%d positions agreed", checked)
}

func BenchmarkEvalPesto(b *testing.B) {
	ps := evalPositions(b)
	b.ResetTimer()

	var sink int
	for i := 0; i < b.N; i++ {
		sink += EvalPesto(ps[i%len(ps)])
	}
	_ = sink
}

func BenchmarkEvalPestoSquareScan(b *testing.B) {
	ps := evalPositions(b)
	b.ResetTimer()

	var sink int
	for i := 0; i < b.N; i++ {
		sink += evalPestoBySquareScan(ps[i%len(ps)])
	}
	_ = sink
}
