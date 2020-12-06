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
	got := Day5(lines, part1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay5Samples(t *testing.T) {
	for _, tt := range day5Samples {
		t.Run(tt.in, func(t *testing.T) {
			want := tt.out
			got := Day5([]string{tt.in}, true)
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}

}

func TestDay5Example(t *testing.T) {
	const want = 13
	testDay5(t, exampleFilename(5), true, 820)
}

func TestDay5(t *testing.T) {
	const want = 904
	testDay5(t, filename(5), true, 904)
}

func BenchmarkDay5(b *testing.B) {
	lines, err := linesFromFilename(filename(5))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day5(lines, true)
	}
}

func TestDay5Part2(t *testing.T) {
	const want = 669
	testDay5(t, filename(5), false, 904)
}
