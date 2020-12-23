package aoc2020

import (
	"fmt"
)

const Dim = 10 // number of digits in input, avoids allocation

// parses n into its digits, returning number of digits
func parse(n int, buf *[Dim]int) int {
	var digits int
	for n > 0 {
		digits++
		(*buf)[digits] = n % 10
		n /= 10
	}
	// reverse
	for i := 0; i < Dim/2; i++ {
		(*buf)[i], (*buf)[Dim-1-i] = (*buf)[Dim-1-i], (*buf)[i]
	}
	return digits
}

func Day23(input int, moves int) int {
	var cups, picked [Dim]int

	ncups, npicked := parse(input, &cups), 0

	// some helper

	labelRange := func() (int, int) {
		min, max := cups[0], cups[0]
		for i := 1; i < ncups; i++ {
			if cups[i] < min {
				min = cups[i]
			}
			if cups[i] > max {
				max = cups[i]
			}
		}
		return min, max
	}

	pickUp := func() { // move from cups to picked
		n := 3
		for i := 0; i < n; i++ {
			picked[i] = cups[1+i]
		}
		npicked = n

		// 'remove' from cups by shifting tail to the left
		l := ncups - n - 1
		for i := 0; i < l; i++ {
			cups[1+i] = cups[1+n+i]
		}
		ncups -= n
	}
	pickDown := func(idx int) { // move from picked to cups
		// move everything after idx to the right
		n := npicked
		l := ncups - idx
		for i := l - 1; i >= 0; i-- { // iterate in reverse to avoid stepping on our own foot
			cups[idx+n+i] = cups[idx+i]
		}
		ncups += n
		// move picked to cups
		for i := 0; i < npicked; i++ {
			cups[idx+i] = picked[i]
		}
		npicked = 0
	}
	isPicked := func(label int) bool {
		for i := range picked {
			if picked[i] == label {
				return true
			}
		}
		return false
	}
	cupIdx := func(label int) int { // label -> index
		for i := 0; i < ncups; i++ {
			if cups[i] == label {
				return i
			}
		}
		msg := "internal error, cannot find cup %d in cups %+v"
		panic(fmt.Sprintf(msg, label, cups))
	}

	revolve := func() {
		tmp := cups[0]
		n := ncups - 1
		for i := 0; i < n; i++ {
			cups[i] = cups[i+1]
		}
		cups[ncups-1] = tmp
	}
	min, max := labelRange()
	for move := 0; move < moves; move++ {
		// fmt.Printf("\n-- move %d --\n", move+1)
		// fmt.Printf(cupsRep())

		pickUp()
		// fmt.Printf(pickedRep())

		dstLabel := cups[0] - 1
		for isPicked(dstLabel) {
			dstLabel--
			if dstLabel < min {
				dstLabel = max
			}
		}
		// fmt.Printf("destination: %d\n", dstLabel)
		idx := cupIdx(dstLabel)
		pickDown(idx + 1) // pick down to the right

		// pick the next cup, the one to the right. Instead of moving
		// 'current' around a ring buffer, we rotate the ring buffer
		revolve()
	}

	nextCup := func(i int) int {
		return (i + 1) % ncups
	}

	idx := nextCup(cupIdx(1))
	labels := 0
	for i := 0; i < ncups-1; i++ { // -1 leaves out first cup '1'
		labels *= 10
		labels += cups[idx]
		idx = nextCup(idx)
	}
	return labels
}
