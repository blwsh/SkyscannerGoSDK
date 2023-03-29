package stats

type ItineraryStats struct {
	MinDuration int `json:"minDuration"`
	MaxDuration int `json:"maxDuration"`
	Total       struct {
		Count    int `json:"count"`
		MinPrice struct {
			Amount       string `json:"amount"`
			Unit         string `json:"unit"`
			UpdateStatus string `json:"updateStatus"`
		} `json:"minPrice"`
	} `json:"total"`
	Stops struct {
		Direct struct {
			Total struct {
				Count    int `json:"count"`
				MinPrice struct {
					Amount       string `json:"amount"`
					Unit         string `json:"unit"`
					UpdateStatus string `json:"updateStatus"`
				} `json:"minPrice"`
			} `json:"total"`
			TicketTypes struct {
				SingleTicket struct {
					Count    int `json:"count"`
					MinPrice struct {
						Amount       string `json:"amount"`
						Unit         string `json:"unit"`
						UpdateStatus string `json:"updateStatus"`
					} `json:"minPrice"`
				} `json:"singleTicket"`
				MultiTicketNonNpt struct {
					Count    int `json:"count"`
					MinPrice struct {
						Amount       string `json:"amount"`
						Unit         string `json:"unit"`
						UpdateStatus string `json:"updateStatus"`
					} `json:"minPrice"`
				} `json:"multiTicketNonNpt"`
				MultiTicketNpt struct {
					Count    int `json:"count"`
					MinPrice struct {
						Amount       string `json:"amount"`
						Unit         string `json:"unit"`
						UpdateStatus string `json:"updateStatus"`
					} `json:"minPrice"`
				} `json:"multiTicketNpt"`
			} `json:"ticketTypes"`
		} `json:"direct"`
		OneStop struct {
			Total struct {
				Count    int `json:"count"`
				MinPrice struct {
					Amount       string `json:"amount"`
					Unit         string `json:"unit"`
					UpdateStatus string `json:"updateStatus"`
				} `json:"minPrice"`
			} `json:"total"`
			TicketTypes struct {
				SingleTicket struct {
					Count    int `json:"count"`
					MinPrice struct {
						Amount       string `json:"amount"`
						Unit         string `json:"unit"`
						UpdateStatus string `json:"updateStatus"`
					} `json:"minPrice"`
				} `json:"singleTicket"`
				MultiTicketNonNpt struct {
					Count    int `json:"count"`
					MinPrice struct {
						Amount       string `json:"amount"`
						Unit         string `json:"unit"`
						UpdateStatus string `json:"updateStatus"`
					} `json:"minPrice"`
				} `json:"multiTicketNonNpt"`
				MultiTicketNpt struct {
					Count    int `json:"count"`
					MinPrice struct {
						Amount       string `json:"amount"`
						Unit         string `json:"unit"`
						UpdateStatus string `json:"updateStatus"`
					} `json:"minPrice"`
				} `json:"multiTicketNpt"`
			} `json:"ticketTypes"`
		} `json:"oneStop"`
		TwoPlusStops struct {
			Total struct {
				Count    int `json:"count"`
				MinPrice struct {
					Amount       string `json:"amount"`
					Unit         string `json:"unit"`
					UpdateStatus string `json:"updateStatus"`
				} `json:"minPrice"`
			} `json:"total"`
			TicketTypes struct {
				SingleTicket struct {
					Count    int `json:"count"`
					MinPrice struct {
						Amount       string `json:"amount"`
						Unit         string `json:"unit"`
						UpdateStatus string `json:"updateStatus"`
					} `json:"minPrice"`
				} `json:"singleTicket"`
				MultiTicketNonNpt struct {
					Count    int `json:"count"`
					MinPrice struct {
						Amount       string `json:"amount"`
						Unit         string `json:"unit"`
						UpdateStatus string `json:"updateStatus"`
					} `json:"minPrice"`
				} `json:"multiTicketNonNpt"`
				MultiTicketNpt struct {
					Count    int `json:"count"`
					MinPrice struct {
						Amount       string `json:"amount"`
						Unit         string `json:"unit"`
						UpdateStatus string `json:"updateStatus"`
					} `json:"minPrice"`
				} `json:"multiTicketNpt"`
			} `json:"ticketTypes"`
		} `json:"twoPlusStops"`
	} `json:"stops"`
	HasChangeAirportTransfer bool `json:"hasChangeAirportTransfer"`
}
