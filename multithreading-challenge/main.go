package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
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

type ResponseDto struct {
	Data  interface{}
	Error error
}

func main() {
	log.Println("> multithreading challenge")
	cep := "90010300"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	brasilApiChannel := make(chan ResponseDto)
	viaCepChannel := make(chan ResponseDto)

	go executeBrasilApi(ctx, cep, brasilApiChannel)
	go executeViaCep(ctx, cep, viaCepChannel)

	select {
	case resp := <-brasilApiChannel:
		log.Printf("brasilApi: %v\n", resp.Data)
		cancel()
	case resp := <-viaCepChannel:
		log.Printf("viaCep: %v\n", resp.Data)
		cancel()
	case <-time.After(1 * time.Second):
		log.Println("timeout")
		cancel()
	}
}

func executeBrasilApi(ctx context.Context, cep string, ch chan<- ResponseDto) {
	log.Println("calling brasil api")
	address, err := getAddressByBrasilApi(cep)
	response := ResponseDto{
		Data:  address,
		Error: err,
	}
	log.Println("brasil api called")

	select {
	case ch <- response:
		log.Println("brasil api: emit to channel")
	case <-ctx.Done():
	}
}

func executeViaCep(ctx context.Context, cep string, ch chan<- ResponseDto) {
	log.Println("calling via cep")
	address, err := getAddressByViaCep(cep)
	response := ResponseDto{
		Data:  address,
		Error: err,
	}
	log.Println("via cep called")

	select {
	case ch <- response:
		log.Println("via cep: emit to channel")
	case <-ctx.Done():
	}
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
