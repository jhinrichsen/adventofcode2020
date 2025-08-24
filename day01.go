package aoc2020

import (
	"fmt"
	"strconv"
)

type Puzzle01 map[uint]bool

func NewDay01(lines []string) (Puzzle01, error) {
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

// Day01 returns product of two numbers that add up to 2020.
func Day01(ns Puzzle01, part1 bool) uint {
	if part1 {
		for k := range ns {
			rest := 2020 - k
			if ns[rest] {
				return k * rest
			}
		}
	} else {
		for k := range ns {
			for l := range ns {
				rest := 2020 - k - l
				if ns[rest] {
					return k * l * rest
				}
			}
		}
	}
	return 0
}
