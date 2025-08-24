package aoc2020

import (
	"testing"
)

var day5Samples = []struct {
	in  string
	out uint
}{
	{"FBFBBFFRLR", 357},
	{"BFFFBBFRRR", 567},
	{"FFFBBBFRRR", 119},
	{"BBFFBBFRLL", 820},
}

func testDay5(t *testing.T, filename string, part1 bool, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got := Day05(lines, part1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay05Samples(t *testing.T) {
	for _, tt := range day5Samples {
		t.Run(tt.in, func(t *testing.T) {
			want := tt.out
			got := Day05([]string{tt.in}, true)
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}

}

func TestDay05Example(t *testing.T) {
	testDay5(t, exampleFilename(5), true, 820)
}

func TestDay05(t *testing.T) {
	testDay5(t, filename(5), true, 904)
}

func BenchmarkDay05Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(5))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day5Part1(lines)
	}
}

func BenchmarkDay05Part1IncludingInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lines, err := linesFromFilename(filename(5))
		if err != nil {
			b.Fatal(err)
		}
		Day5Part1(lines)
	}
}

func BenchmarkDay05(b *testing.B) {
	lines, err := linesFromFilename(filename(5))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day05(lines, true)
	}
}

func BenchmarkDay05IncludingInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lines, err := linesFromFilename(filename(5))
		if err != nil {
			b.Fatal(err)
		}
		Day05(lines, true)
	}
}

func TestDay05Part2(t *testing.T) {
	testDay5(t, filename(5), false, 669)
}
