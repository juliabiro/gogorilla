package gorilla

import (
	"github.com/hajimehoshi/ebiten"
)

const (
	ScreenWidth  = 1200
	ScreenHeight = 700
)

func DrawGorilla(screen *ebiten.Image, g Gorilla) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(g.img.scaleX, g.img.scaleY)
	op.GeoM.Translate(float64(g.X), float64(g.Y))
	screen.DrawImage(g.img.Image, op)
}
