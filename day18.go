package aoc2020

import (
	"fmt"
	"strings"
)

var part1Cfg = OperatorConfiguration{
	"+": {1, false},
	"*": {1, false},
}

var part2Cfg = OperatorConfiguration{
	"+": {2, false},
	"*": {1, false},
}

// Day18 calculates the result of each line, and returns the sum of all lines.
func Day18(lines []string, part1 bool) (int, error) {
	var cfg OperatorConfiguration
	if part1 {
		cfg = part1Cfg
	} else {
		cfg = part2Cfg
	}
	var sum int
	for i, line := range lines {
		// make sure parentheses have whitespace to the left and right
		line = strings.ReplaceAll(line, "(", " ( ")
		line = strings.ReplaceAll(line, ")", " ) ")

		line = ShuntingYard(line, cfg)
		n, err := evalRPN(line)
		if err != nil {
			return 0, fmt.Errorf("line %d: %q results in %w",
				i, line, err)
		}
		sum += n
	}
	return sum, nil
}
