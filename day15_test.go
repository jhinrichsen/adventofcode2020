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
