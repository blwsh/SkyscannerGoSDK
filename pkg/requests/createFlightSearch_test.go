package requests

import (
	"github.com/blwsh/SkyscannerGoSDK/internal/utils/test"
	"github.com/blwsh/SkyscannerGoSDK/test/payloads"
	"testing"
)

func init() {
	test.Setup()
}

func TestCreateFlightSearch(t *testing.T) {
	search, err := CreateFlightSearch(payloads.CreateQueryPayload)

	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if !search.Status.IsValid() {
		t.Fatalf("Status is not valid: %v", search.Status)
	}

	t.Log(search)
}
