package aoc2020

import (
	"strconv"
	"strings"
)

// Day4 returns number of valid passports.
func Day4(lines []string, part1 bool) uint {
	n := uint(0)
	newPass := func() map[string]string {
		return make(map[string]string)
	}
	pass := newPass()
	has := func(key string) bool {
		var ok bool
		_, ok = pass[key]
		return ok
	}
	between := func(s string, min, max uint) bool {
		n, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		un := uint(n)
		return min <= un && un <= max
	}
	keyBetween := func(key string, min, max uint) bool {
		s, ok := pass[key]
		if !ok {
			return false
		}
		return between(s, min, max)
	}
	validHgt := func(s string) bool {
		isCm := strings.HasSuffix(s, "cm")
		isIn := strings.HasSuffix(s, "in")
		if !(isCm || isIn) {
			return false
		}
		s = s[:len(s)-2]
		if isCm {
			return between(s, 150, 193)
		}
		return between(s, 59, 76)
	}
	validHcl := func(s string) bool {
		// one #
		if !strings.HasPrefix(s, "#") {
			return false
		}
		// ...followed by exactly six characters...
		if len(s) != 7 {
			return false
		}
		// ...0-9 or a-f
		num := s[1:]
		for i := range num {
			b := num[i]
			isNum := '0' <= b && b <= '9'
			isAf := 'a' <= b && b <= 'f'
			isHex := isNum || isAf
			if !isHex {
				return false
			}
		}
		return true
	}
	validEcl := func(s string) bool {
		if s == "amb" || s == "blu" || s == "brn" || s == "gry" || s == "grn" || s == "hzl" || s == "oth" {
			return true
		}
		return false
	}
	validPid := func(s string) bool {
		if len(s) != 9 {
			return false
		}
		for i := range s {
			if s[i] < '0' || s[i] > '9' {
				return false
			}
		}
		return true
	}

	var valid func() bool
	if part1 {
		valid = func() bool {
			return has("byr") &&
				has("iyr") &&
				has("eyr") &&
				has("hgt") &&
				has("hcl") &&
				has("ecl") &&
				has("pid")
		}
	} else {
		valid = func() bool {
			return keyBetween("byr", 1920, 2002) &&
				keyBetween("iyr", 2010, 2020) &&
				keyBetween("eyr", 2020, 2030) &&
				validHgt(pass["hgt"]) &&
				validHcl(pass["hcl"]) &&
				validEcl(pass["ecl"]) &&
				validPid(pass["pid"])
		}
	}
	add := func(s string) { // add all keys to current pass
		fields := strings.Fields(s)
		for i := range fields {
			kv := strings.Split(fields[i], ":")
			pass[kv[0]] = kv[1]
		}
	}

	// append an empty line to make sure last pass is validated
	lines = append(lines, "")
	for _, line := range lines {
		if len(line) == 0 {
			if valid() {
				n++
			}
			pass = newPass()
		} else {
			add(line)
		}
	}
	return n
}
