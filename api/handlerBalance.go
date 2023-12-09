package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"

	"github.com/Victor-Acrani/backend-test-klever/outils"
	"github.com/Victor-Acrani/backend-test-klever/types"
	"github.com/gorilla/mux"
)

func (s *Server) HandlerBalance(w http.ResponseWriter, r *http.Request) {
	// get address
	vars := mux.Vars(r)
	address := vars["address"]

	// check address
	isValid := outils.IsValidBitcoinAddress(address)
	if !isValid {
		log.Println("HandlerBalance(): invalid address")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message": "invalid address"}`)
		return
	}

	// get balance
	balance, err := getBalance(address)
	if err != nil {
		log.Println("HandlerBalance -> getBalance(): ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// convert binary data
	jsonResponse, err := json.Marshal(balance)
	if err != nil {
		log.Println("HandlerBalance(): ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func getBalance(address string) (types.BalanceInfo, error) {
	// create client
	client := &http.Client{}

	// set url
	url := fmt.Sprintf("https://bitcoin.explorer.klever.io/api/v2/utxo/%s", address)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.SetBasicAuth("support", "Fg+GJKDACKIEOD3XVps=")

	// do request
	resp, err := client.Do(req)
	if err != nil {
		return types.BalanceInfo{}, err
	}
	defer resp.Body.Close()

	// get response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.BalanceInfo{}, err
	}

	// decode response body
	var utxos []types.UTXO
	err = json.Unmarshal(body, &utxos)
	if err != nil {
		return types.BalanceInfo{}, err
	}

	// create big int values for calculating balance
	confirmedBalance := big.NewInt(0)
	unconfirmedBalance := big.NewInt(0)

	// calculate balance
	for _, utxo := range utxos {
		value := utxo.Value

		amount := new(big.Int)
		amount, ok := amount.SetString(value, 10)
		if !ok {
			return types.BalanceInfo{}, fmt.Errorf("error converting value to big.Int")
		}

		// check number of confirmations
		const minimumConfirmations = 2
		if utxo.Confirmations < minimumConfirmations {
			unconfirmedBalance.Add(unconfirmedBalance, amount)
		} else {
			confirmedBalance.Add(confirmedBalance, amount)
		}
	}

	// create balance struct
	balance := types.BalanceInfo{
		Confirmed:   confirmedBalance.String(),
		Unconfirmed: unconfirmedBalance.String(),
	}

	return balance, nil
}
