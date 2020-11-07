package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/juliabiro/gogorilla/game"
	"log"
)

func main() {
	game := game.Game{}
	// Sepcify the window size as you like. Here, a doulbed size is specified.
	sw, sh := game.Layout(0, 0)
	ebiten.SetWindowSize(sw, sh)
	ebiten.SetWindowTitle("Gorillas")
	game.Setup()
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
