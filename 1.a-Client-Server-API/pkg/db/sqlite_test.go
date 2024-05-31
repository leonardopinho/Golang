package db

import (
	"github.com/leonardopinho/GoLang/1.a-Client-Server-API/pkg/db/models"
	"testing"
)

func TestStartDB(t *testing.T) {
	err := InitDb()
	if err != nil {
		t.Errorf("Error init db %v", err)
	}
}

func TestSaveUSDBRL(t *testing.T) {
	err := InitDb()
	if err != nil {
		t.Errorf("Error init db %v", err)
	}

	data := models.USDBRL{
		Code:       "USD",
		Codein:     "BRL",
		Name:       "Dólar Americano/Real Brasileiro",
		High:       "5.1684",
		Low:        "5.1603",
		VarBid:     "-0.0004",
		PctChange:  "-0.01",
		Bid:        "5.1602",
		Ask:        "5.1612",
		Timestamp:  "1716935404",
		CreateDate: "2024-05-28 19:30:04",
	}

	reg, err := SaveUSDBRL(&data)
	if err != nil {
		t.Errorf("Error saving USDBRL %v", err)
	}

	if reg != nil && reg.Id != 0 {
		_, err := DeleteUSDBRL(reg)
		if err != nil {
			t.Errorf("Error deleting USDBRL %v", err)
		}
	}
}
