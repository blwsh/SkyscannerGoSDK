package requests

import (
	"fmt"
	client "github.com/blwsh/SkyscannerGoSDK/pkg/client"
	pollQuery "github.com/blwsh/SkyscannerGoSDK/pkg/types/requests"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/responses/search"
)

type SessionTokenErr struct{}

func (e SessionTokenErr) Error() string {
	return "session token is empty - call CreateFlightSearch first for session token"
}

// PollLiveFlights
//
// Used to retrieve the complete list of results. This usually takes some amount of time as our backend makes calls to
// our full list of supply partners for inventories. The /poll endpoint is invoked with a sessionToken which is returned
// in the result of the /create call.
//
// @see https://developers.skyscanner.net/docs/flights-live-prices/overview#endpoints
func PollLiveFlights(request pollQuery.PollRequest) (search.PollResponse, error) {
	var pollResponse search.PollResponse

	if request.SessionToken == "" {
		return pollResponse, SessionTokenErr{}
	}

	err := client.Json("POST", fmt.Sprintf("flights/live/search/poll/%s", request.SessionToken), request, &pollResponse)

	return pollResponse, err
}
