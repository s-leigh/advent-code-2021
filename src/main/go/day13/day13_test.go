package day13

import "testing"

func TestDay13Question01TestInput(t *testing.T) {
	expected := 17
	res := day13Question01("testInput.txt")
	if res != expected {
		t.Error("Expected", expected, "got", res)
	}
}

func TestDay13Question01(t *testing.T) {
	expected := 737
	res := day13Question01("input.txt")
	if res != expected {
		t.Error("Expected", expected, "got", res)
	}
}