package day06

import (
	"github.com/s-leigh/advent-code-2021/utils"
)

const separator = ","

func day06Question01(filepath string, numberOfDays int) int {
	fish := utils.SplitBy(filepath, separator)
	breedFish(&fish, numberOfDays)
	return len(fish)
}

func breedFish(fish *[]int, days int) {
	for d := 0; d < days; d++ {
		for i, f := range *fish {
			if f == 0 {
				(*fish)[i] = 6
				*fish = append(*fish, 8)
			} else {
				(*fish)[i]--
			}
		}
	}
}
