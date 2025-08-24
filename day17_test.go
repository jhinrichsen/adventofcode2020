package aoc2020

import (
	"testing"
)

func testDay17(t *testing.T, filename string, cycles uint, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	d, err := NewDay17(lines)
	if err != nil {
		t.Fatal(err)
	}
	for i := uint(0); i < cycles; i++ {
		d.Cycle()
	}
	got := d.ActiveCubes()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay17Neighbours(t *testing.T) {
	const want = 26
	cubes := cube{1, 1, 1}.Neighbours() // coordinate itself doesn't matter
	got := len(cubes)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay17InitialState(t *testing.T) {
	const (
		want = 5
	)
	initialState := []string{
		".#.\n",
		"..#\n",
		"###\n",
	}
	d, err := NewDay17(initialState)
	if err != nil {
		t.Fatal(err)
	}
	got := d.ActiveCubes()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay17Example(t *testing.T) {
	const (
		cycles = 6
		want   = 112
	)
	testDay17(t, exampleFilename(17), cycles, want)
}

func TestDay17Part1(t *testing.T) {
	const (
		cycles = 6
		want   = 353
	)
	testDay17(t, filename(17), cycles, want)
}

func BenchmarkDay17Part1(b *testing.B) {
	const cycles = 6
	lines, err := linesFromFilename(filename(17))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d, err := NewDay17(lines)
		if err != nil {
			b.Fatal(err)
		}
		for i := uint(0); i < cycles; i++ {
			d.Cycle()
		}
		_ = d.ActiveCubes()
	}
}

func testDay17Hyper(t *testing.T, filename string, cycles uint, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	d, err := NewDay17Hyper(lines)
	if err != nil {
		t.Fatal(err)
	}
	for i := uint(0); i < cycles; i++ {
		d.Cycle()
	}
	got := d.ActiveCubes()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay17Part2Example(t *testing.T) {
	const (
		cycles = 6
		want   = 848
	)
	testDay17Hyper(t, exampleFilename(17), cycles, want)
}

func TestDay17Part2(t *testing.T) {
	const (
		cycles = 6
		want   = 2472
	)
	testDay17Hyper(t, filename(17), cycles, want)
}
