package day07

import "testing"

func TestDay07Question01SampleInput(t *testing.T) {
	expected := uint(37)
	res := day07Question01("testInput.txt")
	if res != expected {
		t.Error("Expected", expected, "got", res)
	}
}

func TestDay07Question01(t *testing.T) {
	expected := uint(329389)
	res := day07Question01("input.txt")
	if res != expected {
		t.Error("Expected", expected, "got", res)
	}
}