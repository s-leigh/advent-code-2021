package day06

import "testing"

func TestDay06Question01SampleInput01(t *testing.T) {
	days := 18
	expected := 26
	res := day06Question01("testInput.txt", days)
	if res != expected {
		t.Error("Expected", expected, "got", res)
	}
}

func TestDay06Question01SampleInput02(t *testing.T) {
	days := 80
	expected := 5934
	res := day06Question01("testInput.txt", days)
	if res != expected {
		t.Error("Expected", expected, "got", res)
	}
}

func TestDay06Question01(t *testing.T) {
	days := 80
	expected := 353079
	res := day06Question01("input.txt", days)
	if res != expected {
		t.Error("Expected", expected, "got", res)
	}
}