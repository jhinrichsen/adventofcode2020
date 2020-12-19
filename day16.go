package aoc2020

import (
	"fmt"
	"strconv"
	"strings"
)

// Day16 returns ticket scanning error rate.
// no internal error checking this time, feels weird.
func Day16(lines []string) (uint, error) {
	m := make(map[uint]bool)
	addRange := func(s string) error {
		fs := strings.Split(s, "-")
		min, err := strconv.Atoi(fs[0])
		if err != nil {
			return fmt.Errorf("error parsing range start %q: %w",
				fs[0], err)
		}
		max, err := strconv.Atoi(fs[1])
		if err != nil {
			return fmt.Errorf("error parsing range end %q: %w",
				fs[1], err)
		}
		for i := uint(min); i <= uint(max); i++ {
			m[i] = true
		}
		return nil
	}
	valid := func(field uint) bool {
		_, ok := m[field]
		return ok
	}

	hasFields := func(line string) bool {
		return strings.Contains(line, ",")
	}

	var rate uint
	var section byte
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
			fs := strings.Fields(keyValue[1])
			if len(fs) != 3 {
				msg := "want %d fields in attribute section " +
					"but got %d: %q"
				return 0, fmt.Errorf(msg, 3, len(fs), line)
			}
			if err := addRange(fs[0]); err != nil {
				msg := "error adding first range %q: %w"
				return 0, fmt.Errorf(msg, fs[0], err)
			}
			if err := addRange(fs[2]); err != nil {
				msg := "error adding second range %q: %w"
				return 0, fmt.Errorf(msg, fs[2], err)
			}
		case 1: // my ticket section
		// ignore for part 1
		case 2: // nearby ticket section
			if hasFields(line) { // ignore section comments
				fs := strings.Split(line, ",")
				for i := 0; i < len(fs); i++ {
					n, _ := strconv.Atoi(fs[i])
					un := uint(n)
					if !valid(un) {
						rate += un
					}
				}
			}
		}
	}
	return rate, nil
}
