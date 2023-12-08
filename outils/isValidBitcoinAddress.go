package outils

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

// IsValidBitcoinAddress check if a Bitcoin address exists in Mainnet or in TestNet.
func IsValidBitcoinAddress(address string) bool {
	decodedAddr, err := btcutil.DecodeAddress(address, nil)
	if err != nil {
		return false
	}
	return decodedAddr.IsForNet(&chaincfg.MainNetParams) || decodedAddr.IsForNet(&chaincfg.TestNet3Params)
}
