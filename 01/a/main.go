package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func getSum(cals []int) int {
	sum := 0
	for cal := range cals {
		sum += cals[cal]
	}
	return sum
}

func getPartOne(numbers [][]int) int {
	max := 0
	for elf := range numbers {
		sum := getSum(numbers[elf])
		if sum > max {
			max = sum
		}
	}
	return max
}

func getPartTwo(numbers [][]int) int {
	aggregated := make([]int, 0)
	for elf := range numbers {
		sum := getSum(numbers[elf])
		aggregated = append(aggregated, sum)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(aggregated)))
	return aggregated[0] + aggregated[1] + aggregated[2]
}

func getNumbers(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	numbers := make([][]int, 1)
	index := 0
	for scanner.Scan() {
		row := scanner.Text()
		val, err := strconv.Atoi(row)
		if err != nil {
			index++
			numbers = append(numbers, make([]int, 0))
			continue
		}
		numbers[index] = append(numbers[index], val)
	}
	return numbers
}

func main() {
	fmt.Println("Part one:", getPartOne(getNumbers("../input.txt")))
	fmt.Println("Part two:", getPartTwo(getNumbers("../input.txt")))
}
