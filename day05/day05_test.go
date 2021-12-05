package day05

import "testing"

func TestDay05Question01SampleInput(t *testing.T) {
	res := day05Question01("testInput.txt")
	if res != 5 {
		t.Error("Got", res)
	}
}

func TestDay05Question01(t *testing.T) {
	res := day05Question01("input.txt")
	if res != 5167 {
		t.Fail()
	}
}