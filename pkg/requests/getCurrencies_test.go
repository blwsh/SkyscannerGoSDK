package requests

import "testing"

func TestGetCurrencies(t *testing.T) {
	currencies, err := GetCurrencies()

	if !(currencies.Status.IsValid() && currencies.Status.IsComplete()) {
		t.Fatalf("Status is not valid: %v", currencies.Status)
	}

	if len(currencies.Currencies) == 0 {
		t.Fatalf("No currencies returned")
	}

	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}
