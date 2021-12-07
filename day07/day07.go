package day07

import (
	"github.com/s-leigh/advent-code-2021/utils"
)

func day07Question01(filepath string) uint {
	crabPositions := utils.SplitBy(filepath, ",")
	fuelUseStats := findFuelUseForPositions(&crabPositions, linearFuelUseToGetToPosition)
	shortestJourneyPosition := findShortestJourney(&fuelUseStats)
	return linearFuelUseToGetToPosition(shortestJourneyPosition, &crabPositions)
}

func day07Question02(filepath string) uint {
	crabPositions := utils.SplitBy(filepath, ",")
	fuelUseStats := findFuelUseForPositions(&crabPositions, exponentialFuelUseToGetToPosition)
	shortestJourneyPosition := findShortestJourney(&fuelUseStats)
	return exponentialFuelUseToGetToPosition(shortestJourneyPosition, &crabPositions)
}

func linearFuelUseToGetToPosition(position int, crabPositions *[]int) uint {
	total := uint(0)
	for _, crabPosition := range *crabPositions {
		total += absoluteValue(crabPosition - position)
	}
	return total
}

func exponentialFuelUseToGetToPosition(position int, crabPositions *[]int) uint {
	total := uint(0)
	for _, crabPosition := range *crabPositions {
		steps := absoluteValue(crabPosition - position)
		total += exponentialFuelGauge(steps)
	}
	return total
}

func exponentialFuelGauge(steps uint) uint {
	if steps == 0 {
		return uint(0)
	}
	return steps + exponentialFuelGauge(steps - 1)
}

func findFuelUseForPositions(positions *[]int, fuelUseFn func(int, *[]int) uint) map[int] uint {
	uniquePositions := unique(positions)
	positionsAndFuelUse := make(map[int]uint, len(uniquePositions))
	for i := range *positions {
		positionsAndFuelUse[i] += fuelUseFn(i, positions)
	}
	return positionsAndFuelUse
}

func findShortestJourney(fuelUse *map[int]uint) int {
	shortestJourneyPos := -1
	shortestJourneyFuel := ^uint(0) // max int
	for pos, fuel := range *fuelUse {
		if fuel < shortestJourneyFuel {
			shortestJourneyPos = pos
			shortestJourneyFuel = fuel
		}
	}
	return shortestJourneyPos
}

func unique(intSlice *[]int) []int {
	keys := make(map[int]bool)
	var list []int
	for _, entry := range *intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func absoluteValue(i int) uint {
	if i < 0 {
		return uint(-i)
	}
	return uint(i)
}
