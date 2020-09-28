package main

import "testing"

func TestCalculate(t *testing.T) {
	if Calculate(3) != 6 {
		t.Error("Expected 6")
	}
}

