package aoc2020

import (
	"testing"
)

func testDay20(t *testing.T, filename string, part1 bool, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	d, err := NewDay20(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := d.CornerProduct()
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay20Example(t *testing.T) {
	const (
		part1 = true
		want  = 20899048083289
	)
	testDay20(t, exampleFilename(20), part1, want)
}

func TestDay20(t *testing.T) {
	const (
		part1 = true
		want  = 45079100979683
	)
	testDay20(t, filename(20), part1, want)
}

func BenchmarkDay20Part1(b *testing.B) {
	const (
		part1 = true
		want  = 45079100979683
	)
	for i := 0; i < b.N; i++ {
		lines, err := linesFromFilename(filename(20))
		if err != nil {
			b.Fatal(err)
		}
		d, err := NewDay20(lines)
		if err != nil {
			b.Fatal(err)
		}
		got, err := d.CornerProduct()
		if want != got {
			b.Fatalf("want %d but got %d", want, got)
		}
	}
}
