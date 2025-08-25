package aoc2020

import (
	"fmt"
	"iter"
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

// Day01Iter liefert einen Iterator (iter.Seq[uint]) über alle Zahlen im Buffer.
// Annahmen: Die Eingabe besteht aus Ziffern und '\n' und endet mit Newline.
func Day01Iter(buf []byte) iter.Seq[uint] {
	return func(yield func(uint) bool) {
		var n uint
		for i := range buf {
			b := buf[i]
			if b == '\n' {
				if !yield(n) {
					return
				}
				n = 0
			} else {
				n = n*10 + uint(b-'0')
			}
		}
	}
}

// Day01Pull uses streamable iter to return early on first match without parsing complete input.
func Day01Pull(seq iter.Seq[uint], part1 bool) uint {
    next, stop := iter.Pull(seq)
	defer stop()

	if part1 {
		seen := make(map[uint]struct{})
		for {
			n, ok := next()
			if !ok {
				break
			}
			rest := 2020 - n
			if _, ok := seen[rest]; ok {
				return n * rest
			}
			seen[n] = struct{}{}
		}
		return 0
	}

	// Part 2: streaming three-sum mittels laufender Paarsummen
	pairs := make(map[uint]uint) // sum -> produkt der beiden Werte
	vals := make([]uint, 0, 128)
	for {
		n, ok := next()
		if !ok {
			break
		}
		rest := 2020 - n
		if prod, ok := pairs[rest]; ok {
			return prod * n
		}
		for _, v := range vals {
			pairs[n+v] = n * v
		}
		vals = append(vals, n)
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
