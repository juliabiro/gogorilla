package game_test

import (
	gomock "github.com/golang/mock/gomock"
	"github.com/juliabiro/gogorilla/game"
	"github.com/juliabiro/gogorilla/mocks"
	"testing"
)

func TestOut(t *testing.T) {
	g := game.Game{}
	screenWidth, screenHeight := g.Layout(0, 0)
	var bananaLocationTestCases = []struct {
		coordinates [2]float64
		out         bool
	}{
		// inside the image
		{[2]float64{0.0, 0.0}, false},
		{[2]float64{5.0, 4.0}, false},
		{[2]float64{float64(screenWidth / 2), float64(screenHeight / 2)}, false},
		// out left, but just a bit
		{[2]float64{-0.0001, 0.0}, true},
		// out left, a lot
		{[2]float64{-10.0, 0.0}, true},
		// out right
		{[2]float64{float64(screenWidth + 5), 0.0}, true},
		// out up
		{[2]float64{0, -100.0}, false},
		// out down
		{[2]float64{0, float64(screenHeight + 5)}, true},
	}

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	for _, tc := range bananaLocationTestCases {
		m1 := mocks.NewMockCollisionDetection(ctrl)
		m1.EXPECT().Center().Return(tc.coordinates[0], tc.coordinates[1])
		isOut := game.IsOut(m1)
		if isOut != tc.out {
			t.Fatalf("coordinates %f,%f should be out: %t, but I got %t", tc.coordinates[0], tc.coordinates[1], tc.out, isOut)
		}
	}
}
