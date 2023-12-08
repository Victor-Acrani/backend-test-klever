package outils

import "math/big"

// BigIntSubtraction subtracts two values using big int and returns the result.
func BigIntSubtraction(a string, b string) string {
	bigA := new(big.Int)
	bigB := new(big.Int)

	// set values to big int
	bigA.SetString(a, 10)
	bigB.SetString(b, 10)

	// subtraction
	result := new(big.Int).Sub(bigA, bigB)
	return result.String()
}
