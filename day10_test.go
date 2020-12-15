package aoc2020

import (
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
	var f func([]int) uint
	if part1 {
		f = Day10Part1
	} else {
		f = Day10Part2
	}
	got := f(numbers)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay10Example1Part1(t *testing.T) {
	const (
		part1 = true
		want  = 7 * 5
	)
	testDay10(t, exampleFilename(10), part1, want)
}

func TestDay10Example2Part1(t *testing.T) {
	const (
		part1 = true
		want  = 22 * 10
	)
	testDay10(t, "testdata/day10_example2.txt", part1, want)
}

func TestDay10Part1(t *testing.T) {
	const (
		part1 = true
		want  = 1656
	)
	testDay10(t, filename(10), part1, want)
}

func TestDay10Example1Part2(t *testing.T) {
	const (
		part1 = false
		want  = 8
	)
	testDay10(t, exampleFilename(10), part1, want)
}

func TestDay10Example2Part2(t *testing.T) {
	const (
		part1 = false
		want  = 19208
	)
	testDay10(t, "testdata/day10_example2.txt", part1, want)
}

func TestDay10Part2(t *testing.T) {
	const (
		part1 = false
		want  = 56693912375296
	)
	testDay10(t, filename(10), part1, want)
}

func BenchmarkDay10Part2IncludingInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lines, err := linesFromFilename(filename(10))
		if err != nil {
			b.Fatal(err)
		}
		numbers, err := linesAsNumbers(lines)
		if err != nil {
			b.Fatal(err)
		}
		_ = Day10Part2(numbers)
	}
}
