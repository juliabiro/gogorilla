package gorilla

import (
	"github.com/hajimehoshi/ebiten"
)

const (
	ScreenWidth  = 1200
	ScreenHeight = 700
)

func (g *Gorilla) DrawingParamaters() (*ebiten.Image, *ebiten.DrawImageOptions) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(g.img.scaleX, g.img.scaleY)
	op.GeoM.Translate(float64(g.X), float64(g.Y))
	return g.img.Image, op
}
