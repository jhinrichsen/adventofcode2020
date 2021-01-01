package aoc2020

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func testDay13(t *testing.T, filename string, part1 bool, want uint) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	got := Day13(buf, part1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay13Example(t *testing.T) {
	const (
		part1 = true
		want  = 295
	)
	testDay13(t, exampleFilename(13), part1, want)
}

func TestDay13(t *testing.T) {
	const (
		part1 = true
		want  = 3865
	)
	testDay13(t, filename(13), part1, want)
}

func BenchmarkDay13Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf, err := ioutil.ReadFile(filename(13))
		if err != nil {
			b.Fatal(err)
		}
		_ = Day13(buf, true)
	}
}

const (
	x = 1 // neutral modulo, n % 1 = 0
)

var day13Part2Examples = []struct {
	input []int
	want  uint
}{
	{[]int{67, 7, 59, 61}, 754018},
	{[]int{7, 13, x, x, 59, x, 31, 19}, 1068781},
	{[]int{17, x, 13, 19}, 3417},
	{[]int{67, 7, 59, 61}, 754018},
	{[]int{67, x, 7, 59, 61}, 779210},
	{[]int{67, 7, x, 59, 61}, 1261476},
	{[]int{1789, 37, 47, 1889}, 1202161486},
}

func TestDay13Part2Examples(t *testing.T) {
	for _, tt := range day13Part2Examples {
		id := fmt.Sprintf("%+v", tt.input)
		t.Run(id, func(t *testing.T) {
			// skip last example in short testing mode
			if testing.Short() && tt.want > 10_000_000 {
				msg := fmt.Sprintf("skipping long running %+v",
					tt.input)
				t.Skip(msg)
			}

			got := Day13Part2BruteForce(tt.input)
			if tt.want != got {
				t.Fatalf("want %d but got %d", tt.want, got)
			}
		})
	}
}

func TestDay13Part2(t *testing.T) {
	const want = 13

	if testing.Short() {
		t.Skip("skipping part #2 because i can only do brute force")
	}

	lines, err := linesFromFilename(filename(13))
	if err != nil {
		t.Fatal(err)
	}

	// convert to numbers
	buses := strings.Split(lines[1], ",")
	idxs := make([]int, len(buses))
	for i := range buses {
		var line int
		if buses[i] == "x" {
			line = 1
		} else {
			n, err := strconv.Atoi(buses[i])
			if err != nil {
				t.Fatalf("error parsing field %d: %+v", i, err)
			}
			line = n
		}
		idxs[i] = line
	}

	got := Day13Part2BruteForce(idxs)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}
