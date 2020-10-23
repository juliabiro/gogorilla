package gorilla

import (
	"github.com/hajimehoshi/ebiten"

	"image/color"
	_ "image/png"
	"math"
	"math/rand"
	"strconv"
	"time"
)

const (
	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameNum    = 8
)

const (
	start = iota
	inputAngle
	inputSpeed
	bananaFlying
	gorillaDead
)

// Game implements ebiten.Game interface.
type Game struct {
	gorilla1   *Gorilla
	gorilla2   *Gorilla
	buildings  []Building
	banana     *Banana
	turn       *Gorilla
	gameState  int
	inputAngle string
	inputSpeed string
	counter    int
	textDrawer *TextDrawer
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {

	for _, b := range g.buildings {
		screen.DrawImage(b.DrawingParameters())
	}
	screen.DrawImage(g.gorilla1.DrawingParamaters())
	screen.DrawImage(g.gorilla2.DrawingParamaters())
	screen.DrawImage(g.banana.DrawingParameters())
	//DrawBanana(screen, g.banana)
	// Write your game's rendering.
	WriteInputDialog(screen, g)
	g.textDrawer.Draw(screen, "Gorilla1: "+strconv.Itoa(g.gorilla1.score), 10, 30)
	g.textDrawer.Draw(screen, "Gorilla2: "+strconv.Itoa(g.gorilla2.score), ScreenWidth-150, 30)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

// Setup

func (g *Game) Setup() {
	rand.Seed(time.Now().UnixNano())
	g.textDrawer = NewTextDrawer(color.White)
	g.gameState = start
	g.setupBuildings()

	g.setupGorillas()
	g.turn = g.gorilla1
	g.setupBanana()

}
func (g *Game) setupGorillas() {
	g.gorilla1 = NewGorilla(right)
	g.gorilla2 = NewGorilla(left)
	g.resetGorillas()
}

func (g *Game) setupBanana() {
	g.banana = NewBanana()
	g.resetBanana()

}
func (g *Game) setupBuildings() {
	g.buildings = nil
	k := 0.0
	for k < ScreenWidth {
		b := NewBuilding(k)
		g.buildings = append(g.buildings, *b)
		k = k + b.width
	}
}

// game logic

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(screen *ebiten.Image) error {
	// Write your game's logical update.

	handleEnter(g)
	g.updateGamestate()

	g.counter++
	return nil
}
func (g *Game) updateGamestate() {
	switch g.gameState {
	case inputAngle:
		g.inputAngle += string(ebiten.InputChars())
		if repeatingKeyPressed(ebiten.KeyBackspace) {
			handleBackspace(&g.inputAngle)
		}
	case inputSpeed:
		g.inputSpeed += string(ebiten.InputChars())
		if repeatingKeyPressed(ebiten.KeyBackspace) {
			handleBackspace(&g.inputSpeed)
		}
	case bananaFlying:
		g.banana.applyGravity(gravity)
		g.banana.move(g.turn.direction)
		//  collision detection
		if detectCollision(g.banana, g.gorilla1) {
			g.gorilla1.kill()
			g.gorilla2.increaseScore()
			g.gameState = gorillaDead
		}
		if detectCollision(g.banana, g.gorilla2) {
			g.gorilla2.kill()
			g.gorilla1.increaseScore()
			g.gameState = gorillaDead
		}

		if g.bananaOut() {
			g.changeTurn()
			g.resetBanana()
			g.gameState = inputAngle
		}

	}
}

func distance(p1 Center, p2 Center) float64 {
	x1, y1 := p1.Center()
	x2, y2 := p2.Center()
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}

func detectCollision(p1 Center, p2 Center) bool {
	return distance(p1, p2) < 25
}

func (g *Game) changeTurn() {
	if g.turn == g.gorilla1 {
		g.turn = g.gorilla2
	} else {
		g.turn = g.gorilla1
	}
}

func (g *Game) resetGorillas() {
	g.gorilla1.reset(g.buildings)
	g.gorilla2.reset(g.buildings)
}

func (g *Game) resetBanana() {
	g.banana.reset()
	g.banana.alignWithGorilla(*g.turn)
}
