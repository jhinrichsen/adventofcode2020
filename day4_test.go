package aoc2020

import "testing"

func testDay4(t *testing.T, filename string, part1 bool, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got := Day4(lines, part1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay4Example(t *testing.T) {
	testDay4(t, exampleFilename(4), true, 2)
}

func TestDay4Part1(t *testing.T) {
	testDay4(t, filename(4), true, 235)
}

func TestDay4Part2(t *testing.T) {
	testDay4(t, filename(4), false, 194)
}
