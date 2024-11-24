package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type BrasilApi struct {
	Cep          string `json:"cep"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	State        string `json:"state"`
	Street       string `json:"street"`
}

type ViaCep struct {
	Cep          string `json:"cep"`
	City         string `json:"localidade"`
	Complement   string `json:"complemento"`
	Neighborhood string `json:"bairro"`
	State        string `json:"estado"`
	StateSymbol  string `json:"uf"`
	Street       string `json:"logradouro"`
}

func main() {
	log.Println("> multithreading challenge")
	cep := "91540000"
	a1, _ := getAddressByBrasilApi(cep)
	fmt.Println(a1)

	a2, _ := getAddressByViaCep(cep)
	fmt.Println(a2)

}

func getAddressByBrasilApi(cep string) (*BrasilApi, error) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var schemaResponse BrasilApi
	err = json.Unmarshal(body, &schemaResponse)
	if err != nil {
		return nil, err
	}

	return &schemaResponse, nil
}

func getAddressByViaCep(cep string) (*ViaCep, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json", cep)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var schemaResponse ViaCep
	err = json.Unmarshal(body, &schemaResponse)
	if err != nil {
		return nil, err
	}

	return &schemaResponse, nil
}
