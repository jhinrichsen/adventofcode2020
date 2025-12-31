package aoc2020

// Day15 returns the n-th number of the starting sequence.
// Optimized to O(idx1) time with O(idx1) memory.
// Pre-allocates full array to avoid bounds checks and function call overhead in hot loop.
func Day15(numbers []uint, idx1 int) uint {
	// lastSeen[v] = last index where value v appeared (1-based to avoid -1 sentinel).
	// 0 means unseen. Pre-allocate for all possible values.
	lastSeen := make([]int32, idx1)

	// Seed initial numbers except last
	for i := 0; i < len(numbers)-1; i++ {
		lastSeen[numbers[i]] = int32(i + 1) // 1-based
	}
	last := int(numbers[len(numbers)-1])

	for i := len(numbers); i < idx1; i++ {
		prev := lastSeen[last]
		lastSeen[last] = int32(i) // store current position (1-based: i = i-1+1)
		if prev == 0 {
			last = 0
		} else {
			last = i - int(prev)
		}
	}
	return uint(last)
}
