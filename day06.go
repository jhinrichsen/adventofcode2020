package aoc2020

// Day6 returns number of distict answers by group.
func Day6(lines []string, part1 bool) uint {
	n := uint(0)

	var m map[byte]byte
	var inGroup byte
	newGroup := func() {
		m = make(map[byte]byte)
		inGroup = 0
	}
	newGroup()

	// make sure to terminate last group
	lines = append(lines, "")
	for _, line := range lines {
		if len(line) == 0 {
			if part1 {
				// superset
				n += uint(len(m))
			} else {
				// intersecting set
				for _, v := range m {
					if v == inGroup {
						n++
					}
				}
			}
			newGroup()
		} else {
			for i := range line {
				m[line[i]]++
			}
			inGroup++
		}
	}
	return n
}
