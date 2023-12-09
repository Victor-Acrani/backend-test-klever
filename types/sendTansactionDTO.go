package types

type SendTransactionDTO struct {
	UTXOs []Transaction `json:"utxos"`
}
