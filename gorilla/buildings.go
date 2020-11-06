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

func NewBuilding(startingX int, screenWidth int, screenHeight int) *Building {
	iw := 100 + rand.Intn(screenWidth/12)
	ih := 150 + rand.Intn(screenHeight/2)
	if startingX+iw >= screenWidth {
		iw = screenWidth - startingX
	}
	img, _ := ebiten.NewImage(iw, ih, ebiten.FilterDefault)
	c := color.RGBA{0, 0, 100 + uint8(rand.Intn(155)), 255}

	w := float64(iw)
	h := float64(ih)

	b := Building{Point{float64(startingX), float64(screenHeight - ih)}, img, h, w, c, newWindows(w, h)}
	return &b
}

func (b *Building) Width() int {
	return int(b.width)
}

func newWindows(w, h float64) *Windows {
	wc := color.RGBA{100 + uint8(rand.Intn(155)), 100 + uint8(rand.Intn(155)), 0, 255}
	locrand := uint8(rand.Intn(55))
	loc := color.RGBA{100 + locrand, 100 + locrand, 100 + locrand, 255}
	// how many windows we have
	ch := 5 + rand.Intn(8)
	cv := 5 + rand.Intn(15)
	// what is the width of a window unit (including frame)
	ww := w / float64(ch)
	wh := h / float64(cv)

	// what are the border thicknesses (in percentge of the whole window)
	th := float64(rand.Intn(40)+15) / 100
	tv := float64(rand.Intn(40)+15) / 100

	bh := ww * th / 2
	bv := wh * tv / 2

	// scale the window dmensions so they represent the drawable dimensions:
	ww = ww * (1 - th)
	wh = wh * (1 - tv)
	wimg, _ := ebiten.NewImage(int(ww), int(wh), ebiten.FilterDefault)
	loff := make(map[string]int)
	for i := 0; i < ch; i++ {
		for j := 0; j < cv; j++ {
			if rand.Intn(10) < 2 {
				loff[fmt.Sprintf("%d,%d", i, j)] = 1
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
	for i := 0; i < w.countHorizontal; i++ {
		for j := 0; j < w.countVertical; j++ {
			op.GeoM.Reset()
			if w.lightsOff[fmt.Sprintf("%d,%d", i, j)] == 1 {
				w.img.Fill(w.lightsOffColor)
			} else {
				w.img.Fill(w.color)
			}
			op.GeoM.Translate(w.borderHorizontal+float64(i)*(w.width+2*w.borderHorizontal), w.borderVertical+float64(j)*(w.height+2*w.borderVertical))
			img.DrawImage(w.img, op)
		}
	}
}
