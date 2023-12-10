package outils

import "regexp"

// IsValidBitcoinTx checks if a transaction is valid.
func IsValidBitcoinTx(txID string) bool {
	txIDPattern := "^[A-Fa-f0-9]{64}$"
	regex := regexp.MustCompile(txIDPattern)

	return regex.MatchString(txID)
}
