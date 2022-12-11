package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	StartItems      []int
	Operation       func(a int, b int) int // new = old * 19
	OpValB          int
	TestDivideByVal int
	TestTrueMonkey  int
	TestFalseMonkey int
	InspectCount    int
}

func mul(old int, b int) int    { return old * b }
func mulOld(old int, b int) int { return old * old }
func add(old int, b int) int    { return old + b }
func addOld(old int, b int) int { return old + old }

func calcTopTwoMonkeyInspections(monkeys []Monkey, rounds int, reduceFactor int) int {
	for c := 0; c < rounds; c++ {
		for i, m := range monkeys {
			for _, val := range monkeys[i].StartItems {
				worryLevel := monkeys[i].Operation(val, m.OpValB)
				if reduceFactor == 3 {
					worryLevel /= reduceFactor
				} else {
					worryLevel %= reduceFactor
				}
				if worryLevel%m.TestDivideByVal == 0 {
					monkeys[m.TestTrueMonkey].StartItems = append(monkeys[m.TestTrueMonkey].StartItems, worryLevel)
				} else {
					monkeys[m.TestFalseMonkey].StartItems = append(monkeys[m.TestFalseMonkey].StartItems, worryLevel)
				}
				monkeys[i].InspectCount++
			}
			monkeys[i].StartItems = monkeys[i].StartItems[:0]
		}
	}

	inspectCounts := make([]int, len(monkeys))
	for _, m := range monkeys {
		inspectCounts = append(inspectCounts, m.InspectCount)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspectCounts)))
	return inspectCounts[0]* inspectCounts[1]
}

func getPartOne(monkeys []Monkey) int {
	return calcTopTwoMonkeyInspections(monkeys, 20, 3)
}

func getPartTwo(monkeys []Monkey) int {
	product := 1
	for _, m := range monkeys {
		product *= m.TestDivideByVal
	}
	return calcTopTwoMonkeyInspections(monkeys, 10000, product)
}

func getRows(filename string) []Monkey {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	monkeys := make([]Monkey, 0)
	monkeyId := 0
	for scanner.Scan() {
		row := scanner.Text()
		if strings.Contains(row, "Monkey") {
			fmt.Sscanf(row, "Monkey %d:",
				&monkeyId)
			m := Monkey{}
			monkeys = append(monkeys, m)
		} else if strings.Contains(row, "Starting items:") {
			numberPart := strings.Split(row, ":")[1]
			numbers := strings.Split(numberPart, ",")
			for _, num := range numbers {
				val, err := strconv.Atoi(strings.TrimSpace(num))
				if err != nil {
					panic(err)
				}
				monkeys[monkeyId].StartItems = append(monkeys[monkeyId].StartItems, val)
			}
		} else if strings.Contains(row, "Operation:") {
			formulaPart := strings.Split(row, ":")[1]
			var f func(a int, b int) int
			if strings.Contains(formulaPart, "new = old + old") {
				f = addOld
			} else if strings.Contains(formulaPart, "new = old * old") {
				f = mulOld
			} else if strings.Contains(formulaPart, "new = old * ") {
				f = mul
				num := strings.Replace(formulaPart, "new = old *", "", 1)
				val, err := strconv.Atoi(strings.TrimSpace(num))
				if err != nil {
					panic(err)
				}
				monkeys[monkeyId].OpValB = val
			} else if strings.Contains(formulaPart, "new = old + ") {
				f = add
				num := strings.Replace(formulaPart, "new = old +", "", 1)
				val, err := strconv.Atoi(strings.TrimSpace(num))
				if err != nil {
					panic(err)
				}
				monkeys[monkeyId].OpValB = val
			}
			monkeys[monkeyId].Operation = f
		} else if strings.Contains(row, "Test: divisible by ") {
			fmt.Sscanf(row, "  Test: divisible by %d", &monkeys[monkeyId].TestDivideByVal)
		} else if strings.Contains(row, "If true: throw to monkey") {
			fmt.Sscanf(row, "    If true: throw to monkey %d", &monkeys[monkeyId].TestTrueMonkey)
		} else if strings.Contains(row, "If false: throw to monkey") {
			fmt.Sscanf(row, "    If false: throw to monkey %d", &monkeys[monkeyId].TestFalseMonkey)
		}
	}
	return monkeys
}

func main() {
	fmt.Println("Part one:", getPartOne(getRows("../input.txt")))
	fmt.Println("Part two:", getPartTwo(getRows("../input.txt")))
}
