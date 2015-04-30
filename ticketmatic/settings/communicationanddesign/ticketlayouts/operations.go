package ticketlayouts

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// Get a list of ticket layouts
func Getlist(client *ticketmatic.Client, params *ticketmatic.TicketLayoutQuery) ([]*ticketmatic.TicketLayout, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/communicationanddesign/ticketlayouts")
	r.AddParameter("includearchived", params.Includearchived)
	r.AddParameter("lastupdatesince", params.Lastupdatesince)
	r.AddParameter("filter", params.Filter)

	var obj []*ticketmatic.TicketLayout
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get a single ticket layout
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.TicketLayout, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/communicationanddesign/ticketlayouts/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.TicketLayout
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new ticket layout
func Create(client *ticketmatic.Client, data *ticketmatic.TicketLayout) (*ticketmatic.TicketLayout, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/communicationanddesign/ticketlayouts")
	r.Body(data)

	var obj *ticketmatic.TicketLayout
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing ticket layout
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.TicketLayout) (*ticketmatic.TicketLayout, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/communicationanddesign/ticketlayouts/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.TicketLayout
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a ticket layout
//
// Ticket layouts are archivable: this call won't actually delete the object from
// the database. Instead, it will mark the object as archived, which means it won't
// show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/communicationanddesign/ticketlayouts/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}