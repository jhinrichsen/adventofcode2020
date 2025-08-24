package aoc2020

const (
	empty    = 'L'
	occupied = '#'
)

// Day11 holds a grid.
type Day11 struct {
	Grid  [][]byte
	part1 bool
}

// NewDay11 parses puzzle input into a Day11 struct.
func NewDay11(lines []string, part1 bool) Day11 {
	buf := make([][]byte, len(lines))
	for y := 0; y < len(lines); y++ {
		buf[y] = []byte(lines[y])

	}
	return Day11{buf, part1}
}

// Adjacents returns number of occupied neighbours.
func (a *Day11) Adjacents(x, y int) (count byte) {
	dimx, dimy := len(a.Grid[0]), len(a.Grid)

	if a.part1 {
		seat := func(y, x int) byte {
			if 0 <= x && x < dimx && 0 <= y && y < dimy && a.Grid[y][x] == occupied {
				return 1
			}
			return 0
		}
		// N
		count += seat(y-1, x)
		// NE
		count += seat(y-1, x+1)
		// E
		count += seat(y, x+1)
		// SE
		count += seat(y+1, x+1)
		// S
		count += seat(y+1, x)
		// SW
		count += seat(y+1, x-1)
		// W
		count += seat(y, x-1)
		// NW
		count += seat(y-1, x-1)
	} else {
		seat := func(y, x, dy, dx int) byte {
			x += dx
			y += dy
			for 0 <= y && y < dimy && 0 <= x && x < dimx {
				switch a.Grid[y][x] {
				case occupied:
					return 1
				case empty: // empty seats break visibility
					return 0
				}
				x += dx
				y += dy
			}
			return 0
		}
		// N
		count += seat(y, x, -1, +0)
		// NE
		count += seat(y, x, -1, +1)
		// E
		count += seat(y, x, +0, +1)
		// SE
		count += seat(y, x, +1, +1)
		// S
		count += seat(y, x, +1, +0)
		// SW
		count += seat(y, x, +1, -1)
		// W
		count += seat(y, x, +0, -1)
		// NW
		count += seat(y, x, -1, -1)
	}
	return
}

// Copy creates a clone of a Day11 structure, game of life operates on a framebuffer.
func (a *Day11) Copy() Day11 {
	cp := make([][]byte, len(a.Grid))
	for i := 0; i < len(a.Grid); i++ {
		cp[i] = make([]byte, len(a.Grid[i]))
		copy(cp[i], a.Grid[i])
	}
	return Day11{cp, a.part1}
}

// Occupied returns number of occupied seats.
func (a *Day11) Occupied() (n uint) {
	for _, s := range a.Grid {
		for _, c := range s {
			if c == occupied {
				n++
			}
		}
	}
	return
}

// Redact returns a String representation of a Day 11 structure.
func (a *Day11) Redact() []string {
	var ss []string
	for i := range a.Grid {
		ss = append(ss, string(a.Grid[i]))
	}
	return ss
}

// NewState calculates the new state based on current state and number of
// neighbours.
func (a *Day11) NewState(state byte, adjacents byte) byte {
	var requiredAdjacents byte
	if a.part1 {
		requiredAdjacents = 4
	} else {
		requiredAdjacents = 5
	}

	if state == empty && adjacents == 0 {
		return occupied
	} else if state == occupied && adjacents >= requiredAdjacents {
		return empty
	}
	return state
}

// Step processes one step in an atomic-like fashion.
func (a *Day11) Step() bool {
	var changed bool
	fb := a.Copy()
	for y := range fb.Grid {
		for x := range fb.Grid[y] {
			st := a.Grid[y][x]
			n := a.Adjacents(x, y)
			st2 := a.NewState(st, n)
			if st != st2 {
				fb.Grid[y][x] = st2
				changed = true
			}
		}
	}
	if changed {
		a.Grid = fb.Grid
	}
	return changed
}
