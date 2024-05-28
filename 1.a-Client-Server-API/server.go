package main

import (
	"encoding/json"
	"github.com/valyala/fastjson"
	"io"
	"log"
	"net/http"
)

type USDBRL struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type DollarPrice struct {
	Value string `json:"value"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /cotacao", getDollarPriceHandle)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func getDollarPriceHandle(w http.ResponseWriter, request *http.Request) {
	price, err := getUSDBRL()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if price == nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// result
	err = json.NewEncoder(w).Encode(DollarPrice{Value: price.Bid})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getUSDBRL() (*USDBRL, error) {
	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var current USDBRL
	err = parseJson(body, "USDBRL", &current)
	if err != nil {
		log.Fatal(err)
	}

	return &current, nil
}

func parseJson[T interface{}](data []byte, key string, result *T) error {
	var fastJson fastjson.Parser
	v, err := fastJson.Parse(string(data))
	if err != nil {
		log.Fatal(err)
		return err
	}

	jsonData := v.Get(key)
	if err := json.Unmarshal([]byte(jsonData.String()), &result); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
