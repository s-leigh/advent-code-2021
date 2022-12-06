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

func TestDay05Question02SampleInput(t *testing.T) {
	res := day05Question02("testInput.txt")
	if res != 12 {
		t.Error("Got", res)
	}
}

func TestDay05Question02(t *testing.T) {
	res := day05Question02("input.txt")
	if res != 17604 {
		t.Fail()
	}
}