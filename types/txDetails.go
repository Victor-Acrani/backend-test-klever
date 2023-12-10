package types

// TxDetails gives a struct for receving data from BlockBook API.
type TxDetails struct {
	Vout  []Vout `json:"vout"` // transaction outputs
	Block int    `json:"blockHeight"`
	TxID  string `json:"txID"`
}

type Vout struct {
	Address []string `json:"addresses"`
	Value   string   `json:"value"`
}
