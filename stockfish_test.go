//go:build stockfish
// +build stockfish

package main_test

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"testing"

	"math/rand"

	pawned "github.com/bluescreen10/pawned"
)

func TestCompareToStockfish(t *testing.T) {
	depth := 8

	// Launch stockfish
	stockfish := exec.Command("stockfish")
	sfIn, _ := stockfish.StdinPipe()
	sfOut, _ := stockfish.StdoutPipe()
	sfOutputReader := bufio.NewReader(sfOut)

	if err := stockfish.Start(); err != nil {
		t.Fatal(err)
	}

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
		pawned.Perft(p, i, &output)
		pawnedMoves := readOutput(bufio.NewReader(&output))

		// Compare the results
		diff := make(map[string]any)
		for move, count := range sfMoves {
			if pawnedCount, ok := pawnedMoves[move]; ok {
				if count != pawnedCount {
					t.Errorf("Fen: %s depth: %d move %s: Stockfish: %d, Pawned: %d\n", fen, i, move, count, pawnedCount)
					t.Logf("\n%s\n", p)
					diff[move] = struct{}{}
				}
			} else {
				t.Errorf("Fen: %s depth: %d move %s not found in Pawned\n", fen, i, move)
				t.Logf("\n%s\n", p)
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

		// pick a random move that is different
		if len(diff) > 0 {
			// obtain the keys
			keys := make([]string, 0, len(diff))
			for k := range diff {
				keys = append(keys, k)
			}
			moves += " " + keys[rand.Intn(len(keys))]
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
