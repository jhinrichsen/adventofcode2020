package aoc2020

import (
	"testing"
)

func testDay8(t *testing.T, part1 bool, filename string, want int) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got := Day8(lines, part1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay8Example(t *testing.T) {
	testDay8(t, true, exampleFilename(8), 5)
}

func TestDay8Part1(t *testing.T) {
	testDay8(t, true, filename(8), 2058)
}

func TestDay8Part2(t *testing.T) {
	t.Fatalf("2085 seems correct, program terminates, but is too low")
	testDay8(t, false, filename(8), 2085)
}
