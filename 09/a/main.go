package main

import (
	"bufio"
	"fmt"
	"os"
)

type Instruction struct {
	Direction string
	Count     int
}

type Pos struct {
	X int
	Y int
}

func up(y int, x int, i int) (int, int)    { return y + i, x }
func down(y int, x int, i int) (int, int)  { return y - i, x }
func left(y int, x int, i int) (int, int)  { return y, x - i }
func right(y int, x int, i int) (int, int) { return y, x + i }

func executeSnakeMap(instructions []Instruction, snakeSize int) map[string]int {
	snake := make([]Pos, snakeSize)
	head := 0
	tail := snakeSize - 1

	tailMap := make(map[string]int, 0)
	updateVisit(tailMap, snake[tail].Y, snake[tail].X)
	for _, inst := range instructions {
		for m := 0; m < inst.Count; m++ {
			switch inst.Direction {
			case "U":
				snake[head].Y, snake[head].X = up(snake[head].Y, snake[head].X, 1)
				break
			case "D":
				snake[head].Y, snake[head].X = down(snake[head].Y, snake[head].X, 1)
				break
			case "R":
				snake[head].Y, snake[head].X = right(snake[head].Y, snake[head].X, 1)
				break
			case "L":
				snake[head].Y, snake[head].X = left(snake[head].Y, snake[head].X, 1)
				break
			}
			for h := 0; h < len(snake)-1; h++ {
				diffX := snake[h].X - snake[h+1].X
				diffY := snake[h].Y - snake[h+1].Y
				if abs(diffX) <= 1 && abs(diffY) <= 1 {
					continue
				}
				snake[h+1].X += moveMaxOne(diffX)
				snake[h+1].Y += moveMaxOne(diffY)
				if h+1 == tail {
					updateVisit(tailMap, snake[h+1].Y, snake[h+1].X)
				}
			}
		}
	}
	return tailMap
}

func updateVisit(visit map[string]int, y int, x int) {
	visit[fmt.Sprintf("%d,%d", y, x)]++
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func moveMaxOne(i int) int {
	if i == 0 {
		return 0
	} else if i < 0 {
		return -1
	} else {
		return 1
	}
}

func getPartOne(instructions []Instruction) int {
	snakeSize := 2
	tailMap := executeSnakeMap(instructions, snakeSize)
	return len(tailMap)
}

func getPartTwo(instructions []Instruction) int {
	snakeSize := 10
	tailMap := executeSnakeMap(instructions, snakeSize)
	return len(tailMap)
}

func getRows(filename string) []Instruction {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	moves := make([]Instruction, 0)
	for scanner.Scan() {
		var headMove Instruction
		fmt.Sscanf(scanner.Text(), "%s %d",
			&headMove.Direction, &headMove.Count)
		moves = append(moves, headMove)
	}
	return moves
}

func main() {
	fmt.Println("Part one:", getPartOne(getRows("../input.txt")))
	fmt.Println("Part two:", getPartTwo(getRows("../input.txt")))
}
