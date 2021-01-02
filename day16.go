package aoc2020

import (
	"fmt"
	"strconv"
	"strings"
)

// Day16 returns ticket scanning error rate for part1, and product of all fields
// that start with "departure" for part #2.
func Day16(lines []string, part1 bool) (uint, error) {
	rules := make(map[string]map[uint]bool)
	minMax := func(s string) (uint, uint, error) {
		fs := strings.Split(s, "-")
		min, err := strconv.Atoi(fs[0])
		if err != nil {
			msg := "error parsing range start %q: %w"
			return 0, 0, fmt.Errorf(msg, fs[0], err)
		}
		max, err := strconv.Atoi(fs[1])
		if err != nil {
			msg := "error parsing range end %q: %w"
			return 0, 0, fmt.Errorf(msg, fs[1], err)
		}
		return uint(min), uint(max), nil
	}
	setRange := func(m map[uint]bool, min, max uint) {
		for i := uint(min); i <= uint(max); i++ {
			m[i] = true
		}
	}
	validField := func(value uint) bool {
		for _, v := range rules {
			if _, ok := v[value]; ok {
				return true
			}
		}
		return false
	}
	// check if ticket is valid, return false and first illegal field
	validTicket := func(fields []uint) (uint, bool) {
		for i := 0; i < len(fields); i++ {
			if !validField(fields[i]) {
				return fields[i], false
			}
		}
		return 0, true
	}
	hasFields := func(line string) bool {
		return strings.Contains(line, ",")
	}
	parseTicket := func(line string) ([]uint, error) {
		fs := strings.Split(line, ",")
		var is []uint
		for i := range fs {
			n, err := strconv.Atoi(fs[i])
			if err != nil {
				msg := "error converting value #%d: %w"
				return is, fmt.Errorf(msg, err)
			}
			is = append(is, uint(n))
		}
		return is, nil
	}

	var rate uint
	var section byte
	var myTicket []uint
	var tickets [][]uint
	for _, line := range lines {
		if len(line) == 0 {
			section++
			continue
		}
		switch section {
		case 0: // attribute section
			keyValue := strings.Split(line, ":")
			if len(keyValue) != 2 {
				msg := "want %q but got %q"
				return 0, fmt.Errorf(msg, "key: value", line)
			}
			key := keyValue[0]
			fs := strings.Fields(keyValue[1])
			if len(fs) != 3 {
				msg := "want %d fields in attribute section " +
					"but got %d: %q"
				return 0, fmt.Errorf(msg, 3, len(fs), line)
			}
			min1, max1, err := minMax(fs[0])
			if err != nil {
				msg := "error adding first range %q: %w"
				return 0, fmt.Errorf(msg, fs[0], err)
			}
			min2, max2, err := minMax(fs[2])
			if err != nil {
				msg := "error adding second range %q: %w"
				return 0, fmt.Errorf(msg, fs[2], err)
			}
			m2 := make(map[uint]bool)
			setRange(m2, min1, max1)
			setRange(m2, min2, max2)
			rules[key] = m2
		case 1: // my ticket section
			if hasFields(line) { // ignore section comments
				t, err := parseTicket(line)
				if err != nil {
					return 0, err
				}
				myTicket = t
			}
		case 2: // nearby ticket section
			if hasFields(line) { // ignore section comments
				t, err := parseTicket(line)
				if err != nil {
					return 0, err
				}
				field, ok := validTicket(t)
				if ok {
					tickets = append(tickets, t)
				} else {
					rate += field
				}
			}
		}
	}
	if part1 {
		return rate, nil
	}

	// create n copies of rule names, one for each column
	var positions []map[string]bool
	for range rules {
		m2 := make(map[string]bool, len(rules))
		for kk := range rules {
			m2[kk] = true
		}
		positions = append(positions, m2)
	}

	// based on values, remove impossible rules in all positions
	for _, ticket := range tickets {
		for j, value := range ticket {
			// check if possible field, otherwise remove
			for k, v := range rules {
				if !v[value] {
					delete(positions[j], k)
				}
			}
		}
	}

	// for any position with length = 1, remove this rule from any other
	// position
	var finished bool
	for !finished {
		finished = true
		for i := 0; i < len(positions); i++ {
			if len(positions[i]) == 1 {
				// get first field name (which is the only one)
				var field string
				for k := range positions[i] {
					field = k
					break
				}
				for j := 0; j < len(positions); j++ {
					// do not remove from my positions
					if i == j {
						continue
					}
					// only toggle change if existing key has been removed
					if _, ok := positions[j][field]; ok {
						delete(positions[j], field)
						finished = false
					}
				}
			}
		}
	}

	// make sure any position has exactly one field
	for i := range positions {
		if len(positions[i]) != 1 {
			msg := "position %d has no unique key: %+v"
			return 0, fmt.Errorf(msg, positions[i])
		}
	}

	// multiply all fields starting with "departure" of my ticket
	product := uint(1)
	for i := range positions {
		var rule string
		for k := range positions[i] {
			rule = k
			break
		}
		if strings.HasPrefix(rule, "departure") {
			product *= myTicket[i]
		}
	}
	return product, nil
}
