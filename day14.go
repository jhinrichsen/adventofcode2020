package aoc2020

import (
	"fmt"
	"strconv"
	"strings"
)

// Day14 returns sum of all memory cells.
func Day14(lines []string, part1 bool) (int, error) {
	mem := make(map[uint]int)

	sum := func() int {
		sum := 0
		for _, v := range mem {
			sum += v
		}
		return sum
	}

	parseMem := func(line string) (uint, int, error) {
		fs := strings.Split(line, " = ")
		if len(fs) != 2 {
			return 0, 0, fmt.Errorf("want %q but got %q", "a = b", line)
		}
		// memory cell index
		s := fs[0][4 : len(fs[0])-1]
		idx, err := strconv.Atoi(s)
		if err != nil {
			return 0, 0, fmt.Errorf("bad mem index %q", s)
		}
		// value
		value, err := strconv.Atoi(fs[1])
		if err != nil {
			return 0, 0, fmt.Errorf("want value, got %q", fs[1])
		}
		return uint(idx), value, nil
	}
	var mask string
	for i, line := range lines {
		if strings.HasPrefix(line, "mask = ") {
			mask = line[7:]
			if len(mask) != 36 {
				return 0, fmt.Errorf("line %d: want mask length %d but got %d", i, 36, len(mask))
			}
		} else if strings.HasPrefix(line, "mem[") {
			idx, value, err := parseMem(line)
			if err != nil {
				return 0, fmt.Errorf("line %d: %w", i, err)
			}

			// apply mask
			for i := 0; i < len(mask); i++ {
				bitValue := 1 << i
				if mask[35-i] == '0' && (value&bitValue > 0) { // when clearing bit, make sure is set
					value -= bitValue
				} else if mask[35-i] == '1' && (value&bitValue == 0) { // when setting bit, make sure it is clear
					value += bitValue
				}
			}

			// store
			mem[idx] = value
		}
	}
	return sum(), nil
}
