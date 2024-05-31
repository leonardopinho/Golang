package test

import (
	"github.com/leonardopinho/GoLang/1.a-Client-Server-API/cmd/client"
	"github.com/leonardopinho/GoLang/1.a-Client-Server-API/cmd/server"
	"testing"
)

func TestChallenge(t *testing.T) {
	// server
	go func() {

		err := server.Start()
		if err != nil {
			t.Error(err)
		}

	}()

	// client
	err := client.Start()
	if err != nil {
		t.Error(err)
	}
}
