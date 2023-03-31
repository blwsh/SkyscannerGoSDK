package requests

import "testing"

func TestGetMarkets(t *testing.T) {
	markets, err := GetMarkets()

	if !(markets.Status.IsValid() && markets.Status.IsComplete()) {
		t.Fatalf("Status is not valid: %v", markets.Status)
	}

	if len(markets.Markets) == 0 {
		t.Fatalf("No markets returned")
	}

	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}
