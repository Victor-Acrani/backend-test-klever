package types

// TotalInfo gives a struct for the sent and received amount of an address.
type TotalInfo struct {
	Sent     string `json:"sent"`
	Received string `json:"received"`
}
