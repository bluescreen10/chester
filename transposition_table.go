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
}

// Bit layout of the packed payload word. The eight bits given to the flag
// leave room for a generation counter without changing the encoding again.
const (
	ttFlagBits  = 8
	ttDepthBits = 8
	ttMoveBits  = 16

	ttFlagShift  = 0
	ttDepthShift = ttFlagShift + ttFlagBits
	ttMoveShift  = ttDepthShift + ttDepthBits
	ttScoreShift = ttMoveShift + ttMoveBits
)

// pack folds an entry into a single 64-bit word.
func (e ttEntry) pack() uint64 {
	return uint64(uint32(e.score))<<ttScoreShift |
		uint64(e.move)<<ttMoveShift |
		uint64(uint8(e.depth))<<ttDepthShift |
		uint64(e.flag)<<ttFlagShift
}

// unpackEntry is the inverse of pack.
func unpackEntry(data uint64) ttEntry {
	return ttEntry{
		score: int32(data >> ttScoreShift),
		move:  Move(data >> ttMoveShift),
		depth: int8(data >> ttDepthShift),
		flag:  ttFlag(data >> ttFlagShift),
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

// TranspositionTable is a lockless hash table used to store and retrieve
// search results for previously visited positions. It helps avoid redundant
// work by providing instant lookups for known positions at equal or greater
// depth.
type TranspositionTable struct {
	slots []ttSlot

	// mask indexes the table. The slot count is a power of two, so a bitwise
	// and replaces a modulo -- which would otherwise be a hardware division
	// on every probe, in the hottest lookup the engine has.
	mask uint64
}

// NewTranspositionTable creates a new transposition table with the given
// maximum size in bytes. The actual number of entries is rounded down to
// the nearest power of two to allow for efficient indexing.
func NewTranspositionTable(maxSize uint64) *TranspositionTable {
	count := maxSize / uint64(unsafe.Sizeof(ttSlot{}))

	// Round down to power of 2
	size := uint64(1)
	for size*2 <= count {
		size *= 2
	}

	return &TranspositionTable{
		slots: make([]ttSlot, size),
		mask:  size - 1,
	}
}

// get retrieves the entry stored for the given hash, and reports whether one
// was found. A slot holding a different position, an empty slot, and a slot
// caught mid-write all report false.
func (tt *TranspositionTable) get(hash uint64) (ttEntry, bool) {
	slot := &tt.slots[hash&tt.mask]

	key := slot.key.Load()
	data := slot.data.Load()

	if key^data != hash {
		return ttEntry{}, false
	}

	return unpackEntry(data), true
}

// set stores an entry for the given hash, overwriting whatever occupied the
// slot. The key is written first so that the two words are only ever
// consistent with each other when both belong to the same record.
func (tt *TranspositionTable) set(hash uint64, entry ttEntry) {
	slot := &tt.slots[hash&tt.mask]

	data := entry.pack()
	slot.key.Store(hash ^ data)
	slot.data.Store(data)
}
