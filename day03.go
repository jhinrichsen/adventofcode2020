package aoc2020

type slopeFn = func(x, y uint) (uint, uint)

// slope returns number of trees hit by a sled through the woods.
func slope(lines []string, f slopeFn) uint {
	x, y := uint(0), uint(0)
	dimy := uint(len(lines))
	onTree := func() bool {
		return lines[y][x] == '#'
	}
	trees := uint(0)
	for y < dimy {
		if onTree() {
			trees++
		}
		x, y = f(x, y)
	}
	return trees
}

// Day03 returns number of trees encountered for part1, otherwise product of
// slopes (1,1), (3,1), (5,1), (7,1) and (1,2).
func Day03(lines []string, part1 bool) uint {
	border := uint(len(lines[0]))
	if part1 {
		return slope(lines, func(x, y uint) (uint, uint) {
			return (x + 3) % border, y + 1
		})
	}
	return slope(lines, func(x, y uint) (uint, uint) {
		return (x + 1) % border, y + 1
	}) * slope(lines, func(x, y uint) (uint, uint) {
		return (x + 3) % border, y + 1
	}) * slope(lines, func(x, y uint) (uint, uint) {
		return (x + 5) % border, y + 1
	}) * slope(lines, func(x, y uint) (uint, uint) {
		return (x + 7) % border, y + 1
	}) * slope(lines, func(x, y uint) (uint, uint) {
		return (x + 1) % border, y + 2
	})
}
