package outils_test

import (
	"testing"

	"github.com/Victor-Acrani/backend-test-klever/outils"
)

func TestBigIntSubtraction(t *testing.T) {
	a := "1000000"
	b := "540321"
	expected := "459679"

	actual := outils.BigIntSubtraction(a, b)

	if actual != expected {
		t.Errorf("big int subtraction result should be: expected %v, actual %v",
			expected, actual)
	}
}
