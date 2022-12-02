package main

import "testing"

func TestFilePartOneTestCases(t *testing.T) {
	want := 15
	got := getPartOne(getMoves("../test.txt"))
	if got != want {
		t.Errorf("test file one = %v; want %v", got, want)
	}
}

func TestFilePartOneTestCasesInput(t *testing.T) {
	want := 8933
	got := getPartOne(getMoves("../input.txt"))
	if got != want {
		t.Errorf("test file one = %d; want %d", got, want)
	}
}

func TestFilePartTwoTestCasesTest(t *testing.T) {
	want := 12
	got := getPartTwo(getMoves("../test.txt"))
	if got != want {
		t.Errorf("test file two = %v; want %v", got, want)
	}
}

func TestFilePartTwoTestCasesInput(t *testing.T) {
	want := 11998
	got := getPartTwo(getMoves("../input.txt"))
	if got != want {
		t.Errorf("test file two = %d; want %d", got, want)
	}
}
