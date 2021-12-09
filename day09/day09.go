package day09

import (
	"github.com/s-leigh/advent-code-2021/utils"
	"strings"
)

const DEFAULT_VALUE = 999999

type Point struct {
	value int
	up int
	down int
	left int
	right int
}

func day09Question01(filepath string) int {
	points := convertToPoints(parseRawInput(filepath))
	lowest := lowestPoints(points)
	sum := 0
	for _, lp := range lowest {
		sum += lp.value + 1
	}
	return sum
}

func lowestPoints(points []Point) []Point {
	var lowest []Point
	for _, p := range points {
		if p.value < p.up && p.value < p.down && p.value < p.left && p.value < p.right {
			lowest = append(lowest, p)
		}
	}
	return lowest
}

func convertToPoints(rawInput [][]int) []Point {
	oneDimensionalSliceValueOrDefault := func(slice []int, index int) int {
		if index < len(slice) && index >= 0 {
			return slice[index]
		}
		return DEFAULT_VALUE
	}
	twoDimensionalSliceValueOrDefault := func(slice [][]int, firstIndex int, secondIndex int) int {
		if firstIndex < len(slice) && firstIndex >= 0 {
			if secondIndex <= len(slice[firstIndex]) && secondIndex >= 0 {
				return slice[firstIndex][secondIndex]
			}
		}
		return DEFAULT_VALUE
	}
	var points []Point
	for yIndex, line := range rawInput {
		for xIndex, value := range line {
			points = append(points, Point{
				value,
				twoDimensionalSliceValueOrDefault(rawInput, yIndex-1, xIndex),
				twoDimensionalSliceValueOrDefault(rawInput, yIndex+1, xIndex),
				oneDimensionalSliceValueOrDefault(line, xIndex-1),
				oneDimensionalSliceValueOrDefault(line, xIndex+1),
			})
		}
	}
	return points
}

func parseRawInput(filepath string) [][]int {
	var ints [][]int
	lines := utils.SplitByNewLine(filepath)
	for _, l := range lines {
		stringChars := strings.Split(l, "")
		lineInts := make([]int, len(stringChars))
		for i, char := range stringChars {
			lineInts[i] = utils.StringToInt(char)
		}
		ints = append(ints, lineInts)
	}
	return ints
}