package sprites

import (
	gomock "github.com/golang/mock/gomock"
	"github.com/hajimehoshi/ebiten"
	"github.com/juliabiro/gogorilla/mocks"

	"testing"
)

func TestNewSprite(t *testing.T) {

	s := NewSprite(5, 6, 10, 11)
	if s.X != 5 || s.Y != 6 {
		t.Errorf("sprite location not initialized")
	}

	if s.Width != 10 || s.Height != 11 {
		t.Errorf("sprite dimensions not initialized")
	}
}

func TestDrawImage(t *testing.T) {
	iw, ih := 20, 20
	img, _ := ebiten.NewImage(iw, ih, ebiten.FilterDefault)

	s := NewSprite(0, 0, 40, 40)
	s.SetImage(img)

	drawnimage, op := s.DrawingParameters()

	if drawnimage != s.Img {
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

func TestIsTouching(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	m1 := mocks.NewMockCollisionDetection(ctrl)
	m2 := mocks.NewMockCollisionDetection(ctrl)
	m1.EXPECT().Center().Return(50.0, 50.0)
	m2.EXPECT().Center().Return(100.0, 100.0)

	m1.EXPECT().IsInside(100.0, 100.0).Return(false)
	m2.EXPECT().IsInside(50.0, 50.0).Return(false)

	IsTouching(m1, m2)
}
