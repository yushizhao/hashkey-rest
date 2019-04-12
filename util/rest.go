package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type ApiTest struct {
	ApiPath string
	Method  string
	Body    map[string]interface{}
}

// generate the message to sign
// 生成待签消息
func GenerateMessage(timestamp, method, apiPath string) []byte {
	originData := []byte(timestamp + method + apiPath)
	return originData
}

func (test ApiTest) PreSend() error {
	switch test.ApiPath {
	case "/v1/order/insert":
		if _, found := test.Body["instrumentID"]; !found {
			test.Body["instrumentID"] = config.Symbol
		}

		orderLocalID = append(orderLocalID, test.Body["orderLocalID"].(string))
	}
	return nil
}

func (test ApiTest) Send() (string, error) {

	test.PreSend()

	var body []byte
	if test.Body != nil {
		body, _ = json.Marshal(test.Body)
	} else {
		body = nil
	}

	timestamp := strconv.FormatInt(time.Now().Unix()*1000+config.TimeOffset, 10)
	hashedDataHex := GenerateMessage(timestamp, test.Method, test.ApiPath+string(body))

	log.Println(string(hashedDataHex))

	// hmac
	// 消息验证码
	hmacStr := SHA256HMAC(hashedDataHex, config.SecretKey)

	// fmt.Println(hmacStr)

	req, err := http.NewRequest(test.Method, baseURL+test.ApiPath, bytes.NewReader(body))
	if err != nil {
		return "", err
	}

	// request header
	req.Header["API-KEY"] = []string{config.ApiKeyHMAC}
	req.Header["API-SIGNATURE"] = []string{hmacStr}
	req.Header["API-TIMESTAMP"] = []string{timestamp}
	req.Header["AUTH-TYPE"] = []string{"HMAC"}
	req.Header["Content-Type"] = []string{"application/json"}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, b, "", "    ")
	if err != nil {
		return "", err
	}
	log.Println(string(prettyJSON.Bytes()))

	return string(b), nil
}
