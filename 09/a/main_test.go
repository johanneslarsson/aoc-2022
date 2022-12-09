package main

import "testing"

func TestFilePartOneTestCasesTest(t *testing.T) {
	want := 13
	got := getPartOne(getRows("../test.txt"))
	if got != want {
		t.Errorf("test file one = %v; want %v", got, want)
	}
}

func TestFilePartOneTestCasesLarger(t *testing.T) {
	want := 88
	got := getPartOne(getRows("../larger.txt"))
	if got != want {
		t.Errorf("test file one = %v; want %v", got, want)
	}
}

func TestFilePartOneTestCasesInput(t *testing.T) {
	want := 5874 // 1005501 too low, 1359916 too high
	got := getPartOne(getRows("../input.txt"))
	if got != want {
		t.Errorf("test file one = %v; want %v", got, want)
	}
}

func TestFilePartTwoTestCasesTest(t *testing.T) {
	want := 1
	got := getPartTwo(getRows("../test.txt"))
	if got != want {
		t.Errorf("test file two = %v; want %v", got, want)
	}
}
func TestFilePartTwoTestCasesLarger(t *testing.T) {
	want := 36
	got := getPartTwo(getRows("../larger.txt"))
	if got != want {
		t.Errorf("test file two = %v; want %v", got, want)
	}
}

func TestFilePartTwoTestCasesInput(t *testing.T) {
	want := 2467
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
