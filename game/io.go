package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"strconv"
)

type IOHandler struct {
	inputString     string
	prevInputString string
}

func NewIOHandler() *IOHandler {
	i := IOHandler{}
	i.reset()
	return &i
}

func (i *IOHandler) InputString() string {
	return i.inputString
}

func (i *IOHandler) PrevInputString() string {
	return i.prevInputString
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

func (i *IOHandler) reset() {
	i.inputString = ""
}
