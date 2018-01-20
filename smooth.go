package smooth

import (
	"math"
	"math/big"
)

var (
	int0    = big.NewInt(0)
	int1    = big.NewInt(1)
	intMU64 = new(big.Int).SetUint64(math.MaxUint64)
)

// IsSmooth reports whether n is k-smooth.
// An integer is k-smooth if it has no prime factors greater than k.
// IsSmooth panics if k is negative.
func IsSmooth(n, k *big.Int) bool {
	if k.Cmp(int0) == -1 {
		panic("IsSmooth: negative value for k")
	}
	if n.Cmp(k) != 1 {
		return true
	}
	// ProbablyPrime is deterministic for inputs less than 2^64
	if n.Cmp(intMU64) != 1 && n.ProbablyPrime(0) {
		return false
	}

	m := new(big.Int).Set(n)
	for p := big.NewInt(2); p.Cmp(k) != 1; p.Add(p, int1) {
		if !p.ProbablyPrime(0) {
			continue
		}
		for q, r := new(big.Int).Set(m), new(big.Int); r.Cmp(int0) == 0; {
			m.Set(q)
			q, r = new(big.Int).QuoRem(m, p, r)
		}
		if m.Cmp(int1) == 0 {
			return true
		}
	}
	return false
}
