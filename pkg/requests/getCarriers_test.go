package requests

import (
	"github.com/blwsh/SkyscannerGoSDK/internal/utils/test"
	"testing"
)

func init() {
	test.Setup()
}

func TestGetCarriers(t *testing.T) {
	carriers, err := GetCarriers()

	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if carriers.Status.IsFailed() {
		t.Fatalf("Failed to retrive all results")
	}

	if len(carriers.Carriers) == 0 {
		t.Fatalf("No carriers found")
	}

	t.Log(carriers)
}
