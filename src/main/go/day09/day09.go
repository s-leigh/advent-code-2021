package day09

import (
	"github.com/s-leigh/advent-code-2021/utils"
	"strings"
)

type Point struct {
	value int
	up *Point
	down *Point
	left *Point
	right *Point
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

func lowestPoints(points [][]Point) []Point {
	var lowest []Point
	for _, pRow := range points {
		for _, p := range pRow {
			if p.valueLessThanNeighbours() {
				lowest = append(lowest, p)
			}
		}
	}
	return lowest
}

func (p Point)valueLessThanNeighbours() bool {
	outcome := true
	for _, neighbour := range []*Point{p.up, p.down, p.left, p.right} {
		if neighbour != nil && neighbour.value <= p.value {
			outcome = false
		}
	}
	return outcome
}

func convertToPoints(rawInput [][]int) [][]Point {
	twoDimensionalSliceValueOrDefault := func(slice [][]Point, firstIndex int, secondIndex int) *Point {
		if firstIndex < len(slice) && firstIndex >= 0 {
			if secondIndex < len(slice[firstIndex]) && secondIndex >= 0 {
				return &slice[firstIndex][secondIndex]
			}
		}
		return nil
	}
	var points [][]Point
	// Initialise with nil directions first
	for _, line := range rawInput {
		pointLine := make([]Point, len(line))
		for xIndex, value := range line {
			pointLine[xIndex] = Point{
				value,
				nil,
				nil,
				nil,
				nil,
			}
		}
		points = append(points, pointLine)
	}
	// Now we can create pointers to adjacent values
	for yIndex, line := range points {
		for xIndex, point := range line {
			points[yIndex][xIndex] = Point{
				point.value,
				twoDimensionalSliceValueOrDefault(points, yIndex-1, xIndex),
				twoDimensionalSliceValueOrDefault(points, yIndex+1, xIndex),
				twoDimensionalSliceValueOrDefault(points, yIndex, xIndex-1),
				twoDimensionalSliceValueOrDefault(points, yIndex, xIndex+1),
			}
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