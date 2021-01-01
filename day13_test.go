package aoc2020

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"reflect"
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

func BenchmarkDay13Part1ExcludingReading(b *testing.B) {
	buf, err := ioutil.ReadFile(filename(13))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day13(buf, true)
	}
}

const (
	x = 1 // neutral element for modulo operations, n % 1 = 0, n ∈ ℕ
)

var day13Part2Examples = []struct {
	input []int
	want  uint
}{
	{[]int{67, 7, 59, 61}, 754018},
	{[]int{7, 13, x, x, 59, x, 31, 19}, 1068781},
	//     0  12        55     25  12   (remainder)
	{[]int{17, x, 13, 19}, 3417},
	{[]int{67, 7, 59, 61}, 754018},
	{[]int{67, x, 7, 59, 61}, 779210},
	{[]int{67, 7, x, 59, 61}, 1261476},
	{[]int{1789, 37, 47, 1889}, 1202161486},
}

func TestDay13Part2ExamplesBruteForce(t *testing.T) {
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

/*
func TestDay13Part2ExamplesCRT(t *testing.T) {
	for _, tt := range day13Part2Examples {
		id := fmt.Sprintf("%+v", tt.input)
		t.Run(id, func(t *testing.T) {
			// skip last example in short testing mode
			if testing.Short() && tt.want > 10_000_000 {
				msg := fmt.Sprintf("skipping long running %+v",
					tt.input)
				t.Skip(msg)
			}

			got, err := Day13Part2CRT(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if tt.want != got {
				t.Fatalf("want %d but got %d", tt.want, got)
			}
		})
	}
}
*/

// convert bus lines into remainders.
func remainders(buses []int) []int {
	rs := make([]int, len(buses))
	// convert from timeline to remainder
	for i := 1; i < len(rs); i++ {
		if buses[i] == 1 {
			rs[i] = 1
		} else {
			rs[i] = buses[i] - i
		}
	}
	return rs
}

func TestDay13Part2Remainders(t *testing.T) {
	want := []int{0, 12, 1, 1, 55, 1, 25, 12}
	got := remainders([]int{7, 13, x, x, 59, x, 31, 19})
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %+v but got %+v", want, got)
	}
}

// Signature makes sure the calling conventions for CRT are set right.
func TestDay13Part2CheckCRTCallingConvention(t *testing.T) {
	tt := day13Part2Examples[1]
	var as []*big.Int
	var ks []*big.Int
	for i, bus := range tt.input {
		// ignore x
		if bus == x {
			continue
		}
		// convert bus line to remainder
		rem := bus - i
		as = append(as, big.NewInt(int64(rem)))
		ks = append(ks, big.NewInt(int64(bus)))
	}
	got, err := CRT(as, ks)
	if err != nil {
		t.Fatal(err)
	}
	if uint64(tt.want) != got.Uint64() {
		t.Fatalf("want %d but got %d", tt.want, got)
	}
}

func TestDay13Part2(t *testing.T) {
	const want = 415579909629976

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

	// got := Day13Part2BruteForce(idxs)
	got, err := Day13Part2CRT(idxs)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func BenchmarkDay13Part2(b *testing.B) {
	const want = 415579909629976

	for i := 0; i < b.N; i++ {
		lines, err := linesFromFilename(filename(13))
		if err != nil {
			b.Fatal(err)
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
					b.Fatalf("error parsing field %d: %+v",
						i, err)
				}
				line = n
			}
			idxs[i] = line
		}

		// got := Day13Part2BruteForce(idxs)
		got, err := Day13Part2CRT(idxs)
		if err != nil {
			b.Fatal(err)
		}
		if want != got {
			b.Fatalf("want %d but got %d\n", want, got)
		}
	}
}

func BenchmarkDay13Part2ExcludingReading(b *testing.B) {
	const want = 415579909629976

	lines, err := linesFromFilename(filename(13))
	if err != nil {
		b.Fatal(err)
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
				b.Fatalf("error parsing field %d: %+v",
					i, err)
			}
			line = n
		}
		idxs[i] = line
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got, err := Day13Part2CRT(idxs)
		if err != nil {
			b.Fatal(err)
		}
		if want != got {
			b.Fatalf("want %d but got %d\n", want, got)
		}
	}
}
