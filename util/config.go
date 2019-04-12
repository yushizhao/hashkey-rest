package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	SecretKey  string
	ApiKeyHMAC string
	Host       string
	Symbol     string
	TimeOffset int64
}

var (
	config   Config
	authType string
	baseURL  string
)

func Init(configPath *string) error {
	jsonFile, err := os.Open(*configPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return err
	}

	authType = "HMAC"
	baseURL = config.Host + "/APITrade"

	return nil
}
