package gorilla_test

import (
	"fmt"
	"github.com/juliabiro/gogorilla/gorilla"
	"math"
	"testing"
)

func TestNewBanana(t *testing.T) {
	b := gorilla.NewBanana()
	w, h := b.Dimensions()
	if w == 0 || h == 0 {
		t.Errorf("banana dimensions not initialized")
	}
}

func TestResetBanana(t *testing.T) {
	b := gorilla.NewBanana()
	b.X = 3
	b.Y = 4
	_, _, g := b.MoveData()

	b.SetMoveData(10, 45, 7)
	b.Reset()
	if b.X != 0 {
		t.Fatalf("banana X coordinate not reset properly")
	}
	if b.Y != 0 {
		t.Fatalf("banana Y coordinate not reset properly")
	}
	if g != 0.0 {
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
		{[2]float64{gorilla.ScreenWidth / 2, gorilla.ScreenHeight / 2}, false},
		// out left
		{[2]float64{-1.0, 0.0}, true},
		// out right
		{[2]float64{gorilla.ScreenWidth + 5, 0.0}, true},
		// out up
		{[2]float64{0, -100.0}, false},
		// out down
		{[2]float64{0, gorilla.ScreenHeight + 5}, true},
	}

	b := gorilla.NewBanana()

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

	b := gorilla.NewBanana()
	for _, tc := range bananaMoveSpeedAngleTestCases {
		b.X, b.Y = 0.0, 0.0
		b.SetMoveData(tc.in.speed, tc.in.angle, 0)
		beforeX := b.X
		beforeY := b.Y
		b.Move(gorilla.Right)
		changeX := b.X - beforeX
		changeY := b.Y - beforeY
		if floatEqual(changeX, tc.out.changeX) != true || floatEqual(changeY, tc.out.changeY) != true {
			t.Fatalf("banana didn't move in the right direction. Speed: %f, angle: %f, new location: %f, %f (moving from 0,0), should be %f, %f", tc.in.speed, tc.in.angle, b.X, b.Y, beforeX+tc.out.changeX, beforeY+tc.out.changeY)
		}
	}
}

const (
	toUp         = "up"
	toDown       = "down"
	toRight      = "right"
	toLeft       = "left"
	noHorizontal = "no horizontal move"
	noVertical   = "no vertical move"
)

func getChangeDirection(beforeX, beforeY, afterX, afterY float64) (string, string) {
	var w, h = noHorizontal, noVertical
	if beforeX < afterX {
		w = toRight
	}
	if beforeX > afterX {
		w = toLeft
	}
	if beforeY < afterY {
		h = toDown
	}
	if beforeY > afterY {
		h = toUp
	}
	return w, h
}

func TestMoveDirection(t *testing.T) {
	type input struct {
		speed, angle float64
		direction    int
	}

	type output struct {
		horizontalMoveDirection, verticalMoveDirection string
	}
	var bananaMoveDirectionTestCase = []struct {
		in  input
		out output
	}{
		// base case: 0, 45 or 90 to the right
		{input{10, 0, gorilla.Right}, output{toRight, noVertical}},
		{input{10, 90, gorilla.Right}, output{noHorizontal, toUp}},
		{input{10, 45, gorilla.Right}, output{toRight, toUp}},

		// base case switched to left
		{input{10, 0, gorilla.Left}, output{toLeft, noVertical}},
		{input{10, 90, gorilla.Left}, output{noHorizontal, toUp}},
		{input{10, 45, gorilla.Left}, output{toLeft, toUp}},

		// negative speed to the gorilla.Right and to thegorilla.Left
		{input{-10, 0, gorilla.Right}, output{toLeft, noVertical}},
		{input{-10, 90, gorilla.Right}, output{noHorizontal, toUp}},
		{input{-10, 45, gorilla.Right}, output{toLeft, toUp}},
		{input{-10, 0, gorilla.Left}, output{toRight, noVertical}},
		{input{-10, 90, gorilla.Left}, output{noHorizontal, toUp}},
		{input{-10, 45, gorilla.Left}, output{toRight, toUp}},

		// negative angle to the gorilla.Right and to thegorilla.Left
		{input{10, 0, gorilla.Right}, output{toRight, noVertical}},
		{input{10, -90, gorilla.Right}, output{noHorizontal, toDown}},
		{input{10, -45, gorilla.Right}, output{toRight, toDown}},
		{input{10, 0, gorilla.Left}, output{toLeft, noVertical}},
		{input{10, -90, gorilla.Left}, output{noHorizontal, toDown}},
		{input{10, -45, gorilla.Left}, output{toLeft, toDown}},

		// both speed and angle negative
		{input{-10, 0, gorilla.Right}, output{toLeft, noVertical}},
		{input{-10, -90, gorilla.Right}, output{noHorizontal, toDown}},
		{input{-10, -45, gorilla.Right}, output{toLeft, toDown}},

		{input{-10, 0, gorilla.Left}, output{toRight, noVertical}},
		{input{-10, -90, gorilla.Left}, output{noHorizontal, toDown}},
		{input{-10, -45, gorilla.Left}, output{toRight, toDown}},

		// angles over 90 degrees
		{input{10, 0, gorilla.Right}, output{toRight, noVertical}},
		{input{10, 90, gorilla.Right}, output{noHorizontal, toUp}},
		{input{10, 45, gorilla.Right}, output{toRight, toUp}},
		{input{10, 135, gorilla.Right}, output{toLeft, toUp}},
		{input{10, 180, gorilla.Right}, output{toLeft, noVertical}},
		{input{10, 225, gorilla.Right}, output{toLeft, toDown}},
		{input{10, 270, gorilla.Right}, output{noHorizontal, toDown}},
		{input{10, 315, gorilla.Right}, output{toRight, toDown}},
		{input{10, 360, gorilla.Right}, output{toRight, noVertical}},
		{input{10, 370, gorilla.Right}, output{toRight, toUp}},
	}

	b := gorilla.NewBanana()
	for _, tc := range bananaMoveDirectionTestCase {
		b.Reset()
		b.SetMoveData(tc.in.speed, tc.in.angle, 0)
		beforeX := b.X
		beforeY := b.Y

		b.Move(tc.in.direction)

		w, h := getChangeDirection(beforeX, beforeY, b.X, b.Y)

		if w != tc.out.horizontalMoveDirection || h != tc.out.verticalMoveDirection {
			dir := "right"
			if tc.in.direction == gorilla.Left {
				dir = "left"
			}
			fmt.Printf("banana: %f,%f", b.X, b.Y)
			t.Fatalf("Banana thrown to the %s with speed %f and angle %f should move %s and %s, but instead it went %s and %s", dir, tc.in.speed, tc.in.angle, tc.out.horizontalMoveDirection, tc.out.verticalMoveDirection, w, h)
		}
	}
}
