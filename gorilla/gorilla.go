package gorilla

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/juliabiro/gogorilla/sprites"

	_ "image/png"
	"math/rand"
)

const (
	Right = iota
	Left
)

type Gorilla struct {
	sprites.Sprite
	score     int
	direction int
	alive     bool
}

func NewGorilla(x, y float64, w, h int, direction int) *Gorilla {
	g := Gorilla{*sprites.NewSprite(x, y, w, h), 0, direction, true}

	return &g
}

func (g *Gorilla) Alive() bool {
	return g.alive
}

func (g *Gorilla) Score() int {
	return g.score
}

func (g *Gorilla) Direction() int {
	return g.direction
}

func (g *Gorilla) Reset(b []Building, maxX int) {

	minx := 0
	if g.direction == Left {
		minx = maxX / 2
	}

	g.X = float64(minx + rand.Intn(3*maxX/10))

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
	g.Y = bb.Y - float64(g.Height)
	if g.X < bb.X || g.X+g.Width > bb.X+bb.width {
		g.X = bb.X + float64(rand.Intn(int(bb.width-g.Width)))
	}

}

func (g *Gorilla) Kill() {
	g.alive = false
}

func (g *Gorilla) revive() {
	g.alive = true
}

func (g *Gorilla) IncreaseScore() {
	g.score++
}

func (g *Gorilla) DrawingParamaters() (*ebiten.Image, *ebiten.DrawImageOptions) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(g.scaleX, g.scaleY)
	op.GeoM.Translate(float64(g.X), float64(g.Y))
	return g.Img.Image, op
}
