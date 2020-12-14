package aoc2020

import (
	"sort"
	"testing"
)

func testDay10(t *testing.T, filename string, part1 bool, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	numbers, err := linesAsNumbers(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day10(numbers)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay10Example1(t *testing.T) {
	const (
		part1 = true
		want  = 7 * 5
	)
	testDay10(t, exampleFilename(10), part1, want)
}

func TestDay10Example2(t *testing.T) {
	const (
		part1 = true
		want  = 22 * 10
	)
	testDay10(t, "testdata/day10_example2.txt", part1, want)
}

func TestDay10(t *testing.T) {
	const (
		part1 = true
		want  = 1656
	)
	testDay10(t, filename(10), part1, want)
}
