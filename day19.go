package aoc2020

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Day19 returns the number of messages that match rule 0.
// If part1 is false, apply Part 2 semantics: rules 8 and 11 become recursive.
func Day19(lines []string, part1 bool) (uint, error) {
	rls, msgs, err := parseDay19(lines)
	if err != nil {
		return 0, err
	}

	buildRegex := func() (string, error) {
		if part1 {
			rule, err := resolve(rls, "0")
			if err != nil {
				return "", err
			}
			return strings.Join(rule, ""), nil
		}
		// Part 2: compute regex for rules 42 and 31, then enforce m>n>=1 where
		// message matches ^(?:42){m}(?:31){n}$.
		r42toks, err := resolve(rls, "42")
		if err != nil {
			return "", err
		}
		r31toks, err := resolve(rls, "31")
		if err != nil {
			return "", err
		}
		r42 := strings.Join(r42toks, "")
		r31 := strings.Join(r31toks, "")
		// Build bounded alternation over n = 1..K: (?: (?:42){n+1,}(?:31){n} )
		const K = 10
		alts := make([]string, 0, K)
		for n := 1; n <= K; n++ {
			// at least n+1 of 42, then exactly n of 31
			alt := "(?:" + "(?:" + r42 + "){" + strconv.Itoa(n+1) + ",}" + "(?:" + r31 + "){" + strconv.Itoa(n) + "}" + ")"
			alts = append(alts, alt)
		}
		return "(?:" + strings.Join(alts, "|") + ")", nil
	}

	s, err := buildRegex()
	if err != nil {
		return 0, err
	}
	// build regular expression
	s = "^" + s + "$"
	r, err := regexp.Compile(s)
	if err != nil {
		return 0, fmt.Errorf("error creating regexp for %q: %w", s, err)
	}

	// count matches
	var n uint
	for _, s := range msgs {
		if r.MatchString(s) {
			n++
		}
	}
	return n, nil
}

// rules maps a numerical index to a sequence of token.
type rules map[string][]string

func parseDay19(lines []string) (rules, []string, error) {
	rls := make(map[string][]string)
	var msgs []string
	isRule := func(s string) bool { // 0: 4 1 5
		return strings.Contains(s, ":")
	}
	isCharRule := func(fs []string) bool { // "4"
		return len(fs) == 1 && fs[0][0] == '"' && fs[0][2] == '"'
	}
	for i, line := range lines {
		if isRule(line) {
			fs := strings.Split(line, ":")
			if len(fs) != 2 {
				msg := "line %d: want %q but got %q"
				return nil, nil, fmt.Errorf(msg,
					i, "key: value", line)
			}
			key := fs[0]
			// tokenize rule values
			fs = strings.Fields(fs[1])

			// optionally unquote character rules from surrounding
			// quotes
			if isCharRule(fs) {
				s, err := strconv.Unquote(fs[0])
				if err != nil {
					msg := "line %d: error unquoting %q: %w"
					return nil, nil,
						fmt.Errorf(msg, i, fs[0], err)
				}
				fs[0] = s
			}
			rls[key] = fs
		} else if line > "" {
			msgs = append(msgs, line)
		}
		// ignore lines other than rules or messages
	}
	return rls, msgs, nil
}

func resolve(rls rules, root string) ([]string, error) {
	symbolic := func(s string) bool {
		return '0' <= s[0] && s[0] <= '9'
	}
	anySymbolic := func(ss []string) bool {
		for _, s := range ss {
			if symbolic(s) {
				return true
			}
		}
		return false
	}
	needsBraces := func(ss []string) bool { // complex expressions, | e.a.
		for _, s := range ss {
			if s == "|" {
				return true
			}
		}
		return false
	}
	brace := func(operands []string) []string { // operands -> ( operands )
		braced := make([]string, len(operands)+2)
		braced[0] = "("
		copy(braced[1:], operands)
		braced[len(braced)-1] = ")"
		return braced
	}
	r := rls[root]
	for anySymbolic(r) {
		for i, op := range r {
			if symbolic(op) {
				resolved := rls[op]
				// braces for complex replacements
				if needsBraces(resolved) {
					resolved = brace(resolved)
				}
				r = replace(r, i, resolved)

				// replace one by one, otherwise need to keep
				// track where indices move after replacemnt
				break
			}
		}
	}
	return r, nil
}
