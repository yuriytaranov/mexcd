package mexc

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
)

type API interface{}

type SpotClient struct {
	key    string
	secret string
}

func NewSpotClient(key, secret string) *SpotClient {
	return &SpotClient{
		key:    key,
		secret: secret,
	}
}

func signRequest(params, secret string) string {
	nh := hmac.New(sha256.New, []byte(secret))

	nh.Write([]byte(params))
	dh := nh.Sum(nil)

	return hex.EncodeToString(dh)
}

func PostRequest[T any](endpoint string, params map[string]any, key string, secret string) (*T, error) {
	var queryString string
	addAmpersand := false
	for k, v := range params {
		if addAmpersand {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", k, v)

		addAmpersand = true
	}

	signature := signRequest(queryString, secret)
	queryString += fmt.Sprintf("&signature=%s", signature)

	client := http.Client{}

	r, err := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(queryString)))
	if err != nil {
		return nil, fmt.Errorf("request create error %w", err)
	}
	r.Header.Add("X-MEXC-APIKEY", key)
	r.Header.Add("Content-Type", "application/json")

	_, err = client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("client do error %w", err)
	}
	return nil, nil
}
