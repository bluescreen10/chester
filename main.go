package main

const (
	DefaultFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	BotName    = "Chester"
	Author     = "Mariano Wahlmann"
)

func main() {
	//fmt.Println(unsafe.Sizeof(Position{}))
	startUCI()
}
