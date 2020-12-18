package aoc2020

import (
	"fmt"
	"strconv"
	"testing"
)

func testDay12(t *testing.T, filename string, part1 bool, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	d, err := NewDay12(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := d.ManhattanDistance()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay12ExamplePart1(t *testing.T) {
	const (
		part1 = true
		want  = 25
	)
	testDay12(t, exampleFilename(12), part1, want)
}

func TestDay12Part1(t *testing.T) {
	const (
		part1 = true
		want  = 2057
	)
	testDay12(t, filename(12), part1, want)
}

func BenchmarkDay12Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(12))
	if err != nil {
		b.Fatal(err)
	}
	d, err := NewDay12(lines)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = d.ManhattanDistance()
	}
}

type Day12 struct {
	Direction    complex128
	Instructions []Instruction
}

type Instruction struct {
	Action byte
	Value  uint
}

func (a Instruction) String() string {
	return fmt.Sprintf("%c%d", a.Action, a.Value)
}

func NewDay12(lines []string) (Day12, error) {
	direction := 1 + 0i // "The ship starts by facing east."
	var is []Instruction
	for i := range lines {
		action := lines[i][0]
		s := lines[i][1:]
		value, err := strconv.Atoi(s)
		if err != nil {
			return Day12{}, fmt.Errorf("line %d: error parsing number %q", i, s)
		}
		is = append(is, Instruction{action, uint(value)})
	}
	return Day12{direction, is}, nil
}

func (a Day12) ManhattanDistance() uint {
	pos := 0 + 0i
	for _, in := range a.Instructions {
		f := float64(in.Value)

		switch in.Action {
		case 'N':
			pos += complex(0.0, f)
		case 'S':
			pos += complex(0.0, -f)
		case 'E':
			pos += complex(f, 0.0)
		case 'W':
			pos += complex(-f, 0.0)
		case 'L':
			// turn left n times
			for n := in.Value / 90; n > 0; n-- {
				a.Direction *= 0 + 1i
			}
		case 'R':
			// turn right n times
			for n := in.Value / 90; n > 0; n-- {
				a.Direction *= 0 - 1i
			}
		case 'F':
			pos += a.Direction * complex(f, 0)
		}
	}
	dst := func(f float64) int {
		n := int(f)
		if n < 0 {
			n = -n
		}
		return n
	}
	return uint(dst(real(pos)) + dst(imag(pos)))
}
