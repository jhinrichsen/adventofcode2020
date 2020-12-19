package aoc2020

import "fmt"

// Day15 returns the i-th number of the starting sequence.
func Day15(numbers []uint, i int) uint {
	// index is 1-based, transform to 0-based
	i--

	idxBefore := func(i int) (int, error) { // search for highest index j < i, numbers[j] == numbers[i]
		n := numbers[i-1]
		for j := i - 2; j >= 0; j-- {
			if numbers[j] == n {
				return j, nil
			}
		}
		return 0, fmt.Errorf("not found: %d", n)
	}
	for j := len(numbers); j <= i; j++ {
		last, err := idxBefore(j)
		if err != nil {
			numbers = append(numbers, 0)
		} else {
			lastLast, _ := idxBefore(last)
			numbers = append(numbers, uint(lastLast-last))
		}
	}
	return numbers[i]
}
