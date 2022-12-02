package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	table := map[string]string{}

	table["A"] = "rock"
	table["B"] = "paper"
	table["C"] = "scissor"

	table["X"] = "rock"
	table["Y"] = "paper"
	table["Z"] = "scissor"

	for scanner.Scan() {
		row := scanner.Text()
		fields := strings.Fields(row)

		p1Move := table[fields[0]]
		p2Move := table[fields[1]]

		if fields[1] == "Y" {
			p2Move = p1Move
		} else if fields[1] == "X" {
			if p1Move == "paper" {
				p2Move = "rock"
			} else if p1Move == "rock" {
				p2Move = "scissor"
			} else if p1Move == "scissor" {
				p2Move = "paper"
			}
		} else if fields[1] == "Z" {
			if p1Move == "scissor" {
				p2Move = "rock"
			} else if p1Move == "rock" {
				p2Move = "paper"
			} else if p1Move == "paper" {
				p2Move = "scissor"
			}
		}

		if p1Move == p2Move {
			sum += 3
		} else if p1Move == "paper" && p2Move == "scissor" || p1Move == "rock" && p2Move == "paper" || p1Move == "scissor" && p2Move == "rock" {
			sum += 6
		}

		// selected point
		if p2Move == "paper" {
			sum += 2
		} else if p2Move == "rock" {
			sum += 1
		} else if p2Move == "scissor" {
			sum += 3
		}
	}
	fmt.Println("Part two:", sum)
}
