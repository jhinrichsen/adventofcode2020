package aoc2020

import "fmt"

// Day15 returns the n-th number of the starting sequence.
func Day15(numbers []uint, idx1 int) uint {
	// a number, and a list of its occurences as index
	m := make(map[uint][]int)

	// store a number, and its index
	store := func(n uint, i int) {
		if ns, ok := m[n]; ok {
			ns = append(ns, i)
			m[n] = ns
		} else {
			m[n] = []int{i}
		}
	}

	load := func(n uint) (int, int, error) {
		// if ns, ok := m[n]; ok {
		ns := m[n]
		l := len(ns)
		if l < 2 {
			return 0, 0, fmt.Errorf("need %d but got %d", 2, len(ns))
		}
		return ns[l-1], ns[l-2], nil
	}

	for i := range numbers {
		store(numbers[i], i)
	}

	var n uint
	for i := len(numbers); i < idx1; i++ {
		ultimate, penultimate, err := load(n)
		if err == nil {
			n = uint(ultimate - penultimate) // for deltas, we do not worry about 0 or 1 based indices
		} else {
			n = 0
		}
		store(n, i)
	}
	return n
}
