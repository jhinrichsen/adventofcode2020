package aoc2020

import (
	"testing"
)

func testDay7(t *testing.T, filename string, part1 bool, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	bags, err := NewDay07(lines)
	if err != nil {
		t.Fatal(err)
	}
	var got uint
	if part1 {
		got = Day7Part1(bags)
	} else {
		got = Day7Part2(bags)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay07Part1Example(t *testing.T) {
	testDay7(t, exampleFilename(7), true, 4)
}

func TestDay07Part1(t *testing.T) {
	testDay7(t, filename(7), true, 252)
}

func BenchmarkDay07Part1(b *testing.B) {
	lines := linesFromFilenameTB(b, filename(7))
	b.ResetTimer()
	for b.Loop() {
		bags, _ := NewDay07(lines)
		_ = Day7Part1(bags)
	}
}

func BenchmarkDay07Part2(b *testing.B) {
	lines := linesFromFilenameTB(b, filename(7))
	b.ResetTimer()
	for b.Loop() {
		bags, _ := NewDay07(lines)
		_ = Day7Part2(bags)
	}
}

func TestDay07Part2Example1(t *testing.T) {
	const (
		part1 = false
		want  = 32
	)
	testDay7(t, exampleFilename(7), part1, want)
}

func TestDay07Part2Example2(t *testing.T) {
	const (
		part1 = false
		want  = 126
	)
	testDay7(t, example2Filename(7), part1, want)
}

func TestDay07Part2(t *testing.T) {
	const (
		part1 = false
		want  = 35487
	)
	testDay7(t, filename(7), part1, want)
}
