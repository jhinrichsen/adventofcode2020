package aoc2020

import (
	"reflect"
	"testing"
)

func TestReplace(t *testing.T) {
	want := []string{"1", "2", "3", "4", "5", "6"}
	a := []string{"1", "2", "2", "5", "6"}
	b := []string{"3", "4"}
	got := replace(a, 2, b)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %+v but got %+v", want, got)
	}
}
