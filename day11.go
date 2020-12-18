package aoc2020

const (
	empty    = 'L'
	occupied = '#'
	floor    = '.'
)

// Day11 holds a grid.
type Day11 struct {
	Grid [][]byte
}

// NewDay11 parses puzzle input into a Day11 struct.
func NewDay11(lines []string) Day11 {
	buf := make([][]byte, len(lines))
	for y := 0; y < len(lines); y++ {
		buf[y] = []byte(lines[y])

	}
	return Day11{buf}
}

// Adjacents returns number of occupied neighbours.
func (a *Day11) Adjacents(x, y int) (count uint) {
	dimx, dimy := len(a.Grid[0]), len(a.Grid)
	seat := func(y, x int) bool {
		return 0 <= x && x < dimx && 0 <= y && y < dimy && a.Grid[y][x] == occupied
	}

	// N
	if seat(y-1, x) {
		count++
	}
	// NE
	if seat(y-1, x+1) {
		count++
	}
	// E
	if seat(y, x+1) {
		count++
	}
	// SE
	if seat(y+1, x+1) {
		count++
	}
	// S
	if seat(y+1, x) {
		count++
	}
	// SW
	if seat(y+1, x-1) {
		count++
	}
	// W
	if seat(y, x-1) {
		count++
	}
	// NW
	if seat(y-1, x-1) {
		count++
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
	return Day11{cp}
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

// Step processes one step in an atomic-like fashion.
func (a *Day11) Step() bool {
	var changed bool
	fb := a.Copy()
	for y := range fb.Grid {
		for x := range fb.Grid[y] {
			n := a.Adjacents(x, y)
			if a.Grid[y][x] == empty && n == 0 {
				fb.Grid[y][x] = occupied
				changed = true
			} else if a.Grid[y][x] == occupied && n >= 4 {
				fb.Grid[y][x] = empty
				changed = true
			}
		}
	}
	if changed {
		a.Grid = fb.Grid
	}
	return changed
}
