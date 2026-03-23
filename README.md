# Chester

<p align="center">
  <img src="logo/logo.png"/>
</p>

A chess library and engine written in Go.

## Library Features

- Fast legal move generation using bitboards
- FEN (Forsyth–Edwards Notation) parsing and serialization
- Magic bitboard sliding piece attack lookup
- Zobrist hashing (Polyglot-compatible)
- Perft for move generation testing and benchmarking

## Engine Features

- Universal Chess Interface (UCI)
- Minimax with Alpha-Beta pruning
- Opening book support (Polyglot `.bin` format)

## Installation

```bash
go get github.com/bluescreen10/chester
```

## Usage

### As a library

```go
import "github.com/bluescreen10/chester"

// Parse a position from FEN
pos, err := chester.ParseFEN(chester.DefaultFEN)
if err != nil {
    panic(err)
}

// Generate legal moves
var moves []chester.Move
moves, inCheck := chester.LegalMoves(moves, &pos)

// Apply a move
pos.Do(moves[0])

// Serialize back to FEN
fmt.Println(pos.FEN())

// Perft
results := chester.Perft(&pos, 6)
var nodes int
for r := range results {
    nodes += r.Count
}
fmt.Printf("nodes: %d\n", nodes)
```

### As an engine (UCI)

```bash
go run ./cmd
```

```
position startpos
perft 6

a2a3: 4463267
b2b3: 5310358
c2c3: 5417640
d2d3: 8073082
e2e3: 9726018
f2f3: 4404141
g2g3: 5346260
h2h3: 4463070
a2a4: 5363555
b2b4: 5293555
c2c4: 5866666
d2d4: 8879566
e2e4: 9771632
f2f4: 4890429
g2g4: 5239875
h2h4: 5385554
b1a3: 4856835
b1c3: 5708064
g1f3: 5723523
g1h3: 4877234
perft 119060324 in 363.214708ms
```

(The above results were produced in a Macbook Pro M2)

## References
- [Chess Programming Wiki](https://www.chessprogramming.org)
- [Magic Bitboards](https://www.chessprogramming.org/Magic_Bitboards)
- [Polyglot Opening Book Format](http://hgm.nubati.net/book_format.html)