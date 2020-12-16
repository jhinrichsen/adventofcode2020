package aoc2020

import (
	"fmt"
	"reflect"
	"testing"
)

func testDay11(t *testing.T, filename string, part1 bool, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	d := NewDay11(lines)
	for d.Step() {
	}
	got := d.Occupied()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func genFilename(i uint) string {
	return fmt.Sprintf("testdata/day11_example%d.txt", i)
}

func gen(i uint) ([]string, error) {
	lines, err := linesFromFilename(genFilename(i))
	if err != nil {
		return nil, err
	}
	return lines, nil
}

func TestDay11ExamplePart1(t *testing.T) {
	const (
		part1 = true
		want  = 37
	)
	testDay11(t, genFilename(1), part1, want)
}

func TestDay11GenerationsPart1(t *testing.T) {
	const want = 37
	lines, err := gen(1)
	if err != nil {
		t.Fatal(err)
	}
	d := NewDay11(lines)
	for i := uint(2); i < 7; i++ {
		changed := d.Step()
		if !changed {
			t.Fatalf("step %d: expecting change but got identity", i)
		}
		wantRep, err := gen(i)
		if err != nil {
			t.Fatal(err)
		}
		gotRep := d.Redact()
		if !reflect.DeepEqual(wantRep, gotRep) {
			t.Fatalf("gen %d: want \n%s\n but got \n%s\n", i, wantRep, gotRep)
		}
	}
	got := d.Occupied()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay11Part1(t *testing.T) {
	const (
		part1 = true
		want  = 2361
	)
	testDay11(t, filename(11), part1, want)
}

const (
	empty    = 'L'
	occupied = '#'
	floor    = '.'
)

type Day11 struct {
	Grid [][]byte
}

func NewDay11(lines []string) Day11 {
	buf := make([][]byte, len(lines))
	for y := 0; y < len(lines); y++ {
		buf[y] = []byte(lines[y])

	}
	return Day11{buf}
}

// Adjacents returns number of occupied neighbours.
func (a *Day11) Adjacents(x, y int) (count uint) {
	dimx, dimy := len(a.Grid[0]), len(a.Grid)
	seat := func(y, x int) bool {
		return 0 <= x && x < dimx && 0 <= y && y < dimy && a.Grid[y][x] == occupied
	}

	// N
	if seat(y-1, x) {
		count++
	}
	// NE
	if seat(y-1, x+1) {
		count++
	}
	// E
	if seat(y, x+1) {
		count++
	}
	// SE
	if seat(y+1, x+1) {
		count++
	}
	// S
	if seat(y+1, x) {
		count++
	}
	// SW
	if seat(y+1, x-1) {
		count++
	}
	// W
	if seat(y, x-1) {
		count++
	}
	// NW
	if seat(y-1, x-1) {
		count++
	}
	return
}

func (a *Day11) Copy() Day11 {
	cp := make([][]byte, len(a.Grid))
	for i := 0; i < len(a.Grid); i++ {
		cp[i] = make([]byte, len(a.Grid[i]))
		copy(cp[i], a.Grid[i])
	}
	return Day11{cp}
}

func (a *Day11) Occupied() (n uint) {
	for _, s := range a.Grid {
		for _, c := range s {
			if c == occupied {
				n++
			}
		}
	}
	return
}

func (a *Day11) Redact() []string {
	var ss []string
	for i := range a.Grid {
		ss = append(ss, string(a.Grid[i]))
	}
	return ss
}

func (a *Day11) Step() bool {
	var changed bool
	fb := a.Copy()
	for y := range fb.Grid {
		for x := range fb.Grid[y] {
			n := a.Adjacents(x, y)
			if a.Grid[y][x] == empty && n == 0 {
				fb.Grid[y][x] = occupied
				changed = true
			} else if a.Grid[y][x] == occupied && n >= 4 {
				fb.Grid[y][x] = empty
				changed = true
			}
		}
	}
	if changed {
		a.Grid = fb.Grid
	}
	return changed
}
