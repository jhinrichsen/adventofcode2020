package aoc2020

import (
	"fmt"
	"strconv"
)

type Puzzle01 map[uint]struct{}

func NewDay01(lines []string) (Puzzle01, error) {
	is := make(map[uint]struct{})
	for i, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return is, fmt.Errorf("error converting number #%d: %w", i, err)
		}
		is[uint(n)] = struct{}{}
	}
	return is, nil
}

// Day01 returns product of two numbers that add up to 2020.
// func Day01(ns Puzzle01, part1 bool) uint {
func Day01(buf []byte, part1 bool) uint {
	/* Baseline
	if part1 {
		for k := range ns {
			rest := 2020 - k
			if _, ok := ns[rest]; ok {
				return k * rest
			}
		}
	} else {
		for k := range ns {
			for l := range ns {
				rest := 2020 - k - l
				if _, ok := ns[rest]; ok {
					return k * l * rest
				}
			}
		}
	}
	*/
	ns := make(map[uint]struct{}, len(buf)/2)
	var n uint
	for i := range buf {
		b := buf[i]
		if b == '\n' {
			ns[n] = struct{}{}
			n = 0
		} else {
			n = n*10 + uint(b-'0')
		}
	}

	if part1 {
		for k := range ns {
			rest := 2020 - k
			if _, ok := ns[rest]; ok {
				return k * rest
			}
		}
	} else {
		for k := range ns {
			for l := range ns {
				rest := 2020 - k - l
				if _, ok := ns[rest]; ok {
					return k * l * rest
				}
			}
		}
	}
	return 0
}

// Day01Concurrent has no error checking on valid digits 0..9.
// garbage in, garbage out
func Day01Concurrent(buf []byte, part1 bool) uint {
	c := make(chan uint)
	go func() {
		var n uint
		for i := range buf {
			b := buf[i]
			if b >= '0' && b <= '9' {
				n = n*10 + uint(b-'0')
			} else if b == '\n' {
				c <- n
				n = 0
			}
		}
		// send last number if needed
		if n != 0 || len(buf) > 0 && buf[len(buf)-1] != '\n' {
			c <- n
		}
		close(c)
	}()

	ns := make(map[uint]struct{}, len(buf)/2) // single digit, \n
	for n := range c {
		ns[n] = struct{}{}
	}

	if part1 {
		for k := range ns {
			rest := 2020 - k
			if _, ok := ns[rest]; ok {
				return k * rest
			}
		}
	} else {
		for k := range ns {
			for l := range ns {
				rest := 2020 - k - l
				if _, ok := ns[rest]; ok {
					return k * l * rest
				}
			}
		}
	}
	return 0
}
