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
