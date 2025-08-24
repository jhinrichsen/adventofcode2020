package aoc2020

import (
	"testing"
)

func testDay22(t *testing.T, filename string, part1 bool, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	d1, d2, err := NewDay22(lines)
	if err != nil {
		t.Fatal(err)
	}
	var got uint
	if part1 {
		got, err = Day22Part1(d1, d2)
	} else {
		got = Day22Part2(d1, d2, 1)
	}
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func TestDay22Part1Example(t *testing.T) {
	const (
		want  = 306
		part1 = true
	)
	testDay22(t, exampleFilename(22), part1, want)
}

func TestDay22Part1(t *testing.T) {
	const (
		want  = 29764
		part1 = true
	)
	testDay22(t, filename(22), part1, want)
}

func TestDay22Part2Example(t *testing.T) {
	const (
		want  = 291
		part1 = false
	)
	testDay22(t, exampleFilename(22), part1, want)
}

func TestDay22Part2(t *testing.T) {
	const (
		want  = 32588
		part1 = false
	)
	testDay22(t, filename(22), part1, want)
}

func BenchmarkDay22Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lines, err := linesFromFilename(filename(22))
		if err != nil {
			b.Fatal(err)
		}
		d1, d2, err := NewDay22(lines)
		if err != nil {
			b.Fatal(err)
		}
		_, _ = Day22Part1(d1, d2)
	}
}

func BenchmarkDay22Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lines, err := linesFromFilename(filename(22))
		if err != nil {
			b.Fatal(err)
		}
		d1, d2, err := NewDay22(lines)
		if err != nil {
			b.Fatal(err)
		}
		_ = Day22Part2(d1, d2, 1)
	}
}
