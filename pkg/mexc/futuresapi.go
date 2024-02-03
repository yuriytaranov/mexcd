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
	"net/url"
	"sort"
	"time"
)

type FuturesAPI interface {
	OrderSubmit(request OrderSubmitRequest) (*OrderSubmitResponse, error)
	HistoryPositions(request HistoryPositionsRequest) (*HistoryPositionsResponse, error)
}

type FuturesClient struct {
	base   string
	key    string
	secret string
}

func NewFuturesClient(base, key, secret string) *FuturesClient {
	return &FuturesClient{
		base:   base,
		key:    key,
		secret: secret,
	}
}

func signFuturesRequest(timestamp int64, params, key, secret string) string {
	nh := hmac.New(sha256.New, []byte(secret))

	nh.Write([]byte(fmt.Sprintf("%s%d%s", key, timestamp, params)))
	dh := nh.Sum(nil)

	return hex.EncodeToString(dh)
}

func postSignedFuturesRequest[T any](endpoint string, params map[string]string, key string, secret string) (*T, error) {
	sortKeys := make([]string, 0, len(params))
	for k := range params {
		sortKeys = append(sortKeys, k)
	}
	sort.Strings(sortKeys)
	var queryString string
	addAmpersand := false
	for _, k := range sortKeys {
		v, ok := params[k]
		if !ok {
			continue
		}

		if k == "timestamp" {
			continue
		}

		if addAmpersand {
			queryString += "&"
		}

		queryString += fmt.Sprintf("%s=%s", k, url.QueryEscape(v))
		addAmpersand = true
	}

	timestamp := time.Now().UnixMilli()

	signature := signFuturesRequest(timestamp, queryString, key, secret)

	client := http.Client{}

	r, err := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(queryString)))
	if err != nil {
		return nil, fmt.Errorf("request create error %w", err)
	}
	r.Header.Add("ApiKey", key)
	r.Header.Add("Request-Time", fmt.Sprintf("%d", timestamp))
	r.Header.Add("Signature", signature)
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
