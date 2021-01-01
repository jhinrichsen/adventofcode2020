package aoc2020

import (
	"fmt"
	"math/big"
)

// CRT solves simultaneous congruences, also known as the Chinese Remainder
// Theoreom.
//
// x ≡ a₁ ( mod n₁ )
// x ≡ a₂ ( mod n₂ )
// ⋮
// x ≡ aₖ ( mod nₖ )

//
func CRT(a, n []*big.Int) (*big.Int, error) {
	var one = big.NewInt(1)

	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}
