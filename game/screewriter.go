package game

import (
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image/color"
	"log"
)

const (
	gorilla1InputLoc = 10
	gorilla2InputLoc = screenWidth - 120
	center           = screenWidth/2 - 100
	leftSide         = 10
	rightSide        = screenWidth - 150
)

var (
	mplusNormalFont font.Face
)

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

type ScreenWriter struct {
	textDrawer  *TextDrawer
	counter     int
	inputDialog string
}

func NewScreenWriter() *ScreenWriter {
	sw := ScreenWriter{}
	sw.textDrawer = NewTextDrawer(color.White)
	return &sw
}

func (i *ScreenWriter) WriteMessage(screen *ebiten.Image, message string, location int) {
	i.textDrawer.Draw(screen, message, location, 30)
}

func (i *ScreenWriter) cursorBlink() {
	if i.counter%30 < 15 {
		i.inputDialog += "_"
	}
	i.counter++
}

func (i *ScreenWriter) reset() {
	i.inputDialog = ""
	i.counter = 0
}

func (i *ScreenWriter) WriteInputDialog(screen *ebiten.Image, inputString string, prevInputString string, gameState int, side int) {
	switch gameState {
	case inputAngle:
		i.inputDialog = "angle: " + inputString
		i.cursorBlink()

	case inputSpeed:
		i.inputDialog = "angle: " + prevInputString + "\nspeed: " + inputString
		i.cursorBlink()
	}

	i.textDrawer.Draw(screen, i.inputDialog, side, 60)
}
