package requests

import (
	"github.com/blwsh/SkyscannerGoSDK/internal/utils/test"
	searchQuery "github.com/blwsh/SkyscannerGoSDK/pkg/types/requests"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/responses/search"
	"github.com/blwsh/SkyscannerGoSDK/test/payloads"
	"testing"
	"time"
)

func init() {
	test.Setup()
}

func TestPollFlights(t *testing.T) {
	test.Setup()

	createResponse, err := CreateFlightSearch(payloads.CreateQueryPayload)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	t.Logf("Session token: %s", createResponse.SessionToken)

	pollResponse, err := search.PollResponse{}, nil

	for !pollResponse.Status.IsComplete() {
		pollResponse, err = PollLiveFlights(searchQuery.PollRequest{SessionToken: createResponse.SessionToken})

		t.Logf("Status: %s", pollResponse.Status)

		if err != nil {
			t.Fatalf("Error: %v", err)
		}

		if !pollResponse.Status.IsValid() {
			t.Fatalf("Status is not valid: %v", pollResponse.Status)
		}

		if pollResponse.Status.IsFailed() {
			t.Fatalf("Failed to retrive all results")
		}

		if pollResponse.Status.IsComplete() {
			break
		}

		time.Sleep(1 * time.Second)
	}

	t.Log(pollResponse)
}
