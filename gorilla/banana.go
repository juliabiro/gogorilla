package gorilla

import (
	"github.com/juliabiro/gogorilla/sprites"
)

type Banana struct {
	sprites.Sprite
	// Point
	// img         ScaledImage
	// orientation float64
	// width       float64
	// height      float64
	// angle       float64
	// speed       float64
	// gravity     float64
}

// func (b *Banana) Dimensions() (width, height float64) {
// 	return b.width, b.height
// }

// func (b *Banana) MoveData() (speed, angle, gravity float64) {
// 	return b.speed, b.angle, b.gravity
// }

// func (b *Banana) SetMoveData(speed, angle, gravity float64) {
// 	b.speed = speed
// 	b.angle = angle
// 	b.gravity = gravity
// }

// func (b *Banana) Center() (X, Y float64) {
// 	return b.X + b.width/2, b.Y + b.height/2
// }

// func (b *Banana) Move(direction int) {
// 	if direction == Right {
// 		b.X += b.speed * math.Cos(b.angle*math.Pi/180)
// 	} else {
// 		b.X -= b.speed * math.Cos(b.angle*math.Pi/180)
// 	}

// 	b.Y -= b.speed * math.Sin(b.angle*math.Pi/180)
// 	b.orientation += b.speed / 100

// 	// apply gravity
// 	b.Y += b.gravity
// }

// func (b *Banana) SetImage(img *ebiten.Image) {
// 	b.img = ScaledImage{img, float64(b.width) / float64(img.Bounds().Dx()), float64(b.height) / float64(img.Bounds().Dy())}

// }

func NewBanana(x, y float64, w, h int) *Banana {
	b := Banana{*sprites.NewSprite(x, y, w, h)}
	return &b
}

// func (b *Banana) DrawingParameters() (*ebiten.Image, *ebiten.DrawImageOptions) {
// 	op := &ebiten.DrawImageOptions{}
// 	// the order is important here: the image needs to be scaled before it is moved
// 	op.GeoM.Rotate(float64(b.orientation))
// 	op.GeoM.Scale(b.img.scaleX, b.img.scaleY)
// 	op.GeoM.Translate(float64(b.X), float64(b.Y))

// 	return b.img.Image, op
// }

// func (b *Banana) Reset() {
// 	b.X = 0
// 	b.Y = 0
// 	b.gravity = 0.0
// }

func (b *Banana) AlignWithGorilla(g Gorilla) {
	b.X = g.X
	b.Y = g.Y
	if g.direction == Right {
		b.Y += g.Width
	}
}

// func (b *Banana) ApplyGravity(gravity float64) {
// 	b.gravity += gravity
// }

// func (b *Banana) SetAngle(angle float64) {
// 	b.angle = angle
// }

// func (b *Banana) SetSpeed(speed float64) {
// 	b.speed = speed
// }
