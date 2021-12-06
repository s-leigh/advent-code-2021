package utils

import (
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// copied from https://gobyexample.com/reading-files
func ReadFile(filepath string) string {
	dat, err := os.ReadFile(filepath)
	check(err)
	return string(dat)
}

func SplitByNewLine(filepath string) []string {
	return strings.Split(ReadFile(filepath), "\n")
}

func SplitBy(filepath string, separator string) []int {
	stringSlice := strings.Split(ReadFile(filepath), separator)
	intSlice := make([]int, len(stringSlice))
	for i, s := range stringSlice {
		intSlice[i] = StringToInt(s)
	}
	return intSlice
}

func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	check(err)
	return i
}