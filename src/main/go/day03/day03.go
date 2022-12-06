package day03

import (
	"github.com/s-leigh/advent-code-2021/utils"
	"strconv"
	"strings"
)

func day03Question01(filePath string) int {
	input := utils.SplitByNewLine(filePath)
	mostCommonBits := mostCommonBits(input)
	gamma := binarySliceToInt(mostCommonBits)
	epsilon := binarySliceToInt(flipBits(mostCommonBits))
	return gamma * epsilon
}

func binarySliceToInt(binary []int) int {
	var stringSlice = make([]string, len(binary))
	for i, bin := range binary {
		binAsString := strconv.Itoa(bin)
		stringSlice[i] = binAsString
	}
	asString := strings.Join(stringSlice, "")
	retInt, _ := strconv.ParseInt(asString, 2, 16)
	return int(retInt)
}

func mostCommonBits(input []string) []int {
	binaryNumberLength := len(input[0])
	bitSum := make([]int, binaryNumberLength)
	for _, binary := range input {
		for j, bit := range []rune(binary) {
			bit, _ := strconv.Atoi(string(bit))
			bitSum[j] += bit
		}
	}

	commonestBits := make([]int, binaryNumberLength)
	for i, sum := range bitSum {
		if sum > len(input) / 2 {
			commonestBits[i] = 1
		} else {
			commonestBits[i] = 0
		}
	}
	return commonestBits
}

func flipBits(binary []int) []int {
	flipped := make([]int, len(binary))
	for i, bit := range binary {
		if bit == 1 {
			flipped[i] = 0
		} else {
			flipped[i] = 1
		}
	}
	return flipped
}