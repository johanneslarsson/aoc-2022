package main

import (
	"bufio"
	"fmt"
	"os"
)

func isOuterTree(size int, y int, x int) bool {
	return y == 0 || y == size-1 || x == 0 || x == size-1
}

func up(y int, x int, i int) (int, int) {
	return y - i, x
}

func down(y int, x int, i int) (int, int) {
	return y + i, x
}

func left(y int, x int, i int) (int, int) {
	return y, x - i
}

func right(y int, x int, i int) (int, int) {
	return y, x + i
}

func checkAllTheyWayOut(rows [][]rune, y int, x int, move func(int, int, int) (int, int)) bool {
	isOuter, _ := checkNeighbourTree(rows, y, x, move)
	return isOuter
}

func countLowerTrees(rows [][]rune, y int, x int, move func(int, int, int) (int, int)) int {
	_, count := checkNeighbourTree(rows, y, x, move)
	return count
}

func checkNeighbourTree(rows [][]rune, y int, x int, move func(int, int, int) (int, int)) (outerTree bool, count int) {
	count = 1
	for y1, x1 := move(y, x, count); rows[y][x] > rows[y1][x1]; {
		if isOuterTree(len(rows), y1, x1) {
			return true, count
		}
		count++
		y1, x1 = move(y, x, count)
	}
	return false, count
}


func getPartOne(rows [][]rune) int {
	count := 0
	for y := 0; y < len(rows); y++ {
		for x := 0; x < len(rows[y]); x++ {
			if isOuterTree(len(rows), y, x) {
				count++
			} else {
				if checkAllTheyWayOut(rows, y, x, up) ||
					checkAllTheyWayOut(rows, y, x, down) ||
					checkAllTheyWayOut(rows, y, x, left) ||
					checkAllTheyWayOut(rows, y, x, right) {
					count++
				}
			}
		}
	}
	return count
}

func getPartTwo(rows [][]rune) int {
	max := 0
	for y := 1; y < len(rows)-1; y++ {
		for x := 1; x < len(rows[y])-1; x++ {
			upScore := countLowerTrees(rows, y, x, up)
			downScore := countLowerTrees(rows, y, x, down)
			leftScore := countLowerTrees(rows, y, x, left)
			rightScore := countLowerTrees(rows, y, x, right)
			sum := upScore*downScore*leftScore*rightScore
			if sum > max {
				max = sum
			}
		}
	}
	return max
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
