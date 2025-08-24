package aoc2020

import "testing"

func testDay2(t *testing.T, filename string, part1 bool, want uint) {
    lines, err := linesFromFilename(filename)
    if err != nil {
        t.Fatal(err)
    }
    p, err := NewDay02(lines)
    if err != nil {
        t.Fatal(err)
    }
    got := Day02(p, part1)
    if want != got {
        t.Fatalf("want %d but got %d", want, got)
    }
}

func TestDay02Part1Example(t *testing.T) {
	testDay2(t, exampleFilename(2), true, 2)
}

func TestDay02Part1(t *testing.T) {
	testDay2(t, filename(2), true, 536)
}

func TestDay02Part2Example(t *testing.T) {
	testDay2(t, exampleFilename(2), false, 1)
}

func TestDay02Part2(t *testing.T) {
	testDay2(t, filename(2), false, 558)
}

func BenchmarkDay02Part1(b *testing.B) {
    lines := linesFromFilenameTB(b, filename(2))
    p, err := NewDay02(lines)
    if err != nil { b.Fatal(err) }
    b.ResetTimer()
    for b.Loop() {
        _ = Day02(p, true)
    }
}

func BenchmarkDay02Part2(b *testing.B) {
    lines := linesFromFilenameTB(b, filename(2))
    p, err := NewDay02(lines)
    if err != nil { b.Fatal(err) }
    b.ResetTimer()
    for b.Loop() {
        _ = Day02(p, false)
    }
}
