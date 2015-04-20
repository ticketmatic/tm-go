package orderfees

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// Get a list of order fees
func Getlist(client *ticketmatic.Client, params *ticketmatic.OrderFeeQuery) ([]*ticketmatic.OrderFee, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/ticketsales/orderfees")
	r.AddParameter("includearchived", params.Includearchived)
	r.AddParameter("lastupdatesince", params.Lastupdatesince)
	r.AddParameter("filter", params.Filter)

	var obj []*ticketmatic.OrderFee
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get a single order fee
func Get(client *ticketmatic.Client, id int) (*ticketmatic.OrderFee, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/ticketsales/orderfees/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.OrderFee
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new order fee
func Create(client *ticketmatic.Client, data *ticketmatic.OrderFee) (*ticketmatic.OrderFee, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/ticketsales/orderfees")
	r.Body(data)

	var obj *ticketmatic.OrderFee
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing order fee
func Update(client *ticketmatic.Client, id int, data *ticketmatic.OrderFee) (*ticketmatic.OrderFee, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/ticketsales/orderfees/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.OrderFee
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove an order fee
//
// Order fees are archivable: this call won't actually delete the object from the
// database. Instead, it will mark the object as archived, which means it won't
// show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/ticketsales/orderfees/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
