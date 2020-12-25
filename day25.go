package aoc2020

// Day25 calculates the encryption key for public keys.
func Day25(pk1, pk2 int) int {
	loopSize := func(pk int) int {
		const subject = 7
		n := 1
		for i := 1; n != pk; i++ {
			n *= subject
			n %= 20201227
			if n == pk {
				return i
			}
		}
		return 0
	}

	encryptionKey := func(subject, ls2 int) int {
		n := 1
		for i := 0; i < ls2; i++ {
			n *= subject
			n %= 20201227
		}
		return n

	}

	l1 := loopSize(pk1)
	l2 := loopSize(pk2)

	// use smaller loop size to calculate encryption key
	l0 := l1
	pk0 := pk2
	if l2 < l0 {
		l0 = l2
		pk0 = pk1
	}

	return encryptionKey(pk0, l0)
}
