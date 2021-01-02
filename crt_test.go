package aoc2020

import (
	"math/big"
	"testing"
)

// TestJaysExample solves an example system of simultaneous congruence.
// Taken from "Maths with Jay", https://youtu.be/zIFehsBHB8o
// math (american english), maths (british english)
func TestCRTJaysExample(t *testing.T) {
	want := big.NewInt(78)
	got, err := CRT([]*big.Int{
		big.NewInt(3),
		big.NewInt(1),
		big.NewInt(6),
	}, []*big.Int{
		big.NewInt(5),
		big.NewInt(7),
		big.NewInt(8),
	})
	if err != nil {
		t.Fatal(err)
	}
	if want.Cmp(got) != 0 {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}
