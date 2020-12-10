package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/juliabiro/gogorilla/gorilla"
	_ "image/png"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"
)

const (
	frameOX      = 0
	frameOY      = 32
	frameWidth   = 32
	frameHeight  = 32
	frameNum     = 8
	screenWidth  = 1200
	screenHeight = 700
)

const (
	imageDir = "./images/"
)
const (
	start = iota
	newLap
	inputAngle
	inputSpeed
	bananaFlying
	bananaOut
	gorillaDead
	gravity = 0.5
)

// Game implements ebiten.Game interface.
type Game struct {
	gorilla1     *gorilla.Gorilla
	gorilla2     *gorilla.Gorilla
	buildings    []gorilla.Building
	banana       *gorilla.Banana
	turn         *gorilla.Gorilla
	gameState    int
	iohandler    *IOHandler
	screenwriter *ScreenWriter
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	for _, b := range g.buildings {
		screen.DrawImage(b.DrawingParameters())
	}
	switch g.gameState {
	case start:
		g.screenwriter.WriteMessage(screen, "Welcome to the Gorilla Game!\nPress Enter to start!", center)

	case gorillaDead:
		t := ""
		if g.gorilla1.Alive() {
			t = "Green Gorilla wins!"
		} else {
			t = "Red Gorilla wins!"
		}
		t = t + "\nPress Enter for a new round!"
		g.screenwriter.WriteMessage(screen, t, center)
	}

	screen.DrawImage(g.gorilla1.DrawingParamaters())
	screen.DrawImage(g.gorilla2.DrawingParamaters())
	screen.DrawImage(g.banana.DrawingParameters())

	side := 0
	if g.turn == g.gorilla1 {
		side = leftSide
	} else {
		side = rightSide
	}

	g.screenwriter.WriteMessage(screen, "Green Gorilla: "+strconv.Itoa(g.gorilla1.Score()), leftSide)
	g.screenwriter.WriteMessage(screen, "Red Gorilla: "+strconv.Itoa(g.gorilla2.Score()), rightSide)
	g.screenwriter.WriteInputDialog(screen, g.iohandler.InputString(), g.iohandler.PrevInputString(), g.gameState, side)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// Setup

func LoadImage(file string) *ebiten.Image {

	var err error

	img, _, err := ebitenutil.NewImageFromFile(file, ebiten.FilterDefault)

	if err != nil {
		log.Fatal(err)
	}

	return img
}

func (g *Game) Setup() {
	rand.Seed(time.Now().UnixNano())
	g.iohandler = NewIOHandler()
	g.screenwriter = NewScreenWriter()
	g.gameState = start

	g.setupBuildings()
	g.setupGorillas()
	g.turn = g.gorilla1
	g.setupBanana()

}
func (g *Game) setupGorillas() {
	g.gorilla1 = gorilla.NewGorilla(gorilla.Right)
	g.gorilla2 = gorilla.NewGorilla(gorilla.Left)
	g.gorilla1.SetImage(LoadImage(imageDir + "gorilla1.png"))
	g.gorilla2.SetImage(LoadImage(imageDir + "gorilla2.png"))
	g.resetGorillas()
}

func (g *Game) setupBanana() {
	img := LoadImage(imageDir + "banana.png")
	g.banana = gorilla.NewBanana(0, 0, 20, 20)
	g.banana.SetImage(img)
	g.banana.SetGravity(gravity)
	g.resetBanana()

}
func (g *Game) setupBuildings() {
	g.buildings = nil
	k := 0
	for k < screenWidth {
		b := gorilla.NewBuilding(k, screenWidth, screenHeight)
		g.buildings = append(g.buildings, *b)
		k = k + b.Width()
	}
}

// game logic

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(screen *ebiten.Image) error {
	// Write your game's logical update.

	parsedValue, enterPressed := g.iohandler.handleInput()

	g.updateSprites(parsedValue, enterPressed)

	// game state change needs to be either at the beginning or the end of this function
	g.updateGamestate(enterPressed)

	return nil
}

// only sprite updates here
func (g *Game) updateSprites(parsedValue float64, enterPressed bool) {
	switch g.gameState {
	case newLap:
		g.setupBuildings()
		g.resetGorillas()
		g.changeTurn()
		g.resetBanana()
	case inputAngle:
		if enterPressed {
			// TODO here the Direction needs to be offset based on the gorillas direction
			if g.turn == g.gorilla1 {
				g.banana.SetDirection(parsedValue)
			} else {

				g.banana.SetDirection(-1*parsedValue + 180)
			}
		}
	case inputSpeed:
		if enterPressed {
			g.banana.SetSpeed(parsedValue)
		}
	case bananaFlying:
		g.banana.Move()
		//  collision detection
		if detectCollision(g.banana, g.gorilla1) {
			g.gorilla1.Kill()
			g.gorilla2.IncreaseScore()
		}
		if detectCollision(g.banana, g.gorilla2) {
			g.gorilla2.Kill()
			g.gorilla1.IncreaseScore()
		}

	case bananaOut:
		g.changeTurn()
		g.resetBanana()
		// this is just cosmetic: make sure we don't writ e out the previous input numbers
		g.iohandler.reset()
		g.screenwriter.reset()
	}
}

// only state transitions here
func (g *Game) updateGamestate(enterPressed bool) {
	switch g.gameState {
	case start:
		// this is a special state that we visit only once
		// we just use it to make sure the opening screen is not bank
		if enterPressed {
			g.gameState = inputAngle
		}
	case newLap:
		g.gameState = inputAngle
	case inputAngle:
		if enterPressed {
			g.gameState = inputSpeed
		}
	case inputSpeed:
		if enterPressed {
			g.gameState = bananaFlying
		}
	case bananaFlying:
		if IsOut(g.banana) {
			g.gameState = bananaOut
		}

		if !g.gorilla1.Alive() || !g.gorilla2.Alive() {
			g.gameState = gorillaDead
		}

	case bananaOut:
		g.gameState = inputAngle

	case gorillaDead:
		if enterPressed {
			g.gameState = newLap
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
	g.gorilla1.Reset(g.buildings, screenWidth)
	g.gorilla2.Reset(g.buildings, screenWidth)
}

func (g *Game) resetBanana() {
	gg := g.turn
	g.banana.Reset(gg.X, gg.Y)
	g.banana.AlignWithGorilla(*gg)
}

func IsOut(b Center) bool {
	x, y := b.Center()
	return x < 0 || x > screenWidth || y > screenHeight
}
