package day02

import (
	"fmt"
	"testing"
)

func TestDay02Question01SampleInput(t *testing.T) {
	res := day02Question01("./testInput01.txt")
	if res != 150 {
		t.Fail()
	}
}

func TestDay02Question01(t *testing.T) {
	res := day02Question01("./input.txt")
	fmt.Println(res)
	if res != 1728414 {
		t.Fail()
	}
}

func TestDay02Question02SampleInput(t *testing.T) {
	res := day02Question02("./testInput02.txt")
	if res != 900 {
		t.Fail()
	}
}

func TestDay02Question02(t *testing.T) {
	res := day02Question02("./input.txt")
	fmt.Println(res)
	if res != 1765720035 {
		t.Fail()
	}
}
