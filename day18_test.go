package aoc2020

import (
	"testing"
)

var day18Examples = []struct {
	in   string
	want int
}{
	{"2 * 3 + ( 4 * 5 )", 26},
	{"5 + ( 8 * 3 + 9 + 3 * 4 * 3 )", 437},
	{"5 * 9 * ( 7 * 3 * 3 + 9 * 3 + ( 8 + 6 * 4 ) )", 12240},
	{"( ( 2 + 4 * 9 ) * ( 6 + 9 * 8 + 6 ) + 6 ) + 2 + 4 * 2", 13632},
}

func TestDay18Examples(t *testing.T) {
	for _, tt := range day18Examples {
		id := tt.in
		t.Run(id, func(t *testing.T) {
			rpn := ShuntingYard(tt.in, opCfg)
			got, err := eval(rpn)
			if err != nil {
				t.Fatal(err)
			}
			if tt.want != got {
				t.Fatalf("want %d but got %d", tt.want, got)
			}
		})
	}
}

func testDay18(t *testing.T, filename string, want int) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day18(lines)
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
	testDay18(t, exampleFilename(18), want)
}

func TestDay18(t *testing.T) {
	const (
		want = 21993583522852
	)
	testDay18(t, filename(18), want)
}
