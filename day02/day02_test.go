package day02

import "testing"

func TestDay02Question01SampleInput(t *testing.T) {
	res := day02Question01("./testInput.txt")
	if res != 150 {
		t.Fail()
	}
}

func TestDay02Question01(t *testing.T) {
	res := day02Question01("./input.txt")
	println(res)
	if res != 1728414 {
		t.Fail()
	}
}