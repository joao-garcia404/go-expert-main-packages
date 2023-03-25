package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	http.HandleFunc("/", SearchZipcodeHandler)
	http.HandleFunc("/user", SearchZipcodeHandler)
	http.ListenAndServe(":8080", nil)
}

func SearchZipcodeHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	zipcodeParam := req.URL.Query().Get("cep")

	if zipcodeParam == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	address, searchError := searchZipcode(zipcodeParam)

	if searchError != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	json.NewEncoder(res).Encode(address)
}

func searchZipcode(zipcode string) (*Address, error) {
	reqUrl := fmt.Sprintf("https://viacep.com.br/ws/%v/json/", zipcode)

	request, reqError := http.Get(reqUrl)

	if reqError != nil {
		return nil, reqError
	}

	defer request.Body.Close()

	response, resError := ioutil.ReadAll(request.Body)

	if resError != nil {
		return nil, resError
	}

	var address Address

	parseError := json.Unmarshal(response, &address)

	if parseError != nil {
		return nil, parseError
	}

	return &address, nil
}
