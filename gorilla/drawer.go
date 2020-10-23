package gorilla

import (
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image/color"
	"log"
)

type Drawable interface {
	DrawingParameters() (img *ebiten.Image, op ebiten.DrawImageOptions)
}

type TextDrawer struct {
	font  font.Face
	color color.Color
}

func NewTextDrawer(color color.Color) *TextDrawer {
	td := TextDrawer{}
	td.color = color

	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	td.font = truetype.NewFace(tt, &truetype.Options{
		Size:    20,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	return &td
}

func (td *TextDrawer) Draw(screen *ebiten.Image, s string, x int, y int) {

	text.Draw(screen, s, td.font, x, y, td.color)
}
