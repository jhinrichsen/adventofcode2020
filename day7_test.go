package aoc2020

import (
	"testing"
)

func testDay7(t *testing.T, filename string, part1 bool, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	bags, err := parseDay7(lines)
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

func TestDay7ExamplePart1(t *testing.T) {
	testDay7(t, exampleFilename(7), true, 4)
}

func TestDay7(t *testing.T) {
	testDay7(t, filename(7), true, 252)
}

func BenchmarkDay7Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(7))
	if err != nil {
		b.Fatal(err)
	}
	bags, err := parseDay7(lines)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day7Part1(bags)
	}
}

/*
func TestDay7ExamplePart2(t *testing.T) {
	testDay7(t, exampleFilename(7), false, 32)
}

func TestDay7Example2Part2(t *testing.T) {
	testDay7(t, "testdata/day7_example2.txt", false, 126)
}
*/
