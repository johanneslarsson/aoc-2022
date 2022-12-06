package main

import "testing"

func TestFilePartOneTestCasesExample(t *testing.T) {
	want := 7
	got := getPartOne([]rune("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
	if got != want {
		t.Errorf("test file one = %v; want %v", got, want)
	}
}

func TestFilePartOneTestCasesFirst(t *testing.T) {
	want := 5
	got := getPartOne([]rune("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	if got != want {
		t.Errorf("test file one = %v; want %v", got, want)
	}
}

func TestFilePartOneTestCasesSecond(t *testing.T) {
	want := 6
	got := getPartOne([]rune("nppdvjthqldpwncqszvftbrmjlhg"))
	if got != want {
		t.Errorf("test file one = %v; want %v", got, want)
	}
}

func TestFilePartOneTestCasesInput(t *testing.T) {
	want := 1262
	got := getPartOne(getRows("../input.txt"))
	if got != want {
		t.Errorf("test file one = %v; want %v", got, want)
	}
}

func TestFilePartTwoTestCasesTestFirst(t *testing.T) {
	want := 19
	got := getPartTwo([]rune("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
	if got != want {
		t.Errorf("test file two = %v; want %v", got, want)
	}
}

func TestFilePartTwoTestCasesTestSecond(t *testing.T) {
	want := 23
	got := getPartTwo([]rune("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	if got != want {
		t.Errorf("test file two = %v; want %v", got, want)
	}
}

func TestFilePartTwoTestCasesTestThird(t *testing.T) {
	want := 23
	got := getPartTwo([]rune("nppdvjthqldpwncqszvftbrmjlhg"))
	if got != want {
		t.Errorf("test file two = %v; want %v", got, want)
	}
}

func TestFilePartTwoTestCasesInput(t *testing.T) {
	want := 3444
	got := getPartTwo(getRows("../input.txt"))
	if got != want {
		t.Errorf("test file two = %v; want %v", got, want)
	}
}

func BenchmarkPartOne(b *testing.B) {
	// run the Fib function b.N times
	row := getRows("../input.txt")
	for n := 0; n < b.N; n++ {
		getPartOne(row)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	// run the Fib function b.N times
	row := getRows("../input.txt")
	for n := 0; n < b.N; n++ {
		getPartTwo(row)
	}
}
