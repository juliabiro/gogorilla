package gorilla

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"golang.org/x/image/font"
	"image/color"
	"strconv"
)

const (
	gorilla1InputLoc = 10
	gorilla2InputLoc = ScreenWidth - 120
	center           = ScreenWidth/2 - 100
	leftSide         = 10
	rightSide        = ScreenWidth - 150
)

type IOHandler struct {
	inputString     string
	textDrawer      *TextDrawer
	counter         int
	inputDialog     string
	prevInputString string
}

func NewIOHandler() *IOHandler {
	i := IOHandler{}
	i.textDrawer = NewTextDrawer(color.White)
	i.reset()
	return &i
}

func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}
func handleBackspace(s *string) {
	if len(*s) >= 1 {
		*s = (*s)[:len(*s)-1]
	}
}

func (i *IOHandler) handleInput() (parsedValue float64, enterPressed bool) {
	i.inputString += string(ebiten.InputChars())
	if repeatingKeyPressed(ebiten.KeyBackspace) {
		handleBackspace(&i.inputString)
	}

	return i.handleEnter()
}

func (i *IOHandler) handleEnter() (parsedValue float64, enterPressed bool) {
	var val float64
	if repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyKPEnter) {
		var err error
		val, err = strconv.ParseFloat(i.inputString, 64)
		if err != nil {
			// just ignore invalid input
			i.inputString = ""
		}
		i.prevInputString = i.inputString
		i.inputString = ""
		return val, true
	}
	return val, false
}

var (
	mplusNormalFont font.Face
)

func (i *IOHandler) WriteMessage(screen *ebiten.Image, message string, location int) {
	i.textDrawer.Draw(screen, message, location, 30)
}

func (i *IOHandler) cursorBlink() {
	if i.counter%30 < 15 {
		i.inputDialog += "_"
	}
	i.counter++
}

func (i *IOHandler) reset() {
	i.inputString = ""
	i.prevInputString = ""
	i.inputDialog = ""
	i.counter = 0
}

func (i *IOHandler) WriteInputDialog(screen *ebiten.Image, gameState int, side int) {
	switch gameState {
	case inputAngle:
		i.inputDialog = "angle: " + i.inputString
		i.cursorBlink()

	case inputSpeed:
		i.inputDialog = "angle: " + i.prevInputString + "\nspeed: " + i.inputString
		i.cursorBlink()
	}

	i.textDrawer.Draw(screen, i.inputDialog, side, 60)
}
