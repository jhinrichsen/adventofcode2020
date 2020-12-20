package aoc2020

import "strings"

type OperatorConfiguration map[string]OpCfg

type OpCfg struct {
	precedence       int
	rightAssociative bool
}

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
