package aoc2020

import (
	"fmt"
	"strings"
	"testing"
)

func testDay23(t *testing.T, input uint, moves uint, want uint) {
	got := Day23(input, moves)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay23Example(t *testing.T) {
	const (
		input = 389125467
		moves = 10
		want  = 13
	)
	testDay23(t, input, moves, want)
}

/*
func TestDay23(t *testing.T) {
	const (
		input = 962713854
		want = 13
	)
	testDay23(t, filename(23), want)
}
*/

func Day23(input uint, moves uint) uint {
	parse := func(input uint) (cups []byte) {
		for input > 0 {
			rem := input / 10
			lastDigit := byte(input - (rem * 10))
			cups = append(cups, lastDigit)
			input = rem
		}

		// reverse
		n := len(cups)
		for i := 0; i < n/2; i++ {
			cups[i], cups[n-1-i] = cups[n-1-i], cups[i]
		}
		return
	}
	current := 0
	cups := parse(input)
	pcups := make([]byte, len(cups)) // pickup cups
	labelRange := func() (byte, byte) {
		min, max := cups[0], cups[0]
		for i := 1; i < len(cups); i++ {
			if cups[i] < min {
				min = cups[i]
			}
			if cups[i] > max {
				max = cups[i]
			}
		}
		return min, max
	}

	cupsRep := func() string {
		var sb strings.Builder
		r := sb.WriteRune
		sb.WriteString("cups: ")
		for i := 0; i < len(cups); i++ {
			if i == current {
				r('(')
			}
			r(rune(cups[i] + '0'))
			if i == current {
				r(')')
			}
			if i < len(cups)-1 {
				r(' ')
			}
		}
		r('\n')
		return sb.String()
	}
	pcupsRep := func() string {
		var sb strings.Builder
		r := sb.WriteRune
		sb.WriteString("pick up: ")
		for i := 0; i < len(pcups); i++ {
			r(rune(pcups[i] + '0'))
			if i < len(pcups)-1 {
				r(',')
				r(' ')
			}
		}
		r('\n')
		return sb.String()
	}
	pickUp := func() {
		n := int(cups[current])
		from := current + 1
		pcups = cups[from : from+n]
		fmt.Printf("\tpcups: %+d\n", pcups)

		// remove from cups
		c := make([]byte, len(cups)-n)
		copy(c, cups[0:current+1])
		for i, j := current+1, current+n+1; j < len(cups); i, j = i+1, j+1 {
			fmt.Printf("\tc[%d] <- cups[%d] = %d\n", i, j, cups[j])
			c[i] = cups[j]
		}
		fmt.Printf("\tc: %+d\n", c)
		cups = c
		fmt.Printf("\tpcups2: %+d\n", pcups)
	}
	pickDown := func(idx byte) { // copy from pcups to cups
		fmt.Printf("\tcups before pickDown: %+v\n", cups)
		l := len(cups)
		n := len(pcups)

		c := make([]byte, l+n)
		copy(c, cups[:idx]) // first part of original cups
		for i := int(idx) + n; i < l-int(idx); i++ {
			fmt.Printf("\tcups[%d] <- cups[%d] = %d\n", i, i-n, cups[i-n])
			cups[i] = cups[i-n]
		}
		fmt.Printf("\tcups after pickDown: %+v\n", cups)
	}
	picked := func(label byte) bool {
		for i := range pcups {
			if pcups[i] == label {
				return true
			}
		}
		return false
	}
	cupIdx := func(label byte) byte { // return index of label
		for i := 0; i < len(cups); i++ {
			if cups[i] == label {
				return byte(i)
			}
		}
		msg := "internal error, cannot find cup %d in cups %+v"
		panic(fmt.Sprintf(msg, label, cups))
	}

	min, max := labelRange()
	for move := uint(0); move < moves; move++ {
		fmt.Printf("\n-- move %d --\n", move+1)
		fmt.Printf(cupsRep())

		pickUp()
		fmt.Printf(pcupsRep())

		dstLabel := cups[current] - 1
		for picked(dstLabel) {
			dstLabel--
			if dstLabel < min {
				dstLabel = max
			}
		}
		fmt.Printf("destination: %d\n", dstLabel)
		idx := cupIdx(dstLabel)
		pickDown(idx)
	}
	return 0
}
