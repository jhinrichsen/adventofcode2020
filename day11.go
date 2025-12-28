package aoc2020

const (
	empty    = 'L'
	occupied = '#'
)

// Day11 holds a grid using flat array storage.
type Day11 struct {
	grid  []byte
	next  []byte
	w, h  int
	part1 bool
}

// NewDay11 parses puzzle input into a Day11 struct.
func NewDay11(lines []string, part1 bool) Day11 {
	h := len(lines)
	w := len(lines[0])
	size := w * h

	grid := make([]byte, size)
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			grid[y*w+x] = line[x]
		}
	}

	return Day11{
		grid:  grid,
		next:  make([]byte, size),
		w:     w,
		h:     h,
		part1: part1,
	}
}

// countNeighborsPart2 counts visible occupied seats using ray-casting.
func (d *Day11) countNeighborsPart2(idx int) int {
	x, y := idx%d.w, idx/d.w

	var count int
	for _, dir := range [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}} {
		dy, dx := dir[0], dir[1]
		ny, nx := y+dy, x+dx
		for ny >= 0 && ny < d.h && nx >= 0 && nx < d.w {
			cell := d.grid[ny*d.w+nx]
			if cell == occupied {
				count++
				break
			}
			if cell == empty {
				break
			}
			ny += dy
			nx += dx
		}
	}
	return count
}

// Step processes one step. Returns true if any seat changed.
func (d *Day11) Step() bool {
	changed := false
	threshold := 4
	if !d.part1 {
		threshold = 5
	}

	if d.part1 {
		// Use C8Indices for Part 1 - no bounds checking needed
		g := Grid{W: d.w, H: d.h}
		for idx, nbrs := range g.C8Indices() {
			cell := d.grid[idx]
			if cell == '.' {
				d.next[idx] = '.'
				continue
			}

			var count int
			for nidx := range nbrs {
				if d.grid[nidx] == occupied {
					count++
				}
			}

			newCell := cell
			if cell == empty && count == 0 {
				newCell = occupied
			} else if cell == occupied && count >= threshold {
				newCell = empty
			}

			if newCell != cell {
				changed = true
			}
			d.next[idx] = newCell
		}
	} else {
		// Part 2: ray-casting visibility
		for idx := range d.grid {
			cell := d.grid[idx]
			if cell == '.' {
				d.next[idx] = '.'
				continue
			}

			count := d.countNeighborsPart2(idx)

			newCell := cell
			if cell == empty && count == 0 {
				newCell = occupied
			} else if cell == occupied && count >= threshold {
				newCell = empty
			}

			if newCell != cell {
				changed = true
			}
			d.next[idx] = newCell
		}
	}

	d.grid, d.next = d.next, d.grid
	return changed
}

// Occupied returns number of occupied seats.
func (d *Day11) Occupied() uint {
	var count uint
	for _, c := range d.grid {
		if c == occupied {
			count++
		}
	}
	return count
}

// Redact returns a String representation of a Day11 structure.
func (d *Day11) Redact() []string {
	ss := make([]string, d.h)
	for y := 0; y < d.h; y++ {
		ss[y] = string(d.grid[y*d.w : (y+1)*d.w])
	}
	return ss
}
