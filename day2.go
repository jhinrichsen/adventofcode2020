package aoc2020

import (
	"fmt"
	"strconv"
	"strings"
)

type day2 struct {
	min      uint
	max      uint
	char     byte
	password string
}

func (a day2) Valid(part1 bool) bool {
	if part1 {
		n := uint(0)
		for i := range a.password {
			if a.char == a.password[i] {
				n++
			}
		}
		return a.min <= n && n <= a.max
	}
	// Toboggan Corporate Policies have no concept of "index zero"
	idx := func(i uint) uint {
		return i - 1
	}
	p1 := a.password[idx(a.min)] == a.char
	p2 := a.password[idx(a.max)] == a.char
	return p1 != p2
}

// Day2 return number of valid passwords.
func Day2(lines []string, part1 bool) (uint, error) {
	n := uint(0)
	for i, line := range lines {
		d, err := parseDay2(line)
		if err != nil {
			msg := "error parsing line %d: %q"
			return 0, fmt.Errorf(msg, i, line)
		}
		if d.Valid(part1) {
			n++
		}
	}
	return n, nil
}

// parseDay2 parses a line in format '1-3 a: abcde'.
func parseDay2(s string) (day2, error) {
	var d day2
	ps := strings.Fields(s)
	if len(ps) != 3 {
		msg := "want %d fields but got %d: %q"
		return d, fmt.Errorf(msg, 3, len(ps), s)
	}
	p1s := strings.Split(ps[0], "-")
	if len(p1s) != 2 {
		msg := "want %d fields in first field but got %d: %q"
		return d, fmt.Errorf(msg, 2, len(p1s), ps[0])
	}
	min, err := strconv.Atoi(p1s[0])
	if err != nil {
		msg := "cannot convert min value to number: %q"
		return d, fmt.Errorf(msg, p1s[0])
	}
	max, err := strconv.Atoi(p1s[1])
	if err != nil {
		msg := "cannot convert max value to number: %q"
		return d, fmt.Errorf(msg, p1s[1])
	}
	d.min, d.max = uint(min), uint(max)
	d.char = byte(ps[1][0])
	d.password = ps[2]
	return d, nil
}
