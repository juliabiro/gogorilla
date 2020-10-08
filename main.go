package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	_ "image/png"
	"math/rand"
)

const (
	screenWidth  = 640
	screenHeight = 480

	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameNum    = 8
)

// Game implements ebiten.Game interface.
type Game struct {
	gorilla1 Gorilla
	gorilla2 Gorilla
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(screen *ebiten.Image) error {
	// Write your game's logical update.
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.gorilla1.Draw(screen)
	g.gorilla2.Draw(screen)
	// Write your game's rendering.
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	game := &Game{}
	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Gorillas")
	setup(game)
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type Point struct {
	X int
	Y int
}
type Gorilla struct {
	Point
	alive  bool
	img    *ebiten.Image
	height int
	width  int
}

func (g *Gorilla) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.1, 0.1)
	op.GeoM.Translate(float64(g.X), float64(g.Y))
	screen.DrawImage(g.img, op)
}

func setup(g *Game) {
	g.gorilla1 = Gorilla{Point{rand.Intn(200), 200 + rand.Intn(200)}, true, nil, 50, 50}
	g.gorilla2 = Gorilla{Point{300 + rand.Intn(200), 200 + rand.Intn(200)}, true, nil, 50, 50}

	var err error
	g.gorilla1.img, _, err = ebitenutil.NewImageFromFile("/Users/juliabiro/go/gorilla/gorilla.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	g.gorilla2.img, _, err = ebitenutil.NewImageFromFile("/Users/juliabiro/go/gorilla/gorilla.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}
