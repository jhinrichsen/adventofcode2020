package aoc2020

import (
	"testing"
)

func testDay12(t *testing.T, filename string, part1 bool, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	d, err := NewDay12(lines)
	if err != nil {
		t.Fatal(err)
	}
	if part1 {
		d.Part1()
	} else {
		d.Part2()
	}
	got := d.ManhattanDistance()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay12ExamplePart1(t *testing.T) {
	const (
		part1 = true
		want  = 25
	)
	testDay12(t, exampleFilename(12), part1, want)
}

func TestDay12Part1(t *testing.T) {
	const (
		part1 = true
		want  = 2057
	)
	testDay12(t, filename(12), part1, want)
}

func TestDay12ExamplePart2(t *testing.T) {
	const (
		part1 = false
		want  = 286
	)
	testDay12(t, exampleFilename(12), part1, want)
}

func TestDay12Part2(t *testing.T) {
	const (
		part1 = false
		want  = 71504
	)
	testDay12(t, filename(12), part1, want)
}

func BenchmarkDay12Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(12))
	if err != nil {
		b.Fatal(err)
	}
	d, err := NewDay12(lines)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = d.ManhattanDistance()
	}
}
