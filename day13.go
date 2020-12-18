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
