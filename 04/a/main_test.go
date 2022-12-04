package main

import "testing"

func TestFilePartOneTestCases(t *testing.T) {
	want := 2
	got := getPartOne(getRows("../test.txt"))
	if got != want {
		t.Errorf("test file one = %v; want %v", got, want)
	}
}

func TestFilePartOneTestCasesInput(t *testing.T) {
	want := 526
	got := getPartOne(getRows("../input.txt"))
	if got != want {
		t.Errorf("test file one = %d; want %d", got, want)
	}
}

func TestFilePartTwoTestCasesTest(t *testing.T) {
	want := 4
	got := getPartTwo(getRows("../test.txt"))
	if got != want {
		t.Errorf("test file two = %v; want %v", got, want)
	}
}

func TestFilePartTwoTestCasesInput(t *testing.T) {
	want := 886
	got := getPartTwo(getRows("../input.txt"))
	if got != want {
		t.Errorf("test file two = %d; want %d", got, want)
	}
}

func TestFilePartTwoTestCasesInputUsingArray(t *testing.T) {
	want := 886
	got := getPartTwoUsingArray(getRows("../input.txt"))
	if got != want {
		t.Errorf("test file two = %d; want %d", got, want)
	}
}

func TestFilePartTwoTestCasesInputUsingSet(t *testing.T) {
	want := 886
	got := getPartTwoUsingSet(getRows("../input.txt"))
	if got != want {
		t.Errorf("test file two = %d; want %d", got, want)
	}
}


func BenchmarkPartOne(b *testing.B) {
	// run the Fib function b.N times
	rows := getRows("../input.txt")
	for n := 0; n < b.N; n++ {
		getPartOne(rows)
	}
}


func BenchmarkPartTwo(b *testing.B) {
	// run the Fib function b.N times
	rows := getRows("../input.txt")
	for n := 0; n < b.N; n++ {
		getPartTwo(rows)
	}
}

func BenchmarkPartTwoUsingArray(b *testing.B) {
	// run the Fib function b.N times
	rows := getRows("../input.txt")
	for n := 0; n < b.N; n++ {
		getPartTwoUsingArray(rows)
	}
}

func BenchmarkPartTwoUsingSet(b *testing.B) {
	// run the Fib function b.N times
	rows := getRows("../input.txt")
	for n := 0; n < b.N; n++ {
		getPartTwoUsingSet(rows)
	}
}