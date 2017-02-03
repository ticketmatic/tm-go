package contactfields

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.ContactField `json:"data"`

	// The total number of results that are available without considering limit and offset, useful for paging.
	NbrOfResults int `json:"nbrofresults"`
}

// Get the contact fields
//
// Gets the contact fields
func Getlist(client *ticketmatic.Client) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/system/contactfields", "")

	var obj *List
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get the contact field
//
// Gets the contact field
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.ContactField, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/system/contactfields/{id}", "")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.ContactField
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get the translations for this field
//
// Get the translations for this field
func Translations(client *ticketmatic.Client, id int64) (map[string]string, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/system/contactfields/{id}/translate", "")
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

// Update the translations for this field
//
// Update the translations for this field
func Translate(client *ticketmatic.Client, id int64, data map[string]string) (map[string]string, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/system/contactfields/{id}/translate", "")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj map[string]string
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
