package aoc2020

import "fmt"

const (
	open = '.'
	tree = '#'
)

// Day3 returns number of trees hit by a sled through the woods.
func Day3(lines []string, part1 bool) (uint, error) {
	area, err := parseDay3(lines)
	if err != nil {
		return 0, err
	}
	rightBorder := complex(float64(len(lines[0])), 0)
	pos := 0 + 0i
	hop := func() {
		pos += 3 + 1i
		// x world repeats forever
		if real(pos) > real(rightBorder) {
			pos -= rightBorder
		}
	}
	bottomBorder := float64(len(lines))
	trees := uint(0)
	for imag(pos) < bottomBorder {
		if area[pos] {
			trees++
		}
		hop()
	}
	return trees, nil
}

func parseDay3(lines []string) (map[complex128]bool, error) {
	trees := make(map[complex128]bool)
	for y, line := range lines {
		for x := range line {
			if line[x] == open {
				// NOP
			} else if line[x] == tree {
				trees[complex(float64(x), float64(y))] = true
			} else {
				return trees, fmt.Errorf("unknown type %q",
					string(line[x]))
			}
		}
	}
	return trees, nil
}
