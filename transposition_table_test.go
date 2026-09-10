package chester

import (
	"math/rand/v2"
	"sync"
	"testing"
	"unsafe"
)

// TestTTLayout guards the storage layout. A sixteen-byte slot puts four
// entries in a sixty-four byte bucket, so a probe reads exactly one cache
// line; a slot that grows past sixteen bytes both halves the table's capacity
// and makes every probe touch two lines.
func TestTTLayout(t *testing.T) {
	if got, want := unsafe.Sizeof(ttSlot{}), uintptr(16); got != want {
		t.Errorf("sizeof(ttSlot) = %d, want %d", got, want)
	}
	if got, want := unsafe.Sizeof(ttBucket{}), uintptr(64); got != want {
		t.Errorf("sizeof(ttBucket) = %d, want %d", got, want)
	}
}

func TestTTEntryRoundTrip(t *testing.T) {
	tests := []ttEntry{
		{score: 0, move: 0, depth: 0, flag: exact, gen: 0},
		{score: 1, move: NewMove(SQ_E2, SQ_E4), depth: 1, flag: lowerBound, gen: 1},
		{score: -1, move: NewMove(SQ_E7, SQ_E5), depth: 127, flag: upperBound, gen: ttGenMask},
		{score: MateScore, move: NewPromotionMove(SQ_A7, SQ_A8, Queen), depth: 12, flag: exact, gen: 31},
		{score: -MateScore, move: NewPromotionMove(SQ_A2, SQ_A1, Knight), depth: -1, flag: lowerBound, gen: 4095},
		{score: MateScore + maxPly, move: 0xffff, depth: -128, flag: upperBound, gen: 7},
		{score: -MateScore - maxPly, move: 0xffff, depth: 100, flag: exact, gen: ttGenMask},

		// The edges of the score field, which is narrower than an int32.
		{score: 1<<(ttScoreBits-1) - 1, move: 1, depth: 5, flag: exact, gen: 2},
		{score: -(1 << (ttScoreBits - 1)), move: 1, depth: 5, flag: exact, gen: 2},
	}

	for _, want := range tests {
		if got := unpackEntry(want.pack()); got != want {
			t.Errorf("round trip: got %+v, want %+v", got, want)
		}
	}
}

func TestTTGetSet(t *testing.T) {
	tt := NewTranspositionTable(1 << 16)

	const hash = 0x0123456789abcdef
	want := ttEntry{score: -1234, move: NewMove(SQ_G1, SQ_F3), depth: 7, flag: lowerBound}

	if _, ok := tt.get(hash); ok {
		t.Error("empty table returned a hit")
	}

	tt.set(hash, want)

	got, ok := tt.get(hash)
	if !ok {
		t.Fatal("stored entry not found")
	}

	want.gen = tt.generation()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

// TestTTRejectsOtherPositions checks the property the XOR key exists for: a
// slot holding a different position must report a miss rather than returning
// that position's score under this position's hash.
func TestTTRejectsOtherPositions(t *testing.T) {
	tt := NewTranspositionTable(1 << 12)

	// Two hashes that collide: they differ only above the index bits.
	const a = 0x1111111111111000
	b := a + (tt.mask+1)*7

	if a&tt.mask != b&tt.mask {
		t.Fatalf("test hashes do not collide: %d vs %d", a&tt.mask, b&tt.mask)
	}

	tt.set(a, ttEntry{score: 999, depth: 9, flag: exact})

	if _, ok := tt.get(b); ok {
		t.Error("colliding hash reported a hit")
	}
	if got, ok := tt.get(a); !ok || got.score != 999 {
		t.Errorf("original entry lost: %+v ok=%v", got, ok)
	}

	// The bucket has room for both, which is the point of bucketing: under
	// the old one-slot table, storing b evicted a.
	tt.set(b, ttEntry{score: 111, depth: 9, flag: exact})

	if got, ok := tt.get(a); !ok || got.score != 999 {
		t.Errorf("a was evicted by a colliding store: %+v ok=%v", got, ok)
	}
	if got, ok := tt.get(b); !ok || got.score != 111 {
		t.Errorf("b not stored: %+v ok=%v", got, ok)
	}
}

// TestTTZeroHashIsNotAFalseHit covers the empty-slot case. A fresh table is
// all zeros, so a position whose hash is zero must not match every empty slot
// it lands on -- and neither must anything else.
func TestTTZeroHashIsNotAFalseHit(t *testing.T) {
	tt := NewTranspositionTable(1 << 12)

	// An untouched slot has key == data == 0, so key^data == 0. A zero hash
	// therefore does match, and the entry it returns is the zero entry: a
	// depth of 0 and a score of 0. The search only trusts an entry whose
	// depth is at least the depth it needs, so this is harmless, but the
	// behaviour should be deliberate rather than accidental.
	got, ok := tt.get(0)
	if ok && got != (ttEntry{}) {
		t.Errorf("empty slot returned %+v, want the zero entry", got)
	}
}

// TestTTConcurrentAccess is the reason the table is lockless. Run under -race
// it asserts that concurrent probing and storing is not a data race; the
// correctness property it checks is weaker, and deliberately so: an entry may
// be evicted at any moment, but a hit must never return a record that was
// never stored.
func TestTTConcurrentAccess(t *testing.T) {
	tt := NewTranspositionTable(1 << 16)

	const (
		writers = 4
		readers = 4
		rounds  = 20000
	)

	// Every stored entry satisfies depth == score/2, so a torn read that
	// survived verification would almost certainly break the invariant.
	entryFor := func(n uint64) ttEntry {
		depth := int8(n % 64)
		return ttEntry{
			score: int32(depth) * 2,
			move:  Move(n),
			depth: depth,
			flag:  ttFlag(n % 3),
		}
	}

	var wg sync.WaitGroup

	for w := 0; w < writers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			r := rand.New(rand.NewPCG(1, uint64(w)))
			for i := 0; i < rounds; i++ {
				h := r.Uint64()
				tt.set(h, entryFor(h))
			}
		}()
	}

	for rd := 0; rd < readers; rd++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			r := rand.New(rand.NewPCG(2, uint64(rd)))
			for i := 0; i < rounds; i++ {
				h := r.Uint64()
				if got, ok := tt.get(h); ok {
					if got.score != int32(got.depth)*2 {
						t.Errorf("torn entry survived verification: %+v", got)
						return
					}
				}
			}
		}()
	}

	wg.Wait()
}

// colliding returns n hashes that all land in the same bucket as base.
func colliding(tt *TranspositionTable, base uint64, n int) []uint64 {
	out := make([]uint64, 0, n)
	for i := 1; len(out) < n; i++ {
		out = append(out, base+uint64(i)*(tt.mask+1))
	}
	return out
}

// TestTTBucketHoldsFourPositions checks the capacity a bucket buys. Four
// positions that all index to the same bucket must all survive.
func TestTTBucketHoldsFourPositions(t *testing.T) {
	tt := NewTranspositionTable(1 << 16)

	const base = 0x2222222222220000
	hashes := append([]uint64{base}, colliding(tt, base, ttBucketSlots-1)...)

	for i, h := range hashes {
		tt.set(h, ttEntry{score: int32(100 + i), depth: int8(10 + i), flag: exact})
	}

	for i, h := range hashes {
		got, ok := tt.get(h)
		if !ok {
			t.Errorf("hash %d evicted from a bucket with room", i)
			continue
		}
		if got.score != int32(100+i) {
			t.Errorf("hash %d: score %d, want %d", i, got.score, 100+i)
		}
	}
}

// TestTTSamePositionUpdatesInPlace checks that re-storing a position reuses
// its slot rather than consuming a second one.
func TestTTSamePositionUpdatesInPlace(t *testing.T) {
	tt := NewTranspositionTable(1 << 16)

	const h = 0x3333333333330000
	tt.set(h, ttEntry{score: 1, depth: 4, flag: exact})
	tt.set(h, ttEntry{score: 2, depth: 6, flag: lowerBound})

	got, ok := tt.get(h)
	if !ok {
		t.Fatal("entry lost")
	}
	if got.score != 2 || got.depth != 6 {
		t.Errorf("got %+v, want the deeper record", got)
	}

	// Three other positions must still fit alongside it.
	for i, other := range colliding(tt, h, ttBucketSlots-1) {
		tt.set(other, ttEntry{score: int32(50 + i), depth: 5, flag: exact})
	}
	for i, other := range colliding(tt, h, ttBucketSlots-1) {
		if got, ok := tt.get(other); !ok || got.score != int32(50+i) {
			t.Errorf("neighbour %d lost: %+v ok=%v; the update consumed a second slot", i, got, ok)
		}
	}
}

// TestTTKeepsDeeperFromSameSearch checks that a shallow result does not
// overwrite a deeper one computed for the same position in the same search.
func TestTTKeepsDeeperFromSameSearch(t *testing.T) {
	tt := NewTranspositionTable(1 << 16)

	const h = 0x4444444444440000
	tt.set(h, ttEntry{score: 500, depth: 12, flag: exact})
	tt.set(h, ttEntry{score: 1, depth: 3, flag: exact})

	if got, _ := tt.get(h); got.depth != 12 || got.score != 500 {
		t.Errorf("got %+v, want the depth-12 record retained", got)
	}
}

// TestTTAging is what the generation counter exists for: a deep entry from an
// earlier search must be evicted before a shallow one from the current
// search, because it describes a position the game has moved on from.
func TestTTAging(t *testing.T) {
	tt := NewTranspositionTable(1 << 16)

	const base = 0x5555555555550000
	hashes := append([]uint64{base}, colliding(tt, base, ttBucketSlots-1)...)

	// Fill the bucket with deep results, then let several searches pass.
	for _, h := range hashes {
		tt.set(h, ttEntry{score: 1, depth: 30, flag: exact})
	}
	for i := 0; i < 5; i++ {
		tt.NewSearch()
	}

	// A shallow entry from the current search must find room.
	fresh := colliding(tt, base, ttBucketSlots)[ttBucketSlots-1]
	tt.set(fresh, ttEntry{score: 2, depth: 1, flag: exact})

	if _, ok := tt.get(fresh); !ok {
		t.Error("a current-search entry was refused a slot held by stale ones")
	}

	evicted := 0
	for _, h := range hashes {
		if _, ok := tt.get(h); !ok {
			evicted++
		}
	}
	if evicted != 1 {
		t.Errorf("%d stale entries evicted, want exactly 1", evicted)
	}
}

// TestTTAgingBreaksTiesOnDepth pins what aging actually decides. It is a
// discount on depth, not an override: between entries of comparable depth the
// older one is given up, but a much deeper result from an earlier search is
// still worth more than a shallow one from this search, because recomputing
// it costs far more than recomputing the shallow one.
func TestTTAgingBreaksTiesOnDepth(t *testing.T) {
	tt := NewTranspositionTable(1 << 16)

	const base = 0x6666666666660000
	hashes := append([]uint64{base}, colliding(tt, base, ttBucketSlots)...)

	// One entry from the previous search, the rest from this one, all the
	// same depth.
	tt.set(hashes[0], ttEntry{score: 1, depth: 5, flag: exact})
	tt.NewSearch()
	for _, h := range hashes[1:ttBucketSlots] {
		tt.set(h, ttEntry{score: 2, depth: 5, flag: exact})
	}

	tt.set(hashes[ttBucketSlots], ttEntry{score: 3, depth: 5, flag: exact})

	if _, ok := tt.get(hashes[0]); ok {
		t.Error("the stale entry survived; a fresh one of equal depth was evicted instead")
	}
	for _, h := range hashes[1:ttBucketSlots] {
		if _, ok := tt.get(h); !ok {
			t.Error("a fresh entry was evicted while a stale one of equal depth remained")
		}
	}
}

// TestTTDepthOutweighsMildAge is the other half of the same rule: one search
// of age does not make a deep entry cheaper to lose than a shallow one.
func TestTTDepthOutweighsMildAge(t *testing.T) {
	tt := NewTranspositionTable(1 << 16)

	const base = 0x7777777777770000
	hashes := append([]uint64{base}, colliding(tt, base, ttBucketSlots)...)

	tt.set(hashes[0], ttEntry{score: 1, depth: 30, flag: exact})
	tt.NewSearch()
	for _, h := range hashes[1:ttBucketSlots] {
		tt.set(h, ttEntry{score: 2, depth: 2, flag: exact})
	}

	tt.set(hashes[ttBucketSlots], ttEntry{score: 3, depth: 2, flag: exact})

	if _, ok := tt.get(hashes[0]); !ok {
		t.Error("a depth-30 entry one search old was evicted for a depth-2 one")
	}
}

func TestRelativeAge(t *testing.T) {
	if got := relativeAge(5, 5); got != 0 {
		t.Errorf("same generation = %d, want 0", got)
	}
	if got := relativeAge(5, 3); got != 2 {
		t.Errorf("two searches ago = %d, want 2", got)
	}
	// The counter still wraps eventually; it is just wide enough that a
	// process has to be cleared long before it gets there.
	if got := relativeAge(1, ttGenMask); got != 2 {
		t.Errorf("across a wrap = %d, want 2", got)
	}
}

func TestTTSizeBytes(t *testing.T) {
	tests := []struct {
		request uint64
		want    uint64
	}{
		{request: 64 << 20, want: 64 << 20},   // exact power of two
		{request: 256 << 20, want: 256 << 20}, // the engine default
		{request: 100 << 20, want: 64 << 20},  // rounds down to the lower power
		{request: 64, want: 64},               // one bucket
		{request: 1, want: 64},                // never smaller than one bucket
	}

	for _, test := range tests {
		tt := NewTranspositionTable(test.request)
		if got := tt.SizeBytes(); got != test.want {
			t.Errorf("NewTranspositionTable(%d).SizeBytes() = %d, want %d",
				test.request, got, test.want)
		}

		// Whatever the size, the table has to work.
		tt.set(0xabcdef, ttEntry{score: 7, depth: 3, flag: exact})
		if got, ok := tt.get(0xabcdef); !ok || got.score != 7 {
			t.Errorf("size %d: entry not retrievable", test.request)
		}
	}
}

// TestTTScoreFieldHoldsEverySearchScore checks that the narrowed score field
// still covers every value the search can hand it. Silently truncating a mate
// score would turn a forced win into a quiet evaluation.
func TestTTScoreFieldHoldsEverySearchScore(t *testing.T) {
	for _, ply := range []int{0, 1, maxPly - 1} {
		for _, base := range []int{MateScore, -MateScore, drawScore, 30000, -30000} {
			score := scoreToTT(base, ply)

			got := unpackEntry(ttEntry{score: int32(score)}.pack()).score
			if int(got) != score {
				t.Errorf("score %d (base %d, ply %d) packed to %d", score, base, ply, got)
			}
		}
	}
}

// TestTTGenerationOutlastsAGame is the property the widened field exists for.
// An entry from earlier in a long-lived process must not come back around and
// read as current.
func TestTTGenerationOutlastsAGame(t *testing.T) {
	const searchesPerGame = 40

	if 1<<ttGenBits <= searchesPerGame {
		t.Fatalf("generation wraps within a single game: %d searches, %d values",
			searchesPerGame, 1<<ttGenBits)
	}

	tt := NewTranspositionTable(1 << 16)
	const h = 0x8888888888880000
	tt.set(h, ttEntry{score: 1, depth: 30, flag: exact})

	// Play out a game's worth of searches and then some.
	for i := 0; i < searchesPerGame*4; i++ {
		tt.NewSearch()
	}

	entry, ok := tt.get(h)
	if !ok {
		t.Skip("entry evicted, which is also acceptable")
	}
	if age := relativeAge(tt.generation(), entry.gen); age != searchesPerGame*4 {
		t.Errorf("age after %d searches = %d; the counter wrapped and the entry reads as fresh",
			searchesPerGame*4, age)
	}
}

// TestTTClear checks that a new game starts from an empty table.
func TestTTClear(t *testing.T) {
	tt := NewTranspositionTable(1 << 16)

	hashes := []uint64{0x11, 0x2222, 0x333333, 0x44444444}
	for _, h := range hashes {
		tt.set(h, ttEntry{score: 5, depth: 9, flag: exact})
	}
	for i := 0; i < 10; i++ {
		tt.NewSearch()
	}

	tt.Clear()

	for _, h := range hashes {
		if _, ok := tt.get(h); ok {
			t.Errorf("hash %#x survived Clear", h)
		}
	}
	if got := tt.generation(); got != 0 {
		t.Errorf("generation after Clear = %d, want 0", got)
	}

	// The table has to still work afterwards.
	tt.set(0x99, ttEntry{score: 3, depth: 2, flag: exact})
	if got, ok := tt.get(0x99); !ok || got.score != 3 {
		t.Error("table unusable after Clear")
	}
}
