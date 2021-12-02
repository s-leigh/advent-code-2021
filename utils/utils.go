package utils

import (
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// copied from https://gobyexample.com/reading-files
func readFile(filepath string) string {
	dat, err := os.ReadFile(filepath)
	check(err)
	return string(dat)
}

func SplitByNewLine(filepath string) []string {
	return strings.Split(readFile(filepath), "\n")
}