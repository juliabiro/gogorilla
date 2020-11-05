package gorilla

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	_ "image/png"
	"log"
	"math/rand"
)

const (
	right = iota
	left
)

type Gorilla struct {
	Point
	alive     bool
	img       ScaledImage
	height    float64
	width     float64
	score     int
	direction int
}

func (g *Gorilla) LoadImage() {
	file := imageDir + "gorilla1.png"
	if g.direction == right {
		file = imageDir + "gorilla2.png"
	}
	var err error

	img, _, err := ebitenutil.NewImageFromFile(file, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	g.img = ScaledImage{img, float64(g.width) / float64(img.Bounds().Dx()), float64(g.height) / float64(img.Bounds().Dy())}
}

func NewGorilla(direction int) *Gorilla {
	g := Gorilla{}
	g.alive = true
	g.width = 50
	g.height = 50
	g.direction = direction

	return &g
}

func (g *Gorilla) Center() (X, Y float64) {
	return g.X + g.width/2, g.Y + g.height/2
}

func (g *Gorilla) reset(b []Building) {

	minx := 0
	if g.direction == left {
		minx = ScreenWidth / 2
	}

	g.X = float64(minx + rand.Intn(0.6*ScreenWidth/2))

	g.sitOnRooftop(b)
	g.revive()
}

func (g *Gorilla) sitOnRooftop(b []Building) {
	// find my rooftop
	i := 0
	for g.X > b[i].X+float64(b[i].width) {
		i++
	}
	bb := b[i]

	// make sure I sit on it
	g.Y = bb.Y - float64(g.height)
	if g.X < bb.X || g.X+g.width > bb.X+bb.width {
		g.X = bb.X + float64(rand.Intn(int(bb.width-g.width)))
	}

}

func (g *Gorilla) kill() {
	g.alive = false
}

func (g *Gorilla) revive() {
	g.alive = true
}

func (g *Gorilla) increaseScore() {
	g.score++
}
