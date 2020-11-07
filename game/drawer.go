package game

import (
	"github.com/hajimehoshi/ebiten"
)

type Drawable interface {
	DrawingParameters() (img *ebiten.Image, op ebiten.DrawImageOptions)
}
