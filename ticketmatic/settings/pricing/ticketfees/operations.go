package ticketfees

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// Get a list of ticket fees
func Getlist(client *ticketmatic.Client, params *ticketmatic.TicketFeeParameters) ([]*ticketmatic.ListTicketFee, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/pricing/ticketfees")
	r.AddParameter("includearchived", params.Includearchived)
	r.AddParameter("lastupdatesince", params.Lastupdatesince)
	r.AddParameter("filter", params.Filter)

	var obj []*ticketmatic.ListTicketFee
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get a single ticket fee
func Get(client *ticketmatic.Client, id int) (*ticketmatic.TicketFee, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/pricing/ticketfees/{id}")
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
func Create(client *ticketmatic.Client, data *ticketmatic.CreateTicketFee) (*ticketmatic.TicketFee, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/pricing/ticketfees")
	r.Body(data)

	var obj *ticketmatic.TicketFee
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing ticket fee
func Update(client *ticketmatic.Client, id int, data *ticketmatic.UpdateTicketFee) (*ticketmatic.TicketFee, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/pricing/ticketfees/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

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
func Delete(client *ticketmatic.Client, id int) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/pricing/ticketfees/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
