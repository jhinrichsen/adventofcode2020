package aoc2020

import "fmt"

// Day9 returns the first number that is not the sum of any two numbers in
// preamble for part 1.
// For part 2, it returns the sum of the smallest and largest contiguous range
// whose sum is the number of part 1.
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

	var invalid int
	for i := preamble; i < len(numbers); i++ {
		if !valid(i) {
			invalid = numbers[i]
			break
		}
	}
	if part1 {
		return invalid, nil
	}

	for i := 0; i < len(numbers); i++ {
		min, max := numbers[i], numbers[i]
		sum := 0
		for j := i; j < len(numbers); j++ {
			sum += numbers[j]
			if sum > invalid {
				// busted, next
				break
			}
			if numbers[j] < min {
				min = numbers[j]
			}
			if numbers[j] > max {
				max = numbers[j]
			}
			if sum == invalid {
				return min + max, nil
			}
		}
	}
	return 0, fmt.Errorf("nothing found")
}
