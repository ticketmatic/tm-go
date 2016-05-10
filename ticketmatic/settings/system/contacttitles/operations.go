package contacttitles

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.ContactTitle `json:"data"`

	// The total number of results that are available without considering limit and offset, useful for paging.
	NbrOfResults int `json:"nbrofresults"`
}

// Get a list of contact titles
func Getlist(client *ticketmatic.Client, params *ticketmatic.ContactTitleQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/system/contacttitles")
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

// Get a single contact title
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.ContactTitle, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/system/contacttitles/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.ContactTitle
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new contact title
func Create(client *ticketmatic.Client, data *ticketmatic.ContactTitle) (*ticketmatic.ContactTitle, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/system/contacttitles")
	r.Body(data)

	var obj *ticketmatic.ContactTitle
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing contact title
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.ContactTitle) (*ticketmatic.ContactTitle, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/system/contacttitles/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.ContactTitle
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a contact title
//
// Contact titles are archivable: this call won't actually delete the object from
// the database. Instead, it will mark the object as archived, which means it won't
// show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/system/contacttitles/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
