package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number  int `json:"n"`
	Balance int `json:"s"`
}

func main() {
	account := Account{Number: 1, Balance: 1000}

	res, err := json.Marshal(account)

	if err != nil {
		panic(err)
	}

	println(string(res))

	encodeError := json.NewEncoder(os.Stdout).Encode(account)

	if encodeError != nil {
		panic(encodeError)
	}

	pureJson := []byte(`{"n": 1, "s": 10000}`)
	var accountX Account

	unmarshalError := json.Unmarshal(pureJson, &accountX)

	if unmarshalError != nil {
		panic(unmarshalError)
	}

	println(accountX.Balance)
}
