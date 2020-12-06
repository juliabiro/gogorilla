package sprites

import (
	"github.com/hajimehoshi/ebiten"
)

type Point struct {
	X float64
	Y float64
}

type ScaledImage struct {
	*ebiten.Image
	scaleX, scaleY float64
}

type LoadImage interface {
	LoadImage()
}

type Reset interface {
	Reset()
}

type Drawable interface {
	DrawingParameters(*ebiten.Image, *ebiten.ImageOptions)
}

type Dimensions interface {
	Width() float64
	Height() float64
	Center() Point
}

type Move interface {
	Move()
}
