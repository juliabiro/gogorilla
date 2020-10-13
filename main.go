package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/juliabiro/gogorilla/gorilla"
	"log"
)

func main() {
	game := gorilla.Game{}
	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(gorilla.ScreenWidth, gorilla.ScreenHeight)
	ebiten.SetWindowTitle("Gorillas")
	game.Setup()
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
