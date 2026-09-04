package chester

import "testing"

func TestIsRepetition(t *testing.T) {
	const (
		a uint64 = 0xaaaa
		b uint64 = 0xbbbb
		c uint64 = 0xcccc
	)

	tests := []struct {
		name      string
		stack     []uint64
		hash      uint64
		halfMoves uint8
		want      bool
	}{
		{
			name:      "same position two plies back",
			stack:     []uint64{a, b},
			hash:      a,
			halfMoves: 2,
			want:      true,
		},
		{
			name:      "same position four plies back",
			stack:     []uint64{a, b, c, b},
			hash:      a,
			halfMoves: 4,
			want:      true,
		},
		{
			name: "odd distance is the other side to move",
			// b is one ply back, so it cannot be the same position even
			// though the hashes would have to differ anyway.
			stack:     []uint64{a, b},
			hash:      b,
			halfMoves: 2,
			want:      false,
		},
		{
			name:      "unseen position",
			stack:     []uint64{a, b, c, b},
			hash:      0xdddd,
			halfMoves: 4,
			want:      false,
		},
		{
			name: "irreversible move puts the match out of reach",
			// The half-move clock was reset since a was played, so a can
			// never occur again and must not be scanned.
			stack:     []uint64{a, b, c, b},
			hash:      a,
			halfMoves: 2,
			want:      false,
		},
		{
			name:      "empty stack",
			stack:     nil,
			hash:      a,
			halfMoves: 8,
			want:      false,
		},
		{
			name: "clock longer than the stack",
			// A search started from a FEN has a half-move clock but no
			// preceding positions; the scan must stop at the stack.
			stack:     []uint64{a},
			hash:      a,
			halfMoves: 50,
			want:      false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := &searchCtx{stack: test.stack}
			p := &Position{hash: test.hash, halfMoves: test.halfMoves}

			if got := ctx.isRepetition(p); got != test.want {
				t.Errorf("isRepetition() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestMateScoreRoundTrip(t *testing.T) {
	tests := []struct {
		name  string
		score int
		ply   int
	}{
		{"positive mate", MateScore - 5, 7},
		{"negative mate", -MateScore + 5, 7},
		{"ordinary score", 250, 7},
		{"negative ordinary score", -250, 7},
		{"draw", drawScore, 7},
		{"root ply", MateScore - 5, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			stored := scoreToTT(test.score, test.ply)
			if got := scoreFromTT(stored, test.ply); got != test.score {
				t.Errorf("round trip = %d, want %d", got, test.score)
			}
		})
	}
}

// TestMateScoreIsPlyRelative checks the property the round trip exists for: a
// mate stored at one ply and read back at another must describe the distance
// from the node reading it, not from the node that stored it.
func TestMateScoreIsPlyRelative(t *testing.T) {
	// A mate found five plies below a node at ply 3 is "mate in 8" from the
	// root, and must read back as "mate in 6" at a node two plies shallower.
	stored := scoreToTT(MateScore-8, 3)

	if got, want := scoreFromTT(stored, 1), MateScore-6; got != want {
		t.Errorf("score at ply 1 = %d, want %d", got, want)
	}
}

// TestOrdinaryScoresAreNotAdjusted guards the threshold: a large but
// non-mate score must survive the table untouched at any ply.
func TestOrdinaryScoresAreNotAdjusted(t *testing.T) {
	score := mateThreshold - 1
	if got := scoreToTT(score, 40); got != score {
		t.Errorf("scoreToTT(%d, 40) = %d, want unchanged", score, got)
	}
}
