package aoc2020

import (
	"testing"
)

func testDay3(t *testing.T, filename string, part1 bool, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got := Day03(lines, part1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay03Part1Example(t *testing.T) {
	testDay3(t, exampleFilename(3), true, 7)
}

func TestDay03Part1(t *testing.T) {
	testDay3(t, filename(3), true, 294)
}

func TestDay03Part2(t *testing.T) {
	testDay3(t, filename(3), false, 5774564250)
}

func BenchmarkDay03Part1(b *testing.B) {
	lines := linesFromFilenameTB(b, filename(3))
	b.ResetTimer()
	for b.Loop() {
		_ = Day03(lines, true)
	}
}

func BenchmarkDay03Part2(b *testing.B) {
	lines := linesFromFilenameTB(b, filename(3))
	b.ResetTimer()
	for b.Loop() {
		_ = Day03(lines, false)
	}
}
