package types

// Transaction gives a struct for handling HandlerSend transaction data.
type Transaction struct {
	TxID   string `json:"txid"`   // unique identifier of the transaction on the blockchain.
	Amount string `json:"amount"` // UTXO associated value
}
