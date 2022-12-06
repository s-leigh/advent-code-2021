package day01

import (
	"strconv"
	"strings"
)

func day01Question01(rawInput string) int {
	input := convertRawInput(rawInput)
	return numberOfIncreases(input)
}

func day01Question02(rawInput string) int {
	input := convertRawInput(rawInput)
	windows := splitIntoWindowsOfThree(input, [][]int{})
	windowsSums := sumsOfWindows(windows, []int{})
	return numberOfIncreases(windowsSums)
}

func convertRawInput(rawInput string) []int {
	rawInputSlice := strings.Split(rawInput, "\n")
	var input = make([]int, len(rawInputSlice))
	for i, s := range rawInputSlice {
		input[i], _ = strconv.Atoi(s)
	}
	return input
}

func numberOfIncreases(arr []int) int {
	count := 0
	for i, elem := range arr[1:] {
		if elem > arr[i] {
			count++
		}
	}
	return count
}

func splitIntoWindowsOfThree(arr []int, builtArray [][]int) [][]int {
	if len(arr) == 3 {
		return append(builtArray, arr)
	}
	return splitIntoWindowsOfThree(arr[1:], append(builtArray, arr[0:3]))
}

func sumsOfWindows(arr [][]int, builtArray []int) []int {
	if len(arr) == 1 {
		return append(builtArray, sumWindow(arr[0]))
	}
	return sumsOfWindows(arr[1:], append(builtArray, sumWindow(arr[0])))
}

func sumWindow(window []int) int {
	total := 0
	for _, val := range window {
		total += val
	}
	return total
}
