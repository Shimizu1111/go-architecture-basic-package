package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(5, 6)
	if result != 11 {
		t.Errorf("Expected 11, but got %d", result)
	}
}
