package filterdefinitions

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.FilterDefinition `json:"data"`
}

// Get a list of filter definitions
func Getlist(client *ticketmatic.Client, params *ticketmatic.FilterDefinitionQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/system/filterdefinitions")
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

// Get a single filter definition
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.FilterDefinition, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/system/filterdefinitions/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.FilterDefinition
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new filter definition
func Create(client *ticketmatic.Client, data *ticketmatic.FilterDefinition) (*ticketmatic.FilterDefinition, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/system/filterdefinitions")
	r.Body(data)

	var obj *ticketmatic.FilterDefinition
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing filter definition
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.FilterDefinition) (*ticketmatic.FilterDefinition, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/system/filterdefinitions/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.FilterDefinition
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a filter definition
//
// Filter definitions are archivable: this call won't actually delete the object
// from the database. Instead, it will mark the object as archived, which means it
// won't show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/system/filterdefinitions/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
