package aoc2020

import (
	"fmt"
	"reflect"
	"testing"
)

func testDay11(t *testing.T, filename string, part1 bool, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	d := NewDay11(lines, part1)
	for d.Step() {
	}
	got := d.Occupied()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func genFilename(i uint, part1 bool) string {
	var part string
	if part1 {
		part = ""
	} else {
		part = "part2_"
	}
	return fmt.Sprintf("testdata/day11_%sexample%d.txt", part, i)
}

func gen(i uint, part1 bool) ([]string, error) {
	return linesFromFilename(genFilename(i, part1))
}

func TestDay11Part1Example(t *testing.T) {
	const (
		part1 = true
		want  = 37
	)
	testDay11(t, genFilename(1, part1), part1, want)
}

func TestDay11Part2Example(t *testing.T) {
	const (
		part1 = false
		want  = 26
	)
	testDay11(t, genFilename(1, part1), part1, want)
}

func testDay11Generations(t *testing.T, part1 bool, generations uint) {
	lines, err := gen(1, part1)
	if err != nil {
		t.Fatal(err)
	}
	d := NewDay11(lines, part1)
	for i := uint(2); i < generations; i++ {
		changed := d.Step()
		if !changed {
			t.Fatalf("step %d: expecting change but got identity", i)
		}
		wantRep, err := gen(i, part1)
		if err != nil {
			t.Fatal(err)
		}
		gotRep := d.Redact()
		if !reflect.DeepEqual(wantRep, gotRep) {
			t.Fatalf("gen %d: want \n%s\n but got \n%s\n", i, wantRep, gotRep)
		}
	}
}

func TestDay11GenerationsPart1(t *testing.T) {
	const (
		part1       = true
		generations = 7
	)
	testDay11Generations(t, part1, generations)
}

func TestDay11GenerationsPart2(t *testing.T) {
	const (
		part1       = false
		generations = 8
	)
	testDay11Generations(t, part1, generations)
}

func TestDay11Part1(t *testing.T) {
	const (
		part1 = true
		want  = 2361
	)
	testDay11(t, filename(11), part1, want)
}

func TestDay11Part2(t *testing.T) {
	const (
		part1 = false
		want  = 2119
	)
	testDay11(t, filename(11), part1, want)
}

func BenchmarkDay11Part1(b *testing.B) {
	const (
		part1 = true
		want  = 2361
	)
	for i := 0; i < b.N; i++ {
		lines, err := linesFromFilename(filename(11))
		if err != nil {
			b.Fatal(err)
		}
		d := NewDay11(lines, part1)
		for d.Step() {
		}
		got := d.Occupied()
		if want != got {
			b.Fatalf("want %d but got %d", want, got)
		}
	}
}

func BenchmarkDay11Part2(b *testing.B) {
	const (
		part1 = false
		want  = 2119
	)
	for i := 0; i < b.N; i++ {
		lines, err := linesFromFilename(filename(11))
		if err != nil {
			b.Fatal(err)
		}
		d := NewDay11(lines, part1)
		for d.Step() {
		}
		got := d.Occupied()
		if want != got {
			b.Fatalf("want %d but got %d", want, got)
		}
	}
}

