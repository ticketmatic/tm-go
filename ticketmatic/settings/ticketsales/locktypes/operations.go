package locktypes

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// Get a list of lock types
func Getlist(client *ticketmatic.Client, params *ticketmatic.LockTypeQuery) ([]*ticketmatic.LockType, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/ticketsales/locktypes")
	r.AddParameter("includearchived", params.Includearchived)
	r.AddParameter("lastupdatesince", params.Lastupdatesince)
	r.AddParameter("filter", params.Filter)

	var obj []*ticketmatic.LockType
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get a single lock type
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.LockType, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/ticketsales/locktypes/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.LockType
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new lock type
func Create(client *ticketmatic.Client, data *ticketmatic.LockType) (*ticketmatic.LockType, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/ticketsales/locktypes")
	r.Body(data)

	var obj *ticketmatic.LockType
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing lock type
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.LockType) (*ticketmatic.LockType, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/ticketsales/locktypes/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.LockType
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a lock type
//
// Lock types are archivable: this call won't actually delete the object from the
// database. Instead, it will mark the object as archived, which means it won't
// show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/ticketsales/locktypes/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
