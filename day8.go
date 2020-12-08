package aoc2020

import (
	"strconv"
	"strings"
)

// Day8 returns accumulator on first line executed second time.
func Day8(lines []string) int {
	pc := 0
	acc := 0
	visited := make(map[int]bool)
	for {
		if visited[pc] {
			return acc
		}

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
}
