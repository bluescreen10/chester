package chester

import (
	"sync/atomic"
	"unsafe"
)

// ttFlag represents the type of score stored in a transposition table entry.
type ttFlag uint8

const (
	exact      ttFlag = iota // Score is an exact value (PV-node)
	upperBound               // Score is an upper bound (Alpha-node)
	lowerBound               // Score is a lower bound (Beta-node)
)

// ttEntry is a decoded transposition table record. It is what callers work
// with; the table itself stores each record packed into two 64-bit words, see
// ttSlot.
type ttEntry struct {
	// score is the evaluation score found during search. Mate scores are
	// stored relative to the node they were found at, see scoreToTT.
	score int32

	// move is the best move found at this position, or the move that caused
	// a beta cutoff. It is used to order moves even when the stored depth is
	// too shallow to allow a cutoff, which is where most of its value lies.
	move Move

	// depth is the remaining search depth when this score was recorded.
	depth int8

	// flag indicates whether the score is exact, an upper bound, or a lower bound.
	flag ttFlag

	// gen is the search generation the entry was written in. It is what makes
	// an entry from three moves ago cheaper to evict than a shallow one from
	// the search now running.
	gen uint16
}

// Bit layout of the packed payload word.
//
// The generation gets fourteen bits rather than the handful it needs to tell
// one search from the next, because the counter wrapping is worse than
// useless: an entry exactly one wrap old computes an age of zero and reads as
// brand new, so the rule meant to evict stale entries protects them instead.
// An engine process can live for hundreds of games, so the field has to
// outlast one.
//
// The bits come from the score, which was wildly over-provisioned at
// thirty-two. Stored scores never exceed a mate score plus a ply count, about
// a million, and twenty-four bits signed reaches eight times that.
const (
	ttGenBits   = 14
	ttFlagBits  = 2
	ttDepthBits = 8
	ttMoveBits  = 16
	ttScoreBits = 24

	ttGenShift   = 0
	ttFlagShift  = ttGenShift + ttGenBits
	ttDepthShift = ttFlagShift + ttFlagBits
	ttMoveShift  = ttDepthShift + ttDepthBits
	ttScoreShift = ttMoveShift + ttMoveBits

	ttGenMask   = 1<<ttGenBits - 1
	ttFlagMask  = 1<<ttFlagBits - 1
	ttScoreMask = 1<<ttScoreBits - 1
)

// pack folds an entry into a single 64-bit word.
func (e ttEntry) pack() uint64 {
	return uint64(uint32(e.score))&ttScoreMask<<ttScoreShift |
		uint64(e.move)<<ttMoveShift |
		uint64(uint8(e.depth))<<ttDepthShift |
		uint64(e.flag&ttFlagMask)<<ttFlagShift |
		uint64(e.gen&ttGenMask)<<ttGenShift
}

// unpackEntry is the inverse of pack. The score is sign extended out of its
// twenty-four bit field.
func unpackEntry(data uint64) ttEntry {
	score := int32(data >> ttScoreShift & ttScoreMask)
	if score >= 1<<(ttScoreBits-1) {
		score -= 1 << ttScoreBits
	}

	return ttEntry{
		score: score,
		move:  Move(data >> ttMoveShift),
		depth: int8(data >> ttDepthShift),
		flag:  ttFlag(data>>ttFlagShift) & ttFlagMask,
		gen:   uint16(data>>ttGenShift) & ttGenMask,
	}
}

// ttSlot is one record as it is actually stored.
//
// The table is lockless. Rather than guarding a slot, a writer stores the
// position hash XORed with the payload, and a reader recovers the hash by
// XORing the two words back together. A reader that catches a writer halfway
// through -- seeing one word from the old record and one from the new --
// computes a hash that matches neither position, so the probe reports a miss.
// A torn read is therefore self-detecting and costs nothing but a lookup.
//
// This is what makes the table safe to share between search threads, and it
// is why the entry carries no hash field of its own: the key word doubles as
// both the identity and the checksum.
type ttSlot struct {
	// key is hash ^ data.
	key atomic.Uint64

	// data is the packed ttEntry.
	data atomic.Uint64
}

// read returns the hash the slot belongs to and its decoded entry. An empty
// slot reads back as hash zero, which no probe will match in practice.
func (s *ttSlot) read() (uint64, ttEntry) {
	key := s.key.Load()
	data := s.data.Load()

	return key ^ data, unpackEntry(data)
}

func (s *ttSlot) write(hash uint64, e ttEntry) {
	data := e.pack()
	s.key.Store(hash ^ data)
	s.data.Store(data)
}

// ttBucketSlots is how many entries share a bucket. Four sixteen-byte entries
// fill one sixty-four byte cache line, so probing all of them costs a single
// memory access -- the same price the old one-entry-per-index table paid for
// looking at one.
const ttBucketSlots = 4

// ttBucket is the unit the table indexes. Giving each hash four places to live
// is what lets a deep entry survive a collision: with a single slot, an entry
// that cost a million nodes to compute was evicted by whatever landed on it
// next, however shallow.
type ttBucket struct {
	slots [ttBucketSlots]ttSlot
}

// TranspositionTable is a lockless hash table used to store and retrieve
// search results for previously visited positions. It helps avoid redundant
// work by providing instant lookups for known positions at equal or greater
// depth.
type TranspositionTable struct {
	buckets []ttBucket

	// mask indexes the table. The bucket count is a power of two, so a
	// bitwise and replaces a modulo -- which would otherwise be a hardware
	// division on every probe, in the hottest lookup the engine has.
	mask uint64

	// gen counts searches. Only its low ttGenBits are stored per entry.
	gen atomic.Uint32
}

// NewTranspositionTable creates a new transposition table with the given
// maximum size in bytes. The bucket count is rounded down to the nearest
// power of two to allow for efficient indexing.
func NewTranspositionTable(maxSize uint64) *TranspositionTable {
	count := maxSize / uint64(unsafe.Sizeof(ttBucket{}))

	// Round down to power of 2
	size := uint64(1)
	for size*2 <= count {
		size *= 2
	}

	return &TranspositionTable{
		buckets: make([]ttBucket, size),
		mask:    size - 1,
	}
}

// NewSearch advances the generation counter, marking everything already in
// the table as belonging to an earlier search. Call it once before each root
// search: without it the table cannot tell a result computed for the position
// on the board from one computed twenty moves ago.
func (tt *TranspositionTable) NewSearch() {
	if tt != nil {
		tt.gen.Add(1)
	}
}

func (tt *TranspositionTable) generation() uint16 {
	return uint16(tt.gen.Load()) & ttGenMask
}

// relativeAge returns how many searches ago gen was current. The subtraction
// is taken modulo the field's range, which is wide enough that a wrap needs
// more searches than an engine process will see between table clears.
func relativeAge(now, gen uint16) int {
	return int((now - gen) & ttGenMask)
}

// worth ranks an entry for retention: the higher the number, the more it
// costs to lose. Depth is what an entry is worth, and age is what discounts
// it -- a deep result for a position three moves ago is less use than a
// shallow one for the position actually on the board.
func worth(e ttEntry, now uint16) int {
	return int(e.depth) - 8*relativeAge(now, e.gen)
}

// get retrieves the entry stored for the given hash, and reports whether one
// was found. A bucket holding only other positions, an empty bucket, and a
// slot caught mid-write all report false.
func (tt *TranspositionTable) get(hash uint64) (ttEntry, bool) {
	bucket := &tt.buckets[hash&tt.mask]

	for i := range bucket.slots {
		if h, entry := bucket.slots[i].read(); h == hash {
			return entry, true
		}
	}

	return ttEntry{}, false
}

// set stores an entry for the given hash, choosing which of the bucket's
// slots to give up.
//
// A slot already holding this position is always the one to reuse; keeping
// two records for one position would waste a quarter of the bucket. That
// record is left alone only when it comes from the current search and was
// computed deeper, since replacing it would throw away the better result.
//
// Otherwise the least valuable slot is evicted, where value is depth
// discounted by age. Empty slots read back as depth zero from generation
// zero, so they are evicted first without needing a special case.
func (tt *TranspositionTable) set(hash uint64, entry ttEntry) {
	bucket := &tt.buckets[hash&tt.mask]

	now := tt.generation()
	entry.gen = now

	victim := 0
	victimWorth := 1 << 30

	for i := range bucket.slots {
		h, stored := bucket.slots[i].read()

		if h == hash {
			if stored.gen == now && stored.depth > entry.depth {
				return
			}
			bucket.slots[i].write(hash, entry)
			return
		}

		if w := worth(stored, now); w < victimWorth {
			victim, victimWorth = i, w
		}
	}

	bucket.slots[victim].write(hash, entry)
}

// Clear empties the table. A new game shares no positions with the last one,
// so its entries are only crowding out the ones about to be needed -- and
// clearing also puts a bound on how far the generation counter can run.
func (tt *TranspositionTable) Clear() {
	for i := range tt.buckets {
		for j := range tt.buckets[i].slots {
			tt.buckets[i].slots[j].key.Store(0)
			tt.buckets[i].slots[j].data.Store(0)
		}
	}
	tt.gen.Store(0)
}

// SizeBytes returns the memory the table actually occupies. This is the size
// it was asked for rounded down to a power-of-two bucket count, so it can be
// as little as half the request.
func (tt *TranspositionTable) SizeBytes() uint64 {
	return uint64(len(tt.buckets)) * uint64(unsafe.Sizeof(ttBucket{}))
}
