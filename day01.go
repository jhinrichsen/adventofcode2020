package aoc2020

// Day01 returns product of two numbers that add up to 2020.
func Day01(buf []byte, part1 bool) uint {
	var ns = make(map[uint]struct{}, len(buf)/2)
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

// day01Concurrent has no error checking on valid digits 0..9.
// garbage in, garbage out
func day01Concurrent(buf []byte, part1 bool) uint {
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

	ns := make(map[uint]struct{}, len(buf)/2)
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

// Zero-cost pull iterator (no goroutines, no callback).
type numIter struct {
	b    []byte
	i    int
	done bool
}

func newNumIter(b []byte) numIter { return numIter{b: b} }

// Next returns the next number from the input buffer.
// The last line must be newline terminated.
func (it *numIter) Next() (uint, bool) {
	if it.done {
		return 0, false
	}
	var n uint
	for it.i < len(it.b) {
		c := it.b[it.i]
		it.i++
		if c == '\n' {
			return n, true
		}
		n = n*10 + uint(c-'0')
	}
	return 0, false
}

// day01Pull uses streamable iter to return early on first match without parsing complete input.
func day01Pull(buf []byte, part1 bool) uint {
	var maxDim = len(buf) / 2 // single digit + newline
	seq := newNumIter(buf)
	if part1 {
		seen := make(map[uint]struct{}, maxDim)
		for n, ok := seq.Next(); ok; n, ok = seq.Next() {
			rest := 2020 - n
			if _, ok := seen[rest]; ok {
				return n * rest
			}
			seen[n] = struct{}{}
		}
		return 0
	}

	// Part 2: streaming three-sum mittels laufender Paarsummen
	pairs := make(map[uint]uint, maxDim) // sum -> product
	vals := make([]uint, 0, maxDim)
	for n, ok := seq.Next(); ok; n, ok = seq.Next() {
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
