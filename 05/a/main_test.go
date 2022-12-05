package main

import "testing"

func TestFilePartOneTestCases(t *testing.T) {
	want := "CMZ"
	got := getPartOne(getRows("../test.txt"))
	if got != want {
		t.Errorf("test file one = %v; want %v", got, want)
	}
}

func TestFilePartOneTestCasesInput(t *testing.T) {
	want := "WCZTHTMPS"
	got := getPartOne(getRows("../input.txt"))
	if got != want {
		t.Errorf("test file one = %v; want %v", got, want)
	}
}

func TestFilePartTwoTestCasesTest(t *testing.T) {
	want := "MCD"
	got := getPartTwo(getRows("../test.txt"))
	if got != want {
		t.Errorf("test file two = %v; want %v", got, want)
	}
}

func TestFilePartTwoTestCasesInput(t *testing.T) {
	want := "BLSGJSDTS"
	got := getPartTwo(getRows("../input.txt"))
	if got != want {
		t.Errorf("test file two = %v; want %v", got, want)
	}
}

func BenchmarkPartOne(b *testing.B) {
	// run the Fib function b.N times
	crates, inst := getRows("../input.txt")
	for n := 0; n < b.N; n++ {
		lCrates := make([][]rune, len(crates))
		copy(lCrates, crates)
		getPartOne(lCrates, inst)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	// run the Fib function b.N times
	crates, inst := getRows("../input.txt")
	for n := 0; n < b.N; n++ {
		lCrates := make([][]rune, len(crates))
		copy(lCrates, crates)
		getPartTwo(lCrates, inst)
	}
}
