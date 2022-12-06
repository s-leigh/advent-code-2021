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

func TestDay07Question02SampleInput(t *testing.T) {
	expected := uint(168)
	res := day07Question02("testInput.txt")
	if res != expected {
		t.Error("Expected", expected, "got", res)
	}
}

func TestDay07Question02(t *testing.T) {
	expected := uint(86397080)
	res := day07Question02("input.txt")
	if res != expected {
		t.Error("Expected", expected, "got", res)
	}
}
