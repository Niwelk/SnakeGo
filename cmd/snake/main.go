package main

import (
	"snake_go/internal/game"
)

func main() {
	newGame := game.Create()
	newGame.Run()
}
