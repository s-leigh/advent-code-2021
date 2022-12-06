package day04

import (
	"github.com/s-leigh/advent-code-2021/utils"
	"reflect"
	"strconv"
	"strings"
)

const BoardSize = 5
const WinningNumberReplacement = -1

type board struct {
	rows [][]int
	columns [][]int
}

func day04Question01(boardsFilepath string, numbersFilepath string) int {
	boards := parseBoards(boardsFilepath)
	numbers := parseNumbers(numbersFilepath)
	var winningBoards []board
	var winningNumber int
	for _, number := range numbers {
		markOffNumber(number, &boards)
		boardHasWon, boardsThatWon := checkWin(boards)
		if boardHasWon {
			winningBoards = boardsThatWon
			winningNumber = number
			break
		}
	}
	if len(winningBoards) > 1 {
		panic("More than one winner!")
	}
	return sumBoard(winningBoards[0].rows) * winningNumber
}

func day04Question02(boardsFilepath string, numbersFilepath string) int {
	boards := parseBoards(boardsFilepath)
	numbers := parseNumbers(numbersFilepath)
	losingBoard, losingNumber := markOffUntilLoser(&boards, numbers)
	return sumBoard(losingBoard.rows) * losingNumber
}

func markOffUntilLoser(boards *[]board, numbers []int) (board, int) {
	markOffNumber(numbers[0], boards)
	boardHasWon, boardsThatWon := checkWin(*boards)
	if boardHasWon {
		if len(*boards) == 1 {
			return (*boards)[0], numbers[0]
		}
		for i, b := range *boards {
			for _, wb := range boardsThatWon {
				if reflect.DeepEqual(b, wb) {
					if i+1 > len(*boards) {
						*boards = (*boards)[:len(*boards)-1]
					} else {
						*boards = append((*boards)[:i], (*boards)[i+1:]...)
					}
				}
			}
		}
	}
	return markOffUntilLoser(boards, numbers[1:])
}

func sumBoard(boardRows [][]int) int {
	total := 0
	for _, row := range boardRows {
		for _, num := range row {
			if num != WinningNumberReplacement {
				total += num
			}
		}
	}
	return total
}

func markOffNumber(number int, boards *[]board) {
	for _, board := range *boards {
		for _, column := range board.columns {
			for k, numberOnBoard := range column {
				if numberOnBoard == number {
					column[k] = WinningNumberReplacement
				}
			}
		}
		for _, row := range board.rows {
			for k, numberOnBoard := range row {
				if numberOnBoard == number {
					row[k] = WinningNumberReplacement
				}
			}
		}
	}
}

func checkWin(boards []board) (bool, []board) {
	hasWon := func(rowOrColumn []int) bool {
		counter := 0
		for _, entry := range rowOrColumn {
			if entry == WinningNumberReplacement {
				counter++
			}
		}
		return counter == len(rowOrColumn)
	}
	var winners []board // more than one board can win at a time
	for _, board := range boards {
		for _, col := range board.columns {
			if hasWon(col) {
				winners = append(winners, board)
			}
		}
		for _, row := range board.rows {
			if hasWon(row) {
				winners = append(winners, board)
			}
		}
	}
	if len(winners) > 0 {
		return true, winners
	}
	return false, []board{}
}

func parseNumbers(filepath string) []int {
	rawNumbers := utils.ReadFile(filepath)
	stringNumbers := strings.Split(rawNumbers, ",")
	numbers := make([]int, len(stringNumbers))
	for i, strNum := range stringNumbers {
		numbers[i], _ = strconv.Atoi(strNum)
	}
	return numbers
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
	subsliceFrom := boardIndex * BoardSize
	subsliceTo := subsliceFrom + BoardSize
	return allRows[subsliceFrom:subsliceTo]
}

func getColumns(rowsSubset [][]int) [][]int {
	columns := make([][]int, BoardSize)
	for columnNumber := 0; columnNumber < BoardSize; columnNumber++ {
		for _, row := range rowsSubset {
			columns[columnNumber] = append(columns[columnNumber], row[columnNumber])
		}
	}
	return columns
}
