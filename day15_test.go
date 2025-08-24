package aoc2020

import (
	"fmt"
	"testing"
)

func TestDay15Examples(t *testing.T) {
	for _, tt := range []struct {
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
	} {
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

func BenchmarkDay15Part1(b *testing.B) {
	for b.Loop() {
		_ = Day15([]uint{2, 0, 6, 12, 1, 3}, 2020)
	}
}

func BenchmarkDay15Part2(b *testing.B) {
	for b.Loop() {
		_ = Day15([]uint{2, 0, 6, 12, 1, 3}, 30000000)
	}
}
