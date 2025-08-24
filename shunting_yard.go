package aoc2020

import (
	"fmt"
	"strconv"
	"strings"
)

// OperatorConfiguration holds custom operator precedence rules.
type OperatorConfiguration map[string]OpCfg

// OpCfg configures operator metadata.
type OpCfg struct {
	precedence       int
	rightAssociative bool
}

// DefaultOperatorConfiguration configures the standard math set such as *
// before +.
var DefaultOperatorConfiguration = OperatorConfiguration{
	"^": {4, true},
	"*": {3, false},
	"/": {3, false},
	"+": {2, false},
	"-": {2, false},
}

// ShuntingYard will transform an infix expression into reverse polish notation
// (RPN).
func ShuntingYard(infix string, cfg OperatorConfiguration) (rpn string) {
	var stack []string // holds operators and left parenthesis
	for _, tok := range strings.Fields(infix) {
		switch tok {
		case "(":
			stack = append(stack, tok)
		case ")":
			var op string
			for {
				// pop item ("(" or operator) from stack
				op, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if op == "(" {
					break // discard "("
				}
				rpn += " " + op // add operator to result
			}
		default:
			if o1, isOp := cfg[tok]; isOp {
				// token is an operator
				for len(stack) > 0 {
					// consider top item on stack
					op := stack[len(stack)-1]
					if o2, isOp := cfg[op]; !isOp || o1.precedence > o2.precedence ||
						o1.precedence == o2.precedence && o1.rightAssociative {
						break
					}
					// top item is an operator that needs to come off
					stack = stack[:len(stack)-1] // pop it
					rpn += " " + op              // add it to result
				}
				// push operator (the new one) to stack
				stack = append(stack, tok)
			} else { // token is an operand
				if rpn > "" {
					rpn += " "
				}
				rpn += tok // add operand to result
			}
		}
	}
	// drain stack to result
	for len(stack) > 0 {
		rpn += " " + stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return
}

// evalRPN computes a result for an RPN expression.
// Only '+' and '*' are supported for now.
func evalRPN(rpn string) (int, error) {
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
		switch op {
		case "+":
			push(pop() + pop())
		case "*":
			push(pop() * pop())
		default:
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
