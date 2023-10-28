package calculator

import (
	"testing"
)

func TestSum(t *testing.T) {
	result := Sum(1, 2, 3, 4, 5)
	if result != 15 {
		t.Errorf("Expected 15, but got %d", result)
	}
}
