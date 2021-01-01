package aoc2020

import (
	"fmt"
	"strconv"
	"strings"
)

// NewDay20 parses text lines into a Day 20 struct.
func NewDay20(lines []string) (Day20, error) {
	var d Day20
	d.grid = make(map[uint][]string)

	isTileLine := func(s string) bool {
		return strings.HasPrefix(s, "Tile")
	}

	// make sure lines is an empty line
	if lines[len(lines)-1] > "" {
		lines = append(lines, "")
	}

	var ID uint
	var grid []string
	for i, line := range lines {
		if isTileLine(line) {
			s := line[5 : len(line)-1] // ignore trailing colon
			n, err := strconv.Atoi(s)
			if err != nil {
				msg := "line %d: error parsing tile ID in %q: %w"
				return d, fmt.Errorf(msg, i, line, err)
			}
			ID = uint(n)
		} else if line == "" {
			d.grid[ID] = grid
			grid = nil
		} else {
			// append regular grid line
			grid = append(grid, line)
		}
	}
	return d, nil
}

// Day20 represents a grid.
type Day20 struct {
	grid map[uint][]string
}

// CornerProduct returns the product of all corner tile IDs.
func (a Day20) CornerProduct() (uint, error) {
	bis := a.borders()

	asSet := func(ns []uint) map[uint]bool {
		m := make(map[uint]bool, len(ns))
		for i := range ns {
			m[ns[i]] = true
		}
		return m
	}

	corresponding := func(i int) int {
		// borders come in pairs, return index of partner
		// 0 -- north
		// 1 -- north, flipped
		// 2 -- east
		// 3 -- east, flipped
		// ...
		if i%2 == 0 {
			return i + 1
		}
		return i - 1
	}

	var centers = make(map[uint]bool)
	var middles = make(map[uint]bool)
	var corners = make(map[uint]bool)

	// search for unique borders
	for i := range bis {
		my := asSet(bis[i].borders[:])
		for j := range bis {
			if i == j { // do not remove my borders
				continue
			}
			for k := range bis[j].borders {
				if _, ok := my[bis[j].borders[k]]; ok {
					delete(my, bis[j].borders[k])

					// need to delete flipped border as well
					delete(my, bis[j].borders[corresponding(k)])
				}
			}
		}
		uniqueBorders := len(my) >> 1 // ignore flipped duplicates
		switch uniqueBorders {
		case 0:
			centers[bis[i].tileID] = true
		case 1:
			middles[bis[i].tileID] = true
		case 2:
			corners[bis[i].tileID] = true
		}
	}

	product := uint(1)
	for k := range corners {
		product *= k
	}
	return product, nil
}

func northBorder(lines []string) string {
	return lines[0]
}

func southBorder(lines []string) string {
	return lines[len(lines)-1]
}

func verticalBorder(lines []string, l int) string {
	var sb strings.Builder
	for i := range lines {
		sb.WriteByte(lines[i][l])
	}
	return sb.String()
}

func eastBorder(lines []string) string {
	return verticalBorder(lines, len(lines[0])-1)
}

func westBorder(lines []string) string {
	return verticalBorder(lines, 0)
}

// BorderInfo keeps a numerical representation of a tile's borders.
type BorderInfo struct {
	tileID  uint
	borders [8]uint // north, east, south, west, both regular and flipped
}

// Borders sets up border information for each tile.
func (a *Day20) borders() []BorderInfo {
	var bis []BorderInfo
	for k, v := range a.grid {
		bi := BorderInfo{
			k,
			[8]uint{
				BorderID(northBorder(v)),
				BorderID(Reverse(northBorder(v))),
				BorderID(eastBorder(v)),
				BorderID(Reverse(eastBorder(v))),
				BorderID(southBorder(v)),
				BorderID(Reverse(southBorder(v))),
				BorderID(westBorder(v)),
				BorderID(Reverse(westBorder(v))),
			},
		}
		bis = append(bis, bi)
	}
	return bis
}

// BorderID generates a unique ID for a textual border.
func BorderID(s string) uint {
	var n uint
	for i := 0; i < len(s); i++ {
		n = n << 1
		if s[i] == '#' {
			n++
		}
	}
	return n
}

// Reverse implements the missing strings.Reverse (string -> string)
func Reverse(s string) string {
	var sb strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}
