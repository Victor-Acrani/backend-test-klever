package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Victor-Acrani/backend-test-klever/outils"
	"github.com/Victor-Acrani/backend-test-klever/types"
	"github.com/gorilla/mux"
)

func (s *Server) HandlerTransaction(w http.ResponseWriter, r *http.Request) {
	// get tx
	vars := mux.Vars(r)
	tx := vars["tx"]

	// check transaction
	isValid := outils.IsValidBitcoinTx(tx)
	if !isValid {
		log.Println("HandlerTransaction(): invalid transaction")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message": "invalid transaction"}`)
		return
	}

	// get tx details
	txDetails, err := getTransactionDetails(tx)
	if err != nil {
		log.Println("HandlerTransaction -> getTransactionDetails(): ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// toTxDTO
	txDTO := toTxDTO(txDetails)

	// create json reponse
	jsonData, err := json.Marshal(&txDTO)
	if err != nil {
		log.Println("HandlerTransaction(): ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func getTransactionDetails(txID string) (types.TxDetails, error) {
	// create client
	client := &http.Client{}

	// set url
	url := fmt.Sprintf("https://bitcoin.explorer.klever.io/api/v2/tx/%s", txID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return types.TxDetails{}, err
	}
	req.SetBasicAuth("support", "Fg+GJKDACKIEOD3XVps=")

	// do request
	resp, err := client.Do(req)
	if err != nil {
		return types.TxDetails{}, err
	}
	defer resp.Body.Close()

	// decode binary body
	var txDetails types.TxDetails
	err = json.NewDecoder(resp.Body).Decode(&txDetails)
	if err != nil {
		return txDetails, err
	}

	return txDetails, nil
}

func toTxDTO(tx types.TxDetails) types.TxDTO {
	addressValue := make([]types.TxAddressValue, len(tx.Vout))
	for i := 0; i < len(tx.Vout); i++ {
		addressValue[i].Address = tx.Vout[i].Address[0]
		addressValue[i].Value = tx.Vout[i].Value
	}

	return types.TxDTO{
		Addresses: addressValue,
		Block:     tx.Block,
		TxID:      tx.TxID,
	}
}
