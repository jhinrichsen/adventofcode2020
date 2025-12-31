package aoc2020

import (
	"fmt"
	"strconv"
)

const (
	// leftTurn rotates a complex number by 90 degree anticlockwise when multiplied.
	leftTurn = 0 + 1i

	// rightTurn rotates a complex number by 90 degree clockwise when multiplied.
	rightTurn = 0 - 1i
)

// day12State holds a parsed puzzle input.
type day12State struct {
	Position     complex128
	Direction    complex128
	Instructions []instruction
}

// instruction holds one line of puzzle input, e.g. "N10".
type instruction struct {
	Action byte
	Value  uint
}

func (a instruction) String() string {
	return fmt.Sprintf("%c%d", a.Action, a.Value)
}

// NewDay12 parses puzzle input for day 12.
func NewDay12(lines []string) (day12State, error) {
	position := 0 + 0i
	direction := 1 + 0i // "The ship starts by facing east."
	var is []instruction
	for i := range lines {
		action := lines[i][0]
		s := lines[i][1:]
		value, err := strconv.Atoi(s)
		if err != nil {
			return day12State{}, fmt.Errorf("line %d: error parsing number %q", i, s)
		}
		is = append(is, instruction{action, uint(value)})
	}
	return day12State{position, direction, is}, nil
}

// Part1 processes a puzzle according to the rules of first part.
func (a *day12State) Part1() {
	for _, in := range a.Instructions {
		f := float64(in.Value)

		switch in.Action {
		case 'N':
			a.Position += complex(0.0, f)
		case 'S':
			a.Position += complex(0.0, -f)
		case 'E':
			a.Position += complex(f, 0.0)
		case 'W':
			a.Position += complex(-f, 0.0)
		case 'L':
			// turn left n times
			for n := in.Value / 90; n > 0; n-- {
				a.Direction *= leftTurn
			}
		case 'R':
			// turn right n times
			for n := in.Value / 90; n > 0; n-- {
				a.Direction *= rightTurn
			}
		case 'F':
			a.Position += a.Direction * complex(f, 0)
		}
	}
}

// Part2 processes a puzzle according to the rules of second part.
func (a *day12State) Part2() {
	waypoint := 10 + 1i // "waypoint starts 10 units east and 1 unit north"
	for _, in := range a.Instructions {
		f := float64(in.Value)
		switch in.Action {
		case 'N':
			waypoint += complex(0.0, f)
		case 'S':
			waypoint += complex(0.0, -f)
		case 'E':
			waypoint += complex(f, 0.0)
		case 'W':
			waypoint += complex(-f, 0.0)
		case 'L':
			// turn left n times
			for n := in.Value / 90; n > 0; n-- {
				waypoint *= leftTurn
			}
		case 'R':
			// turn right n times
			for n := in.Value / 90; n > 0; n-- {
				waypoint *= rightTurn
			}
		case 'F':
			scale := complex(float64(in.Value), 0)
			a.Position += scale * waypoint
		}
	}
}

// ManhattanDistance starts at (0,0), executes all instructions and return the
// resulting distance.
func (a day12State) ManhattanDistance() uint {
	return abs(real(a.Position)) + abs(imag(a.Position))
}

// abs returns absolute value.
func abs(f float64) uint {
	n := int(f)
	if n < 0 {
		return uint(-n)
	}
	return uint(n)
}
