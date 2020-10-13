package gorilla

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	_ "image/png"
	"log"
	"math"
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

func (b *Banana) detectCollisionGorilla(g Gorilla) bool {
	return math.Sqrt(math.Pow(b.X+b.width/2-(g.X+g.width/2), 2)+math.Pow(b.Y+b.height/2-(g.Y+g.height/2), 2)) < 25
}

// TODO turn into factory function
func (g *Gorilla) setup() {
	g.alive = true
	g.width = 50
	g.height = 50
	var err error
	img, _, err := ebitenutil.NewImageFromFile("./gorilla.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	g.img = ScaledImage{img, float64(g.width) / float64(img.Bounds().Dx()), float64(g.height) / float64(img.Bounds().Dy())}
}
func (g *Gorilla) reset(minx int, b []Building, direction int) {

	g.X = float64(minx + rand.Intn(0.6*ScreenWidth/2))

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
	g.direction = direction
	g.alive = true
}
