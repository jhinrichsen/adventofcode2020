package aoc2020

import (
	"io/ioutil"
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
	buf, err := ioutil.ReadFile(filename(13))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day13(buf, true)
	}
}

func BenchmarkDay13Part1IncludingInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf, err := ioutil.ReadFile(filename(13))
		if err != nil {
			b.Fatal(err)
		}
		_ = Day13(buf, true)
	}
}
