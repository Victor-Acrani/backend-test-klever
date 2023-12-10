package types

// SendTransactionDTO gives a struct for HandlerSend response body.
type SendTransactionDTO struct {
	UTXOs []Transaction `json:"utxos"`
}
