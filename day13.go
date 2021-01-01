package aoc2020

// look mom - no imports

// Day13  returns product of minutes to wait and bus ID.
// lets roll our own stateful parser, just because we can...
func Day13(buf []byte, part1 bool) uint {
	var n, busID, timestamp, minBusID uint
	minWait := ^uint(0)
	for i := 0; i < len(buf); i++ {
		b := buf[i]
		if b == '\n' && timestamp == 0 { // first line complete
			timestamp = n
			n = 0
		} else if b == '\n' || b == ',' { // field in second line complete
			busID = n
			n = 0
			wait := busID - (timestamp % busID) // modulo is time waited since departure
			if wait < minWait {
				minWait = wait
				minBusID = busID
			}
		} else if b == 'x' { // skip 'x'
			i++ // ignore next ,
		} else if '0' <= b && b <= '9' { // build a number
			n = 10*n + uint(b-'0')
		}
	}
	return minWait * minBusID
}

// Day13Part2BruteForce returns the point in time when all buses depart at offsets
// matching their positions in the list.
// Based on my 2016 day 15 solution.
func Day13Part2BruteForce(input []int) uint {
	// transform into functions where index is t and value is modulo
	var fns []func(int) bool
	for i := 0; i < len(input); i++ {
		modulus := input[i] // remember Go kills iterator index in closures
		fns = append(fns, func(t int) bool {
			p := t%modulus == 0
			return p
		})
	}

	// return true if bus index n departs one minute before bus index n+1
	nextMinute := func(idx, t int) bool {
		p1 := fns[idx](t)
		p2 := fns[idx+1](t + 1)
		return p1 && p2
	}

	match := func(t uint) bool {
		for j := 0; j < len(fns)-1; j++ {
			p := nextMinute(j, int(t)+j)
			if !p {
				return false
			}
		}
		return true
	}

	for t := uint(0); ; t++ {
		if match(t) {
			return t
		}
	}
}
