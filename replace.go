package aoc2020

// replace removes element at index i and inserts another array into this
// position.
func replace(dst []string, i int, src []string) []string {
	a := make([]string, len(dst)+len(src)-1)
	copy(a, dst[:i])
	copy(a[i:], src)
	copy(a[i+len(src):], dst[i+1:])
	return a
}
