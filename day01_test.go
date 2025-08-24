package aoc2020

import (
	"fmt"
	"strconv"
	"testing"
)

func testDay1(t *testing.T, filename string, want uint, part1 bool) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	m, err := parseUints(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day01(m, part1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func parseUints(lines []string) (map[uint]bool, error) {
	is := make(map[uint]bool)
	for i, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return is, fmt.Errorf("error converting number #%d: %w", i, err)
		}
		is[uint(n)] = true
	}
	return is, nil
}

func TestSampleDay1Part1(t *testing.T) {
	const (
		want  = 514579
		part1 = true
	)
	testDay1(t, exampleFilename(1), want, part1)
}

func TestDay1Part1(t *testing.T) {
	const (
		want  = 485739
		part1 = true
	)
	testDay1(t, filename(1), want, part1)
}

func TestSampleDay1Part2(t *testing.T) {
	const (
		want  = 241861950
		part1 = false
	)
	testDay1(t, exampleFilename(1), want, part1)
}

func TestDay1Part2(t *testing.T) {
	const (
		want  = 161109702
		part1 = false
	)
	testDay1(t, filename(1), want, part1)
}
