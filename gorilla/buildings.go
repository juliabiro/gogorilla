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
	countHorizontal  int
	countVertical    int
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
	windows *Windows
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

		g.buildings = append(g.buildings, Building{Point{float64(k), float64(ScreenHeight - h)}, img, h, w, c, newWindows(w, h)})
		k = k + w
	}
}

func newWindows(w, h float64) *Windows {
	wc := color.RGBA{100 + uint8(rand.Intn(155)), 100 + uint8(rand.Intn(155)), 0, 255}
	locrand := uint8(rand.Intn(55))
	loc := color.RGBA{100 + locrand, 100 + locrand, 100 + locrand, 255}
	ww := w / float64((5 + rand.Intn(8)))
	bh := ww * float64(rand.Intn(15)+8) / 50
	wh := h / float64((5 + rand.Intn(15)))
	bv := wh * float64(rand.Intn(15)+8) / 50
	ch := w / ww
	cv := h / wh
	wimg, _ := ebiten.NewImage(int(ww), int(wh), ebiten.FilterDefault)
	loff := make(map[string]int)
	for i := 0.0; i*ww < w; i++ {
		for j := 0.0; j*wh < h; j++ {
			if rand.Intn(10) < 2 {
				loff[fmt.Sprintf("%s,%s", i, j)] = 1
			}
		}
	}
	windows := Windows{wc, ww, wh, bh, bv, int(ch), int(cv), wimg, loc, loff}
	return &windows
}

func (b *Building) DrawingParameters() (*ebiten.Image, *ebiten.DrawImageOptions) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(float64(b.X), float64(b.Y))
	b.img.Fill(b.color)
	b.windows.Draw(b.img)

	return b.img, op
}

func (w *Windows) Draw(img *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	//scaleX := float64(w.width-2*w.borderHorizontal) / float64(w.width)
	//scaleY := float64(w.height-2*w.borderVertical) / float64(w.height)
	for i := 0; i < w.countHorizontal; i++ {
		for j := 0; j < w.countVertical; j++ {
			op.GeoM.Reset()
			//op.GeoM.Scale(scaleX, scaleY)
			if w.lightsOff[fmt.Sprintf("%s,%s", i, j)] == 1 {
				w.img.Fill(w.lightsOffColor)
			} else {
				w.img.Fill(w.color)
			}
			op.GeoM.Translate(float64(i)*(w.width+w.borderHorizontal), float64(j)*(w.height+w.borderVertical))
			img.DrawImage(w.img, op)
		}
	}
}
