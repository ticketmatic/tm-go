package events

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.Event `json:"data"`

	// Lookup data
	Lookups *Lookups `json:"lookup"`
}

type Lookups struct {
	// Event locations
	Locations map[string]*ticketmatic.EventLocation `json:"locations"`

	// Price types
	Pricetypes map[string]*ticketmatic.PriceType `json:"pricetypes"`

	// Seat ranks
	Seatranks map[string]interface{} `json:"seatranks"`
}

// Get a list of events
func Getlist(client *ticketmatic.Client, params *ticketmatic.EventQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/events")
	r.AddParameter("filter", params.Filter)
	r.AddParameter("includearchived", params.Includearchived)
	r.AddParameter("lastupdatesince", params.Lastupdatesince)
	r.AddParameter("limit", params.Limit)
	r.AddParameter("offset", params.Offset)
	r.AddParameter("orderby", params.Orderby)
	r.AddParameter("output", params.Output)
	r.AddParameter("searchterm", params.Searchterm)
	r.AddParameter("simplefilter", params.Simplefilter)
	r.AddParameter("context", params.Context)

	var obj *List
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get a single event
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.Event, error) {
	r := client.NewRequest("GET", "/{accountname}/events/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.Event
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
