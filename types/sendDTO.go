package types

import (
	"fmt"

	"github.com/Victor-Acrani/backend-test-klever/outils"
	"github.com/go-playground/validator/v10"
)

type SendRequestDTO struct {
	Address string `json:"address" validate:"required"`
	Amount  string `json:"amount" validate:"numeric,required"`
}

func (s *SendRequestDTO) Validate() error {
	v := validator.New()
	err := v.Struct(s)
	if err != nil {
		return err
	}

	ok := outils.IsValidBitcoinAddress(s.Address)
	if !ok {
		return fmt.Errorf("invalid address")
	}

	return nil
}
