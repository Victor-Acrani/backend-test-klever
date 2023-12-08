package outils_test

import (
	"testing"

	"github.com/Victor-Acrani/backend-test-klever/outils"
)

func TestIsValidBitcoinAddress(t *testing.T) {
	address := "bc1qyzxdu4px4jy8gwhcj82zpv7qzhvc0fvumgnh0r"
	result := outils.IsValidBitcoinAddress(address)

	if !result {
		t.Errorf("bitcoin address should be valid: expected %v, actual %v",
			true, result)
	}
}
