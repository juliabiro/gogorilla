package gorilla

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image/color"
	"strconv"
)

const (
	gorilla1InputLoc = 10
	gorilla2InputLoc = ScreenWidth - 120
)

var (
	mplusNormalFont font.Face
)

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

func handleEnter(g *Game) {
	if repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyKPEnter) {
		var err error
		switch g.gameState {
		case start:
			g.resetBanana()
			g.gameState = inputAngle
		case inputAngle:
			g.banana.angle, err = strconv.ParseFloat(g.inputAngle, 64)
			if err != nil {
				g.inputAngle = ""
			}
			g.gameState = inputSpeed
		case inputSpeed:
			g.banana.speed, err = strconv.ParseFloat(g.inputSpeed, 64)
			if err != nil {
				g.inputSpeed = ""
			}
			g.inputAngle = ""
			g.inputSpeed = ""
			g.gameState = bananaFlying
		case gorillaDead:
			setupBuildings(g)
			g.resetGorillas()
			g.changeTurn()
			g.resetBanana()
			g.gameState = start
		}
	}

}

func WriteInputDialog(screen *ebiten.Image, g *Game) {
	t := ""
	switch g.gameState {
	case start:
		text.Draw(screen, "game start: press Enter to continue", mplusNormalFont, ScreenWidth/2-100, 30, color.White)
		return
	case gorillaDead:
		if g.gorilla1.alive {
			t = "Gorilla1 wins!"
		} else {
			t = "Gorilla2 wins!"
		}
		t = t + "Press Enter to continue."
		text.Draw(screen, t, mplusNormalFont, ScreenWidth/2-100, 30, color.White)
		return
	case inputAngle:
		t = "angle: " + g.inputAngle
		if g.counter%30 < 15 {
			t += "_"
		}
	case inputSpeed:
		t = "angle: " + g.inputAngle + "\nspeed: " + g.inputSpeed
		if g.counter%30 < 15 {
			t += "_"

		}
	}
	loc := gorilla1InputLoc
	if g.turn == g.gorilla2 {
		loc = gorilla2InputLoc
	}
	text.Draw(screen, t, mplusNormalFont, loc, 60, color.White)

}
