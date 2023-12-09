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

func (s *Server) HandlerDetails(w http.ResponseWriter, r *http.Request) {
	// get address
	vars := mux.Vars(r)
	address := vars["address"]

	// check address
	isValid := outils.IsValidBitcoinAddress(address)
	if !isValid {
		log.Println("HandlerDetails(): invalid address")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message": "invalid address"}`)
		return
	}

	// get details
	details, err := getBitcoinAddressDetails(address)
	if err != nil {
		log.Println("HandlerDetails(): ", err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"message": "internal server error"}`)
		return
	}

	// create json data
	dataAddress := toAddressData(details)
	jsonData, err := json.Marshal(&dataAddress)
	if err != nil {
		log.Println("HandlerDetails(): ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// getBitcoinAddressDetails gets the details of a bitcoin address.
func getBitcoinAddressDetails(address string) (*types.DetailsDTO, error) {
	// create client
	client := &http.Client{}

	// set url
	url := fmt.Sprintf("https://bitcoin.explorer.klever.io/api/v2/address/%s", address)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.SetBasicAuth("support", "Fg+GJKDACKIEOD3XVps=")

	// do request
	resp, err := client.Do(req)
	if err != nil {
		return &types.DetailsDTO{}, err
	}
	defer resp.Body.Close()

	// get response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return &types.DetailsDTO{}, err
	}

	// convert binary data
	var details types.DetailsDTO
	err = json.Unmarshal(bodyBytes, &details)
	if err != nil {
		return &types.DetailsDTO{}, err
	}

	return &details, nil
}

// toAddressData returns AddressData from a DetailsDTO.
func toAddressData(details *types.DetailsDTO) types.AddressData {
	return types.AddressData{
		Address:     details.Address,
		Balance:     details.Balance,
		TotalTx:     details.TotalTx,
		BalanceInfo: types.BalanceInfo{Confirmed: details.Balance, Unconfirmed: details.BalanceUnconfirmed},
		Total:       types.TotalInfo{Sent: details.TotalSent, Received: details.TotalReceived},
	}
}
