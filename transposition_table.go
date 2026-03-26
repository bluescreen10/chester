package chester

import (
	"sync"
	"unsafe"
)

// ttFlag represents the type of score stored in a transposition table entry.
type ttFlag uint8

const (
	exact      ttFlag = iota // Score is an exact value (PV-node)
	upperBound               // Score is an upper bound (Alpha-node)
	lowerBound               // Score is a lower bound (Beta-node)
)

// ttEntry represents a single record in the transposition table.
type ttEntry struct {
	// hash is the full Zobrist hash of the position to verify no collisions.
	hash uint64

	// score is the evaluation score found during search.
	score int

	// depth is the remaining search depth when this score was recorded.
	depth int

	// flag indicates whether the score is exact, an upper bound, or a lower bound.
	flag ttFlag
}

// TranspositionTable is a thread-safe hash table used to store and retrieve
// search results for previously visited positions. It helps avoid redundant
// work by providing instant lookups for known positions at equal or greater
// depth.
type TranspositionTable struct {
	entries []ttEntry
	size    uint64
	mu      sync.RWMutex
}

// NewTranspositionTable creates a new transposition table with the given
// maximum size in bytes. The actual number of entries is rounded down to
// the nearest power of two to allow for efficient indexing.
func NewTranspositionTable(maxSize uint64) *TranspositionTable {
	count := maxSize / uint64(unsafe.Sizeof(ttEntry{}))

	// Round down to power of 2
	size := uint64(1)
	for size*2 <= uint64(count) {
		size *= 2
	}

	return &TranspositionTable{
		entries: make([]ttEntry, size),
		size:    size,
	}
}

// get retrieves the entry associated with the given hash from the table.
// If no entry is found or a collision occurs, it returns a zero-valued entry.
func (tt *TranspositionTable) get(hash uint64) ttEntry {
	tt.mu.RLock()
	entry := tt.entries[hash%tt.size]
	tt.mu.RUnlock()

	return entry
}

// set stores a new entry in the transposition table, overwriting any
// existing entry at the same index.
func (tt *TranspositionTable) set(entry ttEntry) {
	tt.mu.Lock()
	tt.entries[entry.hash%tt.size] = entry
	tt.mu.Unlock()
}
