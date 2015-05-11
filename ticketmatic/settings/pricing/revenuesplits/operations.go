package revenuesplits

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	Data []*ticketmatic.RevenueSplit `json:"data"`
}

// Get a list of revenue splits
func Getlist(client *ticketmatic.Client, params *ticketmatic.RevenueSplitQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/pricing/revenuesplits")
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

// Get a single revenue split
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.RevenueSplit, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/pricing/revenuesplits/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.RevenueSplit
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new revenue split
func Create(client *ticketmatic.Client, data *ticketmatic.RevenueSplit) (*ticketmatic.RevenueSplit, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/pricing/revenuesplits")
	r.Body(data)

	var obj *ticketmatic.RevenueSplit
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing revenue split
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.RevenueSplit) (*ticketmatic.RevenueSplit, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/pricing/revenuesplits/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.RevenueSplit
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a revenue split
//
// Revenue splits are archivable: this call won't actually delete the object from
// the database. Instead, it will mark the object as archived, which means it won't
// show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/pricing/revenuesplits/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
