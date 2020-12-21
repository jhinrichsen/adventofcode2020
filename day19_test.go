package aoc2020

import (
	"testing"
)

func testDay19(t *testing.T, filename string, part1 bool, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day19(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay19Example(t *testing.T) {
	const (
		part1 = true
		want  = 2
	)
	testDay19(t, exampleFilename(19), part1, want)
}

func TestDay19(t *testing.T) {
	const (
		part1 = true
		want  = 203
	)
	testDay19(t, filename(19), part1, want)
}
