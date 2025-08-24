package aoc2020

import (
	"testing"
)

func testDay1(t *testing.T, filename string, want uint, part1 bool) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	m, err := NewDay01(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day01(m, part1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay01Part1Example(t *testing.T) {
	const (
		want  = 514579
		part1 = true
	)
	testDay1(t, exampleFilename(1), want, part1)
}

func TestDay01Part1(t *testing.T) {
	const (
		want  = 485739
		part1 = true
	)
	testDay1(t, filename(1), want, part1)
}

func TestDay01Part2Example(t *testing.T) {
	const (
		want  = 241861950
		part1 = false
	)
	testDay1(t, exampleFilename(1), want, part1)
}

func TestDay01Part2(t *testing.T) {
	const (
		want  = 161109702
		part1 = false
	)
	testDay1(t, filename(1), want, part1)
}
