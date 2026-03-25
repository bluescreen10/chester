package chester

import (
	"sync"
	"unsafe"
)

type ttFlag uint8

const (
	exact      ttFlag = iota // Exact score (PV-Node)
	upperBound               // Score is <= alpha (All-Node)
	lowerBound               // Score is >= beta (Cut-Node)
)

type ttEntry struct {
	// Full hash to verify no collisions
	hash uint64

	// The evaluation score
	score int

	// How deep the search was
	depth int

	// Type of node (Exact, Alpha, or Beta)
	flag ttFlag
}

type TranspositionTable struct {
	entries []ttEntry
	size    uint64
	mu      sync.RWMutex
}

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

func (tt *TranspositionTable) get(hash uint64) ttEntry {
	tt.mu.RLock()
	entry := tt.entries[hash%tt.size]
	tt.mu.RUnlock()

	return entry
}

func (tt *TranspositionTable) set(entry ttEntry) {
	tt.mu.Lock()
	tt.entries[entry.hash%tt.size] = entry
	tt.mu.Unlock()
}
