package main

import (
	"tiktaktoe/game"
)

type Session struct {
	game    game.Game
	players []string
}

func main() {
	game := game.Game{}
	game.Init(3)
}
