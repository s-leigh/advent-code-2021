package day03

import "testing"

func TestDay03Question01SampleInput(t *testing.T) {
	res := day03Question01("./testInput01.txt")
	if res != 198 {
		t.Error("Expected 198, got", res)
	}
}

func TestDay03Question01(t *testing.T) {
	res := day03Question01("./input.txt")
	if res != 3374136 {
		t.Error("Expected 3374136, got", res)
	}
}