package sprites

import (
	gomock "github.com/golang/mock/gomock"
	"github.com/hajimehoshi/ebiten"
	"github.com/juliabiro/gogorilla/mocks"
	"math"
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

const (
	toUp         = "up"
	toDown       = "down"
	toRight      = "right"
	toLeft       = "left"
	noHorizontal = "no horizontal move"
	noVertical   = "no vertical move"
)

func floatEqual(a, b float64) bool {
	return math.Abs(b-a) < 0.0001 // absolutaley arbitrary tolerance for imprecision
}

func getChangeDirection(beforeX, beforeY, afterX, afterY float64) (string, string) {
	var w, h = noHorizontal, noVertical

	if !floatEqual(beforeX, afterX) {
		if beforeX < afterX {
			w = toRight
		}
		if beforeX > afterX {
			w = toLeft
		}
	}

	if !floatEqual(beforeY, afterY) {
		if beforeY < afterY {
			h = toDown
		}
		if beforeY > afterY {
			h = toUp
		}
	}
	return w, h
}

func TestMove(t *testing.T) {
	type input struct {
		speed, angle float64
	}

	type output struct {
		horizontalMoveDirection, verticalMoveDirection string
	}
	var bananaMoveDirectionTestCase = []struct {
		in  input
		out output
	}{
		// base case: 0, 45 or 90 to the right
		{input{10, 0}, output{toRight, noVertical}},
		{input{10, 90}, output{noHorizontal, toUp}},
		{input{10, 45}, output{toRight, toUp}},

		// negative speed to the gorilla.Right and to thegorilla.Left
		{input{-10, 0}, output{toLeft, noVertical}},
		{input{-10, 90}, output{noHorizontal, toDown}},
		{input{-10, 45}, output{toLeft, toDown}},

		// negative angle to the gorilla.Right and to thegorilla.Left
		{input{10, 0}, output{toRight, noVertical}},
		{input{10, -90}, output{noHorizontal, toDown}},
		{input{10, -45}, output{toRight, toDown}},

		// both speed and angle negative
		{input{-10, 0}, output{toLeft, noVertical}},
		{input{-10, -90}, output{noHorizontal, toUp}},
		{input{-10, -45}, output{toLeft, toUp}},

		// angles over 90 degrees
		{input{10, 0}, output{toRight, noVertical}},
		{input{10, 90}, output{noHorizontal, toUp}},
		{input{10, 45}, output{toRight, toUp}},
		{input{10, 135}, output{toLeft, toUp}},
		{input{10, 180}, output{toLeft, noVertical}},
		{input{10, 225}, output{toLeft, toDown}},
		{input{10, 270}, output{noHorizontal, toDown}},
		{input{10, 315}, output{toRight, toDown}},
		{input{10, 360}, output{toRight, noVertical}},
		{input{10, 370}, output{toRight, toUp}},
	}

	b := NewSprite(0, 0, 20, 20)
	for _, tc := range bananaMoveDirectionTestCase {
		b.Reset(0, 0)
		b.SetSpeed(tc.in.speed)
		b.SetDirection(tc.in.angle)
		beforeX := b.X
		beforeY := b.Y

		b.Move()

		w, h := getChangeDirection(beforeX, beforeY, b.X, b.Y)

		if w != tc.out.horizontalMoveDirection || h != tc.out.verticalMoveDirection {
			t.Fatalf("Banana thrown with speed %f and angle %f should move %s and %s, but instead it went %s and %s", tc.in.speed, tc.in.angle, tc.out.horizontalMoveDirection, tc.out.verticalMoveDirection, w, h)
		}
	}
}
