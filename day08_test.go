package aoc2020

import (
	"testing"
)

func testDay8(t *testing.T, part1 bool, filename string, want int) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got, b := Day08(lines, part1)
	// ignore ran-to-end for part1
	if !part1 && !b {
		t.Fatal("did not run to end, terminated abnormally")
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay08Example(t *testing.T) {
	testDay8(t, true, exampleFilename(8), 5)
}

func TestDay08Part1(t *testing.T) {
	testDay8(t, true, filename(8), 2058)
}

func TestDay08Part2(t *testing.T) {
	testDay8(t, false, filename(8), 1000)
}
