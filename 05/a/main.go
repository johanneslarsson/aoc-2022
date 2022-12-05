package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Instruction struct {
	Count int
	From  int
	To    int
}

func getResult(crates [][]rune) string {
	result := ""
	for _, crate := range crates {
		result += string(crate[len(crate)-1:])
	}
	return result
}

func reverse(r []rune) {
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
}

func moveCrates(crates [][]rune, instructions []Instruction, preservePickOrder bool) {
	for _, inst := range instructions {
		fromCrate := crates[inst.From]
		index := len(fromCrate) - inst.Count
		lastElements := fromCrate[index:]     // get last
		crates[inst.From] = fromCrate[:index] // remove last
		if preservePickOrder {
			reverse(lastElements)
		}
		crates[inst.To] = append(crates[inst.To], lastElements...)
	}
}

func getPartOne(crates [][]rune, instructions []Instruction) string {
	moveCrates(crates, instructions, true)

	result := getResult(crates)
	return result
}

func getPartTwo(crates [][]rune, instructions []Instruction) string {
	moveCrates(crates, instructions, false)

	result := getResult(crates)
	return result
}

func getRows(filename string) ([][]rune, []Instruction) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	crates := make([][]rune, 0)
	instructions := make([]Instruction, 0)
	for scanner.Scan() {
		row := scanner.Text()
		if strings.Contains(row, "[") {
			for i := 1; i < len(row); i += 4 {
				j := (i - 1) / 4
				if j+1 > len(crates) {
					crates = append(crates, make([]rune, 0))
				}
				if string(row[i]) != " " {
					crates[j] = append(crates[j], int32(row[i]))
				}
			}
		} else if strings.Contains(row, "move") {
			var inst Instruction
			fmt.Sscanf(row, "move %d from %d to %d",
				&inst.Count, &inst.From, &inst.To)
			inst.From -= 1
			inst.To -= 1
			instructions = append(instructions, inst)
		}
	}
	// Reverse the crates
	for _, r := range crates {
		reverse(r)
	}
	return crates, instructions
}

func main() {
	fmt.Println("Part one:", getPartOne(getRows("../input.txt")))
	fmt.Println("Part two:", getPartTwo(getRows("../input.txt")))
}
