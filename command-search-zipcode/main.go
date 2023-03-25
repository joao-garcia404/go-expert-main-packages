package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Address struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, zipcode := range os.Args[1:] {
		reqUrl := fmt.Sprintf("https://viacep.com.br/ws/%v/json", zipcode)

		request, reqError := http.Get(reqUrl)

		if reqError != nil {
			fmt.Fprintf(os.Stderr, "Error during request: %v\n", reqError)
		}

		defer request.Body.Close()

		response, resError := io.ReadAll(request.Body)

		if resError != nil {
			fmt.Fprintf(os.Stderr, "Error during response: %v\n", resError)
		}

		var address Address

		parseError := json.Unmarshal(response, &address)

		if parseError != nil {
			fmt.Fprintf(os.Stderr, "Error during json parse: %v\n", parseError)
		}

		file, fileError := os.Create("city.txt")

		if fileError != nil {
			fmt.Fprintf(os.Stderr, "Error during file creation: %v\n", fileError)
		}

		defer file.Close()

		_, writeError := file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", address.Cep, address.Localidade, address.Uf))

		if writeError != nil {
			fmt.Fprintf(os.Stderr, "Error during file write: %v\n", writeError)
		}
	}
}
