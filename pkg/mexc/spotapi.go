package mexc

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type APISpot interface {
	NewOrder(request NewOrderRequest) (*NewOrderResponse, error)
	TestNewOrder(request NewOrderRequest) (*NewOrderResponse, error)
}

type SpotClient struct {
	base   string
	key    string
	secret string
}

func NewSpotClient(base, key, secret string) *SpotClient {
	return &SpotClient{
		base:   base,
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

func postSignedRequest[T any](endpoint string, params map[string]string, key string, secret string) (*T, error) {
	var queryString string
	addAmpersand := false
	for k, v := range params {
		if k == "timestamp" {
			continue
		}

		if addAmpersand {
			queryString += "&"
		}

		queryString += fmt.Sprintf("%s=%s", k, v)
		addAmpersand = true
	}

	if timestamp, ok := params["timestamp"]; ok {
		queryString += fmt.Sprintf("&timestamp=%s", timestamp)
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

	res, err := client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("client do error %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read body error %w", err)
		}

		return nil, fmt.Errorf("request code is not ok %s (%d) error=%w", res.Status, res.StatusCode, errors.New(string(b)))
	}

	body, err := io.ReadAll(res.Body)
	log.Println(endpoint, string(body))

	var target T
	if err := json.Unmarshal(body, &target); err != nil {
		return nil, fmt.Errorf("failed to decode response %w", err)
	}

	return &target, nil
}
