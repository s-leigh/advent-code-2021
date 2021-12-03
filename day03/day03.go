package day03

import (
	"github.com/s-leigh/advent-code-2021/utils"
	"strconv"
	"strings"
)

func day03Question01(filePath string) int {
	input := utils.SplitByNewLine(filePath)
	binaryNumberLength := len(input[0])

	bitSum := make([]int, binaryNumberLength)
	for _, binary := range input {
		for j, bit := range []rune(binary) {
			bit, _ := strconv.Atoi(string(bit))
			bitSum[j] += bit
		}
	}

	mostCommonBits := make([]string, binaryNumberLength)
	for i, sum := range bitSum {
		if sum > len(input) / 2 {
			mostCommonBits[i] = "1"
		} else {
			mostCommonBits[i] = "0"
		}
	}

	var epsilonString string
	for _, r := range mostCommonBits {
		if string(r) == "1" {
			epsilonString += "0"
		} else {
			epsilonString += "1"
		}
	}

	epsilon, _ := strconv.ParseInt(epsilonString, 2, 16)
	commonBitsString := strings.Join(mostCommonBits, "")
	gamma, _ := strconv.ParseInt(commonBitsString, 2, 16)

	return int(gamma) * int(epsilon)
}