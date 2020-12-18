package aoc2020

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type color = string

// Bag contains 0..n other bags.
type Bag struct {
	N     uint
	Color color
}

// Bags of a certain color contain n other bags.
type Bags map[color]map[Bag]bool

// Day7Part1 returns number of bag colors that can eventually contain at least one
// shiny gold bag.
func Day7Part1(bags Bags) uint {
	allBags := make(map[color]bool)
	allBags["shiny gold"] = true
	changed := true
	for changed {
		changed = false
		for c := range allBags {
			for outer, w := range bags {
				for inner := range w {
					if inner.Color == c {
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

// Day7Part2 returns number of bags embedded in a shiny gold bag.
func Day7Part2(bags Bags) uint {
	// find all leafs, i.e. bags without entries
	leafs := make(map[color]bool)
	for _, v := range bags {
		for bag := range v {
			if _, ok := bags[bag.Color]; !ok {
				leafs[bag.Color] = true
			}
		}
	}
	log.Printf("found %d leaf(s): %+v\n", len(leafs), leafs)

	// from leafs, traverse up each path, multiplying as we go along
	var n uint
	for leaf := range leafs {
		m := uint(1)
		c := leaf
		log.Printf("going up from leaf %+v\n", c)
		// work our way up from leaf to shiny gold
		for c != "shiny gold" {
			for k, v := range bags {
				for w := range v {
					if w.Color == c {
						m *= w.N
						log.Printf("%q contains %d %q, bags: %d\n", k, w.N, w.Color, m)
						log.Printf("traversing %+v -> %+v\n", c, k)
						c = k
					}
				}
			}
		}
		n += m
	}
	return n
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
		allBags[key.Color] = bags
	}
	return allBags, nil
}
