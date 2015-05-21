package paymentmethods

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.PaymentMethod `json:"data"`
}

// Get a list of payment methods
func Getlist(client *ticketmatic.Client, params *ticketmatic.PaymentMethodQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/ticketsales/paymentmethods")
	r.AddParameter("includearchived", params.Includearchived)
	r.AddParameter("lastupdatesince", params.Lastupdatesince)
	r.AddParameter("filter", params.Filter)

	var obj *List
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get a single payment method
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.PaymentMethod, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/ticketsales/paymentmethods/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.PaymentMethod
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new payment method
func Create(client *ticketmatic.Client, data *ticketmatic.PaymentMethod) (*ticketmatic.PaymentMethod, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/ticketsales/paymentmethods")
	r.Body(data)

	var obj *ticketmatic.PaymentMethod
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing payment method
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.PaymentMethod) (*ticketmatic.PaymentMethod, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/ticketsales/paymentmethods/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.PaymentMethod
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a payment method
//
// Payment methods are archivable: this call won't actually delete the object from
// the database. Instead, it will mark the object as archived, which means it won't
// show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/ticketsales/paymentmethods/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
