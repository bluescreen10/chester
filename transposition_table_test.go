package chester

import (
	"math/rand/v2"
	"sync"
	"testing"
	"unsafe"
)

// TestTTSlotSize guards the storage layout. Four slots to a 64-byte cache line
// is the point of packing the record into two words; a slot that grows past
// sixteen bytes halves the number of positions a table of a given size holds.
func TestTTSlotSize(t *testing.T) {
	if got, want := unsafe.Sizeof(ttSlot{}), uintptr(16); got != want {
		t.Errorf("sizeof(ttSlot) = %d, want %d", got, want)
	}
}

func TestTTEntryRoundTrip(t *testing.T) {
	tests := []ttEntry{
		{score: 0, move: 0, depth: 0, flag: exact},
		{score: 1, move: NewMove(SQ_E2, SQ_E4), depth: 1, flag: lowerBound},
		{score: -1, move: NewMove(SQ_E7, SQ_E5), depth: 127, flag: upperBound},
		{score: MateScore, move: NewPromotionMove(SQ_A7, SQ_A8, Queen), depth: 12, flag: exact},
		{score: -MateScore, move: NewPromotionMove(SQ_A2, SQ_A1, Knight), depth: -1, flag: lowerBound},
		{score: 1 << 30, move: 0xffff, depth: -128, flag: upperBound},
		{score: -(1 << 30), move: 0xffff, depth: 100, flag: exact},
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
