package types_test

import (
	"testing"

	"github.com/Victor-Acrani/backend-test-klever/outils"
	"github.com/Victor-Acrani/backend-test-klever/types"
)

func TestSendDtoValidate(t *testing.T) {
	t.Run("valid request", func(t *testing.T) {
		sendDTO := types.SendRequestDTO{
			Address: "bc1qyzxdu4px4jy8gwhcj82zpv7qzhvc0fvumgnh0r",
			Amount:   "12345678912",
		}

		ok := outils.IsValidBitcoinAddress(sendDTO.Address)
		if !ok {
			t.Errorf("ok error should be: expected %v , actual %v", true, ok)
		}

		err := sendDTO.Validate()
		if err != nil {
			t.Errorf("validate error should be: expected %v , actual %v", nil, err)
		}
	})

	t.Run("incomplete body", func(t *testing.T) {
		sendDTO := types.SendRequestDTO{}

		ok := outils.IsValidBitcoinAddress(sendDTO.Address)
		if ok {
			t.Errorf("ok error should be: expected %v , actual %v", false, ok)
		}

		err := sendDTO.Validate()
		if err == nil {
			t.Errorf("validate error should be: expected %v , actual %v",
			 "Error:Field validation for 'Value' failed on the 'numeric' tag", 
			 err)
		}
	})
}
