package utils

import (
	"encoding/json"
	"github.com/valyala/fastjson"
	"log"
)

func ParseJson[T interface{}](data []byte, key string, result *T) error {
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
