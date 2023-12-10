package types

// TxDTO gives a struct for HandlerTransaction response.
type TxDTO struct {
	Addresses []TxAddressValue `json:"addresses"`
	Block     int              `json:"block"`
	TxID      string           `json:"txID"`
}

type TxAddressValue struct {
	Address string `json:"address"`
	Value   string `json:"value"`
}
