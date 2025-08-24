package aoc2020

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
)

func linesFromFilename(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	return linesFromReader(f)
}

func linesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)
	}
	return lines, nil
}

func exampleFilename(day int) string {
	return fmt.Sprintf("testdata/day%02d_example.txt", day)
}

func example2Filename(day int) string {
	return fmt.Sprintf("testdata/day%02d_example2.txt", day)
}

func filename(day int) string {
	return fmt.Sprintf("testdata/day%02d.txt", day)
}

// linesAsNumber converts strings into integer.
func linesAsNumbers(lines []string) ([]int, error) {
	var is []int
	for i := range lines {
		n, err := strconv.Atoi(lines[i])
		if err != nil {
			return is, fmt.Errorf("error in line %d: cannot convert %q to number",
				i, lines[i])
		}
		is = append(is, n)
	}
	return is, nil
}

func linesFromFilenameTB(tb testing.TB, filename string) []string {
	tb.Helper()
	f, err := os.Open(filename)
	if err != nil {
		tb.Fatal(err)
	}
	return linesFromReaderTB(tb, f)
}

func linesFromReaderTB(tb testing.TB, r io.Reader) []string {
	tb.Helper()
	var lines []string
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)
	}
	if err := sc.Err(); err != nil {
		tb.Fatal(err)
	}
	return lines
}

func contentFromFilename(tb testing.TB, filename string) []byte {
	tb.Helper()
	buf, err := os.ReadFile(filename)
	if err != nil {
		tb.Fatal(err)
	}
	return buf
}
