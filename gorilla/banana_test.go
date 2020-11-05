package gorilla

import (
	"testing"
)

func TestNewBanana(t *testing.T) {
	b := NewBanana()
	if b.width == 0 || b.height == 0 {
		t.Errorf("banana dimensions not initialized")
	}
}

func TestResetBanana(t *testing.T) {
	b := NewBanana()
	b.X = 3
	b.Y = 4
	b.gravity = 7
	b.reset()
	if b.X != 0 {
		t.Fatalf("banana X coordinate not reset properly")
	}
	if b.Y != 0 {
		t.Fatalf("banana Y coordinate not reset properly")
	}
	if b.gravity != 0.0 {
		t.Fatalf("banana gravity coordinate not reset properly")
	}
}

func TestOut(t *testing.T) {
	var bananaLocationTestCases = []struct {
		coordinates [2]float64
		out         bool
	}{
		// inside the image
		{[2]float64{0.0, 0.0}, false},
		{[2]float64{5.0, 4.0}, false},
		{[2]float64{ScreenWidth / 2, ScreenHeight / 2}, false},
		// out left
		{[2]float64{-1.0, 0.0}, true},
		// out right
		{[2]float64{ScreenWidth + 5, 0.0}, true},
		// out up
		{[2]float64{0, -100.0}, false},
		// out down
		{[2]float64{0, ScreenHeight + 5}, true},
	}

	b := NewBanana()

	for _, tc := range bananaLocationTestCases {
		b.X = tc.coordinates[0]
		b.Y = tc.coordinates[1]
		if b.Out() != tc.out {
			t.Fatalf("coordinates %f,%f should be out: %t", tc.coordinates[0], tc.coordinates[1], tc.out)
		}
	}
}
