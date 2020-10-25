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

func (b *Banana) Center() (X, Y float64) {
	return b.X + b.width/2, b.Y + b.height/2
}

func (b *Banana) move(direction int) {
	if direction == right {
		b.X += b.speed * math.Cos(b.angle*math.Pi/180)
	} else {
		b.X -= b.speed * math.Cos(b.angle*math.Pi/180)
	}

	b.Y -= b.speed * math.Sin(b.angle*math.Pi/180)
	b.orientation += b.speed / 100

	// apply gravity
	b.Y += b.gravity
}

func (b *Banana) Out() bool {
	return b.X < 0 || b.X > ScreenWidth || b.Y > ScreenHeight
}

func NewBanana() *Banana {
	b := Banana{}
	b.width = 20
	b.height = 20
	var err error
	img, _, err := ebitenutil.NewImageFromFile("./banana.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatal(err)
	}
	b.img = ScaledImage{img, float64(b.width) / float64(img.Bounds().Dx()), float64(b.height) / float64(img.Bounds().Dy())}
	return &b
}

func (b *Banana) DrawingParameters() (*ebiten.Image, *ebiten.DrawImageOptions) {
	op := &ebiten.DrawImageOptions{}
	// the order is important here: the image needs to be scaled before it is moved
	op.GeoM.Rotate(float64(b.orientation))
	op.GeoM.Scale(b.img.scaleX, b.img.scaleY)
	op.GeoM.Translate(float64(b.X), float64(b.Y))

	return b.img.Image, op
}

func (b *Banana) reset() {
	b.X = 0
	b.Y = 0
	b.gravity = 0.0
}

func (b *Banana) alignWithGorilla(g Gorilla) {
	b.X = g.X
	b.Y = g.Y
	if g.direction == right {
		b.X += g.width
	}
}

func (b *Banana) applyGravity(gravity float64) {
	b.gravity += gravity
}

func (b *Banana) setAngle(angle float64) {
	b.angle = angle
}

func (b *Banana) setSpeed(speed float64) {
	b.speed = speed
}
