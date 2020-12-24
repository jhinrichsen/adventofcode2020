package aoc2020

import (
	"fmt"
)

type HexFloor = complex128 // type alias

type Day24 struct {
	directions [][]HexFloor
	tiles      map[HexFloor]bool // all tiles, and a flip indicator
}

func NewDay24(lines []string) (Day24, error) {
	var d Day24
	d.tiles = make(map[HexFloor]bool)
	for i, line := range lines {
		var ds []HexFloor
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

func (a *Day24) Part1() {
	for _, path := range a.directions {
		// ref needs to be flipped?
		ref := 0 + 0i
		for j := range path {
			ref += path[j]
		}
		// only flip last tile, not complete path!
		// only store active tiles
		newState := !a.tiles[ref]
		if newState {
			a.tiles[ref] = true
		} else {
			delete(a.tiles, ref)
		}
	}
}

// Flipped returns number of active tiles.
// As we only store active tiles, the number of tiles on the floor.
func (a Day24) Flipped() uint {
	return uint(len(a.tiles))
}

func (a Day24) Dimension() (HexFloor, HexFloor) { // lower left and upper right
	var minX, minY, maxX, maxY float64

	// cleanly initialize min and max from any random tile
	for k := range a.tiles {
		minX, maxX = real(k), real(k)
		minY, maxY = imag(k), imag(k)
		break
	}

	for k := range a.tiles {
		if real(k) < minX {
			minX = real(k)
		} else if real(k) > maxX {
			maxX = real(k)
		}
		if imag(k) < minY {
			minY = imag(k)
		} else if imag(k) > maxY {
			maxY = imag(k)
		}
	}
	return complex(minX, minY), complex(maxX, maxY)
}

func (a *Day24) Part2(days uint) {
	activeNeighbours := func(tile HexFloor) byte {
		var n byte
		for _, c := range []HexFloor{1 + 0i, 0 + 1i, 0 - 1i, -1 + 1i, 1 - 1i, -1 + 0i} {
			if a.tiles[tile+c] {
				n++
			}
		}
		return n
	}
	newState := func(tile HexFloor, active bool) bool {
		n := activeNeighbours(tile)
		// fmt.Printf("tile %f has %d active neighbours\n", tile, n)
		if active {
			if n == 0 || n > 2 {
				// fmt.Printf("tile %f comes alive\n", tile)
				return false
			}
		} else {
			if activeNeighbours(tile) == 2 {
				// fmt.Printf("tile %f dies\n", tile)
				return true
			}
		}
		return active // no change
	}
	for i := uint(0); i < days; i++ {
		offscreen := make(map[HexFloor]bool)
		min, max := a.Dimension()

		// allow floor to expand at its borders
		grow := 1 + 1i
		min -= grow
		max += grow
		for y := imag(min); y <= imag(max); y++ {
			for x := real(min); x <= real(max); x++ {
				tile := complex(x, y)
				active := newState(tile, a.tiles[tile])
				// only store active tiles
				if active { // change
					offscreen[tile] = true
				} else {
					delete(offscreen, tile)
				}
			}
		}
		a.tiles = offscreen
		// fmt.Printf("Day %d: %d\n", 1+i, a.Flipped())
	}
}
