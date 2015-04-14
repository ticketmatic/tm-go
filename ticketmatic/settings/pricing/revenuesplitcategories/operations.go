package revenuesplitcategories

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// Get a list of revenue split categories
func Getlist(client *ticketmatic.Client, params *ticketmatic.RevenueSplitCategoryParameters) ([]*ticketmatic.ListRevenueSplitCategory, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/pricing/revenuesplitcategories")
	r.AddParameter("includearchived", params.Includearchived)
	r.AddParameter("lastupdatesince", params.Lastupdatesince)
	r.AddParameter("filter", params.Filter)

	var obj []*ticketmatic.ListRevenueSplitCategory
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get a single revenue split category
func Get(client *ticketmatic.Client, id int) (*ticketmatic.RevenueSplitCategory, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/pricing/revenuesplitcategories/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.RevenueSplitCategory
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new revenue split category
func Create(client *ticketmatic.Client, data *ticketmatic.CreateRevenueSplitCategory) (*ticketmatic.RevenueSplitCategory, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/pricing/revenuesplitcategories")
	r.Body(data)

	var obj *ticketmatic.RevenueSplitCategory
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing revenue split category
func Update(client *ticketmatic.Client, id int, data *ticketmatic.UpdateRevenueSplitCategory) (*ticketmatic.RevenueSplitCategory, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/pricing/revenuesplitcategories/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.RevenueSplitCategory
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a revenue split category
//
// Revenue split categories are archivable: this call won't actually delete the
// object from the database. Instead, it will mark the object as archived, which
// means it won't show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/pricing/revenuesplitcategories/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}

// Batch modify revenue split categories
func Batch(client *ticketmatic.Client) error {
	r := client.NewRequest("PUT", "/{accountname}/settings/pricing/revenuesplitcategories")

	return r.Run(nil)
}
