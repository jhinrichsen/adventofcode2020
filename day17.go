package aoc2020

import (
	"strings"
)

const (
	// CubeActive is one of the only two states a cube can have.
	CubeActive = '#'
	// CubeInactive is one of the only two states a cube can have.
	// Inactive sounds so much better than 0x13 0x14 0x10 0x13.
	CubeInactive = '.'
)

type cube struct {
	x, y, z int // coordinates into a 3D world
}

// Neighbours returns a cube's 26 neighbours in 3D space, ignoring that in our
// particular case z-axis is symmetrical to z=0.
func (a cube) Neighbours() []cube {
	cubes := make([]cube, 3*3*3-1)
	var idx int
	for z := a.z - 1; z <= a.z+1; z++ {
		for y := a.y - 1; y <= a.y+1; y++ {
			for x := a.x - 1; x <= a.x+1; x++ {
				c := cube{x, y, z}
				// center cube is not a neighbour
				if c != a {
					cubes[idx] = c
					idx++
				}
			}
		}
	}
	return cubes
}

// Day17 models Conway's game of life in 3D.
type Day17 struct {
	Active           map[cube]struct{}
	DimX, DimY, DimZ int // dimension of our universe
}

// NewDay17 parses a 2D cell grid into a Day17 (z=0).
func NewDay17(lines []string) (Day17, error) {
	d := Day17{}
	d.Active = make(map[cube]struct{})
	d.DimX = len(lines[0])
	d.DimY = len(lines)
	d.DimZ = 1
	z := 0
	for y := 0; y < d.DimY; y++ {
		for x := 0; x < d.DimX; x++ {
			if lines[y][x] == CubeActive {
				d.Active[cube{x, y, z}] = struct{}{}
			}
		}
	}
	return d, nil
}

// ActiveCubes returns number of active cubes in 3D universe.
func (a *Day17) ActiveCubes() (n uint) {
	return uint(len(a.Active))
}

// Cycle runs one atomic generation change.
func (a *Day17) Cycle() {
	a.Expand()
	// all updates into a an offline framebuffer
	fb := make(map[cube]struct{})
	for z := -a.DimZ; z <= a.DimZ; z++ {
		for y := -a.DimY; y <= a.DimY; y++ {
			for x := -a.DimX; x <= a.DimX; x++ {
				here := cube{x, y, z}
				_, oldMe := a.Active[here]
				n := a.ActiveNeighbours(here)
				newMe := NewState(oldMe, n)

				// keep only active cubes
				if newMe {
					fb[here] = struct{}{}
				}
			}
		}
	}
	// new generation comes into existence in one atomic switch
	a.Active = fb
}

// ActiveNeighbours returns the number of neighbours in active state, a number
// between 0 and 26.
func (a Day17) ActiveNeighbours(c cube) (n uint) {
	ns := c.Neighbours()
	for i := 0; i < len(ns); i++ {
		c := ns[i]
		if _, ok := a.Active[c]; ok {
			n++
		}
	}
	return
}

// Expand grows the universe by one unit on each axis.
func (a *Day17) Expand() {
	a.DimX++
	a.DimY++
	a.DimZ++
}

// Rep returns a two-dimensional string representation for axis z.
func (a Day17) Rep(z int) string {
	var sb strings.Builder
	for y := 0; y < a.DimY; y++ {
		for x := 0; x < a.DimX; x++ {
			d := cube{x, y, z}
			if _, ok := a.Active[d]; ok {
				sb.WriteRune(CubeActive)
			} else {
				sb.WriteRune(CubeInactive)
			}
		}
		if y < a.DimY-1 { // new line except for last one
			sb.WriteRune('\n')
		}
	}
	return sb.String()
}

// NewState determines the state upon a generation cycle, depending on the
// current state and the state of immediate neighbours.
func NewState(activeMe bool, activeNeighbours uint) bool {
	if activeMe {
		if activeNeighbours == 2 || activeNeighbours == 3 {
			return true
		}
		return false
	}
	if activeNeighbours == 3 {
		return true
	}
	return false
}

// ===== Part 2: 4D hypercubes =====

// hcube represents a 4D cube (hypercube) at coordinates (x,y,z,w).
type hcube struct {
	x, y, z, w int
}

// Neighbours returns the 80 neighbours in 4D space (3^4 - 1).
func (a hcube) Neighbours() []hcube {
	cubes := make([]hcube, 3*3*3*3-1)
	var idx int
	for w := a.w - 1; w <= a.w+1; w++ {
		for z := a.z - 1; z <= a.z+1; z++ {
			for y := a.y - 1; y <= a.y+1; y++ {
				for x := a.x - 1; x <= a.x+1; x++ {
					c := hcube{x, y, z, w}
					if c != a { // center cube is not a neighbour
						cubes[idx] = c
						idx++
					}
				}
			}
		}
	}
	return cubes
}

// Day17Hyper models Conway's game of life in 4D.
type Day17Hyper struct {
	Active                 map[hcube]struct{}
	DimX, DimY, DimZ, DimW int // dimension of our universe per axis (half-extent)
}

// NewDay17Hyper parses a 2D cell grid into a Day17Hyper (z=0, w=0).
func NewDay17Hyper(lines []string) (Day17Hyper, error) {
	d := Day17Hyper{}
	d.Active = make(map[hcube]struct{})
	d.DimX = len(lines[0])
	d.DimY = len(lines)
	d.DimZ = 1
	d.DimW = 1
	z, w := 0, 0
	for y := 0; y < d.DimY; y++ {
		for x := 0; x < d.DimX; x++ {
			if lines[y][x] == CubeActive {
				d.Active[hcube{x, y, z, w}] = struct{}{}
			}
		}
	}
	return d, nil
}

// ActiveCubes returns number of active cubes in 4D universe.
func (a *Day17Hyper) ActiveCubes() (n uint) {
	return uint(len(a.Active))
}

// Cycle runs one atomic generation change in 4D.
func (a *Day17Hyper) Cycle() {
	a.Expand()
	// offline framebuffer
	fb := make(map[hcube]struct{})
	for w := -a.DimW; w <= a.DimW; w++ {
		for z := -a.DimZ; z <= a.DimZ; z++ {
			for y := -a.DimY; y <= a.DimY; y++ {
				for x := -a.DimX; x <= a.DimX; x++ {
					here := hcube{x, y, z, w}
					_, oldMe := a.Active[here]
					n := a.ActiveNeighbours(here)
					newMe := NewState(oldMe, n)
					if newMe {
						fb[here] = struct{}{}
					}
				}
			}
		}
	}
	a.Active = fb
}

// ActiveNeighbours returns the number of neighbours in active state (0..80).
func (a Day17Hyper) ActiveNeighbours(c hcube) (n uint) {
	ns := c.Neighbours()
	for i := 0; i < len(ns); i++ {
		c := ns[i]
		if _, ok := a.Active[c]; ok {
			n++
		}
	}
	return
}

// Expand grows the universe by one unit on each axis.
func (a *Day17Hyper) Expand() {
	a.DimX++
	a.DimY++
	a.DimZ++
	a.DimW++
}

// Rep returns a two-dimensional string representation for axes z,w fixed (debug).
func (a Day17Hyper) Rep(z, w int) string {
	var sb strings.Builder
	for y := 0; y < a.DimY; y++ {
		for x := 0; x < a.DimX; x++ {
			d := hcube{x, y, z, w}
			if _, ok := a.Active[d]; ok {
				sb.WriteRune(CubeActive)
			} else {
				sb.WriteRune(CubeInactive)
			}
		}
		if y < a.DimY-1 {
			sb.WriteRune('\n')
		}
	}
	return sb.String()
}
