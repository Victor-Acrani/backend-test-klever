package types

type Transaction struct {
	TxID   string `json:"txid"`   // unique identifier of the transaction on the blockchain.
	Amount string `json:"amount"` // UTXO associated value
}
