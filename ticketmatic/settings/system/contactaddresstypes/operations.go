package contactaddresstypes

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.ContactAddressType `json:"data"`
}

// Get a list of contact address types
func Getlist(client *ticketmatic.Client, params *ticketmatic.ContactAddressTypeQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/system/contactaddresstypes")
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

// Get a single contact address type
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.ContactAddressType, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/system/contactaddresstypes/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.ContactAddressType
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new contact address type
func Create(client *ticketmatic.Client, data *ticketmatic.ContactAddressType) (*ticketmatic.ContactAddressType, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/system/contactaddresstypes")
	r.Body(data)

	var obj *ticketmatic.ContactAddressType
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing contact address type
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.ContactAddressType) (*ticketmatic.ContactAddressType, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/system/contactaddresstypes/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.ContactAddressType
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a contact address type
//
// Contact address types are archivable: this call won't actually delete the object
// from the database. Instead, it will mark the object as archived, which means it
// won't show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/system/contactaddresstypes/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}

// Fetch translatable fields
//
// Returns a dictionary with string values in all languages for each translatable
// field.
func Translations(client *ticketmatic.Client, id int64) (map[string]string, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/system/contactaddresstypes/{id}/translate")
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
	r := client.NewRequest("PUT", "/{accountname}/settings/system/contactaddresstypes/{id}/translate")
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
