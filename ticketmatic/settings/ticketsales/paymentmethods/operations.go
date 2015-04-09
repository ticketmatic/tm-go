package paymentmethods

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// Get a list of payment methods
func Getlist(client *ticketmatic.Client, params *ticketmatic.PaymentMethodParameters) ([]*ticketmatic.ListPaymentMethod, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/ticketsales/paymentmethods")
	r.AddParameter("includearchived", params.Includearchived)
	r.AddParameter("lastupdatesince", params.Lastupdatesince)
	r.AddParameter("filter", params.Filter)

	var obj []*ticketmatic.ListPaymentMethod
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get a single payment method
func Get(client *ticketmatic.Client, id int) (*ticketmatic.PaymentMethod, error) {
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
func Create(client *ticketmatic.Client, data *ticketmatic.CreatePaymentMethod) (*ticketmatic.PaymentMethod, error) {
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
func Update(client *ticketmatic.Client, id int, data *ticketmatic.UpdatePaymentMethod) (*ticketmatic.PaymentMethod, error) {
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
// Payment methods are archivable: this call won't actually delete the object from the database.
// Instead, it will mark the object as archived, which means it won't show up anymore in most
// places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure consistency of
// historical data.
func Delete(client *ticketmatic.Client, id int) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/ticketsales/paymentmethods/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}

// Batch modify payment methods
func Batch(client *ticketmatic.Client) error {
	r := client.NewRequest("PUT", "/{accountname}/settings/ticketsales/paymentmethods")

	return r.Run(nil)
}
