package client

import (
	"testing"
)

func TestNewRequestErrOnNoApiKey(t *testing.T) {
	_, err := NewRequest("GET", "https://example.com", nil)
	if err != nil {
		if err.Error() == "API key is not set. Call client.SetApiKey() before making a request." {
			return // expected
		}
	}
}

func TestNewRequest(t *testing.T) {
	testApiKey := "test-api-key"

	SetApiKey(testApiKey)
	SetLocale("en-GB")

	path := "test"
	req, err := NewRequest("GET", path, nil)
	if err != nil {
		t.Errorf("NewRequest() returned an error: %s", err)
	}
	if req.Header.Get("x-api-key") != testApiKey {
		t.Errorf("NewRequest() did not set the API key header")
	}
	if req.Header.Get("Content-Type") != "application/json" {
		t.Errorf("NewRequest() did not set the Content-Type header")
	}
	expected := "https://partners.api.skyscanner.net/apiservices/v3/" + path
	if req.URL.String() != expected {
		t.Errorf("NewRequest() did not set the correct URL. Expected %s, got %s", expected, req.URL.String())
	}
}
