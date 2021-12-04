package day04

import "testing"

func TestDay04Question01SampleInput(t *testing.T) {
	res := day04Question01("./testInputBoards.txt", "./testInputNumbers.txt")
	if res != 4512 {
		t.Fail()
	}
}

func TestDay04Question01(t *testing.T) {
	res := day04Question01("./inputBoards.txt", "./inputNumbers.txt")
	if res != 58838 {
		t.Error("Got", res)
	}
}

func TestDay04Question02SampleInput(t *testing.T) {
	res := day04Question02("./testInputBoards.txt", "./testInputNumbers.txt")
	if res != 1924 {
		t.Error("Expected 1924, got", res)
	}
}

func TestDay04Question02(t *testing.T) {
	res := day04Question02("./inputBoards.txt", "./inputNumbers.txt")
	if res != 6256 {
		t.Fail()
	}
}