package Orderfeedefinitions

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.OrderFeeDefinition `json:"data"`
}

// Get a list of order fee definitions
func Getlist(client *ticketmatic.Client, params *ticketmatic.OrderFeeDefinitionQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/pricing/Orderfeedefinitions")
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

// Get a single order fee definition
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.OrderFeeDefinition, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/pricing/Orderfeedefinitions/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.OrderFeeDefinition
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new order fee definition
func Create(client *ticketmatic.Client, data *ticketmatic.OrderFeeDefinition) (*ticketmatic.OrderFeeDefinition, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/pricing/Orderfeedefinitions")
	r.Body(data)

	var obj *ticketmatic.OrderFeeDefinition
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove an order fee definition
//
// Order fee definitions are archivable: this call won't actually delete the object
// from the database. Instead, it will mark the object as archived, which means it
// won't show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/pricing/Orderfeedefinitions/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
