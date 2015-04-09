package webskins

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// Get a list of web sales skins
func Getlist(client *ticketmatic.Client, params *ticketmatic.WebSalesSkinParameters) ([]*ticketmatic.ListWebSalesSkin, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/communicationanddesign/webskins")
	r.AddParameter("lastupdatesince", params.Lastupdatesince)
	r.AddParameter("filter", params.Filter)

	var obj []*ticketmatic.ListWebSalesSkin
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get a single web sales skin
func Get(client *ticketmatic.Client, id int) (*ticketmatic.WebSalesSkin, error) {
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
func Create(client *ticketmatic.Client, data *ticketmatic.CreateWebSalesSkin) (*ticketmatic.WebSalesSkin, error) {
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
func Update(client *ticketmatic.Client, id int, data *ticketmatic.UpdateWebSalesSkin) (*ticketmatic.WebSalesSkin, error) {
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
func Delete(client *ticketmatic.Client, id int) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/communicationanddesign/webskins/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}

// Batch modify web sales skins
func Batch(client *ticketmatic.Client) error {
	r := client.NewRequest("PUT", "/{accountname}/settings/communicationanddesign/webskins")

	return r.Run(nil)
}
