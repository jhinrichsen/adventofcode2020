package aoc2020

import (
	"testing"
)

func testDay21(t *testing.T, filename string, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	d, err := NewDay21(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := d.Part1()
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func TestDay21Example(t *testing.T) {
	const (
		want = 5
	)
	testDay21(t, exampleFilename(21), want)
}

func TestDay21(t *testing.T) {
	const (
		want = 2170
	)
	testDay21(t, filename(21), want)
}

func BenchmarkDay21(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lines, err := linesFromFilename(filename(21))
		if err != nil {
			b.Fatal(err)
		}
		d, err := NewDay21(lines)
		if err != nil {
			b.Fatal(err)
		}
		_ = d.Part1()
	}
}
