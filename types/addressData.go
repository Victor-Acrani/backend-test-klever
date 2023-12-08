package types

// AddressData gives a struct for the data of a Bitcoin address.
type AddressData struct {
	Address     string      `json:"address"`
	Balance     string      `json:"balance,omitempty"`
	TotalTx     int         `json:"totalTx,omitempty"`
	BalanceInfo BalanceInfo `json:"balanceInfo,omitempty"`
	Total       TotalInfo   `json:"total,omitempty"`
}
