package aoc2020

import (
	"testing"
)

func testDay24(t *testing.T, filename string, days, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	d, err := NewDay24(lines)
	if err != nil {
		t.Fatal(err)
	}
	d.Part1()
	d.Part2(days)
	got := d.Flipped()
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func TestDay24Part1Example(t *testing.T) {
	const (
		days = 0
		want = 10
	)
	testDay24(t, exampleFilename(24), days, want)
}

func TestDay24Part1(t *testing.T) {
	const (
		days = 0
		want = 254 // Dec 24, 98 min after puzzle opened
	)
	testDay24(t, filename(24), days, want)
}

func TestDay24Part2Example(t *testing.T) {
	const (
		days = 100
		want = 2208
	)
	testDay24(t, exampleFilename(24), days, want)
}

func TestDay24Part2(t *testing.T) {
	const (
		days = 100
		want = 3697
	)
	testDay24(t, filename(24), days, want)
}

func BenchmarkDay24Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lines, err := linesFromFilename(filename(24))
		if err != nil {
			b.Fatal(err)
		}
		d, err := NewDay24(lines)
		if err != nil {
			b.Fatal(err)
		}
		d.Part1()
		_ = d.Flipped()
	}
}

func BenchmarkDay24Part2(b *testing.B) {
	const days = 100
	for i := 0; i < b.N; i++ {
		lines, err := linesFromFilename(filename(24))
		if err != nil {
			b.Fatal(err)
		}
		d, err := NewDay24(lines)
		if err != nil {
			b.Fatal(err)
		}
		d.Part1()
		d.Part2(days)
		_ = d.Flipped()
	}
}

