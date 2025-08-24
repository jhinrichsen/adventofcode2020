package aoc2020

import "sort"

// Day5Part1 returns max seat ID for binary space partitioned seats.
func Day5Part1(seats []string) uint {
	max := uint(0)
	for _, seat := range seats {
		n := uint(0)
		for i := range seat {
			n *= 2
			n += ^((uint(seat[i])) >> 2) & 1
		}
		if n > max {
			max = n
		}
	}
	return max
}

// Day5 returns max seat ID for part1 or missing seat for binary space
// partitioned seats.
func Day5(seats []string, part1 bool) uint {
	max := uint(0)
	var ids []uint
	for _, seat := range seats {
		n := uint(0)
		for i := range seat {
			n *= 2
			n += ^((uint(seat[i])) >> 2) & 1
		}
		if n > max {
			max = n
		}
		ids = append(ids, n)
	}
	if part1 {
		return max
	}

	// sort, and return first hole
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	for i := 1; i < len(ids); i++ {
		if ids[i-1]+1 < ids[i] {
			return ids[i] - 1
		}
	}
	return 0
}
