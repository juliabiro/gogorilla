package gorilla

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
