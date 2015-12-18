package products

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.Product `json:"data"`
}

// Get a list of products
func Getlist(client *ticketmatic.Client, params *ticketmatic.ProductQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/products")
	r.AddParameter("includearchived", params.Includearchived)
	r.AddParameter("lastupdatesince", params.Lastupdatesince)
	r.AddParameter("filter", params.Filter)
	r.AddParameter("typeid", params.Typeid)

	var obj *List
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get a single product
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.Product, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/products/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.Product
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new product
func Create(client *ticketmatic.Client, data *ticketmatic.Product) (*ticketmatic.Product, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/products")
	r.Body(data)

	var obj *ticketmatic.Product
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing product
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.Product) (*ticketmatic.Product, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/products/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.Product
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a product
//
// Products are archivable: this call won't actually delete the object from the
// database. Instead, it will mark the object as archived, which means it won't
// show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/products/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
