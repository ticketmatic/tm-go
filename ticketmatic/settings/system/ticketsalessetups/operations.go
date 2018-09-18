package ticketsalessetups

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.Ticketsalessetup `json:"data"`

	// The total number of results that are available without considering limit and offset, useful for paging.
	NbrOfResults int `json:"nbrofresults"`
}

// Get a list of ticketsalessetups
func Getlist(client *ticketmatic.Client, params *ticketmatic.TicketsalessetupQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/system/ticketsalessetups", "json")
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

// Get a single ticketsalessetup
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.Ticketsalessetup, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/system/ticketsalessetups/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.Ticketsalessetup
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new ticketsalessetup
func Create(client *ticketmatic.Client, data *ticketmatic.Ticketsalessetup) (*ticketmatic.Ticketsalessetup, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/system/ticketsalessetups", "json")
	r.Body(data, "json")

	var obj *ticketmatic.Ticketsalessetup
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing ticketsalessetup
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.Ticketsalessetup) (*ticketmatic.Ticketsalessetup, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/system/ticketsalessetups/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	var obj *ticketmatic.Ticketsalessetup
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a ticketsalessetup
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/system/ticketsalessetups/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
