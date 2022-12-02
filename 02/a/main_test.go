package main

import "testing"

func TestFilePartOneTestCases(t *testing.T) {
	want := 24000
	got := getPartOne(getNumbers("../test.txt"))
	if got != want {
		t.Errorf("test file one = %v; want %v", got, want)
	}
}

func TestFilePartOneTestCasesInput(t *testing.T) {
	want := 70509
	got := getPartOne(getNumbers("../input.txt"))
	if got != want {
		t.Errorf("test file one = %d; want %d", got, want)
	}
}

func TestFilePartTwoTestCasesTest(t *testing.T) {
	want := 45000
	got := getPartTwo(getNumbers("../test.txt"))
	if got != want {
		t.Errorf("test file two = %v; want %v", got, want)
	}
}

func TestFilePartTwoTestCasesInput(t *testing.T) {
	want := 208567
	got := getPartTwo(getNumbers("../input.txt"))
	if got != want {
		t.Errorf("test file two = %d; want %d", got, want)
	}
}
