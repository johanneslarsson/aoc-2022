package main

import "testing"

func TestFilePartOneTestCasesTest(t *testing.T) {
	want := 21
	got := getPartOne(getRows("../test.txt"))
	if got != want {
		t.Errorf("test file one = %v; want %v", got, want)
	}
}

func TestFilePartOneTestCasesInput(t *testing.T) {
	want := 1835 // 1005501 too low, 1359916 too high
	got := getPartOne(getRows("../input.txt"))
	if got != want {
		t.Errorf("test file one = %v; want %v", got, want)
	}
}

func TestFilePartTwoTestCasesTest(t *testing.T) {
	want := 8
	got := getPartTwo(getRows("../test.txt"))
	if got != want {
		t.Errorf("test file two = %v; want %v", got, want)
	}
}

func TestFilePartTwoTestCasesInput(t *testing.T) {
	want := 263670
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
