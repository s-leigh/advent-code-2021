package day09

import "testing"

func TestDay09Question01SampleInput(t *testing.T) {
	expected := 15
	res := day09Question01("testInput.txt")
	if res != expected {
		t.Error("Expected", expected, "got", res)
	}
}

func TestDay09Question01(t *testing.T) {
	expected := 489
	res := day09Question01("input.txt")
	if res != expected {
		t.Error("Expected", expected, "got", res)
	}
}
