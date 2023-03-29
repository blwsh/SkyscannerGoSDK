package requests

import (
	"github.com/blwsh/SkyscannerGoSDK/internal/utils/test"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/requests"
	"testing"
)

func init() {
	test.Setup()
}

func TestAutoSuggestCarhire(t *testing.T) {
	carriers, err := AutoSuggestCarhire(requests.AutoSuggestRequest{
		Query: requests.AutoSuggestQuery{
			Market:     "US",
			Locale:     "en-US",
			SearchTerm: "LAX",
		},
		Limit: 20,
	})

	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if len(carriers.Places) == 0 {
		t.Fatalf("No Places found")
	}

	t.Log(carriers)
}
