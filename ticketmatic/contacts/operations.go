package contacts

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.Contact `json:"data"`
}

// Get a list of contacts
func Getlist(client *ticketmatic.Client, params *ticketmatic.ContactQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/contacts")
	if params != nil {
		r.AddParameter("filter", params.Filter)
		r.AddParameter("includearchived", params.Includearchived)
		r.AddParameter("lastupdatesince", params.Lastupdatesince)
		r.AddParameter("limit", params.Limit)
		r.AddParameter("offset", params.Offset)
		r.AddParameter("orderby", params.Orderby)
		r.AddParameter("output", params.Output)
		r.AddParameter("searchterm", params.Searchterm)
	}

	var obj *List
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get a single contact
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.Contact, error) {
	r := client.NewRequest("GET", "/{accountname}/contacts/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.Contact
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new contact
//
// Creates a new contact
func Create(client *ticketmatic.Client, data *ticketmatic.Contact) (*ticketmatic.Contact, error) {
	r := client.NewRequest("POST", "/{accountname}/contacts")
	r.Body(data)

	var obj *ticketmatic.Contact
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Update a contact
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.Contact) (*ticketmatic.Contact, error) {
	r := client.NewRequest("PUT", "/{accountname}/contacts/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.Contact
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a contact
//
// Contacts are archivable: this call won't actually delete the object from the
// database. Instead, it will mark the contact as deleted, which means it won't
// show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/contacts/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
