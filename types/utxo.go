package types

type UTXO struct {
	Confirmed     bool   `json:"confirmed"`     // Indicates whether this transaction has been confirmed on the blockchain.
	TxID          string `json:"txid"`          // unique identifier of the transaction on the blockchain.
	Value         string `json:"value"`         // UTXO associated value
	Confirmations int    `json:"confirmations"` // number of blocks added to the blockchain after this transaction was included.
}
