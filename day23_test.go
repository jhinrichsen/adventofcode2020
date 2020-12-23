package aoc2020

import (
	"testing"
)

func testDay23(t *testing.T, input int, moves int, want int) {
	got := Day23(input, moves)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay23Example10(t *testing.T) {
	const (
		input = 389125467
		moves = 10
		want  = 92658374
	)
	testDay23(t, input, moves, want)
}

func TestDay23Example100(t *testing.T) {
	const (
		input = 389125467
		moves = 100
		want  = 67384529
	)
	testDay23(t, input, moves, want)
}

func TestDay23(t *testing.T) {
	const (
		input = 962713854
		moves = 100
		want  = 65432978
	)
	testDay23(t, input, moves, want)
}

func BenchmarkDay23Example10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		const (
			input = 389125467
			moves = 10
		)
		_ = Day23(input, moves)
	}
}

func BenchmarkDay23Example100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		const (
			input = 389125467
			moves = 100
		)
		_ = Day23(input, moves)
	}
}

func BenchmarkDay23(b *testing.B) {
	for i := 0; i < b.N; i++ {
		const (
			input = 962713854
			moves = 100
		)
		_ = Day23(input, moves)
	}
}

func BenchmarkDay23_10_000_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		const (
			input = 962713854
			moves = 10_000_000
		)
		_ = Day23(input, moves)
	}
}
