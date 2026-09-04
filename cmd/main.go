package main

const (
	BotName = "Chester"
	Author  = "Mariano Wahlmann"
)

// version is the release this binary was built from. build.sh sets it at link
// time with -ldflags "-X main.version=..."; an ordinary go build leaves it as
// "dev".
var version = "dev"

func main() {
	startUCI()
}
