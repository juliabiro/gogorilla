package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
	"time"
)

const (
	screenWidth  = 1200
	screenHeight = 700

	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameNum    = 8
)

// Game implements ebiten.Game interface.
type Game struct {
	gorilla1  Gorilla
	gorilla2  Gorilla
	buildings []Building
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

	for _, b := range g.buildings {
		b.Draw(screen)
	}
	g.gorilla1.Draw(screen)
	g.gorilla2.Draw(screen)
	// Write your game's rendering.
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game := &Game{}
	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(screenWidth, screenHeight)
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

type Building struct {
	Point
	img    *ebiten.Image
	height int
	width  int
	color  color.Color
}

func (b *Building) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(float64(b.X), float64(b.Y))
	b.img.Fill(b.color)

	screen.DrawImage(b.img, op)
}

func setupBuildings(g *Game) {
	k := 0
	for k < screenWidth {
		w := 100 + rand.Intn(screenWidth/12)
		h := 150 + rand.Intn(screenHeight/2)
		if k+w >= screenWidth {
			w = screenWidth - k
		}
		img, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
		g.buildings = append(g.buildings, Building{Point{k, screenHeight - h}, img, h, w, color.RGBA{0, 0, 100 + uint8(rand.Intn(155)), 255}})
		k = k + w
	}
}

func (g *Gorilla) init(minx int, b []Building) {

	g.alive = true
	var err error
	g.img, _, err = ebitenutil.NewImageFromFile("/Users/juliabiro/go/gorilla/gorilla.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	g.width = 50
	g.height = 50
	g.X = minx + rand.Intn(0.6*screenWidth/2)

	// find my rooftop
	i := 0
	for g.X > b[i].X+b[i].width {
		i++
	}
	bb := b[i]

	// make sure I sit on it
	g.Y = bb.Y - g.height
	if g.X < bb.X || g.X+g.width > bb.X+bb.width {
		g.X = bb.X + rand.Intn(bb.width-g.width)
	}
}
func setupGorillas(g *Game) {
	g.gorilla1.init(0, g.buildings)
	g.gorilla2.init(screenWidth/2, g.buildings)
}

func setup(g *Game) {
	setupBuildings(g)
	setupGorillas(g)
}
