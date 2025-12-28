package aoc2020

import (
	"fmt"
)

type hexFloor = complex128 // type alias

// Day24 represents a 2D-list of directions, and a hexagonal floor.
type Day24 struct {
	directions [][]hexFloor
	tiles      map[hexFloor]struct{} // active tiles set (used for Part1 and conversion)
	// Dense array for Part2
	grid   []bool
	next   []bool
	w, h   int
	ox, oy int // origin offset
}

// NewDay24 parses lines of text into a Day24 struct.
func NewDay24(lines []string) (Day24, error) {
	var d Day24
	d.tiles = make(map[hexFloor]struct{})
	for i, line := range lines {
		var ds []hexFloor
		for j := 0; j < len(line); j++ {
			// map hexagonal directions into complex numbers
			// greedy directions first
			if line[j] == 's' && line[j+1] == 'e' {
				ds = append(ds, 1-1i)
				j++
			} else if line[j] == 's' && line[j+1] == 'w' {
				ds = append(ds, 0-1i)
				j++
			} else if line[j] == 'n' && line[j+1] == 'w' {
				ds = append(ds, -1+1i)
				j++
			} else if line[j] == 'n' && line[j+1] == 'e' {
				ds = append(ds, 0+1i)
				j++
			} else if line[j] == 'e' {
				ds = append(ds, 1+0i)
			} else if line[j] == 'w' {
				ds = append(ds, -1+0i)
			} else {
				msg := "line %d: want direction but got %q(%q)"
				return d, fmt.Errorf(msg,
					i, line[j], line[j+1])
			}
		}
		d.directions = append(d.directions, ds)
	}
	return d, nil
}

// Part1 solves day 24, part #1.
func (a *Day24) Part1() {
	for _, path := range a.directions {
		ref := 0 + 0i
		for j := range path {
			ref += path[j]
		}
		if _, ok := a.tiles[ref]; ok {
			delete(a.tiles, ref)
		} else {
			a.tiles[ref] = struct{}{}
		}
	}
}

// Flipped returns number of active tiles.
func (a Day24) Flipped() uint {
	if a.grid != nil {
		var n uint
		for _, v := range a.grid {
			if v {
				n++
			}
		}
		return n
	}
	return uint(len(a.tiles))
}

// initDenseGrid converts sparse tiles to dense array with margin for growth.
func (a *Day24) initDenseGrid(margin int) {
	var minX, minY, maxX, maxY float64
	first := true
	for k := range a.tiles {
		if first {
			minX, maxX = real(k), real(k)
			minY, maxY = imag(k), imag(k)
			first = false
			continue
		}
		if real(k) < minX {
			minX = real(k)
		}
		if real(k) > maxX {
			maxX = real(k)
		}
		if imag(k) < minY {
			minY = imag(k)
		}
		if imag(k) > maxY {
			maxY = imag(k)
		}
	}

	a.ox = int(minX) - margin
	a.oy = int(minY) - margin
	a.w = int(maxX) - int(minX) + 1 + 2*margin
	a.h = int(maxY) - int(minY) + 1 + 2*margin

	a.grid = make([]bool, a.w*a.h)
	for k := range a.tiles {
		x := int(real(k)) - a.ox
		y := int(imag(k)) - a.oy
		a.grid[y*a.w+x] = true
	}
	a.next = make([]bool, a.w*a.h)
}

// step performs one Game of Life step using C6Indices.
func (a *Day24) step() {
	g := Grid{W: a.w, H: a.h}
	for idx, nbrs := range g.C6Indices() {
		var count int
		for nidx := range nbrs {
			if a.grid[nidx] {
				count++
			}
		}

		active := a.grid[idx]
		newActive := active
		if active {
			if count == 0 || count > 2 {
				newActive = false
			}
		} else {
			if count == 2 {
				newActive = true
			}
		}
		a.next[idx] = newActive
	}
	a.grid, a.next = a.next, a.grid
}

// Part2 solves day 24 part #2.
func (a *Day24) Part2(days uint) {
	a.initDenseGrid(int(days) + 1)
	for i := uint(0); i < days; i++ {
		a.step()
	}
}
