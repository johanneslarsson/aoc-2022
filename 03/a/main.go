package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getValue(match []rune) (sum int) {
	for _, m := range match {
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
	for _, row := range rows {
		rowLen := len(row)
		first := row[:rowLen/2]
		second := row[rowLen/2:]
		for _, firstRune := range first {
			if strings.ContainsRune(string(second), firstRune) {
				match = append(match, firstRune)
				break
			}
		}
	}
	return getValue(match)
}

func getPartTwo(rows [][]rune) int {
	match := make([]rune, 0)
	for i := 0; i < len(rows); i += 3 {
		first := rows[i]
		second := rows[i+1]
		third := rows[i+2]
		for _, firstRune := range first {
			if strings.ContainsRune(string(second), firstRune) && strings.ContainsRune(string(third), firstRune) {
				match = append(match, firstRune)
				break
			}
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
