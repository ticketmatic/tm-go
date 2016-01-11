package productcategories

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.Productcategory `json:"data"`
}

// Get a list of productcategories
func Getlist(client *ticketmatic.Client, params *ticketmatic.ProductcategoryQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/productcategories")
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

// Get a single productcategory
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.Productcategory, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/productcategories/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.Productcategory
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new productcategory
func Create(client *ticketmatic.Client, data *ticketmatic.Productcategory) (*ticketmatic.Productcategory, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/productcategories")
	r.Body(data)

	var obj *ticketmatic.Productcategory
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing productcategory
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.Productcategory) (*ticketmatic.Productcategory, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/productcategories/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.Productcategory
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a productcategory
//
// Productcategories are archivable: this call won't actually delete the object
// from the database. Instead, it will mark the object as archived, which means it
// won't show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/productcategories/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
