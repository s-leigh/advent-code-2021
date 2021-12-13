package day13

import (
	"fmt"
	"github.com/s-leigh/advent-code-2021/utils"
	"strings"
)

func day13Question01(filepath string) int {
	coOrds, folds := parseInput(filepath)
	fmt.Printf("%v\n%v\n", coOrds, folds)
	return -1
}

// Returns e.g. [[1,2],[2,3]], [["y", "7"],["x","4"]]
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