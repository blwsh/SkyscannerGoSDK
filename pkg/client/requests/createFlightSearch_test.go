package requests

import (
	"fmt"
	"github.com/blwsh/SkyscannerGoSDK/internal/utils/test"
	"github.com/blwsh/SkyscannerGoSDK/test/payloads"
	"strconv"
	"testing"
)

func init() {
	test.Setup()
}

func TestCreateFlightSearch(t *testing.T) {
	search, err := CreateFlightSearch(payloads.CreateQueryPayload)

	for _, itinerary := range search.Content.Results.Itineraries {
		for _, option := range itinerary.PricingOptions {
			i, _ := strconv.ParseFloat(option.Price.Amount, 64)

			for _, legId := range itinerary.LegIds {
				leg := search.Content.Results.Legs[legId]

				t.Logf("£%v | %v @ %v >-(%v mins)-> %v @ %v",
					fmt.Sprintf("%.2f", i/1000),
					search.Content.Results.Places[leg.OriginPlaceId].Name,
					fmt.Sprintf("%02d:%02d", leg.DepartureDateTime.Hour, leg.DepartureDateTime.Minute),
					leg.DurationInMinutes,
					search.Content.Results.Places[leg.DestinationPlaceId].Name,
					fmt.Sprintf("%02d:%02d", leg.ArrivalDateTime.Hour, leg.ArrivalDateTime.Minute))
			}
		}
	}
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if !search.Status.IsValid() {
		t.Fatalf("Status is not valid: %v", search.Status)
	}

	t.Log(search)
}
