package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	countPartOne, countPartTwo := 0, 0
	aStart, aEnd, bStart, bEnd := 0, 0, 0, 0
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d",
			&aStart, &aEnd, &bStart, &bEnd)
		if  (aStart <= bStart && bEnd <= aEnd) ||
			(bStart <= aStart && aEnd <= bEnd){
			countPartOne++
		}
		if  (aStart <= bStart && aEnd >= bStart) ||
			(bStart <= aStart && bEnd >= aStart){
			countPartTwo++
		}
	}
	fmt.Println("Part 1", countPartOne, "Part 2", countPartTwo)
}
