package aoc2020

import (
	"testing"
)

func testDay16(t *testing.T, filename string, part1 bool, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day16(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay16Part1Example(t *testing.T) {
	const (
		part1 = true
		want  = 71
	)
	testDay16(t, exampleFilename(16), part1, want)
}

func TestDay16Part1(t *testing.T) {
	const (
		part1 = true
		want  = 25984
	)
	testDay16(t, filename(16), part1, want)
}

func TestDay16Part2(t *testing.T) {
	const (
		part1 = false
		want  = 1265347500049
	)
	testDay16(t, filename(16), part1, want)
}

func BenchmarkDay16Part1(b *testing.B) {
	const (
		part1 = true
		want  = 25984
	)
	for i := 0; i < b.N; i++ {
		lines, err := linesFromFilename(filename(16))
		if err != nil {
			b.Fatal(err)
		}
		got, err := Day16(lines, part1)
		if err != nil {
			b.Fatal(err)
		}
		if want != got {
			b.Fatalf("want %d but got %d", want, got)
		}
	}
}

func BenchmarkDay16Part2(b *testing.B) {
	const (
		part1 = false
		want  = 1265347500049
	)
	for i := 0; i < b.N; i++ {
		lines, err := linesFromFilename(filename(16))
		if err != nil {
			b.Fatal(err)
		}
		got, err := Day16(lines, part1)
		if err != nil {
			b.Fatal(err)
		}
		if want != got {
			b.Fatalf("want %d but got %d", want, got)
		}
	}
}
