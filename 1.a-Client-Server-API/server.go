package main

import (
	"context"
	"encoding/json"
	"github.com/leonardopinho/GoLang/1.a-Client-Server-API/core/middleware"
	"github.com/leonardopinho/GoLang/1.a-Client-Server-API/db"
	"github.com/leonardopinho/GoLang/1.a-Client-Server-API/models"
	"github.com/valyala/fastjson"
	"io"
	"log"
	"net/http"
	"time"
)

type BidResponse struct {
	Value string `json:"value"`
}

func main() {
	err := db.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	// server
	mux := http.NewServeMux()
	mux.HandleFunc("GET /cotacao", getDollarPriceHandle)
	log.Fatal(http.ListenAndServe(":8080", middleware.RecoveryMiddleware(mux)))
}

func getDollarPriceHandle(w http.ResponseWriter, _ *http.Request) {
	price, err := getUSDBRL()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if price == nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	// save in database
	err = db.SaveUSDBRL(price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// result
	err = json.NewEncoder(w).Encode(BidResponse{Value: price.Bid})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getUSDBRL() (*models.USDBRL, error) {
	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	// TODO: lazy context in first time
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var current models.USDBRL
	err = parseJson(body, "USDBRL", &current)
	if err != nil {
		log.Println(err)
		return nil, err
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
