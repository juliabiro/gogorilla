package gorilla

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
	"math"
)

const (
	gravity = 0.98
)

type Banana struct {
	Point
	img         ScaledImage
	orientation float64
	width       float64
	height      float64
	angle       float64
	speed       float64
	gravity     float64
}

func (b *Banana) move(direction int) {
	if direction == right {
		b.X += b.speed * math.Cos(b.angle*math.Pi/180)
	} else {
		b.X -= b.speed * math.Cos(b.angle*math.Pi/180)
	}

	b.Y -= b.speed * math.Sin(b.angle*math.Pi/180)
	b.orientation += 0.1

	// apply gravity
	b.Y += b.gravity
}

func (g *Game) bananaOut() bool {
	return g.banana.X < 0 || g.banana.X > ScreenWidth || g.banana.Y > ScreenHeight
}

func (g *Game) resetBanana() {
	g.banana.X = g.turn.X
	if g.turn.direction == right {
		g.banana.X += g.turn.width
	}
	g.banana.Y = g.turn.Y
	g.banana.gravity = 0.0
}
func setupBanana(g *Game) {
	g.banana.width = 20
	g.banana.height = 20
	var err error
	img, _, err := ebitenutil.NewImageFromFile("./banana.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatal(err)
	}
	g.banana.img = ScaledImage{img, float64(g.banana.width) / float64(img.Bounds().Dx()), float64(g.banana.height) / float64(img.Bounds().Dy())}
	g.resetBanana()
}

func (b *Banana) DrawingParameters() (*ebiten.Image, *ebiten.DrawImageOptions) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(b.img.scaleX, b.img.scaleY)
	op.GeoM.Rotate(float64(b.orientation))
	// the order is important here: the image needs to be scaled before it is moved
	op.GeoM.Translate(float64(b.X), float64(b.Y))

	return b.img.Image, op
	//screen.DrawImage(b.img.Image, op)

}
