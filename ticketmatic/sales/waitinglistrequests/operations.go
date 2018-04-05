package waitinglistrequests

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.WaitingListRequest `json:"data"`

	// The total number of results that are available without considering limit and offset, useful for paging.
	NbrOfResults int `json:"nbrofresults"`
}

// Get a list of waiting list requests
func Getlist(client *ticketmatic.Client, params *ticketmatic.WaitingListRequestQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/sales/waitinglistrequests", "json")
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

// Get a single waiting list request
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.WaitingListRequest, error) {
	r := client.NewRequest("GET", "/{accountname}/sales/waitinglistrequests/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.WaitingListRequest
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new waiting list request
func Create(client *ticketmatic.Client, data *ticketmatic.WaitingListRequest) (*ticketmatic.WaitingListRequest, error) {
	r := client.NewRequest("POST", "/{accountname}/sales/waitinglistrequests", "json")
	r.Body(data, "json")

	var obj *ticketmatic.WaitingListRequest
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing waiting list request
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.WaitingListRequest) (*ticketmatic.WaitingListRequest, error) {
	r := client.NewRequest("PUT", "/{accountname}/sales/waitinglistrequests/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	var obj *ticketmatic.WaitingListRequest
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a waiting list request
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/sales/waitinglistrequests/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
