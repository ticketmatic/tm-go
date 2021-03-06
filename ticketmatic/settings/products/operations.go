package products

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.Product `json:"data"`

	// The total number of results that are available without considering limit and offset, useful for paging.
	NbrOfResults int `json:"nbrofresults"`
}

// Get a list of products
func Getlist(client *ticketmatic.Client, params *ticketmatic.ProductQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/products", "json")
	if params != nil {
		r.AddParameter("typeid", params.Typeid)
		r.AddParameter("filter", params.Filter)
		r.AddParameter("includearchived", params.Includearchived)
		r.AddParameter("lastupdatesince", params.Lastupdatesince)
	}

	var obj *List
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get a single product
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.Product, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/products/{id}", "json")
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
	r := client.NewRequest("POST", "/{accountname}/settings/products", "json")
	r.Body(data, "json")

	var obj *ticketmatic.Product
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing product
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.Product) (*ticketmatic.Product, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/products/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

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
	r := client.NewRequest("DELETE", "/{accountname}/settings/products/{id}", "json")
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
// (https://www.ticketmatic.com/docs/api/coreconcepts/translations) for more
// information.
func Translations(client *ticketmatic.Client, id int64) (map[string]string, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/products/{id}/translate", "json")
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
// (https://www.ticketmatic.com/docs/api/coreconcepts/translations) for more
// information.
func Translate(client *ticketmatic.Client, id int64, data map[string]string) (map[string]string, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/products/{id}/translate", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	var obj map[string]string
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
