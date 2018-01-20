package smooth

import (
	"math/big"
	"testing"
)

func TestIsSmooth(t *testing.T) {
	var (
		big2   = new(big.Int).Exp(big.NewInt(2), big.NewInt(96), nil)
		big3   = new(big.Int).Exp(big.NewInt(3), big.NewInt(64), nil)
		bigmul = new(big.Int).MulRange(41, 61)
	)
	for _, test := range []struct {
		n, k *big.Int
		want bool
	}{
		{new(big.Int).Set(big2), big.NewInt(2), true},
		{new(big.Int).Set(big3), big.NewInt(2), false},
		{new(big.Int).Add(big2, int1), big.NewInt(2), false},
		{new(big.Int).Add(big3, int1), big.NewInt(2), false},
		{new(big.Int).Set(bigmul), big.NewInt(81), true},
		{new(big.Int).Set(bigmul), big.NewInt(59), false},
	} {
		if got := IsSmooth(test.n, test.k); got != test.want {
			t.Errorf("IsSmooth(%v, %v): got %v, want %v", test.n, test.k, got, test.want)
		}
	}
}
