package gorilla

import (
	"math"
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

func floatEqual(a, b float64) bool {
	return math.Abs(b-a) < 0.0001 // absolutaley arbitrary tolerance for imprecision
}

func TestMoveSpeedAngle(t *testing.T) {
	var bananaMoveSpeedAngleTestCases = []struct {
		speed, angle     float64
		changeX, changeY float64
	}{
		{0.0, 0.0, 0.0, 0.0},
		{10.0, 0.0, 10.0, 0.0},
		{10.0, 90.0, 0.0, -10.0},
	}

	b := NewBanana()
	for _, tc := range bananaMoveSpeedAngleTestCases {
		b.X, b.Y = 0.0, 0.0
		b.speed = tc.speed
		b.angle = tc.angle
		beforeX := b.X
		beforeY := b.Y
		b.move(right)
		changeX := b.X - beforeX
		if floatEqual(changeX, tc.changeX) != true {
			t.Fatalf("banana didn't move in the right direction. Speed: %f, angle: %f, new location: %f, %f (moving from 0,0), should be %f, %f", tc.speed, tc.angle, b.X, b.Y, beforeX+tc.changeX, beforeY+tc.changeY)
		}
	}
}
