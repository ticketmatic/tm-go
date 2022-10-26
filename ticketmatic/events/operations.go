package events

import (
	"io"

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

type TicketStream struct {
	stream *ticketmatic.Stream
}

func (s *TicketStream) Next() (*ticketmatic.EventTicket, error) {
	var obj *ticketmatic.EventTicket
	err := s.stream.Next(&obj)
	if err != nil {
		if err == io.EOF {
			err = nil
		}
		return nil, err
	}
	return obj, nil
}

func (s *TicketStream) Close() {
	s.stream.Close()

}

// Get a list of events
func Getlist(client *ticketmatic.Client, params *ticketmatic.EventQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/events", "json")
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
	r := client.NewRequest("GET", "/{accountname}/events/{id}", "json")
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
	r := client.NewRequest("POST", "/{accountname}/events", "json")
	r.Body(data, "json")

	var obj *ticketmatic.Event
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Update an event
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.Event) (*ticketmatic.Event, error) {
	r := client.NewRequest("PUT", "/{accountname}/events/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	var obj *ticketmatic.Event
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Batch operations
//
// Apply batch operations to a set of events.
//
// The parameters required are specific to the type of operation.
//
// What will be affected?
//
// The operation will be applied to the events with given IDs. The amount of IDs is
// limited to 1000 per call.
//
//	ids: [1, 2, 3]
//
// This will apply the operation to events with ID 1, 2 and 3.
//
// # Batch operations
//
// The following operations are supported:
//
// * duplicate: Duplicate events
//
// * publish: Publish a selection of events;
//
// * redraft: Redraft a selection of events;
//
// * close: Close a selection of events;
//
// * reopen: Re-open a selection of events;
//
// * delete: Deletes the selection of events.
//
// * update: Update a specific field for the
// selection of events. See BatchEventParameters
// (https://www.ticketmatic.com/docs/api/types/BatchEventParameters) for more info.
func Batch(client *ticketmatic.Client, data *ticketmatic.BatchEventOperation) error {
	r := client.NewRequest("POST", "/{accountname}/events/batch", "json")
	r.Body(data, "json")

	return r.Run(nil)
}

// Delete an event
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/events/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}

// Get all tickets for an event
//
// Returns the list of all tickets that are part of this event.
func Gettickets(client *ticketmatic.Client, id int64, params *ticketmatic.EventTicketQuery) (*TicketStream, error) {
	r := client.NewRequest("GET", "/{accountname}/events/{id}/tickets", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	if params != nil {
		r.AddParameter("simplefilter", params.Simplefilter)
	}

	stream, err := r.Stream()
	if err != nil {
		return nil, err
	}
	return &TicketStream{stream}, nil
}

// Batch update tickets for an event
//
// Update the contents of one or more custom fields for multiple tickets in one
// call. Batch update is limited to 5000 tickets per call.
//
// Warning: Do not change the barcode of a ticket that has been delivered: existing
// printed tickets will no longer work.
func Batchupdatetickets(client *ticketmatic.Client, id int64, data []*ticketmatic.EventTicket) error {
	r := client.NewRequest("PUT", "/{accountname}/events/{id}/tickets/batch", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	return r.Run(nil)
}

// Lock a set of tickets
//
// The lock call is limited to 100 tickets per call.
func Locktickets(client *ticketmatic.Client, id int64, data *ticketmatic.EventLockTickets) error {
	r := client.NewRequest("PUT", "/{accountname}/events/{id}/tickets/lock", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	return r.Run(nil)
}

// Unlock a set of tickets
//
// The unlock call is limited to 100 tickets per call.
func Unlocktickets(client *ticketmatic.Client, id int64, data *ticketmatic.EventUnlockTickets) error {
	r := client.NewRequest("PUT", "/{accountname}/events/{id}/tickets/unlock", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	return r.Run(nil)
}

// Update the seat rank for a set of tickets
//
// Updates the seat rank for tickets, works only for active events.
func Updateseatrankfortickets(client *ticketmatic.Client, id int64, data *ticketmatic.EventUpdateSeatRankForTickets) error {
	r := client.NewRequest("PUT", "/{accountname}/events/{id}/tickets/seatrank", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	return r.Run(nil)
}

// Scan out tickets that are scanned in
//
// Scan out tickets thar are scanned in, filter on tickettypeid if needed.
func Scanticketsout(client *ticketmatic.Client, id int64, data *ticketmatic.EventScanTicketsOut) ([]int64, error) {
	r := client.NewRequest("PUT", "/{accountname}/events/{id}/tickets/scanout", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	var obj []int64
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Fetch translatable fields
//
// Returns a dictionary with string values in all languages for each translatable
// field.
func Translations(client *ticketmatic.Client, id int64) (map[string]string, error) {
	r := client.NewRequest("GET", "/{accountname}/events/{id}/translate", "json")
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
	r := client.NewRequest("PUT", "/{accountname}/events/{id}/translate", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	var obj map[string]string
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Purge event
func Purge(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("PUT", "/{accountname}/events/{id}/purge", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}

// Save an image
func Saveimage(client *ticketmatic.Client, id int64, data string) (string, error) {
	r := client.NewRequest("POST", "/{accountname}/events/{id}/image", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	var obj string
	err := r.Run(&obj)
	if err != nil {
		return "", err
	}
	return obj, nil
}

// Delete the event image
func Deleteimage(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/events/{id}/image", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
