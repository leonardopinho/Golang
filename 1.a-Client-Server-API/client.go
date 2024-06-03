package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/leonardopinho/GoLang/1.a-Client-Server-API/pkg/db/models"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	err := getDollar()
	if err != nil {
		log.Fatal(err)
	}
}

func getDollar() error {
	url := "http://localhost:8080/cotacao"

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: %s", resp.Status)
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}

	var bid models.Bid
	if err := json.Unmarshal(body, &bid); err != nil {
		log.Fatal(err)
		return err
	}

	// save txt
	err = SaveBidLog(bid)
	if err != nil {
		return err
	}

	return nil
}

func SaveBidLog(bid models.Bid) error {
	f, err := os.Create("bid.txt")
	if err != nil {
		log.Fatal(err)
		return err
	}

	txt := fmt.Sprintf("Dólar:%s", bid.Value)
	_, err = f.Write([]byte(txt))
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Println("BID successfully saved in log.")

	return nil
}
