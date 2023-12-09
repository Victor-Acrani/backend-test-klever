package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"

	"github.com/Victor-Acrani/backend-test-klever/types"
)

func (s *Server) HandlerSend(w http.ResponseWriter, r *http.Request) {
	// read body
	var sendR types.SendRequestDTO
	err := json.NewDecoder(r.Body).Decode(&sendR)
	if err != nil {
		log.Println("HandlerSend(): ", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// validate body
	if err := sendR.Validate(); err != nil {
		log.Println("HandlerSend(): ", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get UTXOs
	utxos, err := getUTXOs(sendR.Address)
	if err != nil {
		log.Println("HandlerSend(): ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get first n necessary UTXO for transaction
	desiredAmount, _ := new(big.Int).SetString(sendR.Amount, 10)
	var selectedUTXOs []types.Transaction
	totalAmount := new(big.Int)
	var isAmountEnough bool

	for _, utxo := range utxos {
		amount, ok := new(big.Int).SetString(utxo.Value, 10)
		if !ok {
			continue
		}

		// add amount
		totalAmount.Add(totalAmount, amount)
		// convert UTXO to Transaction struct
		t := types.Transaction{
			TxID:   utxo.TxID,
			Amount: utxo.Value,
		}
		// append transaction
		selectedUTXOs = append(selectedUTXOs, t)

		if totalAmount.Cmp(desiredAmount) >= 0 {
			isAmountEnough = true
			break
		}
	}

	// check if amount is enough
	if !isAmountEnough {
		err := fmt.Errorf("address does not have enough founds")
		log.Println("HandlerSend(): ", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// create json response
	response := types.SendTransactionDTO{UTXOs: selectedUTXOs}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println("HandlerSend(): ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func getUTXOs(address string) ([]types.UTXO, error) {
	// create client
	client := &http.Client{}

	// set url
	url := fmt.Sprintf("https://bitcoin.explorer.klever.io/api/v2/utxo/%s", address)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.SetBasicAuth("support", "Fg+GJKDACKIEOD3XVps=")

	// do request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// convert binary to struct
	var utxos []types.UTXO
	err = json.Unmarshal(body, &utxos)
	if err != nil {
		return nil, err
	}

	return utxos, nil
}
