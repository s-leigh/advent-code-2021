package day05

import (
	"github.com/s-leigh/advent-code-2021/utils"
	"strconv"
	"strings"
)

const StartAndEndSeparator = " -> "
const XYSeparator = ","
const GridSquareSize = 1000

type coordinates struct {
	x int
	y int
}

type line struct {
	from coordinates
	to   coordinates
}

func day05Question01(filepath string) int {
	lines := parseInput(filepath)
	horizontalAndVertical := horizontalAndVerticalLines(lines)
	grid := makeGrid()
	markGrid(grid, horizontalAndVertical)
	return numberOfPointsOver1(grid)
}

func day05Question02(filepath string) int {
	lines := parseInput(filepath)
	grid := makeGrid()
	markGrid(grid, lines)
	return numberOfPointsOver1(grid)
}

func markGrid(grid *[][]int, lines []line) {
	for _, line := range lines {
		if isHorizontal(line) {
			if line.from.x < line.to.x {
				for i := line.from.x; i <= line.to.x; i++ {
					(*grid)[line.from.y][i]++
				}
			} else {
				for i := line.from.x; i >= line.to.x; i-- {
					(*grid)[line.from.y][i]++
				}
			}
		} else if isVertical(line) {
			if line.from.y < line.to.y {
				for i := line.from.y; i <= line.to.y; i++ {
					(*grid)[i][line.from.x]++
				}
			} else {
				for i := line.from.y; i >= line.to.y; i-- {
					(*grid)[i][line.from.x]++
				}
			}
		} else {
			markDiagonal(grid, line)
		}
	}
}

func markDiagonal(grid *[][]int, line line) {
	yCounter := 0
	alterYCounter := func() {
		if line.from.y < line.to.y {
			yCounter++
		} else {
			yCounter--
		}
	}
	if line.from.x < line.to.x {
		for i := line.from.x; i <= line.to.x; i++ {
			(*grid)[line.from.y+yCounter][i]++
			alterYCounter()
		}
	} else {
		for i := line.from.x; i >= line.to.x; i-- {
			(*grid)[line.from.y+yCounter][i]++
			alterYCounter()
		}
	}
}

func numberOfPointsOver1(grid *[][]int) int {
	total := 0
	for _, line := range *grid {
		for _, entry := range line {
			if entry >= 2 {
				total ++
			}
		}
	}
	return total
}

func parseInput(filepath string) []line {
	input := utils.SplitByNewLine(filepath)
	var lines []line
	for _, inputLine := range input {
		stringCoords := strings.Split(inputLine, StartAndEndSeparator) // ["1,2","2,3"]
		for i := 0; i < len(stringCoords); i += 2 {
			stringXYFrom := strings.Split(stringCoords[i], XYSeparator) // ["1","2"]
			stringXYTo := strings.Split(stringCoords[i+1], XYSeparator)
			xFrom := stringToInt(stringXYFrom[0])
			yFrom := stringToInt(stringXYFrom[1])
			xTo := stringToInt(stringXYTo[0])
			yTo := stringToInt(stringXYTo[1])
			lines = append(lines, line{coordinates{xFrom, yFrom}, coordinates{xTo, yTo}})
		}
	}
	return lines
}

func horizontalAndVerticalLines(lines []line) []line {
	var hAndV []line
	for _, l := range lines {
		if isHorizontal(l) || isVertical(l) {
			hAndV = append(hAndV, l)
		}
	}
	return hAndV
}

func isHorizontal(line line) bool {
	return line.from.y == line.to.y
}

func isVertical(line line) bool {
	return line.from.x == line.to.x
}

func makeGrid() *[][]int {
	grid := make([][]int, GridSquareSize)
	for i := range grid {
		grid[i] = make([]int, GridSquareSize)
	}
	return &grid
}

func stringToInt(str string) int {
	i, err := strconv.Atoi(str)
	utils.Check(err)
	return i
}
