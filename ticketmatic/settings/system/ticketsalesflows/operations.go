package ticketsalesflows

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.Ticketsalesflow `json:"data"`

	// The total number of results that are available without considering limit and offset, useful for paging.
	NbrOfResults int `json:"nbrofresults"`
}

// Get a list of ticketsalesflows
func Getlist(client *ticketmatic.Client, params *ticketmatic.TicketsalesflowQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/system/ticketsalesflows", "json")
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

// Get a single ticketsalesflow
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.Ticketsalesflow, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/system/ticketsalesflows/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.Ticketsalesflow
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new ticketsalesflow
func Create(client *ticketmatic.Client, data *ticketmatic.Ticketsalesflow) (*ticketmatic.Ticketsalesflow, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/system/ticketsalesflows", "json")
	r.Body(data, "json")

	var obj *ticketmatic.Ticketsalesflow
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing ticketsalesflow
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.Ticketsalesflow) (*ticketmatic.Ticketsalesflow, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/system/ticketsalesflows/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	var obj *ticketmatic.Ticketsalesflow
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a ticketsalesflow
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/system/ticketsalesflows/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}

// Create a flow session
func Flowsession(client *ticketmatic.Client, data *ticketmatic.Flowsession) (string, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/system/ticketsalesflows/flowsession", "json")
	r.Body(data, "json")

	var obj string
	err := r.Run(&obj)
	if err != nil {
		return "", err
	}
	return obj, nil
}

// Get info on a flow
func Flowinfo(client *ticketmatic.Client, token string) (*ticketmatic.Flowinfo, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/system/ticketsalesflows/flowinfo/{token}", "json")
	r.UrlParameters(map[string]interface{}{
		"token": token,
	})

	var obj *ticketmatic.Flowinfo
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
