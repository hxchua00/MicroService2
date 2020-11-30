package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 2)

	if result != 3 {
		t.Errorf("Result incorrect, got: %d, want: %d", result, 3)
	}
}
