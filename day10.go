package aoc2020

import (
	"fmt"
	"sort"
)

// Day10Part1 returns number of 1 multiplied by number of 3 differences.
func Day10Part1(numbers []int) uint {
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

// combinations returns the number of different path for n consecutive numbers.
// I figured out the values myself, looking up the sequence shows OIS sequence
// A076739.
func combinations(n int) uint {
	switch n {
	case 1:
		return 1
	case 2:
		return 1
	case 3:
		return 2
	case 4:
		return 4
	case 5:
		return 7
	default:
		panic(fmt.Sprintf("unexpected range: %d", n))
	}
}

// Day10Part2 returns number of combinations in O(n).
func Day10Part2(numbers []int) uint {
	n := uint(1)
	// add a leading 0 and a trailing max+3
	numbers = append(numbers, 0)
	sort.Ints(numbers)
	numbers = append(numbers, numbers[len(numbers)-1]+3)

	nextGap := func(idx int) int {
		for i := idx; i < len(numbers)-1; i++ {
			delta := numbers[i+1] - numbers[i]
			if delta == 3 {
				return i + 1
			} else if delta > 3 {
				return i // backtrack to last index where delta < 3
			}
		}
		return len(numbers)
	}
	for i := 0; i < len(numbers)-1; {
		j := nextGap(i)
		n *= combinations(j - i)
		i = j
	}
	return n
}
