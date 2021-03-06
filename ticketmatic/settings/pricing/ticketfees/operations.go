package ticketfees

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.TicketFee `json:"data"`

	// The total number of results that are available without considering limit and offset, useful for paging.
	NbrOfResults int `json:"nbrofresults"`
}

// Get a list of ticket fees
func Getlist(client *ticketmatic.Client, params *ticketmatic.TicketFeeQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/pricing/ticketfees", "json")
	if params != nil {
		r.AddParameter("filter", params.Filter)
		r.AddParameter("includearchived", params.Includearchived)
		r.AddParameter("lastupdatesince", params.Lastupdatesince)
	}

	var obj *List
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get a single ticket fee
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.TicketFee, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/pricing/ticketfees/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.TicketFee
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new ticket fee
func Create(client *ticketmatic.Client, data *ticketmatic.TicketFee) (*ticketmatic.TicketFee, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/pricing/ticketfees", "json")
	r.Body(data, "json")

	var obj *ticketmatic.TicketFee
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing ticket fee
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.TicketFee) (*ticketmatic.TicketFee, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/pricing/ticketfees/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	var obj *ticketmatic.TicketFee
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a ticket fee
//
// Ticket fees are archivable: this call won't actually delete the object from the
// database. Instead, it will mark the object as archived, which means it won't
// show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/pricing/ticketfees/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
