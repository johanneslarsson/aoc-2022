package main

import (
	"bufio"
	"fmt"
	"os"
)

func getPartOne(row []rune) int {
	for i := 0; i < len(row)-3; i++ {
		r0 := row[i]
		r1 := row[i+1]
		r2 := row[i+2]
		r3 := row[i+3]
		if r0 != r1 && r0 != r2 && r0 != r3 &&
			r1 != r2 && r1 != r3 &&
			r2 != r3 {
			return i + 4
		}
	}
	return 0
}

func getPartTwo(row []rune) int {
	for i := 0; i < len(row)-13; i++ {
		r0 := row[i]
		r1 := row[i+1]
		r2 := row[i+2]
		r3 := row[i+3]
		r4 := row[i+4]
		r5 := row[i+5]
		r6 := row[i+6]
		r7 := row[i+7]
		r8 := row[i+8]
		r9 := row[i+9]
		r10 := row[i+10]
		r11 := row[i+11]
		r12 := row[i+12]
		r13 := row[i+13]
		if r0 != r1 && r0 != r2 && r0 != r3 && r0 != r4 && r0 != r5 && r0 != r6 && r0 != r7 && r0 != r8 && r0 != r9 && r0 != r10 && r0 != r11 && r0 != r12 && r0 != r13 &&
			r1 != r2 && r1 != r3 && r1 != r4 && r1 != r5 && r1 != r6 && r1 != r7 && r1 != r8 && r1 != r9 && r1 != r10 && r1 != r11 && r1 != r12 && r1 != r13 &&
			r2 != r3 && r2 != r4 && r2 != r5 && r2 != r6 && r2 != r7 && r2 != r8 && r2 != r9 && r2 != r10 && r2 != r11 && r2 != r12 && r2 != r13 &&
			r3 != r4 && r3 != r5 && r3 != r6 && r3 != r7 && r3 != r8 && r3 != r9 && r3 != r10 && r3 != r11 && r3 != r12 && r3 != r13 &&
			r4 != r5 && r4 != r6 && r4 != r7 && r4 != r8 && r4 != r9 && r4 != r10 && r4 != r11 && r4 != r12 && r4 != r13 &&
			r5 != r6 && r5 != r7 && r5 != r8 && r5 != r9 && r5 != r10 && r5 != r11 && r5 != r12 && r5 != r13 &&
			r6 != r7 && r6 != r8 && r6 != r9 && r6 != r10 && r6 != r11 && r6 != r12 && r6 != r13 &&
			r7 != r8 && r7 != r9 && r7 != r10 && r7 != r11 && r7 != r12 && r7 != r13 &&
			r8 != r9 && r8 != r10 && r8 != r11 && r8 != r12 && r8 != r13 &&
			r9 != r10 && r9 != r11 && r9 != r12 && r9 != r13 &&
			r10 != r11 && r10 != r12 && r10 != r13 &&
			r11 != r12 && r11 != r13 &&
			r12 != r13 {
			return i + 14
		}
	}

	return 0
}

func getRows(filename string) []rune {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	rows := make([][]rune, 0)
	for scanner.Scan() {
		row := scanner.Text()
		rows = append(rows, []rune(row))
	}
	return rows[0]
}

func main() {
	fmt.Println("Part one:", getPartOne(getRows("../input.txt")))
	fmt.Println("Part two:", getPartTwo(getRows("../input.txt")))
}
