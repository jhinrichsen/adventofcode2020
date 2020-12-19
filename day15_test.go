package aoc2020

import (
	"fmt"
	"testing"
)

var day15ExampleInput = []uint{0, 3, 6}
var day15Samples = []struct {
	input []uint
	i     int
	want  uint
}{
	{day15ExampleInput, 4, 0},
	{day15ExampleInput, 5, 3},
	{day15ExampleInput, 6, 3},
	{day15ExampleInput, 7, 1},
	{day15ExampleInput, 8, 0},
	{day15ExampleInput, 9, 4},
	{day15ExampleInput, 10, 0},
	{day15ExampleInput, 2020, 436},

	{[]uint{0, 3, 6}, 30000000, 175594},
}

// those take some bit of work
var day15ExtendedSamples = []struct {
	input []uint
	i     int
	want  uint
}{
	{[]uint{0, 3, 6}, 30000000, 175594}, // for benchmark, repeat first element
	{[]uint{1, 3, 2}, 30000000, 2578},
	{[]uint{2, 1, 3}, 30000000, 3544142},
	{[]uint{1, 2, 3}, 30000000, 261214},
	{[]uint{2, 3, 1}, 30000000, 6895259},
	{[]uint{3, 2, 1}, 30000000, 18},
	{[]uint{3, 1, 2}, 30000000, 362},
}

func TestDay15Examples(t *testing.T) {
	for _, tt := range day15Samples {
		id := fmt.Sprintf("%d", tt.i)
		t.Run(id, func(t *testing.T) {
			got := Day15(tt.input, tt.i)
			if tt.want != got {
				t.Fatalf("want %d but got %d", tt.want, got)
			}
		})
	}
}

func TestDay15ExtendedExamples(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long running test cases")
	}
	for _, tt := range day15ExtendedSamples {
		id := fmt.Sprintf("%d", tt.i)
		t.Run(id, func(t *testing.T) {
			got := Day15(tt.input, tt.i)
			if tt.want != got {
				t.Fatalf("want %d but got %d", tt.want, got)
			}
		})
	}
}

func TestDay15Part1(t *testing.T) {
	const (
		idx  = 2020
		want = 1428
	)
	got := Day15([]uint{2, 0, 6, 12, 1, 3}, idx)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay15Part2(t *testing.T) {
	const (
		idx  = 30000000
		want = 3718541
	)
	got := Day15([]uint{2, 0, 6, 12, 1, 3}, idx)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

// BenchmarkDay15 is a table driven benchmark.
func BenchmarkDay15Part2(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping long running benchmark")
	}
	for _, tb := range day15ExtendedSamples {
		id := fmt.Sprintf("%+v", (tb.input))
		b.Run(id, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Day15(tb.input, tb.i)
			}
		})
	}
}
