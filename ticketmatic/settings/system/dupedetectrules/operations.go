package dupedetectrules

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.DupeDetectRule `json:"data"`

	// The total number of results that are available without considering limit and offset, useful for paging.
	NbrOfResults int `json:"nbrofresults"`
}

// Get a list of dupe detect rules
func Getlist(client *ticketmatic.Client, params *ticketmatic.DupeDetectRuleQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/system/dupedetectrules", "json")
	if params != nil {
		r.AddParameter("filter", params.Filter)
		r.AddParameter("lastupdatesince", params.Lastupdatesince)
	}

	var obj *List
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get a single dupe detect rule
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.DupeDetectRule, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/system/dupedetectrules/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.DupeDetectRule
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new dupe detect rule
func Create(client *ticketmatic.Client, data *ticketmatic.DupeDetectRule) (*ticketmatic.DupeDetectRule, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/system/dupedetectrules", "json")
	r.Body(data, "json")

	var obj *ticketmatic.DupeDetectRule
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing dupe detect rule
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.DupeDetectRule) (*ticketmatic.DupeDetectRule, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/system/dupedetectrules/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	var obj *ticketmatic.DupeDetectRule
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a dupe detect rule
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/system/dupedetectrules/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
