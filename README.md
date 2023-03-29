# SkyscannerGoSDK

Unofficial Go SDK for Skyscanner API

## Installation

```bash
$ go get github.com/andrewslotin/skyscanner-go-sdk
```

## Example

```go
package main

import (
    "fmt"
    "os"
    "strconv"

    "github.com/blwsh/SkyscannerGoSDK/pkg/client"
    "github.com/blwsh/SkyscannerGoSDK/pkg/requests"
    "github.com/blwsh/SkyscannerGoSDK/pkg/types"
    "github.com/blwsh/SkyscannerGoSDK/pkg/types/flight"
    requestTypes "github.com/blwsh/SkyscannerGoSDK/pkg/types/requests"
)

func main() {
    client.SetApiKey(os.Getenv("SKYSCANNER_API_KEY"))
    client.SetLocale("en-US")

    search, err := requests.CreateFlightSearch(
        requestTypes.CreateQueryRequest{
            Query: requestTypes.CreateQuery{
                Adults:               1,
                Market:               "US",
                Currency:             "USD",
                Locale:               "en-US",
                DateTimeGroupingType: types.ByMonth,
                CabinClass:           "CABIN_CLASS_ECONOMY",
                QueryLegs: []requests.LegQuery{
                    {
                        OriginPlaceId:      requests.PlaceIdQuery{Iata: "MAN"},
                        DestinationPlaceId: requests.PlaceIdQuery{Iata: "HAM"},
                        Date: flight.Date{
                            Year:  2023,
                            Month: 12,
                            Day:   01,
                        },
                    },
                },
            },
        },
    )

    if err != nil {
        return
    }

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
}
```

### Output

```
£168.62 | Manchester @ 17:40 >-(240 mins)-> Hamburg International @ 22:40
£86.81 | Manchester @ 07:00 >-(100 mins)-> Hamburg International @ 09:40
£201.28 | Manchester @ 09:20 >-(225 mins)-> Hamburg International @ 14:05
£163.01 | Manchester @ 12:35 >-(210 mins)-> Hamburg International @ 17:05
£157.64 | Manchester @ 15:00 >-(245 mins)-> Hamburg International @ 20:05
£157.64 | Manchester @ 18:00 >-(215 mins)-> Hamburg International @ 22:35
£200.20 | Manchester @ 09:05 >-(245 mins)-> Hamburg International @ 14:10
£155.63 | Manchester @ 12:35 >-(270 mins)-> Hamburg International @ 18:05
```

https://developers.skyscanner.net/docs/getting-started/requests-and-responses
