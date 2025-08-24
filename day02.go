package aoc2020

import (
	"fmt"
	"strconv"
	"strings"
)

// Puzzle02 holds all password policy entries for Day 02.
type Puzzle02 []struct {
	min      uint
	max      uint
	char     byte
	password string
}

// NewDay02 parses input lines into a Puzzle02 structure.
func NewDay02(lines []string) (Puzzle02, error) {
	var p Puzzle02
	for _, line := range lines {
		ps := strings.Fields(line)
		if len(ps) != 3 {
			return nil, fmt.Errorf("want %d fields but got %d: %q", 3, len(ps), line)
		}
		p1s := strings.Split(ps[0], "-")
		if len(p1s) != 2 {
			return nil, fmt.Errorf("want %d fields in first field but got %d: %q", 2, len(p1s), ps[0])
		}
		min, err := strconv.Atoi(p1s[0])
		if err != nil {
			return nil, fmt.Errorf("cannot convert min value to number: %q", p1s[0])
		}
		max, err := strconv.Atoi(p1s[1])
		if err != nil {
			return nil, fmt.Errorf("cannot convert max value to number: %q", p1s[1])
		}
		p = append(p, struct {
			min      uint
			max      uint
			char     byte
			password string
		}{
			min:      uint(min),
			max:      uint(max),
			char:     byte(ps[1][0]),
			password: ps[2],
		})
	}
	return p, nil
}

// Day02 returns number of valid passwords for Part 1 or Part 2.
func Day02(p Puzzle02, part1 bool) uint {
	var n uint
	for _, d := range p {
		if part1 {
			c := uint(0)
			for i := range d.password {
				if d.char == d.password[i] {
					c++
				}
			}
			if d.min <= c && c <= d.max {
				n++
			}
		} else {
			// Toboggan Corporate Policies have no concept of "index zero"
			idx := func(i uint) uint { return i - 1 }
			p1 := d.password[idx(d.min)] == d.char
			p2 := d.password[idx(d.max)] == d.char
			if p1 != p2 {
				n++
			}
		}
	}
	return n
}
