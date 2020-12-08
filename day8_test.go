package aoc2020

import (
	"testing"
)

func testDay8(t *testing.T, filename string, want int) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got := Day8(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay8Example(t *testing.T) {
	testDay8(t, exampleFilename(8), 5)
}

func TestDay8(t *testing.T) {
	testDay8(t, filename(8), 2058)
}
