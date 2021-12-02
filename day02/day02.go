package day02

import (
	"github.com/s-leigh/advent-code-2021/utils"
	"strconv"
	"strings"
)

func directionAndScale(input string) (string, int) {
	split := strings.Split(input, " ")
	scale, _ := strconv.Atoi(split[1])
	return split[0], scale
}

func day02Question01(filePath string) int {
	input := utils.SplitByNewLine(filePath)
	var horiz, depth int
	for _, inputLine := range input {
		direction, scale := directionAndScale(inputLine)
		switch direction {
		case "forward": horiz += scale
		case "down": depth += scale
		case "up": depth -= scale
		}
	}
	return horiz * depth
}
