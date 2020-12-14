package aoc2020

import (
	"sort"
)

func Day10(numbers []int) uint {
	sort.Ints(numbers)
	var ones, threes uint
	ones++   // count from 0 to first joltage
	threes++ // built-in adapter is always 3 higher than highest adapter
	for i := 0; i < len(numbers)-1; i++ {
		delta := numbers[i+1] - numbers[i]
		if delta == 1 {
			ones++
		} else if delta == 3 {
			threes++
		}
	}
	return ones * threes
}
