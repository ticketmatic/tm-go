package phonenumbertypes

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.PhoneNumberType `json:"data"`
}

// Get a list of phone number types
func Getlist(client *ticketmatic.Client, params *ticketmatic.PhoneNumberTypeQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/system/phonenumbertypes")
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

// Get a single phone number type
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.PhoneNumberType, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/system/phonenumbertypes/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.PhoneNumberType
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new phone number type
func Create(client *ticketmatic.Client, data *ticketmatic.PhoneNumberType) (*ticketmatic.PhoneNumberType, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/system/phonenumbertypes")
	r.Body(data)

	var obj *ticketmatic.PhoneNumberType
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing phone number type
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.PhoneNumberType) (*ticketmatic.PhoneNumberType, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/system/phonenumbertypes/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.PhoneNumberType
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a phone number type
//
// Phone number types are archivable: this call won't actually delete the object
// from the database. Instead, it will mark the object as archived, which means it
// won't show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/system/phonenumbertypes/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
