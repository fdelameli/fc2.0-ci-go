package main

import "testing"

func TestSum(t *testing.T) {
	expected := 30
	total := Sum(15, 15)

	if total != expected {
		t.Errorf("The sum result is invalid: Result: %d. Expected: %d", total, expected)
	}
}
