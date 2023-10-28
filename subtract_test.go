package calculator

import (
	"testing"
)

func TestSubtract(t *testing.T) {
	result := Subtract(9, 4)
	if result != 5 {
		t.Errorf("Expected 5, but got %d", result)
	}
}
