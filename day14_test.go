package aoc2020

import "testing"

func testDay14(t *testing.T, filename string, part1 bool, want int) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day14(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay14ExamplePart1(t *testing.T) {
	const (
		part1 = true
		want  = 165
	)
	testDay14(t, exampleFilename(14), part1, want)
}

func TestDay14Part1(t *testing.T) {
	const (
		part1 = true
		want  = 6386593869035
	)
	testDay14(t, filename(14), part1, want)
}
