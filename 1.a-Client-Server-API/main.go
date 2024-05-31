package main

import (
	"github.com/leonardopinho/GoLang/1.a-Client-Server-API/cmd/client"
	server "github.com/leonardopinho/GoLang/1.a-Client-Server-API/cmd/server"
	"log"
)

func main() {
	// server
	go func() {

		err := server.Start()
		if err != nil {
			log.Fatal(err)
		}

	}()

	// client
	err := client.Start()
	if err != nil {
		log.Fatal(err)
	}
}
