package aoc2020

import (
	"reflect"
	"strings"
	"testing"
)

func testDay3(t *testing.T, filename string, part1 bool, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day3(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay3Part1Example(t *testing.T) {
	testDay3(t, exampleFilename(3), true, 7)
}

func TestDay3Part1(t *testing.T) {
	testDay3(t, filename(3), true, 284)
}

// TestDay3Parsing checks if parsing and printing the puzzle input recreates the
// original input.
func TestDay3Reproduction(t *testing.T) {
	original, err := linesFromFilename(filename(3))
	if err != nil {
		t.Fatal(err)
	}
	trees, err := parseDay3(original)
	if err != nil {
		t.Fatal(err)
	}
	dimx, dimy := len(original[0]), len(original)
	var repro []string
	for y := 0; y < dimy; y++ {
		var sb strings.Builder
		for x := 0; x < dimx; x++ {
			c := complex(float64(x), float64(y))
			if trees[c] {
				sb.WriteRune(tree)
			} else {
				sb.WriteRune(open)
			}
		}
		repro = append(repro, sb.String())
	}

	// compare original against reproduction
	if !reflect.DeepEqual(original, repro) {
		t.Fatalf("original and reproduction are different")
	}
}
