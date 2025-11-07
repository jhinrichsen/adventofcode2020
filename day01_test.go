package aoc2020

import (
	"testing"
)

const (
	day01Part1Example = 514579
	day01Part1        = 485739
	day01Part2Example = 241861950
	day01Part2        = 161109702
)

func testDay1(t *testing.T, filename string, want uint, part1 bool) {
	buf := contentFromFilename(t, filename)
	got := Day01(buf, part1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay01Part1Example(t *testing.T) {
	testDay1(t, exampleFilename(1), day01Part1Example, true)
}

func TestDay01Part1(t *testing.T) {
	testDay1(t, filename(1), day01Part1, true)
}

func TestDay01Part2Example(t *testing.T) {
	testDay1(t, exampleFilename(1), day01Part2Example, false)
}

func TestDay01Part2(t *testing.T) {
	testDay1(t, filename(1), day01Part2, false)
}

func BenchmarkDay01Part1(b *testing.B) {
	buf := contentFromFilename(b, filename(1))
	b.ResetTimer()
	for b.Loop() {
		_ = Day01(buf, true)
	}
}

func BenchmarkDay01Part2(b *testing.B) {
	buf := contentFromFilename(b, filename(1))
	b.ResetTimer()
	for b.Loop() {
		_ = Day01(buf, false)
	}
}

// --- Concurrent variant tests ---

func BenchmarkDay01Part1Concurrent(b *testing.B) {
	buf := contentFromFilename(b, filename(1))
	b.ResetTimer()
	for b.Loop() {
		_ = Day01Concurrent(buf, true)
	}
}

func BenchmarkDay01Part2Concurrent(b *testing.B) {
	buf := contentFromFilename(b, filename(1))
	b.ResetTimer()
	for b.Loop() {
		_ = Day01Concurrent(buf, false)
	}
}

// --- Pull-driven variant tests ---

func testDay01Pull(t *testing.T, filename string, want uint, part1 bool) {
	buf := contentFromFilename(t, filename)
	got := Day01Pull(buf, part1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay01Part1Pull(t *testing.T) {
	testDay01Pull(t, filename(1), day01Part1, true)
}

func TestDay01Part2Pull(t *testing.T) {
	testDay01Pull(t, filename(1), day01Part2, false)
}

func BenchmarkDay01Part1Pull(b *testing.B) {
	buf := contentFromFilename(b, filename(1))
	b.ResetTimer()
	for b.Loop() {
		_ = Day01Pull(buf, true)
	}
}

func BenchmarkDay01Part2Pull(b *testing.B) {
	buf := contentFromFilename(b, filename(1))
	b.ResetTimer()
	for b.Loop() {
		_ = Day01Pull(buf, false)
	}
}
