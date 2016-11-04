package events

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.Event `json:"data"`

	// The total number of results that are available without considering limit and offset, useful for paging.
	NbrOfResults int `json:"nbrofresults"`

	// Lookup data
	Lookups *Lookups `json:"lookup"`
}

type Lookups struct {
	// Event locations
	Locations map[string]*ticketmatic.EventLocation `json:"locations"`

	// Price types
	Pricetypes map[string]*ticketmatic.PriceType `json:"pricetypes"`

	// Seat ranks
	Seatranks map[string]*ticketmatic.SeatRank `json:"seatranks"`
}

// List results
type EventTicketList struct {
	// Result data
	Data []*ticketmatic.EventTicket `json:"data"`

	// The total number of results that are available without considering limit and offset, useful for paging.
	NbrOfResults int `json:"nbrofresults"`
}

// Get a list of events
func Getlist(client *ticketmatic.Client, params *ticketmatic.EventQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/events")
	if params != nil {
		r.AddParameter("context", params.Context)
		r.AddParameter("filter", params.Filter)
		r.AddParameter("lastupdatesince", params.Lastupdatesince)
		r.AddParameter("limit", params.Limit)
		r.AddParameter("offset", params.Offset)
		r.AddParameter("orderby", params.Orderby)
		r.AddParameter("output", params.Output)
		r.AddParameter("searchterm", params.Searchterm)
		r.AddParameter("simplefilter", params.Simplefilter)
	}

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

// Create a new event
func Create(client *ticketmatic.Client, data *ticketmatic.Event) (*ticketmatic.Event, error) {
	r := client.NewRequest("POST", "/{accountname}/events")
	r.Body(data)

	var obj *ticketmatic.Event
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Update an event
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.Event) (*ticketmatic.Event, error) {
	r := client.NewRequest("PUT", "/{accountname}/events/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.Event
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Delete an event
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/events/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}

// Get all tickets for an event
//
// Returns the list of all tickets that are part of this event.
func Gettickets(client *ticketmatic.Client, id int64, params *ticketmatic.EventTicketQuery) (*EventTicketList, error) {
	r := client.NewRequest("GET", "/{accountname}/events/{id}/tickets")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	if params != nil {
		r.AddParameter("limit", params.Limit)
		r.AddParameter("offset", params.Offset)
		r.AddParameter("simplefilter", params.Simplefilter)
	}

	var obj *EventTicketList
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Batch update tickets for an event
//
// Update the contents of one or more custom fields for multiple tickets in one
// call. Batch update is limited to 5000 tickets per call.
//
// Warning: Do not change the barcode of a ticket that has been delivered: existing
// printed tickets will no longer work.
func Batchupdatetickets(client *ticketmatic.Client, id int64, data []*ticketmatic.EventTicket) error {
	r := client.NewRequest("PUT", "/{accountname}/events/{id}/tickets/batch")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	return r.Run(nil)
}

// Lock a set of tickets for an event (only for events with seatingplans)
//
// Lock a set of tickets for a seating plan. The lock call is limited to 100
// tickets per call.
func Locktickets(client *ticketmatic.Client, id int64, data *ticketmatic.EventLockTickets) error {
	r := client.NewRequest("PUT", "/{accountname}/events/{id}/tickets/lock")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	return r.Run(nil)
}

// Unlock a set of tickets for an event (only for events with seatingplans)
//
// Unlock a set of tickets for an event with a seating plan. The unlock call is
// limited to 100 tickets per call.
func Unlocktickets(client *ticketmatic.Client, id int64, data *ticketmatic.EventUnlockTickets) error {
	r := client.NewRequest("PUT", "/{accountname}/events/{id}/tickets/unlock")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	return r.Run(nil)
}

// Fetch translatable fields
//
// Returns a dictionary with string values in all languages for each translatable
// field.
func Translations(client *ticketmatic.Client, id int64) (map[string]string, error) {
	r := client.NewRequest("GET", "/{accountname}/events/{id}/translate")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj map[string]string
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Update translations
func Translate(client *ticketmatic.Client, id int64, data map[string]string) (map[string]string, error) {
	r := client.NewRequest("PUT", "/{accountname}/events/{id}/translate")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj map[string]string
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
