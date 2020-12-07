package aoc2020

import (
	"fmt"
	"strconv"
	"strings"
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
	got := Day7(bags, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay7Example(t *testing.T) {
	testDay7(t, exampleFilename(7), true, 4)
}

func TestDay7(t *testing.T) {
	testDay7(t, filename(7), true, 252)
}

// Bag contains 0..n other bags.
type Bag struct {
	N     uint
	Color string
}

// Bags contain other bags.
type Bags map[Bag]map[Bag]bool

// Day7 returns number of bag colors that can eventually contain at least one
// shiny gold bag.
func Day7(bags Bags, part1 bool) uint {
	allBags := make(map[Bag]bool)
	allBags[Bag{0, "shiny gold"}] = true
	changed := true
	for changed {
		changed = false
		for bag := range allBags {
			for outer, w := range bags {
				for inner := range w {
					if inner.Color == bag.Color {
						// already recorded?
						if _, ok := allBags[outer]; !ok {
							allBags[outer] = true
							changed = true
						}
					}
				}
			}
		}
	}
	return uint(len(allBags)) - 1 // don't count shiny gold itself
}

// parseDay7 parses one input line in the form "light red bags contain 1 bright
// white bag, 2 muted yellow bags.".
// Bags that contain no other bags are dismissed.
func parseDay7(lines []string) (Bags, error) {
	allBags := make(Bags)
	for i, line := range lines {
		fs := strings.Fields(line)
		if fs[2] != "bags" {
			return allBags, fmt.Errorf("line %d: want field 3 to be %q but got %q", i, "bags", fs[2])
		}
		key := Bag{1, fs[0] + " " + fs[1]}

		// special case
		if strings.Contains(line, "no other bags") {
			continue
		}

		s := strings.Join(fs[4:], " ")
		ps := strings.Split(s, ",")
		bags := make(map[Bag]bool)
		for j := range ps {
			fs := strings.Fields(ps[j])
			if len(fs) != 4 {
				return allBags, fmt.Errorf("line %d, part %d: want %d fields but got %d: %+v",
					i, j, 4, len(fs), fs)
			}
			n, err := strconv.Atoi(fs[0])
			if err != nil {
				return allBags, fmt.Errorf("line %d: cannot convert number of %d. bag: %q", i, j, ps[0])
			}
			color := fs[1] + " " + fs[2]
			bags[Bag{uint(n), color}] = true
		}
		allBags[key] = bags
	}
	return allBags, nil
}

func BenchmarkDay7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lines, err := linesFromFilename(exampleFilename(7))
		if err != nil {
			b.Fatal(err)
		}
		_, err = parseDay7(lines)
		if err != nil {
			b.Fatal(err)
		}
	}
}
