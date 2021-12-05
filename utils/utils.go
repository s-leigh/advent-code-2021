package utils

import (
	"os"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// copied from https://gobyexample.com/reading-files
func ReadFile(filepath string) string {
	dat, err := os.ReadFile(filepath)
	Check(err)
	return string(dat)
}

func SplitByNewLine(filepath string) []string {
	return strings.Split(ReadFile(filepath), "\n")
}