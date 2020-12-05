package aoc2020

import "testing"

func testDay4(t *testing.T, filename string, part1 bool, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got := Day4(lines, part1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay4Example(t *testing.T) {
	testDay4(t, exampleFilename(4), true, 2)
}

func TestDay4Part1(t *testing.T) {
	testDay4(t, filename(4), true, 235)
}

func TestDay4Part2(t *testing.T) {
	testDay4(t, filename(4), false, 194)
}

func BenchmarkDay4Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(4))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day4(lines, false)
	}
}

func BenchmarkDay4Part2IncludingInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lines, err := linesFromFilename(filename(4))
		if err != nil {
			b.Fatal(err)
		}
		Day4(lines, false)
	}
}
