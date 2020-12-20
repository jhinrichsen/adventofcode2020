package aoc2020

import (
	"testing"
)

var day18Examples = []struct {
	in        string
	wantPart1 int
	wantPart2 int
}{
	{"2 * 3 + ( 4 * 5 )", 26, 46},
	{"5 + ( 8 * 3 + 9 + 3 * 4 * 3 )", 437, 1445},
	{"5 * 9 * ( 7 * 3 * 3 + 9 * 3 + ( 8 + 6 * 4 ) )", 12240, 669060},
	{"( ( 2 + 4 * 9 ) * ( 6 + 9 * 8 + 6 ) + 6 ) + 2 + 4 * 2", 13632, 23340},
}

func TestDay18Examples(t *testing.T) {
	f := func(t *testing.T, part1 bool, infix string, want int) {
		var cfg OperatorConfiguration
		if part1 {
			cfg = part1Cfg
		} else {
			cfg = part2Cfg
		}
		rpn := ShuntingYard(infix, cfg)
		got, err := eval(rpn)
		if err != nil {
			t.Fatal(err)
		}
		if want != got {
			t.Fatalf("want %d but got %d", want, got)
		}
	}
	for _, tt := range day18Examples {
		id := tt.in
		t.Run(id, func(t *testing.T) {
			f(t, true, tt.in, tt.wantPart1)
			f(t, false, tt.in, tt.wantPart2)
		})
	}
}

func testDay18(t *testing.T, filename string, part1 bool, want int) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day18(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Example(t *testing.T) {
	const (
		want = 26 + 437 + 12240 + 13632
	)
	testDay18(t, exampleFilename(18), true, want)
}

func TestDay18Part1(t *testing.T) {
	const (
		part1 = true
		want  = 21993583522852
	)
	testDay18(t, filename(18), part1, want)
}

func TestDay18Part2(t *testing.T) {
	const (
		part1 = false
		want  = 122438593522757
	)
	testDay18(t, filename(18), part1, want)
}
