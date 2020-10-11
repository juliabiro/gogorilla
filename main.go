package main

import (
	"fmt"
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

type Windows struct {
	color            color.Color
	width            int
	height           int
	borderHorizontal int
	borderVertical   int
	img              *ebiten.Image
	lightsOffColor   color.Color
	lightsOff        map[string]int
}

func (w *Windows) Draw(screen *ebiten.Image, b *Building) {
	op := &ebiten.DrawImageOptions{}

	scaleX := float64(w.width-2*w.borderHorizontal) / float64(w.width)
	scaleY := float64(w.height-2*w.borderVertical) / float64(w.height)
	for i := 1; i*w.width-w.borderHorizontal < b.width; i++ {
		for j := 1; j*w.height-w.borderVertical < b.height; j++ {
			op.GeoM.Reset()
			op.GeoM.Scale(scaleX, scaleY)
			if w.lightsOff[fmt.Sprintf("%s,%s", i, j)] == 1 {
				w.img.Fill(w.lightsOffColor)
			} else {
				w.img.Fill(w.color)
			}
			op.GeoM.Translate(float64(b.X+(i-1)*w.width+w.borderHorizontal), float64(b.Y+(j-1)*w.height+w.borderVertical))
			screen.DrawImage(w.img, op)
		}
	}
}

type Building struct {
	Point
	img     *ebiten.Image
	height  int
	width   int
	color   color.Color
	windows Windows
}

func (b *Building) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(float64(b.X), float64(b.Y))
	b.img.Fill(b.color)
	screen.DrawImage(b.img, op)
	b.windows.Draw(screen, b)
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
		c := color.RGBA{0, 0, 100 + uint8(rand.Intn(155)), 255}
		wc := color.RGBA{100 + uint8(rand.Intn(155)), 100 + uint8(rand.Intn(155)), 0, 255}
		locrand := uint8(rand.Intn(55))
		loc := color.RGBA{100 + locrand, 100 + locrand, 100 + locrand, 255}
		ww := w / (5 + rand.Intn(8))
		bh := ww * (rand.Intn(15) + 8) / 50
		wh := h / (5 + rand.Intn(15))
		bv := wh * (rand.Intn(15) + 8) / 50
		wimg, _ := ebiten.NewImage(ww, wh, ebiten.FilterDefault)
		loff := make(map[string]int)
		for i := 0; i*ww < w; i++ {
			for j := 0; j*wh < h; j++ {
				if rand.Intn(10) < 2 {
					loff[fmt.Sprintf("%s,%s", i, j)] = 1
				}
			}
		}
		windows := Windows{wc, ww, wh, bh, bv, wimg, loc, loff}
		g.buildings = append(g.buildings, Building{Point{k, screenHeight - h}, img, h, w, c, windows})
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
