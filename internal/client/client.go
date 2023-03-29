package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const ApiBasePath = "https://partners.api.skyscanner.net/apiservices/v3/"

var client *http.Client
var apiKey *string
var locale *string

type InvalidRequestErr struct {
	StatusCode int
	Reason     string
}

type ErrResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details"`
}

func (e InvalidRequestErr) Error() string {
	return fmt.Sprintf("Invalid request. Status code: %d, Reason: %s", e.StatusCode, e.Reason)
}

func Client() *http.Client {
	if client == nil {
		client = &http.Client{
			Timeout: 10 * time.Second,
		}
	}

	return client
}

func SetApiKey(key string) {
	apiKey = &key
}

func SetLocale(localeString string) {
	locale = &localeString
}

func GetLocale() string {
	return *locale
}

type ErrNoApiKey struct{}

func (e ErrNoApiKey) Error() string {
	return "API key is not set. Call SetApiKey() before making a request."
}

type ErrLocaleMissing struct{}

func (e ErrLocaleMissing) Error() string {
	return "Locale is not set. Call SetLocale() before making a request."
}

func NewRequest(method, url string, body any) (*http.Request, error) {
	if apiKey == nil {
		return nil, ErrNoApiKey{}
	}

	if locale == nil {
		return nil, ErrLocaleMissing{}
	}

	// Trim the leading slash if it exists
	if url[0] == '/' {
		url = url[1:]
	}

	url = ApiBasePath + url

	var bodyReader io.Reader

	if body != nil {
		requestJson, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		bodyReader = bytes.NewReader(requestJson)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-api-key", *apiKey)
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

func Json(method, url string, body any, v any) error {
	request, err := NewRequest(method, url, body)
	if err != nil {
		return err
	}

	response, err := Client().Do(request)
	if err != nil {
		return err
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode >= 400 {
		errResponse := &ErrResponse{}

		err = json.Unmarshal(bytes, &errResponse)
		if err != nil {
			return err
		}

		return InvalidRequestErr{
			StatusCode: response.StatusCode,
			Reason:     errResponse.Message,
		}
	}

	err = json.Unmarshal(bytes, &v)
	if err != nil {
		return err
	}

	return nil
}
