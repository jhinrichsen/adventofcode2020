package aoc2020

import (
	"testing"
)

func testDay9(t *testing.T, filename string, part1 bool, preamble int, want int) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	numbers, err := linesAsNumbers(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day09(numbers, preamble, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay9ExamplePart1(t *testing.T) {
	const (
		part1    = true
		preamble = 5
		want     = 127
	)
	testDay9(t, exampleFilename(9), part1, preamble, want)
}

func TestDay9Part1(t *testing.T) {
	const (
		part1    = true
		preamble = 25
		want     = 3199139634
	)
	testDay9(t, filename(9), part1, preamble, want)
}

func TestDay9ExamplePart2(t *testing.T) {
	const (
		part1    = false
		preamble = 5
		want     = 62
	)
	testDay9(t, exampleFilename(9), part1, preamble, want)
}

func TestDay9Part2(t *testing.T) {
	const (
		part1    = false
		preamble = 25
		want     = 438559930
	)
	testDay9(t, filename(9), part1, preamble, want)
}
