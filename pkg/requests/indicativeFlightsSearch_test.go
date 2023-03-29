package requests

import (
	"github.com/blwsh/SkyscannerGoSDK/internal/utils/test"
	"github.com/blwsh/SkyscannerGoSDK/test/payloads"
	"testing"
)

func init() {
	test.Setup()
}

func TestIndicativeFlightSearch(t *testing.T) {
	indicativeFlights, err := IndicativeFlightSearch(payloads.IndicativeSearchPayload)

	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if !indicativeFlights.Status.IsValid() {
		t.Fatalf("Status is not valid: %v", indicativeFlights.Status)
	}

	if len(indicativeFlights.Content.Results.Quotes) == 0 {
		t.Fatalf("No quotes found")
	}

	t.Log(indicativeFlights)
}
