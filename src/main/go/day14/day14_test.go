package day14

import "testing"

func TestDay14Question01TestInput(t *testing.T) {
	expected := 1588
	res := day14Question01("testInput.txt", 10)
	if res != expected {
		t.Error("Expected", expected, "actual", res)
	}
}

func TestDay14Question01(t *testing.T) {
	expected := 3009
	res := day14Question01("input.txt", 10)
	if res != expected {
		t.Error("Expected", expected, "actual", res)
	}
}
