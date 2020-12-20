package aoc2020

import (
	"fmt"
	"strconv"
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
		n, err := eval(line)
		if err != nil {
			return 0, fmt.Errorf("line %d: %q results in %w",
				i, line, err)
		}
		sum += n
	}
	return sum, nil
}

// eval computes a result for an RPN expression.
// Only '+' and '*' are supported.
func eval(rpn string) (int, error) {
	stack := make([]int, len(rpn)/2)
	var sp int
	pop := func() int {
		sp--
		return stack[sp]
	}
	push := func(n int) {
		stack[sp] = n
		sp++
	}

	for i, op := range strings.Fields(rpn) {
		if op == "+" {
			push(pop() + pop())
		} else if op == "*" {
			push(pop() * pop())
		} else {
			// number
			n, err := strconv.Atoi(op)
			if err != nil {
				msg := "field #%d: want number but got %q"
				return 0, fmt.Errorf(msg, i, op)
			}
			push(n)
		}
	}
	// depth 0 =
	if sp != 1 {
		return 0, fmt.Errorf("want stack depth %d but got %d", 1, sp)
	}
	return pop(), nil
}
