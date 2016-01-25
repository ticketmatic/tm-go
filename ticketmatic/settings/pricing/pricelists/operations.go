package pricelists

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.PriceList `json:"data"`
}

// Get a list of price lists
func Getlist(client *ticketmatic.Client, params *ticketmatic.PriceListQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/pricing/pricelists")
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

// Get a single price list
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.PriceList, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/pricing/pricelists/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.PriceList
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new price list
func Create(client *ticketmatic.Client, data *ticketmatic.PriceList) (*ticketmatic.PriceList, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/pricing/pricelists")
	r.Body(data)

	var obj *ticketmatic.PriceList
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing price list
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.PriceList) (*ticketmatic.PriceList, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/pricing/pricelists/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.PriceList
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a price list
//
// Price lists are archivable: this call won't actually delete the object from the
// database. Instead, it will mark the object as archived, which means it won't
// show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/pricing/pricelists/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
