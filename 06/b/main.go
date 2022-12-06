package main

import (
	"bufio"
	"fmt"
	"os"
)

func getFirstMarkerPosition(row []rune, size int) int {
	for i := 0; i < len(row)-size-1; i++ {
		set := make(map[int32]bool, size)
		set[row[i]] = true
		duplicateFound := false
		for j := i + 1; j < i+size; j++ {
			if _, ok := set[row[j]]; ok {
				duplicateFound = true
				break
			} else {
				set[row[j]] = true
			}
		}
		if !duplicateFound {
			return i + size
		}
	}
	return 0
}

func getPartOne(row []rune) int {
	return getFirstMarkerPosition(row, 4)
}

func getPartTwo(row []rune) int {
	return getFirstMarkerPosition(row, 14)
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
