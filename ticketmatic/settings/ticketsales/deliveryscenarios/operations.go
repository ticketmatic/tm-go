package deliveryscenarios

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// Get a list of delivery scenarios
func Getlist(client *ticketmatic.Client, params *ticketmatic.DeliveryScenarioParameters) ([]*ticketmatic.ListDeliveryScenario, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/ticketsales/deliveryscenarios")
	r.AddParameter("includearchived", params.Includearchived)
	r.AddParameter("lastupdatesince", params.Lastupdatesince)
	r.AddParameter("filter", params.Filter)

	var obj []*ticketmatic.ListDeliveryScenario
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get a single delivery scenario
func Get(client *ticketmatic.Client, id int) (*ticketmatic.DeliveryScenario, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/ticketsales/deliveryscenarios/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.DeliveryScenario
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new delivery scenario
func Create(client *ticketmatic.Client, data *ticketmatic.CreateDeliveryScenario) (*ticketmatic.DeliveryScenario, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/ticketsales/deliveryscenarios")
	r.Body(data)

	var obj *ticketmatic.DeliveryScenario
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing delivery scenario
func Update(client *ticketmatic.Client, id int, data *ticketmatic.UpdateDeliveryScenario) (*ticketmatic.DeliveryScenario, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/ticketsales/deliveryscenarios/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.DeliveryScenario
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a delivery scenario
//
// Delivery scenarios are archivable: this call won't actually delete the object
// from the database. Instead, it will mark the object as archived, which means it
// won't show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/ticketsales/deliveryscenarios/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}

// Batch modify delivery scenarios
func Batch(client *ticketmatic.Client) error {
	r := client.NewRequest("PUT", "/{accountname}/settings/ticketsales/deliveryscenarios")

	return r.Run(nil)
}
