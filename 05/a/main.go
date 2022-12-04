package main

import (
	"bufio"
	"fmt"
	"os"
)

type Section struct {
	Start int
	End   int
}

type Pair struct {
	First  Section
	Second Section
}

func getPartOne(pairs []Pair) int {
	count := 0
	for _, pair := range pairs {
		if (pair.First.Start <= pair.Second.Start && pair.Second.End <= pair.First.End) ||
			(pair.Second.Start <= pair.First.Start && pair.First.End <= pair.Second.End) {
			count++
		}
	}
	return count
}

func getPartTwo(pairs []Pair) int {
	count := 0
	for _, pair := range pairs {
		if (pair.First.Start <= pair.Second.Start && pair.First.End >= pair.Second.Start) ||
			(pair.Second.Start <= pair.First.Start && pair.Second.End >= pair.First.Start) {
			count++
		}
	}
	return count
}

func getPartTwoUsingArray(pairs []Pair) int {
	count := 0
	for _, pair := range pairs {
		numbers := make([]bool, 100)
		for j := pair.First.Start; j <= pair.First.End; j++ {
			numbers[j] = true
		}
		for j := pair.Second.Start; j <= pair.Second.End; j++ {
			if numbers[j] {
				count++
				break
			}
		}
	}
	return count
}

func getPartTwoUsingSet(pairs []Pair) int {
	count := 0
	for _, pair := range pairs {
		numbers := make(map[int]bool, getMax(pair.First.End, pair.Second.End))
		for j := pair.First.Start; j <= pair.First.End; j++ {
			numbers[j] = true
		}
		for j := pair.Second.Start; j <= pair.Second.End; j++ {
			if _, ok := numbers[j]; ok {
				count++
				break
			}
		}
	}
	return count
}

func getMax(a int, b int) int {
	if a > b {
		return a + 1
	}
	return b + 1
}

func getRows(filename string) []Pair {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	rows := make([]Pair, 0)
	for scanner.Scan() {
		pair := Pair{}
		fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d",
			&pair.First.Start, &pair.First.End, &pair.Second.Start, &pair.Second.End)
		rows = append(rows, pair)
	}
	return rows
}

func main() {
	fmt.Println("Part one:", getPartOne(getRows("../input.txt")))
	fmt.Println("Part two:", getPartTwo(getRows("../input.txt")))
}
