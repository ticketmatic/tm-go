package ordermails

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.OrderMailTemplate `json:"data"`

	// The total number of results that are available without considering limit and offset, useful for paging.
	NbrOfResults int `json:"nbrofresults"`
}

// Get a list of order mail templates
func Getlist(client *ticketmatic.Client, params *ticketmatic.OrderMailTemplateQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/communicationanddesign/ordermails")
	if params != nil {
		r.AddParameter("includearchived", params.Includearchived)
		r.AddParameter("lastupdatesince", params.Lastupdatesince)
		r.AddParameter("filter", params.Filter)
	}

	var obj *List
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get a single order mail template
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.OrderMailTemplate, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/communicationanddesign/ordermails/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.OrderMailTemplate
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new order mail template
func Create(client *ticketmatic.Client, data *ticketmatic.OrderMailTemplate) (*ticketmatic.OrderMailTemplate, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/communicationanddesign/ordermails")
	r.Body(data)

	var obj *ticketmatic.OrderMailTemplate
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing order mail template
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.OrderMailTemplate) (*ticketmatic.OrderMailTemplate, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/communicationanddesign/ordermails/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.OrderMailTemplate
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove an order mail template
//
// Order mail templates are archivable: this call won't actually delete the object
// from the database. Instead, it will mark the object as archived, which means it
// won't show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/communicationanddesign/ordermails/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
