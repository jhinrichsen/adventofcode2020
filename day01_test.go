package aoc2020

import (
	"testing"
)

func testDay1(t *testing.T, filename string, want uint, part1 bool) {
	buf := contentFromFilename(t, filename)
	got := Day01(buf, part1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay01Part1Example(t *testing.T) {
	const (
		want  = 514579
		part1 = true
	)
	testDay1(t, exampleFilename(1), want, part1)
}

func TestDay01Part1(t *testing.T) {
	const (
		want  = 485739
		part1 = true
	)
	testDay1(t, filename(1), want, part1)
}

func TestDay01Part2Example(t *testing.T) {
	const (
		want  = 241861950
		part1 = false
	)
	testDay1(t, exampleFilename(1), want, part1)
}

func TestDay01Part2(t *testing.T) {
	const (
		want  = 161109702
		part1 = false
	)
	testDay1(t, filename(1), want, part1)
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
