package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type CepResponse struct {
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

type BrasilApiResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func main() {
	c1 := make(chan *BrasilApiResponse)
	c2 := make(chan *CepResponse)

	go func() {
		resp, err := BuscaBrasilAPI("40155250")
		if err != nil {
			panic(err)
		}
		c1 <- resp
	}()

	go func() {
		resp, err := BuscaViaCep("40155250")
		if err != nil {
			panic(err)
		}
		c2 <- resp
	}()

	select {
	case result := <-c1:
		fmt.Println("[BrasilAPI]", result)
	case result := <-c2:
		fmt.Println("[ViaCep]", result)
	case <-time.After(time.Second):
		fmt.Println("[ERROR] Timeout")
	}
}

func BuscaViaCep(cep string) (*CepResponse, error) {
	if len(cep) != 8 {
		return nil, errors.New("CEP deve conter exatamente 8 caracteres.")
	}

	url := "http://viacep.com.br/ws/" + cep + "/json/"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result CepResponse
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func BuscaBrasilAPI(cep string) (*BrasilApiResponse, error) {
	if len(cep) != 8 {
		return nil, errors.New("CEP deve conter exatamente 8 caracteres.")
	}

	url := "https://brasilapi.com.br/api/cep/v1/" + cep
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result BrasilApiResponse
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func convertToCepResponse(b BrasilApiResponse) (*CepResponse, error) {
	result := &CepResponse{
		Cep:        b.Cep,
		Logradouro: b.Street,
		Bairro:     b.Neighborhood,
		Localidade: b.City,
		Uf:         b.State,
	}
	return result, nil
}
