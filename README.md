# Bitcoin Manipulation API
This project comprises an API for handling Bitcoin-related data, offering endpoints to fetch address details, balance, perform BTC transactions, and view transaction information.

## Challenge Description
The challenge involves creating endpoints that manage Bitcoin address data, including confirmed and unconfirmed balances, transaction information, and UTXO selection for sending BTC.

## Endpoints
### 1 Address Details (GET api/v2/details/{address})
- Receives a Bitcoin address and returns data related to that address.
- JSON structure of the response:
``` json
{
    "address": "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh",
    "balance": "30764266",
    "totalTx": 644,
    "balanceInfo": {
        "confirmed": "30764266",
        "unconfirmed": "0"
    },
    "total": {
        "sent": "1287005839",
        "received": "1317770105"
    }
}
```

### 2 Address Balance  (GET api/v2/balance/{address})
- Receives a Bitcoin address and computes the confirmed/unconfirmed balance based on the UTXO list for that address.
- JSON structure of the response:
``` json
{
    "confirmed": "30764266",
    "unconfirmed": "0"
}
```
### 3 Send BTC (POST api/v2/send)
- Receives a Bitcoin address and the amount of BTC to send, selecting the necessary UTXOs to perform the transaction.
``` json
{
    "address":"bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh",
    "amount":"100000"
}
```
- JSON structure of the response:
``` json
{
    "utxos": [
        {
            "txid": "4c8b6e0f67d2b7d668e8a5b69ae3ae4a092cc840cc37ae979edbbe8545728976",
            "amount": "65000"
        },
        {
            "txid": "d1d152d9980d9ddc3794b47182cf396eaac53485bdf69a56cad75342b16f5f6b",
            "amount": "10000"
        },
        {
            "txid": "a88b03fca41182422ec439ca2732db33e6030aeb5934ba9f1498e3787376e942",
            "amount": "182483"
        }
    ]
}
```

### 4 Transaction Details (GET api/v2/tx/{tx})
- Receives a transaction ID and calculates the addresses and BTC amounts received by each address in the given transaction.
- JSON structure of the response:
``` json
{
    "addresses": [
        {
            "address": "36iYTpBFVZPbcyUs8pj3BtutZXzN6HPNA6",
            "value": "623579"
        },
        {
            "address": "bc1qe29ydjtwyjdmffxg4qwtd5wfwzdxvnap989glq",
            "value": "3283266"
        },
        // ...
    ],
    "block": 675674,
    "txID": "3654d26660dcc05d4cfb25a1641a1e61f06dfeb38ee2279bdb049d018f1830ab"
}
```

### 5 Health check (GET api/v2)
- JSON structure of the response:
``` json
{
    "alive": true
}
```

## Rules and Considerations
- Validation of Bitcoin address and transaction ID is mandatory.
- Choose between HTTP or GRPC for the API implementation.
- Add an API status check endpoint (Health check).
- Unit and integration tests are essential.
- Utilize big int for value manipulations.

## Example Values
 - Address: bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh
- Transaction: 3654d26660dcc05d4cfb25a1641a1e61f06dfeb38ee2279bdb049d018f1830ab

## Blockbook Endpoints
- Address Details: https://bitcoin.blockbook.chains.klever.io/api/v2/address/{address}
- Address UTXO Details: https://bitcoin.blockbook.chains.klever.io/api/v2/utxo/{address}
- Transaction Details: https://bitcoin.blockbook.chains.klever.io/api/v2/tx/{tx}

## Running the Project
To run this project locally, follow these steps:

1. Clone this repository.
```ssh
git clone https://github.com/Victor-Acrani/backend-test-klever.git
```
2. Install the necessary dependencies.
3. Run the application.
```ssh
go run . --listenaddr :8080
```

## Authors
- This challenge was proposed by [Klever Wallet.](https://klever.io/en-us) 