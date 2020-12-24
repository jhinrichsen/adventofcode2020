package aoc2020

import (
	"fmt"
	"testing"
)

func testDay24(t *testing.T, filename string, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	d, err := NewDay24(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := d.Part1()
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func TestDay24Example(t *testing.T) {
	const (
		want = 10
	)
	testDay24(t, exampleFilename(24), want)
}

type direction int

const (
	east direction = iota
	southEast
	southWest
	west
	northWest
	northEast
)

type Day24 struct {
	directions [][]direction
}

func (a Day24) Part1() uint {
	return 0
}

func NewDay24(lines []string) (Day24, error) {
	var d Day24
	for i, line := range lines {
		var ds []direction
		for j := 0; j < len(line); j++ {
			// greedy directions first
			if line[j] == 's' && line[j+1] == 'e' {
				ds = append(ds, southEast)
				j++
			} else if line[j] == 's' && line[j+1] == 'w' {
				ds = append(ds, southWest)
				j++
			} else if line[j] == 'n' && line[j+1] == 'w' {
				ds = append(ds, northWest)
				j++
			} else if line[j] == 'n' && line[j+1] == 'e' {
				ds = append(ds, northEast)
				j++
			} else if line[j] == 'e' {
				ds = append(ds, east)
			} else if line[j] == 'w' {
				ds = append(ds, west)
			} else {
				msg := "line %d: want direction but got %q(%q)"
				return d, fmt.Errorf(msg,
					i, line[j], line[j+1])
			}
		}
		d.directions = append(d.directions, ds)
	}
	return d, nil
}
