package day13

import (
	"github.com/s-leigh/advent-code-2021/utils"
	"strings"
)

type Axis int

const (
	x Axis = iota
	y
)

func (a Axis) String() string {
	if a == x {
		return "x"
	}
	return "y"
}

func day13Question01(filepath string) int {
	coOrds, folds := parseInput(filepath)
	f := folds[0] // Q1 only cares about the first fold
	degreeOfFoldage := utils.StringToInt(f[1])
	if f[0] == x.String() {
		fold(&coOrds, x, degreeOfFoldage)
	} else {
		fold(&coOrds, y, degreeOfFoldage)
	}
	return len(coOrds) - numberOfDuplicates(&coOrds, 0)
}

func fold(coOrdinates *[][]int, axis Axis, degree int) {
	foldingCoordinateIndex := -1
	if axis == x {
		foldingCoordinateIndex = 0
	} else {
		foldingCoordinateIndex = 1
	}
	for _, coOrd := range *coOrdinates {
		if coOrd[foldingCoordinateIndex] > degree {
			coOrd[foldingCoordinateIndex] = coOrd[foldingCoordinateIndex] - ((coOrd[foldingCoordinateIndex] - degree) * 2)
		}
	}
}

func numberOfDuplicates(coOrdinates *[][]int, counter int) int {
	if len(*coOrdinates) == 1 {
		return counter
	}
	x, xs := (*coOrdinates)[0], (*coOrdinates)[1:]
	for _, coOrd := range xs {
		if coOrd[0] == x[0] && coOrd[1] == x[1] {
			counter++
		}
	}
	return numberOfDuplicates(&xs, counter)
}

// Returns e.g. [[1,2],[2,3]], [["y","7"],["x","4"]]
func parseInput(filepath string) ([][]int, [][]string) {
	stringInput := utils.SplitByNewLine(filepath)
	var coOrds [][]int
	newLineIndex := -1
	for i, str := range stringInput {
		if str == "" {
			newLineIndex = i
			break
		}
		xAndY := strings.Split(str, ",")
		coOrds = append(coOrds, []int{utils.StringToInt(xAndY[0]), utils.StringToInt(xAndY[1])})
	}
	foldInstructionsStrings := stringInput[newLineIndex + 1:]
	parsedFoldInstructions := make([][]string, len(foldInstructionsStrings))
	for j, str := range foldInstructionsStrings {
		simplified := strings.Split(str, "fold along ") // " y=7"
		sliced := strings.Split(simplified[1], "=")
		parsedFoldInstructions[j] = []string{sliced[0], sliced[1]}
	}
	return coOrds, parsedFoldInstructions
}