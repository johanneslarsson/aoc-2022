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
		xReg = append(xReg, last)
		if inst.Operation == "addx" {
			xReg = append(xReg, last+inst.Counter)
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
	pixelCount := crtWidth * crtHeight
	crtDisplay := make([][]string, crtHeight)
	for i := range crtDisplay {
		crtDisplay[i] = make([]string, crtWidth)
	}
	for i := range xReg {
		if i == pixelCount {
			break
		}
		row := i / crtWidth
		col := i % crtWidth
		printed := false
		for j := -1; j < 2; j++ {
			if col == (xReg[i] + j) {
				crtDisplay[row][col] = "#"
				printed = true
				break
			}
		}
		if !printed {
			crtDisplay[row][col] = "."
		}
	}

	for i := range crtDisplay {
		fmt.Println(crtDisplay[i])
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
