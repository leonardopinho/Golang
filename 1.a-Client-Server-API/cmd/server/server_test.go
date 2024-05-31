package server

import (
	"testing"
)

func TestGetUSDBRL(t *testing.T) {
	_, err := GetUSDBRL()
	if err != nil {
		t.Error(err)
	}
}
