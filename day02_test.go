package aoc2020

import "testing"

func testDay2(t *testing.T, filename string, part1 bool, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day02(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay02Part1Example(t *testing.T) {
	testDay2(t, exampleFilename(2), true, 2)
}

func TestDay02Part1(t *testing.T) {
	testDay2(t, filename(2), true, 536)
}

func TestDay02Part2Example(t *testing.T) {
	testDay2(t, exampleFilename(2), false, 1)
}

func TestDay02Part2(t *testing.T) {
	testDay2(t, filename(2), false, 558)
}
