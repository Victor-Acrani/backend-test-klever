package outils_test

import (
	"testing"

	"github.com/Victor-Acrani/backend-test-klever/outils"
)

func TestIsValidBitcoinTx(t *testing.T) {
	tx := "3654d26660dcc05d4cfb25a1641a1e61f06dfeb38ee2279bdb049d018f1830ab"
	ok := outils.IsValidBitcoinTx(tx)

	if !ok {
		t.Errorf("transaction should ve valid: expected %v, actual %v", true, ok)
	}
}
