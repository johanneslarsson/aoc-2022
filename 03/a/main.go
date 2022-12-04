package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getValue(match []rune) int {
	sum := 0
	for i := range match {
		m := match[i]
		sub := 0
		if int(m) > 96 {
			sub = 96
		} else {
			sub = 38
		}
		sum += int(m) - sub
	}
	return sum
}

func getPartOne(rows [][]rune) int {
	match := make([]rune, 0)
	for i := range rows {
		row := rows[i]
		rowLen := len(row)
		first := row[:rowLen/2]
		second := row[rowLen/2:]
		for j := range first {
			if strings.ContainsRune(string(second), first[j]) {
				match = append(match, first[j])
				break
			}
		}
	}
	return getValue(match)
}

func getPartTwo(rows [][]rune) int {
	match := make([]rune, 0)
	i := 0
	for range rows {
		row := rows[i]
		next := rows[i+1]
		nextNext := rows[i+2]
		for j := range row {
			if strings.ContainsRune(string(next), row[j]) && strings.ContainsRune(string(nextNext), row[j]) {
				match = append(match, row[j])
				break
			}
		}
		i += 3
		if i == len(rows) {
			break
		}
	}
	return getValue(match)
}

func getRows(filename string) [][]rune {
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
	return rows
}

func main() {
	fmt.Println("Part one:", getPartOne(getRows("../input.txt")))
	fmt.Println("Part two:", getPartTwo(getRows("../input.txt")))
}
