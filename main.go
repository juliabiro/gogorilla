package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"image/color"
	_ "image/png"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"
)

const (
	screenWidth  = 1200
	screenHeight = 700

	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameNum    = 8
)
const (
	right = iota
	left
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
	gorilla1   Gorilla
	gorilla2   Gorilla
	buildings  []Building
	banana     Banana
	turn       Gorilla
	gameState  int
	inputAngle string
	inputSpeed string
	counter    int
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

func (g *Game) bananaOut() bool {
	return g.banana.X < 0 || g.banana.X > screenWidth || g.banana.Y < 0 || g.banana.Y > screenHeight
}

func (g *Game) changeTurn() {
	if g.turn == g.gorilla1 {
		g.turn = g.gorilla2
	} else {
		g.turn = g.gorilla1
	}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(screen *ebiten.Image) error {
	// Write your game's logical update.
	if repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyKPEnter) {
		switch g.gameState {
		case start:
			g.gameState = inputAngle
		case inputAngle:
			g.gameState = inputSpeed
		case inputSpeed:
			g.banana.angle, _ = strconv.ParseFloat(g.inputAngle, 64)
			g.banana.speed, _ = strconv.ParseFloat(g.inputSpeed, 64)
			g.inputAngle = ""
			g.inputSpeed = ""
			g.gameState = bananaFlying
		default:
		}
	}

	switch g.gameState {
	case inputAngle:
		g.inputAngle += string(ebiten.InputChars())
	case inputSpeed:
		g.inputSpeed += string(ebiten.InputChars())
	case bananaFlying:
		if g.turn == g.gorilla1 {
			g.banana.move(right)
		} else {
			g.banana.move(left)
		}
		if g.bananaOut() {
			g.changeTurn()
			g.resetBanana()
			g.gameState = inputAngle
		}
	}
	g.counter++
	return nil
}

func (g *Game) WriteInputDialog(screen *ebiten.Image) {
	t := ""
	switch g.gameState {
	case start:
		ebitenutil.DebugPrint(screen, "game start: press Enter to continue")
	case gorillaDead:
		t = fmt.Sprintf("%s won! press Enter to continue.", g.turn)
		ebitenutil.DebugPrint(screen, t)
	case inputAngle:
		t = "angle: " + g.inputAngle
		if g.counter%60 < 30 {
			t += "_"
		}
	case inputSpeed:
		t = "angle: " + g.inputAngle + "\nspeed: " + g.inputSpeed
		if g.counter%60 < 30 {
			t += "_"
		}
	default:
	}
	ebitenutil.DebugPrint(screen, t)

}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {

	for _, b := range g.buildings {
		b.Draw(screen)
	}
	g.gorilla1.Draw(screen)
	g.gorilla2.Draw(screen)
	g.banana.Draw(screen)
	// Write your game's rendering.
	g.WriteInputDialog(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game := &Game{}
	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Gorillas")
	setup(game)
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type Point struct {
	X float64
	Y float64
}

type scaledImage struct {
	*ebiten.Image
	scaleX, scaleY float64
}

type Gorilla struct {
	Point
	alive  bool
	img    scaledImage
	height float64
	width  float64
}

func (g *Gorilla) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(g.img.scaleX, g.img.scaleY)
	op.GeoM.Translate(float64(g.X), float64(g.Y))
	screen.DrawImage(g.img.Image, op)
}

type Windows struct {
	color            color.Color
	width            float64
	height           float64
	borderHorizontal float64
	borderVertical   float64
	img              *ebiten.Image
	lightsOffColor   color.Color
	lightsOff        map[string]int
}

func (w *Windows) Draw(screen *ebiten.Image, b *Building) {
	op := &ebiten.DrawImageOptions{}

	scaleX := float64(w.width-2*w.borderHorizontal) / float64(w.width)
	scaleY := float64(w.height-2*w.borderVertical) / float64(w.height)
	for i := 1.0; i*w.width-w.borderHorizontal < b.width; i++ {
		for j := 1.0; j*w.height-w.borderVertical < b.height; j++ {
			op.GeoM.Reset()
			op.GeoM.Scale(scaleX, scaleY)
			if w.lightsOff[fmt.Sprintf("%s,%s", i, j)] == 1 {
				w.img.Fill(w.lightsOffColor)
			} else {
				w.img.Fill(w.color)
			}
			op.GeoM.Translate(b.X+float64((i-1)*w.width+w.borderHorizontal), b.Y+float64((j-1)*w.height+w.borderVertical))
			screen.DrawImage(w.img, op)
		}
	}
}

type Building struct {
	Point
	img     *ebiten.Image
	height  float64
	width   float64
	color   color.Color
	windows Windows
}

func (b *Building) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(float64(b.X), float64(b.Y))
	b.img.Fill(b.color)
	screen.DrawImage(b.img, op)
	b.windows.Draw(screen, b)
}

type Banana struct {
	Point
	img         scaledImage
	orientation float64
	width       float64
	height      float64
	angle       float64
	speed       float64
}

func (b *Banana) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(b.img.scaleX, b.img.scaleY)
	op.GeoM.Rotate(float64(b.orientation))
	// the order is important here: the image needs to be scaled before it is moved
	op.GeoM.Translate(float64(b.X), float64(b.Y))
	screen.DrawImage(b.img.Image, op)

}

func (b *Banana) move(direction int) {
	if direction == right {
		b.X += b.speed * math.Cos(b.angle*math.Pi/180)
	} else {
		b.X -= b.speed * math.Cos(b.angle*math.Pi/180)
	}

	b.Y -= b.speed * math.Sin(b.angle)
	b.orientation += 0.1
}

func setupBuildings(g *Game) {
	k := 0.0
	for k < screenWidth {
		w := float64(100 + rand.Intn(screenWidth/12))
		h := float64(150 + rand.Intn(screenHeight/2))
		if k+w >= screenWidth {
			w = screenWidth - k
		}
		img, _ := ebiten.NewImage(int(w), int(h), ebiten.FilterDefault)
		c := color.RGBA{0, 0, 100 + uint8(rand.Intn(155)), 255}

		wc := color.RGBA{100 + uint8(rand.Intn(155)), 100 + uint8(rand.Intn(155)), 0, 255}
		locrand := uint8(rand.Intn(55))
		loc := color.RGBA{100 + locrand, 100 + locrand, 100 + locrand, 255}
		ww := w / float64((5 + rand.Intn(8)))
		bh := ww * float64(rand.Intn(15)+8) / 50
		wh := h / float64((5 + rand.Intn(15)))
		bv := wh * float64(rand.Intn(15)+8) / 50
		wimg, _ := ebiten.NewImage(int(ww), int(wh), ebiten.FilterDefault)
		loff := make(map[string]int)
		for i := 0.0; i*ww < w; i++ {
			for j := 0.0; j*wh < h; j++ {
				if rand.Intn(10) < 2 {
					loff[fmt.Sprintf("%s,%s", i, j)] = 1
				}
			}
		}
		windows := Windows{wc, ww, wh, bh, bv, wimg, loc, loff}
		g.buildings = append(g.buildings, Building{Point{float64(k), float64(screenHeight - h)}, img, h, w, c, windows})
		k = k + w
	}
}

func (g *Gorilla) init(minx int, b []Building) {

	g.alive = true
	g.width = 50
	g.height = 50
	var err error
	img, _, err := ebitenutil.NewImageFromFile("/Users/juliabiro/go/gorilla/gorilla.png", ebiten.FilterDefault)
	g.img = scaledImage{img, float64(g.width) / float64(img.Bounds().Dx()), float64(g.height) / float64(img.Bounds().Dy())}
	if err != nil {
		log.Fatal(err)
	}
	g.X = float64(minx + rand.Intn(0.6*screenWidth/2))

	// find my rooftop
	i := 0
	for g.X > b[i].X+float64(b[i].width) {
		i++
	}
	bb := b[i]

	// make sure I sit on it
	g.Y = bb.Y - float64(g.height)
	if g.X < bb.X || g.X+g.width > bb.X+bb.width {
		g.X = bb.X + float64(rand.Intn(int(bb.width-g.width)))
	}

}
func setupGorillas(g *Game) {
	g.gorilla1.init(0, g.buildings)
	g.gorilla2.init(screenWidth/2, g.buildings)
}

func (g *Game) resetBanana() {
	g.banana.X = g.turn.X + g.turn.width
	g.banana.Y = g.turn.Y
}
func setupBanana(g *Game) {
	g.banana.width = 20
	g.banana.height = 20
	var err error
	img, _, err := ebitenutil.NewImageFromFile("/Users/juliabiro/go/gorilla/banana.png", ebiten.FilterDefault)
	g.banana.img = scaledImage{img, float64(g.banana.width) / float64(img.Bounds().Dx()), float64(g.banana.height) / float64(img.Bounds().Dy())}

	if err != nil {
		log.Fatal(err)
	}
	g.resetBanana()
}

func setup(g *Game) {
	g.gameState = start
	setupBuildings(g)
	setupGorillas(g)
	g.turn = g.gorilla1
	setupBanana(g)
	// log.Printf("%v", g.gorilla1)
	// log.Printf("%v", g.gorilla2)
}
