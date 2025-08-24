package aoc2020

// Day1 returns product of two numbers that add up to 2020.
func Day1(ns map[uint]bool, part1 bool) uint {
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
