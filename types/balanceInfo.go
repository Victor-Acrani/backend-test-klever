package types

// BalanceInfo gives a struct for balance data.
type BalanceInfo struct {
	Confirmed   string `json:"confirmed"`
	Unconfirmed string `json:"unconfirmed"`
}
