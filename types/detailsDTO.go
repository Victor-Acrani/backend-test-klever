package types

// DetailsDTO gives a struct for receiving data from Blockbook API.
type DetailsDTO struct {
	Address            string `json:"address"`
	Balance            string `json:"balance,omitempty"`
	BalanceUnconfirmed string `json:"unconfirmedBalance"`
	TotalTx            int    `json:"txs"`
	TotalReceived      string `json:"totalReceived"`
	TotalSent          string `json:"totalSent"`
}
