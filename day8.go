package aoc2020

import (
	"strconv"
	"strings"
)

// Day8 returns accumulator on first line executed second time.
func Day8(lines []string, part1 bool) int {
	pc, previous := 0, 0
	acc := 0
	visited := make(map[int]bool)
	for {
		if visited[pc] {
			break
		}

		previous = pc
		fs := strings.Fields(lines[pc])
		n, _ := strconv.Atoi(fs[1])
		if fs[0] == "acc" {
			acc += n
		} else if fs[0] == "jmp" {
			pc += n
			continue
		}

		visited[pc] = true
		pc++
	}

	if part1 {
		return acc
	}
	// rewrite program
	lines[previous] = "nop +0"
	return Day8(lines, true)
}
