//go:build stockfish
// +build stockfish

package main_test

import (
	"bufio"
	"bytes"
	"fmt"
	"maps"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"

	"math/rand"

	pawned "github.com/bluescreen10/pawned"
)

func TestCompareToStockfish(t *testing.T) {
	maxDepth := 
	for depth := 1; depth <= maxDepth; depth++ {
		comparePawnedAndStockfish(t, depth, "")
	}
}

func comparePawnedAndStockfish(t *testing.T, depth int, moves string) {
	// Run Stockfish
	sfMoves, fen := sfPerft(depth, moves)

	// Run our engine
	p, err := pawned.Parse(fen)
	if err != nil {
		t.Fatal(err)
	}
	output := bytes.Buffer{}
	pawned.Perft(p, depth, &output)
	pawnedMoves := readOutput(bufio.NewReader(&output))

	// Compare the results
	diff := compareMaps(pawnedMoves, sfMoves)

	if len(diff) > 0 {
		move := pickRandomMove(diff)
		if depth < 2 {
			t.Fatalf(
				"move: %s, stockfish: %d, pawned: %d\n"+
					"moves: %s\n"+
					"fen: %s\n"+
					"%s\n",
				move, len(sfMoves), len(pawnedMoves), movesToString(sfMoves), fen, p)
		} else {
			comparePawnedAndStockfish(t, depth-1, moves+" "+move)
		}
	}
}

func sfPerft(depth int, moves string) (map[string]int, string) {
	stockfish := exec.Command("stockfish")

	sfIn, _ := stockfish.StdinPipe()
	sfOut, _ := stockfish.StdoutPipe()
	sfOutputReader := bufio.NewReader(sfOut)

	if err := stockfish.Start(); err != nil {
		panic(err)
	}
	defer stockfish.Process.Kill()

	if moves != "" {
		sfIn.Write([]byte(fmt.Sprintf("position startpos moves %s\n", moves)))
	} else {
		sfIn.Write([]byte("position startpos\n"))
	}

	sfIn.Write([]byte(fmt.Sprintf("d %d\n", depth)))
	fen := readFen(sfOutputReader)

	sfIn.Write([]byte(fmt.Sprintf("go perft %d\n", depth)))
	return readOutput(sfOutputReader), fen
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

func movesToString(moves map[string]int) string {
	var allMoves []string
	for m := range maps.Keys(moves) {
		allMoves = append(allMoves, m)
	}
	sort.Strings(allMoves)
	return "\"" + strings.Join(allMoves, "\", \"") + "\""
}

func pickRandomMove(moves map[string]int) string {
	var allMoves []string
	for m := range maps.Keys(moves) {
		allMoves = append(allMoves, m)
	}
	return allMoves[rand.Intn(len(allMoves))]
}

func compareMaps(map1, map2 map[string]int) map[string]int {
	diff := make(map[string]int)
	for move, count := range map1 {
		if count2, ok := map2[move]; ok {
			if count != count2 {
				diff[move] = count
			}
		} else {
			diff[move] = count
		}
	}
	for move, count := range map2 {
		if _, ok := map1[move]; !ok {
			diff[move] = count
		}
	}
	return diff
}
