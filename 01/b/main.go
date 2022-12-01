package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	numbers := make([]int, 0)
	sum := 0
	for scanner.Scan() {
		row := scanner.Text()
		val, err := strconv.Atoi(row)
		if err != nil {
			numbers = append(numbers, sum)
			sum = 0
		} else {
			sum += val
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Println("Part one:", numbers[0])
	fmt.Println("Part two:", numbers[0]+numbers[1]+numbers[2])
}
