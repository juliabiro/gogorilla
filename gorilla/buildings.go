package gorilla

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"image/color"
	_ "image/png"
	"math/rand"
)

type Windows struct {
	color            color.Color
	width            float64
	height           float64
	borderHorizontal float64
	borderVertical   float64
	img              *ebiten.Image
	lightsOffColor   color.Color
	lightsOff        map[string]int
}

type Building struct {
	Point
	img     *ebiten.Image
	height  float64
	width   float64
	color   color.Color
	windows Windows
}

// setup
// TODO: turn into factory function
func setupBuildings(g *Game) {
	g.buildings = nil
	k := 0.0
	for k < ScreenWidth {
		w := float64(100 + rand.Intn(ScreenWidth/12))
		h := float64(150 + rand.Intn(ScreenHeight/2))
		if k+w >= ScreenWidth {
			w = ScreenWidth - k
		}
		img, _ := ebiten.NewImage(int(w), int(h), ebiten.FilterDefault)
		c := color.RGBA{0, 0, 100 + uint8(rand.Intn(155)), 255}

		wc := color.RGBA{100 + uint8(rand.Intn(155)), 100 + uint8(rand.Intn(155)), 0, 255}
		locrand := uint8(rand.Intn(55))
		loc := color.RGBA{100 + locrand, 100 + locrand, 100 + locrand, 255}
		ww := w / float64((5 + rand.Intn(8)))
		bh := ww * float64(rand.Intn(15)+8) / 50
		wh := h / float64((5 + rand.Intn(15)))
		bv := wh * float64(rand.Intn(15)+8) / 50
		wimg, _ := ebiten.NewImage(int(ww), int(wh), ebiten.FilterDefault)
		loff := make(map[string]int)
		for i := 0.0; i*ww < w; i++ {
			for j := 0.0; j*wh < h; j++ {
				if rand.Intn(10) < 2 {
					loff[fmt.Sprintf("%s,%s", i, j)] = 1
				}
			}
		}
		windows := Windows{wc, ww, wh, bh, bv, wimg, loc, loff}
		g.buildings = append(g.buildings, Building{Point{float64(k), float64(ScreenHeight - h)}, img, h, w, c, windows})
		k = k + w
	}
}
