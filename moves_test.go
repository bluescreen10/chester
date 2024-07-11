package main_test

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"testing"

	pawned "github.com/bluescreen10/pawned"
)

func TestMoveGen(t *testing.T) {
	p, err := pawned.Parse(pawned.DefaultFEN)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		depth int
		nodes int
	}{
		{1, 20},
		{2, 400},
		{3, 8902},
		{4, 197281},
		{5, 4865609},
		{6, 119060324},
		{7, 3195901860},
	}

	t.
	for _, test := range tests {
		if got := pawned.Perft(p, test.depth, false, os.Stdout); got != test.nodes {
			t.Fatalf("Perft(%d) = %d, want %d", test.depth, got, test.nodes)
		}
	}

	// depth := 5
	// fmt.Printf("Node(depth = %d): %d\n", depth, perft(p, depth, true))
	t.Fatal("not implemented")
}

func TestCompareToStockfish(t *testing.T) {
	depth := 6

	// Launch stockfish
	stockfish := exec.Command("stockfish")
	sfIn, _ := stockfish.StdinPipe()
	sfOut, _ := stockfish.StdoutPipe()
	sfOutputReader := bufio.NewReader(sfOut)

	stockfish.Start()

	//create a buffer to capture the output
	output := bytes.Buffer{}
	var moves string

	for i := depth; i > 0; i-- {
		if moves != "" {
			sfIn.Write([]byte(fmt.Sprintf("position startpos moves %s\n", moves)))
		} else {
			sfIn.Write([]byte("position startpos\n"))
		}

		sfIn.Write([]byte(fmt.Sprintf("d %d\n", i)))
		fen := readFen(sfOutputReader)

		sfIn.Write([]byte(fmt.Sprintf("go perft %d\n", i)))
		sfMoves := readOutput(sfOutputReader)

		// Run our engine
		p, err := pawned.Parse(fen)
		if err != nil {
			t.Fatal(err)
		}
		pawned.Perft(p, i, true, &output)
		pawnedMoves := readOutput(bufio.NewReader(&output))

		// Compare the results
		diff := make(map[string]any)
		for move, count := range sfMoves {
			if pawnedCount, ok := pawnedMoves[move]; ok {
				if count != pawnedCount {
					t.Errorf("Fen: %s depth: %d move %s: Stockfish: %d, Pawned: %d\n", fen, i, move, count, pawnedCount)
					t.Logf("%s\n", p)
					diff[move] = struct{}{}
				}
			} else {
				t.Errorf("Fen: %s depth: %d move %s not found in Pawned\n", fen, i, move)
				t.Logf("%s\n", p)
				diff[move] = struct{}{}
			}
		}

		for move := range pawnedMoves {
			if _, ok := sfMoves[move]; !ok {
				t.Errorf("Fen: %s depth: %d move %s not found in Stockfish\n", fen, i, move)
				t.Logf("\n%s", p)
				diff[move] = struct{}{}
			}
		}

		// pick the first move that is different
		for move := range diff {
			moves += " " + move
			break
		}

	}

	stockfish.Process.Kill()
}

func readOutput(reader *bufio.Reader) map[string]int {
	terminator := regexp.MustCompile(`Node.*: \d+`)
	move := regexp.MustCompile(`([a-h][1-8][a-h][1-8]): (\d+)`)

	moves := make(map[string]int)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		if matches := move.FindAllSubmatch(line, 1); matches != nil {
			count, _ := strconv.Atoi(string(matches[0][2]))
			moves[string(matches[0][1])] = count
		}

		if terminator.Match(line) {
			break
		}
	}
	return moves
}

func readFen(reader *bufio.Reader) string {
	fen := regexp.MustCompile(`Fen:\s+(.*)`)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		if matches := fen.FindAllSubmatch(line, 1); matches != nil {
			return string(matches[0][1])
		}
	}
	panic("Fen not found")
}
