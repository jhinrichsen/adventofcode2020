package aoc2020

import (
	"testing"
)

func testDay19(t *testing.T, filename string, part1 bool, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day19(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay19Example(t *testing.T) {
	const (
		part1 = true
		want  = 2
	)
	testDay19(t, exampleFilename(19), part1, want)
}

func TestDay19(t *testing.T) {
	const (
		part1 = true
		want  = 203
	)
	testDay19(t, filename(19), part1, want)
}

func TestDay19Part2Example(t *testing.T) {
	// Example from the Part Two description (already using updated rules 8 and 11)
	lines := []string{
		"42: 9 14 | 10 1",
		"9: 14 27 | 1 26",
		"10: 23 14 | 28 1",
		"1: \"a\"",
		"11: 42 31 | 42 11 31",
		"5: 1 14 | 15 1",
		"19: 14 1 | 14 14",
		"12: 24 14 | 19 1",
		"16: 15 1 | 14 14",
		"31: 14 17 | 1 13",
		"6: 14 14 | 1 14",
		"2: 1 24 | 14 4",
		"0: 8 11",
		"13: 14 3 | 1 12",
		"15: 1 | 14",
		"17: 14 2 | 1 7",
		"23: 25 1 | 22 14",
		"28: 16 1",
		"4: 1 1",
		"20: 14 14 | 1 15",
		"3: 5 14 | 16 1",
		"27: 1 6 | 14 18",
		"14: \"b\"",
		"21: 14 1 | 1 14",
		"25: 1 1 | 1 14",
		"22: 14 14",
		"8: 42 | 42 8",
		"26: 14 22 | 1 20",
		"18: 15 15",
		"7: 14 5 | 1 21",
		"24: 14 1",
		"",
		"abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa",
		"bbabbbbaabaabba",
		"babbbbaabbbbbabbbbbbaabaaabaaa",
		"aaabbbbbbaaaabaababaabababbabaaabbababababaaa",
		"bbbbbbbaaaabbbbaaabbabaaa",
		"bbbababbbbaaaaaaaabbababaaababaabab",
		"ababaaaaaabaaab",
		"ababaaaaabbbaba",
		"baabbaaaabbaaaababbaababb",
		"abbbbabbbbaaaababbbbbbaaaababb",
		"aaaaabbaabaaaaababaa",
		"aaaabbaaaabbaaa",
		"aaaabbaabbaaaaaaabbbabbbaaabbaabaaa",
		"babaaabbbaaabaababbaabababaaab",
		"aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba",
		"",
	}
	got, err := Day19(lines, false)
	if err != nil {
		t.Fatal(err)
	}
	const want = 12
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay19Part2(t *testing.T) {
    const (
        part1 = false
        want  = 304
    )
    testDay19(t, filename(19), part1, want)
}
