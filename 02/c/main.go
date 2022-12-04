package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// A - 0 - ROCK
// B - 1 - PAPER
// C - 2 - SCISSOR

// X - 0 - ROCK
// Y - 1 - PAPER
// Z - 2 - SCISSOR

const ArraySize = 3

type Move struct {
	Player1 int
	Player2 int
}

func calcSum(m Move) int {
	sum := 0
	diff := (m.Player2 - m.Player1 + ArraySize) % ArraySize
	if diff == 0 {
		sum += 3
	} else if diff == 1 {
		sum += 6
	}
	sum += m.Player2 + 1
	return sum
}

func getPartOne(moves []Move) int {
	sum := 0
	for i := range moves {
		m := moves[i]
		sum += calcSum(m)
	}
	return sum
}

func getPartTwo(moves []Move) int {
	sum := 0
	for i := range moves {
		m := moves[i]
		// adjust
		if m.Player2 == 0 { // prev
			m.Player2 = (m.Player1 + ArraySize - 1) % ArraySize
		} else if m.Player2 == 1 {
			m.Player2 = m.Player1
		} else if m.Player2 == 2 { // next
			m.Player2 = (m.Player1 + 1) % ArraySize
		}
		// calc
		sum += calcSum(m)
	}
	return sum
}

func getMoves(filename string) []Move {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	moves := make([]Move, 0)
	for scanner.Scan() {
		row := scanner.Text()
		fields := strings.Fields(row)
		moves = append(moves, Move{Player1: int([]rune(fields[0])[0]) - 65, Player2: int([]rune(fields[1])[0]) - 88})
	}
	return moves
}

func main() {
	fmt.Println("Part one:", getPartOne(getMoves("../input.txt")))
	fmt.Println("Part two:", getPartTwo(getMoves("../input.txt")))
}
