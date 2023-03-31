package requests

import (
	"github.com/blwsh/SkyscannerGoSDK/internal/utils/test"
	"testing"
)

func init() {
	test.Setup()
}

func TestGetFlightsHierarchy(t *testing.T) {
	flights, err := GetFlightsHierarchy()

	if !(flights.Status.IsValid() && flights.Status.IsComplete()) {
		t.Fatalf("Status is not valid: %v", flights.Status)
	}

	if len(flights.Places) == 0 {
		t.Fatalf("No places returned")
	}

	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}
