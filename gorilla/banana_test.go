package gorilla_test

import (
	"github.com/juliabiro/gogorilla/gorilla"
	"math"
	"testing"
)

func TestNewBanana(t *testing.T) {
	b := gorilla.NewBanana(3, 4, 0, 0)
	w, h := b.Width, b.Height
	if w == 3 || h == 4 {
		t.Errorf("banana dimensions not initialized")
	}
}

func TestResetBanana(t *testing.T) {
	b := gorilla.NewBanana(3, 4, 20, 20)

	b.SetSpeed(10)
	b.SetDirection(45)
	b.SetGravity(8)

	b.Reset(6, 7)
	if b.X != 6 {
		t.Fatalf("banana X coordinate not reset properly")
	}
	if b.Y != 7 {
		t.Fatalf("banana Y coordinate not reset properly")
	}
	if b.Gravity != 0.0 {
		t.Fatalf("banana gravity coordinate not reset properly")
	}
}

func floatEqual(a, b float64) bool {
	return math.Abs(b-a) < 0.0001 // absolutaley arbitrary tolerance for imprecision
}

func TestMoveSpeedAngle(t *testing.T) {
	type input struct {
		speed, angle float64
	}
	type output struct {
		changeX, changeY float64
	}
	var bananaMoveSpeedAngleTestCases = []struct {
		in  input
		out output
	}{
		{input{0.0, 0.0}, output{0.0, 0.0}},
		{input{10.0, 0.0}, output{10.0, 0.0}},
		{input{10.0, 90.0}, output{0.0, -10.0}},
		{input{10.0, 45.0}, output{1.0 / math.Sqrt(2) * 10, -1.0 / math.Sqrt(2) * 10}},
	}

	b := gorilla.NewBanana(0, 0, 20, 20)
	for _, tc := range bananaMoveSpeedAngleTestCases {
		b.SetSpeed(tc.in.speed)
		b.SetDirection(tc.in.angle)
		beforeX := b.X
		beforeY := b.Y
		b.Move()
		changeX := b.X - beforeX
		changeY := b.Y - beforeY
		if floatEqual(changeX, tc.out.changeX) != true || floatEqual(changeY, tc.out.changeY) != true {
			t.Fatalf("banana didn't move in the right direction. Speed: %f, angle: %f, new location: %f, %f (moving from 0,0), should be %f, %f", tc.in.speed, tc.in.angle, b.X, b.Y, beforeX+tc.out.changeX, beforeY+tc.out.changeY)
		}
	}
}
