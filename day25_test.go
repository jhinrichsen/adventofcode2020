package aoc2020

import (
	"testing"
)

func testDay25(t *testing.T, n1, n2, want int) {
	got := Day25(n1, n2)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func TestDay25Example(t *testing.T) {
	const (
		n1   = 5764801
		n2   = 17807724
		want = 14897079
	)
	testDay25(t, n1, n2, want)
}

func TestDay25Part1(t *testing.T) {
	const (
		n1   = 11404017
		n2   = 13768789
		want = 18862163
	)
	testDay25(t, n1, n2, want)
}

func BenchmarkDay25(b *testing.B) {
	const (
		n1 = 11404017
		n2 = 13768789
	)
	for i := 0; i < b.N; i++ {
		_ = Day25(n1, n2)
	}
}
