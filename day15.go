package aoc2020

// Day15 returns the n-th number of the starting sequence.
// Optimized to O(idx1) time with O(U) memory, where U is the range of numbers encountered.
// Uses a dynamically grown slice of last-seen indices to avoid map and slice allocations.
func Day15(numbers []uint, idx1 int) uint {
    // lastSeen[v] = last index where value v appeared (0-based). -1 means unseen.
    // We start with a modest capacity; will grow as needed when values exceed current length.
    lastSeen := make([]int, 1)
    for i := range lastSeen {
        lastSeen[i] = -1
    }

    ensureCap := func(v int) {
        if v < len(lastSeen) {
            return
        }
        // Grow to at least v+1. Double strategy for amortized O(1).
        newLen := len(lastSeen)
        if newLen == 0 {
            newLen = 1
        }
        for newLen <= v {
            newLen *= 2
        }
        // Cap upper bound to idx1+1 to avoid runaway growth (v won't exceed idx1 in AoC15).
        if newLen > idx1+1 {
            newLen = idx1 + 1
        }
        old := lastSeen
        lastSeen = make([]int, newLen)
        copy(lastSeen, old)
        for i := len(old); i < newLen; i++ {
            lastSeen[i] = -1
        }
    }

    // Seed initial numbers except last; track last number spoken
    var last uint
    for i := 0; i < len(numbers)-1; i++ {
        v := int(numbers[i])
        ensureCap(v)
        lastSeen[v] = i
    }
    last = numbers[len(numbers)-1]

    for i := len(numbers); i < idx1; i++ {
        v := int(last)
        ensureCap(v)
        prev := lastSeen[v]
        var next uint
        if prev == -1 {
            next = 0
        } else {
            next = uint((i - 1) - prev)
        }
        lastSeen[v] = i - 1
        last = next
    }
    return last
}
