package paymentmethods

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.PaymentMethod `json:"data"`

	// The total number of results that are available without considering limit and offset, useful for paging.
	NbrOfResults int `json:"nbrofresults"`
}

// Get a list of payment methods
func Getlist(client *ticketmatic.Client, params *ticketmatic.PaymentMethodQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/ticketsales/paymentmethods")
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

// Fetch translatable fields
//
// Returns a dictionary with string values in all languages for each translatable
// field.
//
// See translations
// (https://apps.ticketmatic.com/#/knowledgebase/api/coreconcepts_translations) for
// more information.
func Translations(client *ticketmatic.Client, id int64) (map[string]string, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/ticketsales/paymentmethods/{id}/translate")
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
//
// Sets updated translation strings.
//
// See translations
// (https://apps.ticketmatic.com/#/knowledgebase/api/coreconcepts_translations) for
// more information.
func Translate(client *ticketmatic.Client, id int64, data map[string]string) (map[string]string, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/ticketsales/paymentmethods/{id}/translate")
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
