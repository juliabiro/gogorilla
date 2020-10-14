package gorilla

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
)

const (
	ScreenWidth  = 1200
	ScreenHeight = 700
)

func DrawBuilding(screen *ebiten.Image, b Building) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(float64(b.X), float64(b.Y))
	b.img.Fill(b.color)
	screen.DrawImage(b.img, op)
	DrawWindows(screen, b)
}
func DrawWindows(screen *ebiten.Image, b Building) {
	w := b.windows
	op := &ebiten.DrawImageOptions{}

	scaleX := float64(w.width-2*w.borderHorizontal) / float64(w.width)
	scaleY := float64(w.height-2*w.borderVertical) / float64(w.height)
	for i := 1.0; i*w.width-w.borderHorizontal < b.width; i++ {
		for j := 1.0; j*w.height-w.borderVertical < b.height; j++ {
			op.GeoM.Reset()
			op.GeoM.Scale(scaleX, scaleY)
			if w.lightsOff[fmt.Sprintf("%s,%s", i, j)] == 1 {
				w.img.Fill(w.lightsOffColor)
			} else {
				w.img.Fill(w.color)
			}
			op.GeoM.Translate(b.X+float64((i-1)*w.width+w.borderHorizontal), b.Y+float64((j-1)*w.height+w.borderVertical))
			screen.DrawImage(w.img, op)
		}
	}
}
func DrawGorilla(screen *ebiten.Image, g Gorilla) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(g.img.scaleX, g.img.scaleY)
	op.GeoM.Translate(float64(g.X), float64(g.Y))
	screen.DrawImage(g.img.Image, op)
}
