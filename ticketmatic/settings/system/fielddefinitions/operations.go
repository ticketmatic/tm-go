package fielddefinitions

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.FieldDefinition `json:"data"`

	// The total number of results that are available without considering limit and offset, useful for paging.
	NbrOfResults int `json:"nbrofresults"`
}

// Get a list of field definitions
func Getlist(client *ticketmatic.Client, params *ticketmatic.FieldDefinitionQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/system/fielddefinitions", "json")
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

// Get a single field definition
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.FieldDefinition, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/system/fielddefinitions/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.FieldDefinition
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new field definition
func Create(client *ticketmatic.Client, data *ticketmatic.FieldDefinition) (*ticketmatic.FieldDefinition, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/system/fielddefinitions", "json")
	r.Body(data, "json")

	var obj *ticketmatic.FieldDefinition
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing field definition
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.FieldDefinition) (*ticketmatic.FieldDefinition, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/system/fielddefinitions/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	var obj *ticketmatic.FieldDefinition
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a field definition
//
// Field definitions are archivable: this call won't actually delete the object
// from the database. Instead, it will mark the object as archived, which means it
// won't show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/system/fielddefinitions/{id}", "json")
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
	r := client.NewRequest("GET", "/{accountname}/settings/system/fielddefinitions/{id}/translate", "json")
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
	r := client.NewRequest("PUT", "/{accountname}/settings/system/fielddefinitions/{id}/translate", "json")
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

// Get data for field definitions
func Getdata(client *ticketmatic.Client, data *ticketmatic.FielddefinitionsDataRequest) ([]*ticketmatic.FielddefinitionsDataResult, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/system/fielddefinitions/data", "json")
	r.Body(data, "json")

	var obj []*ticketmatic.FielddefinitionsDataResult
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
