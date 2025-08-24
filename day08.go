package aoc2020

import (
	"strconv"
	"strings"
)

// Day8 returns accumulator, and , and true if the program terminated normally,
// or false if any one line is executed second time (which indicates a loop).
func Day8(lines []string, part1 bool) (int, bool) {
	var pc, acc int
	visited := make(map[int]bool)
	ended := func() bool {
		return pc == len(lines)
	}
	for !ended() {
		if visited[pc] {
			break
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

	if part1 {
		return acc, ended()
	}

	rewrite := func(s string) (string, bool) {
		if strings.HasPrefix(s, "nop") {
			return strings.ReplaceAll(s, "nop", "jmp"), true
		} else if strings.HasPrefix(s, "jmp") {
			return strings.ReplaceAll(s, "jmp", "nop"), true
		}
		return s, false
	}
	var changed bool
	for i := range lines {
		lines[i], changed = rewrite(lines[i])
		if !changed {
			continue
		}

		rc, b := Day8(lines, true)
		if b {
			return rc, true
		}

		// revert to original
		lines[i], _ = rewrite(lines[i])
	}
	return -1, false
}
