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
