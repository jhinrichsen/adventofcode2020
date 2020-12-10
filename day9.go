package aoc2020

import "fmt"

// Day9 returns the first number that is not the sum of any two numbers in
// preamble.
func Day9(numbers []int, preamble int, part1 bool) (int, error) {
	valid := func(idx int) bool {
		start := idx - preamble
		stop := idx
		for i := start; i < stop; i++ {
			for j := start; j < stop; j++ {
				if i == j {
					// any two numbers, but not the same number twice
					continue
				}
				if numbers[i]+numbers[j] == numbers[idx] {
					return true
				}
			}
		}
		return false
	}

	for i := preamble; i < len(numbers); i++ {
		if !valid(i) {
			return numbers[i], nil
		}
	}
	return 0, fmt.Errorf("nothing found")
}
