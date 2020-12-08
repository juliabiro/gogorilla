package sprites

import (
	"github.com/hajimehoshi/ebiten"
	"testing"
)

func TestNewSprite(t *testing.T) {
	iw, ih := 20, 20
	img, _ := ebiten.NewImage(iw, ih, ebiten.FilterDefault)

	s := NewSprite(5, 6, 10, 11, img)
	if s.x != 5 || s.y != 6 {
		t.Errorf("sprite location not initialized")
	}

	if s.width != 10 || s.height != 11 {
		t.Errorf("sprite dimensions not initialized")
	}
}

func TestDrawImage(t *testing.T) {
	iw, ih := 20, 20
	img, _ := ebiten.NewImage(iw, ih, ebiten.FilterDefault)

	s := NewSprite(0, 0, 40, 40, img)

	drawnimage, op := s.DrawingParameters()

	if drawnimage != s.img {
		t.Errorf("Not the right image")
	}
	expected_om := [2][3]float64{{2, 0, 0}, {0, 2, 0}}

	geom := op.GeoM
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			if geom.Element(i, j) != expected_om[i][j] {
				t.Errorf("scaling matrix ain't right")
			}
		}
	}

}
