package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Instruction struct {
	Operation string
	Counter   int
}

func parseXReg(instructions []Instruction) []int {
	xReg := make([]int, 0)
	xReg = append(xReg, 1)
	for _, inst := range instructions {
		last := xReg[len(xReg)-1]
		if inst.Operation == "addx" {
			new := last + inst.Counter
			xReg = append(xReg, []int{last, new}...)
		} else {
			xReg = append(xReg, last)
		}
	}
	return xReg
}

func getPartOne(instructions []Instruction) int {
	xReg := parseXReg(instructions)
	res := 0
	for i := range xReg {
		j := i + 1
		if j == 20 || j == 60 || j == 100 || j == 140 || j == 180 || j == 220 {
			res += j * xReg[i]
		}
	}
	return res
}

func getPartTwo(instructions []Instruction) int {
	xReg := parseXReg(instructions)
	crtWidth := 40
	crtHeight := 6
	tv := make([][]string, crtHeight)
	for i := range tv {
		tv[i] = make([]string, crtWidth)
	}
	for i := range xReg {
		if i == 240 {
			break
		}
		row := i / 40
		col := i % 40
		printed := false
		for j := -1; j < 2; j++ {
			if col == (xReg[i] + j) {
				tv[row][col] = "#"
				printed = true
				break
			}
		}
		if !printed {
			tv[row][col] = "."
		}
	}

	for i := range tv {
		fmt.Println(tv[i])
	}
	return 0
}

func getRows(filename string) []Instruction {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	instructions := make([]Instruction, 0)
	for scanner.Scan() {
		row := scanner.Text()
		var inst Instruction
		if strings.Contains(row, "noop") {
			inst.Operation = row
		}
		fmt.Sscanf(row, "%s %d",
			&inst.Operation, &inst.Counter)
		instructions = append(instructions, inst)
	}
	return instructions
}

func main() {
	fmt.Println("Part one:", getPartOne(getRows("../input.txt")))
	fmt.Println("Part two:", getPartTwo(getRows("../input.txt")))
}
