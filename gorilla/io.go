package gorilla

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"golang.org/x/image/font"
	"log"
	"strconv"
)

const (
	gorilla1InputLoc = 10
	gorilla2InputLoc = ScreenWidth - 120
)

type IOHandler struct {
	inputString string
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
			log.Fatal(err)
		}
		i.inputString = ""
		return val, true
	}
	return val, false
}

var (
	mplusNormalFont font.Face
)

func WriteInputDialog(screen *ebiten.Image, g *Game) {
	t := ""
	switch g.gameState {
	case start:
		g.textDrawer.Draw(screen, "Game start: press Enter to continue", ScreenWidth/2-100, 30)
		return
	case gorillaDead:
		if g.gorilla1.alive {
			t = "Green Gorilla wins!"
		} else {
			t = "Red Gorilla wins!"
		}
		t = t + " Press Enter to continue."
		g.textDrawer.Draw(screen, t, ScreenWidth/2-100, 30)
		return
	case inputAngle:
		t = "angle: " + g.iohandler.inputString
		if g.counter%30 < 15 {
			t += "_"
		}
	case inputSpeed:
		t = "angle: " + g.iohandler.inputString + "\nspeed: " + g.iohandler.inputString
		if g.counter%30 < 15 {
			t += "_"

		}
	}
	loc := gorilla1InputLoc
	if g.turn == g.gorilla2 {
		loc = gorilla2InputLoc
	}
	g.textDrawer.Draw(screen, t, loc, 60)

}
