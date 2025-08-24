package aoc2020

import (
	"fmt"
	"strconv"
	"strings"
)

type color = string

// Bag contains 0..n other bags.
type Bag struct {
	N     uint
	Color color
}

// Bags of a certain color contain n other bags.
type Bags map[color]map[Bag]bool

// Day7Part1 returns number of bag colors that can eventually contain at least one
// shiny gold bag.
func Day7Part1(bags Bags) uint {
	allBags := make(map[color]bool)
	allBags["shiny gold"] = true
	changed := true
	for changed {
		changed = false
		for c := range allBags {
			for outer, w := range bags {
				for inner := range w {
					if inner.Color == c {
						// already recorded?
						if _, ok := allBags[outer]; !ok {
							allBags[outer] = true
							changed = true
						}
					}
				}
			}
		}
	}
	return uint(len(allBags)) - 1 // don't count shiny gold itself
}

// Day7Part2 returns number of bags embedded in a shiny gold bag.
func Day7Part2(bags Bags) uint {
	const theBag = "shiny gold"

	resolve1 := func(symbol string) []string { // resolve 1 symbol into its inner bags
		embedded := bags[symbol]
		n := len(embedded)
		if n == 0 { // nothing embedded
			return []string{"1"}
		}

		var ops []string

		// this bag
		ops = append(ops, "(", "1", "+")

		// inner bags
		i := 0 // i [0..n[
		for k := range embedded {
			ops = append(ops,
				"(",
				strconv.Itoa(int(k.N)),
				"*",
				k.Color, // inner bag
				")")
			if i < n-1 { // join '+'
				ops = append(ops, "+")
				i++
			}
		}

		// closing
		ops = append(ops, ")")
		return ops
	}

	isSymbol := func(s string) bool {
		return 'a' <= s[0] && s[0] <= 'z'
	}

	// keep resolving until no more symbols left
	ops := resolve1(theBag)
poorMansTailCallOptimization:
	for i := 0; i < len(ops); i++ {
		if isSymbol(ops[i]) {
			rs := resolve1(ops[i])
			// replace ops[i] with rs[]
			new := make([]string, len(ops)+len(rs)-1)
			copy(new, ops[:i])
			copy(new[i:], rs[:])
			copy(new[i+len(rs):], ops[i+1:])
			ops = new
			goto poorMansTailCallOptimization
		}
	}

	// infix -> rpn
	rpn := ShuntingYard(strings.Join(ops, " "), DefaultOperatorConfiguration)
	n, err := evalRPN(rpn)
	if err != nil {
		// cannot evaluate an expression we constructed ourself
		e := fmt.Errorf("internal error evaluating %q: %w", rpn, err)
		panic(e)
	}
	// return uint(n - 1) // - do not count our shiny gold bag
	return uint(n - 1)
}

// parseDay7 parses one input line in the form "light red bags contain 1 bright
// white bag, 2 muted yellow bags.".
// Bags that contain no other bags are dismissed.
func parseDay7(lines []string) (Bags, error) {
	allBags := make(Bags)
	for i, line := range lines {
		fs := strings.Fields(line)
		if fs[2] != "bags" {
			return allBags, fmt.Errorf("line %d: want field 3 to be %q but got %q", i, "bags", fs[2])
		}
		key := Bag{1, fs[0] + " " + fs[1]}

		// special case
		if strings.Contains(line, "no other bags") {
			continue
		}

		s := strings.Join(fs[4:], " ")
		ps := strings.Split(s, ",")
		bags := make(map[Bag]bool)
		for j := range ps {
			fs := strings.Fields(ps[j])
			if len(fs) != 4 {
				return allBags, fmt.Errorf("line %d, part %d: want %d fields but got %d: %+v",
					i, j, 4, len(fs), fs)
			}
			n, err := strconv.Atoi(fs[0])
			if err != nil {
				return allBags, fmt.Errorf("line %d: cannot convert number of %d. bag: %q", i, j, ps[0])
			}
			color := fs[1] + " " + fs[2]
			bags[Bag{uint(n), color}] = true
		}
		allBags[key.Color] = bags
	}
	return allBags, nil
}
