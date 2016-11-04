package webskins

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.WebSalesSkin `json:"data"`

	// The total number of results that are available without considering limit and offset, useful for paging.
	NbrOfResults int `json:"nbrofresults"`
}

// Get a list of web sales skins
func Getlist(client *ticketmatic.Client, params *ticketmatic.WebSalesSkinQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/communicationanddesign/webskins")
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

// Get a single web sales skin
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.WebSalesSkin, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/communicationanddesign/webskins/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.WebSalesSkin
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new web sales skin
func Create(client *ticketmatic.Client, data *ticketmatic.WebSalesSkin) (*ticketmatic.WebSalesSkin, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/communicationanddesign/webskins")
	r.Body(data)

	var obj *ticketmatic.WebSalesSkin
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing web sales skin
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.WebSalesSkin) (*ticketmatic.WebSalesSkin, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/communicationanddesign/webskins/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.WebSalesSkin
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a web sales skin
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/communicationanddesign/webskins/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
