package aoc2020

import (
	"testing"
)

func testDay20Part1(t *testing.T, filename string, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	d, err := NewDay20(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := d.CornerProduct()
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay20Part1(b *testing.B) {
	const (
		part1 = true
		want  = 45079100979683
	)
	for i := 0; i < b.N; i++ {
		lines, err := linesFromFilename(filename(20))
		if err != nil {
			b.Fatal(err)
		}
		d, err := NewDay20(lines)
		if err != nil {
			b.Fatal(err)
		}
		got, err := d.CornerProduct()
		if err != nil {
			b.Fatal(err)
		}
		if want != got {
			b.Fatalf("want %d but got %d", want, got)
		}
	}
}

func TestDay20Example(t *testing.T) {
	testDay20Part1(t, exampleFilename(20), 20899048083289)
}

func TestDay20Part1(t *testing.T) {
	testDay20Part1(t, filename(20), 45079100979683)
}

func TestDay20Part2Example(t *testing.T) {
	lines, err := linesFromFilename(exampleFilename(20))
	if err != nil {
		t.Fatal(err)
	}
	d, err := NewDay20(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := d.WaterRoughness()
	if err != nil {
		t.Fatal(err)
	}
	const want = 273
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay20Part2(t *testing.T) {
	lines, err := linesFromFilename(filename(20))
	if err != nil {
		t.Fatal(err)
	}
	d, err := NewDay20(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := d.WaterRoughness()
	if err != nil {
		t.Fatal(err)
	}
	const want = 1946
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay20Part2(b *testing.B) {
	const want = 1946
	for i := 0; i < b.N; i++ {
		lines, err := linesFromFilename(filename(20))
		if err != nil {
			b.Fatal(err)
		}
		d, err := NewDay20(lines)
		if err != nil {
			b.Fatal(err)
		}
		got, err := d.WaterRoughness()
		if err != nil {
			b.Fatal(err)
		}
		if want != got {
			b.Fatalf("want %d but got %d", want, got)
		}
	}
}
