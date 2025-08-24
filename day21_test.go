package aoc2020

import (
	"testing"
)

func testDay21Part1(t *testing.T, filename string, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	d, err := NewDay21(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := d.Part1()
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func testDay21Part2(t *testing.T, filename string, want string) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	d, err := NewDay21(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := d.Part2()
	if want != got {
		t.Fatalf("want %q but got %q\n", want, got)
	}
}

func TestDay21Part1Example(t *testing.T) {
	const (
		want = 5
	)
	testDay21Part1(t, exampleFilename(21), want)
}

func TestDay21Part1(t *testing.T) {
	const (
		want = 2170
	)
	testDay21Part1(t, filename(21), want)
}

func BenchmarkDay21Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lines, err := linesFromFilename(filename(21))
		if err != nil {
			b.Fatal(err)
		}
		d, err := NewDay21(lines)
		if err != nil {
			b.Fatal(err)
		}
		_ = d.Part1()
	}
}

func TestDay21Part2Example(t *testing.T) {
	const (
		want = "mxmxvkd,sqjhc,fvjkl"
	)
	testDay21Part2(t, exampleFilename(21), want)
}

func TestDay21Part2(t *testing.T) {
	const (
		want = "nfnfk,nbgklf,clvr,fttbhdr,qjxxpr,hdsm,sjhds,xchzh"
	)
	testDay21Part2(t, filename(21), want)
}
