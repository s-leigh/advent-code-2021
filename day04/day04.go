package day04

import (
	"github.com/s-leigh/advent-code-2021/utils"
	"strconv"
	"strings"
)

const BOARD_SIZE = 5

type board struct {
	rows [][]int
	columns [][]int
}

func day04Question01(filepath string) int {
	_ = parseBoards(filepath)
	return -1
}

func parseBoards(filepath string) []board {
	stringRows := utils.SplitByNewLine(filepath) // ["1  2  3", "2  3  4"]
	sliceRows := make([][]int, len(stringRows))
	for i, str := range stringRows {
		intermediate := strings.Split(str, " ")
		for _, str := range intermediate {
			if str != "" && str != " " && str != "  " {
				asInt, _ := strconv.Atoi(str)
				sliceRows[i] = append(sliceRows[i], asInt) // [[1,2,...],...]
			}
		}
	}
	numberOfBoards := 1 // fenceposting as we'll count number of newlines
	for i, row := range sliceRows {
		if len(row) == 0 { // new line between boards, but we know how big the boards are
			numberOfBoards++
			sliceRows = append(sliceRows[:i], sliceRows[i+1:]...)
		}
	}

	boards := make([]board, numberOfBoards)
	for boardIndex := 0; boardIndex < numberOfBoards; boardIndex++ {
		rows := getRows(sliceRows, boardIndex)
		columns := getColumns(rows)
		boards[boardIndex].rows = rows
		boards[boardIndex].columns = columns
	}
	return boards
}

func getRows(allRows [][]int, boardIndex int) [][]int {
	// boardIndex = 0, subset = 0-4
	// boardIndex = 1, subset = 5-9
	subsliceFrom := boardIndex * BOARD_SIZE
	subsliceTo := subsliceFrom + BOARD_SIZE
	return allRows[subsliceFrom:subsliceTo]
}

func getColumns(rowsSubset [][]int) [][]int {
	columns := make([][]int, BOARD_SIZE)
	for columnNumber := 0; columnNumber < BOARD_SIZE; columnNumber++ {
		for _, row := range rowsSubset {
			columns[columnNumber] = append(columns[columnNumber], row[columnNumber])
		}
	}
	return columns
}
