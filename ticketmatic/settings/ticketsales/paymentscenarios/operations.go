package paymentscenarios

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// Get a list of payment scenarios
func Getlist(client *ticketmatic.Client, params *ticketmatic.PaymentScenarioQuery) ([]*ticketmatic.PaymentScenario, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/ticketsales/paymentscenarios")
	r.AddParameter("includearchived", params.Includearchived)
	r.AddParameter("lastupdatesince", params.Lastupdatesince)
	r.AddParameter("filter", params.Filter)

	var obj []*ticketmatic.PaymentScenario
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get a single payment scenario
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.PaymentScenario, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/ticketsales/paymentscenarios/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.PaymentScenario
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new payment scenario
func Create(client *ticketmatic.Client, data *ticketmatic.PaymentScenario) (*ticketmatic.PaymentScenario, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/ticketsales/paymentscenarios")
	r.Body(data)

	var obj *ticketmatic.PaymentScenario
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing payment scenario
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.PaymentScenario) (*ticketmatic.PaymentScenario, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/ticketsales/paymentscenarios/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.PaymentScenario
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a payment scenario
//
// Payment scenarios are archivable: this call won't actually delete the object
// from the database. Instead, it will mark the object as archived, which means it
// won't show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/ticketsales/paymentscenarios/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
