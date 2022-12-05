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

func getResult(crateStacks [][]rune) (result string) {
	for _, stack := range crateStacks {
		result += string(stack[len(stack)-1:])
	}
	return result
}

func reverse(r []rune) {
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
}

func moveCrates(crateStacks [][]rune, instructions []Instruction, preservePickOneByOneOrder bool) {
	for _, inst := range instructions {
		fromStack := crateStacks[inst.From]
		stackIndex := len(fromStack) - inst.Count
		lastCrates := fromStack[stackIndex:]            // get last
		crateStacks[inst.From] = fromStack[:stackIndex] // remove last
		if preservePickOneByOneOrder {
			reverse(lastCrates)
		}
		crateStacks[inst.To] = append(crateStacks[inst.To], lastCrates...)
	}
}

func getPartOne(crateStacks [][]rune, instructions []Instruction) string {
	moveCrates(crateStacks, instructions, true)
	result := getResult(crateStacks)
	return result
}

func getPartTwo(crateStacks [][]rune, instructions []Instruction) string {
	moveCrates(crateStacks, instructions, false)
	result := getResult(crateStacks)
	return result
}

func getRows(filename string) ([][]rune, []Instruction) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	crateStacks := make([][]rune, 0)
	instructions := make([]Instruction, 0)
	for scanner.Scan() {
		row := scanner.Text()
		if strings.Contains(row, "[") {
			for i := 1; i < len(row); i += 4 {
				j := (i - 1) / 4
				if j+1 > len(crateStacks) {
					crateStacks = append(crateStacks, make([]rune, 0))
				}
				if string(row[i]) != " " {
					crateStacks[j] = append(crateStacks[j], int32(row[i]))
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
	// Reverse the crateStack
	for _, crateStack := range crateStacks {
		reverse(crateStack)
	}
	return crateStacks, instructions
}

func main() {
	fmt.Println("Part one:", getPartOne(getRows("../input.txt")))
	fmt.Println("Part two:", getPartTwo(getRows("../input.txt")))
}
