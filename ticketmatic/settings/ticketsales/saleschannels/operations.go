package saleschannels

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// Get a list of sales channels
func Getlist(client *ticketmatic.Client, params *ticketmatic.SalesChannelQuery) ([]*ticketmatic.SalesChannel, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/ticketsales/saleschannels")
	r.AddParameter("includearchived", params.Includearchived)
	r.AddParameter("lastupdatesince", params.Lastupdatesince)
	r.AddParameter("filter", params.Filter)

	var obj []*ticketmatic.SalesChannel
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get a single sales channel
func Get(client *ticketmatic.Client, id int) (*ticketmatic.SalesChannel, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/ticketsales/saleschannels/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.SalesChannel
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new sales channel
func Create(client *ticketmatic.Client, data *ticketmatic.SalesChannel) (*ticketmatic.SalesChannel, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/ticketsales/saleschannels")
	r.Body(data)

	var obj *ticketmatic.SalesChannel
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing sales channel
func Update(client *ticketmatic.Client, id int, data *ticketmatic.SalesChannel) (*ticketmatic.SalesChannel, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/ticketsales/saleschannels/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.SalesChannel
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a sales channel
//
// Sales channels are archivable: this call won't actually delete the object from
// the database. Instead, it will mark the object as archived, which means it won't
// show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/ticketsales/saleschannels/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
