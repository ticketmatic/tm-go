package ticketlayouttemplates

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.TicketLayoutTemplate `json:"data"`
}

// Get a list of ticket layout templates
func Getlist(client *ticketmatic.Client, params *ticketmatic.TicketLayoutTemplateQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/communicationanddesign/ticketlayouttemplates")
	r.AddParameter("includearchived", params.Includearchived)
	r.AddParameter("lastupdatesince", params.Lastupdatesince)
	r.AddParameter("filter", params.Filter)
	r.AddParameter("typeid", params.Typeid)

	var obj *List
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get a single ticket layout template
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.TicketLayoutTemplate, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/communicationanddesign/ticketlayouttemplates/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.TicketLayoutTemplate
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new ticket layout template
func Create(client *ticketmatic.Client, data *ticketmatic.TicketLayoutTemplate) (*ticketmatic.TicketLayoutTemplate, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/communicationanddesign/ticketlayouttemplates")
	r.Body(data)

	var obj *ticketmatic.TicketLayoutTemplate
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing ticket layout template
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.TicketLayoutTemplate) (*ticketmatic.TicketLayoutTemplate, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/communicationanddesign/ticketlayouttemplates/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.TicketLayoutTemplate
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a ticket layout template
//
// Ticket layout templates are archivable: this call won't actually delete the
// object from the database. Instead, it will mark the object as archived, which
// means it won't show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/communicationanddesign/ticketlayouttemplates/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
