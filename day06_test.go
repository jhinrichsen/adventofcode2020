package aoc2020

import "testing"

func testDay6(t *testing.T, filename string, part1 bool, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got := Day06(lines, part1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay06Example(t *testing.T) {
	testDay6(t, exampleFilename(6), true, 11)
}

func TestDay06Part1(t *testing.T) {
	testDay6(t, filename(6), true, 6596)
}

func TestDay06Part2Example(t *testing.T) {
	testDay6(t, "testdata/day6_example_part2.txt", false, 6)
}

func TestDay06Part2(t *testing.T) {
	testDay6(t, filename(6), false, 3219)
}

func BenchmarkDay06Part1(b *testing.B) {
	lines := linesFromFilenameTB(b, filename(6))
	b.ResetTimer()
	for b.Loop() {
		_ = Day06(lines, true)
	}
}

func BenchmarkDay06Part2(b *testing.B) {
	lines := linesFromFilenameTB(b, filename(6))
	b.ResetTimer()
	for b.Loop() {
		_ = Day06(lines, false)
	}
}
