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

	mostCommonBits := make([]int, binaryNumberLength)
	for i, sum := range bitSum {
		if sum > len(input) / 2 {
			mostCommonBits[i] = 1
		} else {
			mostCommonBits[i] = 0
		}
	}

	//var commonBits []int
	//for _, binary := range input {
	//	bin, _ := strconv.Atoi(binary)
	//	ones := bits.OnesCount(uint(bin))
	//	if ones > len(binary) - ones {
	//		commonBits = append(commonBits, 1)
	//	} else {
	//		commonBits = append(commonBits, 0)
	//	}
	//}

	var commonBitsStringSlice []string
	for _, bit := range mostCommonBits {
		commonBitsStringSlice = append(commonBitsStringSlice, strconv.Itoa(bit))
	}
	commonBitsString := strings.Join(commonBitsStringSlice, "")
	println(commonBitsString)
	var epsilonString string
	for _, r := range commonBitsString {
		if string(r) == "1" {
			epsilonString += "0"
		} else {
			epsilonString += "1"
		}
	}
	println("eps", epsilonString)

	epsilon, _ := strconv.ParseInt(epsilonString, 2, 16)

	gamma, _ := strconv.ParseInt(commonBitsString, 2, 16)

	return int(gamma) * int(epsilon)
}