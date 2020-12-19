package aoc2020

import (
	"testing"
)

func testDay16(t *testing.T, filename string, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day16(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay16Example(t *testing.T) {
	const (
		want = 71
	)
	testDay16(t, exampleFilename(16), want)
}

func TestDay16(t *testing.T) {
	const (
		want = 25984
	)
	testDay16(t, filename(16), want)
}
