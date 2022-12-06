package day14

import (
	"fmt"
	"github.com/s-leigh/advent-code-2021/utils"
	"sort"
	"strings"
)

type Instruction struct {
	match  string
	insert string
}

type Insert struct {
	index       int
	instruction Instruction
}

func day14Question01(filepath string, iterations int) int {
	polymer, instructions := parseInput(filepath)
	fmt.Println(polymer)
	fmt.Println(instructions)
	var inserts []Insert
	for i := 0; i < iterations; i++ {
		inserts = sortedPolymerInserts(polymer, instructions, []Insert{})
		polymer = applySortedInserts(polymer, inserts)
	}

	mostCommon, leastCommon := mostAndLeastCommonElementCounts(polymer)
	return mostCommon - leastCommon
}

func mostAndLeastCommonElementCounts(polymer string) (int, int) {
	var counts = map[string]int{}
	splitPolymer := strings.Split(polymer, "")
	counts = countElements(splitPolymer, counts)
	var values []int
	for _, val := range counts {
		if val != 0 {
			values = append(values, val)
		}
	}
	sort.SliceStable(values, func(i, j int) bool {
		return values[i] > values[j]
	})
	return values[0], values[len(values)-1]
}

func countElements(input []string, counts map[string]int) map[string]int {
	if len(input) == 1 {
		return counts
	}
	x, xs := input[0], input[1:]
	if _, ok := counts[x]; !ok {
		counts[x] = 1
		for _, s := range xs {
			if s == x {
				counts[s]++
			}
		}
	}
	return countElements(input[1:], counts)
}

func applySortedInserts(polymer string, inserts []Insert) string {
	//fmt.Printf("insert %v into polymer %v\n", inserts[0], polymer)
	polymer = fmt.Sprintf("%s%s%s", polymer[:inserts[0].index], inserts[0].instruction.insert, polymer[inserts[0].index:])
	//fmt.Println(polymer)
	if len(inserts) == 1 {
		return polymer
	}
	return applySortedInserts(polymer, inserts[1:])
}

func sortedPolymerInserts(polymer string, instructions []Instruction, inserts []Insert) []Insert {
	for i := 0; i < len(polymer)-1; i++ {
		if polymer[i:i+2] == instructions[0].match {
			inserts = append(inserts, Insert{i + 1, instructions[0]})
		}
	}
	if len(instructions) == 1 {
		sort.SliceStable(inserts, func(i, j int) bool {
			return inserts[i].index > inserts[j].index
		})
		return inserts
	}
	return sortedPolymerInserts(polymer, instructions[1:], inserts)
}

func parseInput(filepath string) (string, []Instruction) {
	lines := utils.SplitByNewLine(filepath)
	polymer := lines[0]
	instructions := make([]Instruction, len(lines[2:]))
	for i, op := range lines[2:] {
		split := strings.Split(op, " -> ")
		instructions[i].match = split[0]
		instructions[i].insert = split[1]
	}
	return polymer, instructions
}
