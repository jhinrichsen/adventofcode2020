package aoc2020

import (
	"fmt"
	"strconv"
	"strings"
)

// Day14 returns sum of all memory cells.
func Day14(lines []string, part1 bool) (int, error) {
	const bitSize = 36
	mem := make(map[uint]int)

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
			if len(mask) != bitSize {
				return 0, fmt.Errorf("line %d: want mask length %d but got %d", i, bitSize, len(mask))
			}
		} else if strings.HasPrefix(line, "mem[") {
			idx, value, err := parseMem(line)
			if err != nil {
				return 0, fmt.Errorf("line %d: %w", i, err)
			}

			if part1 { // apply mask to value of memory cells
				for i := 0; i < len(mask); i++ {
					bitValue := 1 << i
					if mask[bitSize-1-i] == '0' && (value&bitValue != 0) { // when clearing bit, make sure is set
						value -= bitValue
					} else if mask[bitSize-1-i] == '1' && (value&bitValue == 0) { // when setting bit, make sure it is clear
						value += bitValue
					}
				}
				// store
				mem[idx] = value
			} else { // apply mask to memory cells
				// phase one: clear and set non-fluctuating bits
				buf := []byte(fmt.Sprintf("%036b", idx))
				for j := 0; j < len(mask); j++ {
					if mask[j] == '1' || mask[j] == 'X' {
						buf[j] = mask[j]
					}
				}
				// phase two: permute floating bits in mask
				perms := make(map[string]bool)
				perms[string(buf)] = true
				for {
					more := false
					for k := range perms {
						if strings.ContainsRune(k, 'X') {
							more = true
							perms[strings.Replace(k, "X", "0", 1)] = true
							perms[strings.Replace(k, "X", "1", 1)] = true
							delete(perms, k)
						}
					}
					if !more {
						break
					}
				}
				for mask := range perms {
					idx, err := strconv.ParseUint(mask, 2, bitSize)
					if err != nil {
						return 0, fmt.Errorf("internal error converting %q: %w", mask, err)
					}
					mem[uint(idx)] = value
				}
			}
		}
	}
	sum := func() int {
		sum := 0
		for _, v := range mem {
			sum += v
		}
		return sum
	}

	return sum(), nil
}
