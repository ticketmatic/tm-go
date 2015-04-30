package orders

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// Get a list of orders
func Getlist(client *ticketmatic.Client, params *ticketmatic.OrderQuery) ([]*ticketmatic.Order, error) {
	r := client.NewRequest("GET", "/{accountname}/orders")
	r.AddParameter("filter", params.Filter)
	r.AddParameter("includearchived", params.Includearchived)
	r.AddParameter("lastupdatesince", params.Lastupdatesince)
	r.AddParameter("limit", params.Limit)
	r.AddParameter("offset", params.Offset)
	r.AddParameter("orderby", params.Orderby)
	r.AddParameter("output", params.Output)
	r.AddParameter("searchterm", params.Searchterm)
	r.AddParameter("simplefilter", params.Simplefilter)

	var obj []*ticketmatic.Order
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get a single order
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.Order, error) {
	r := client.NewRequest("GET", "/{accountname}/orders/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.Order
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new order
//
// Creates a new empty order.
//
// Each order is linked to a sales channel
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/SalesChannel), which
// needs to be supplied when creating.
func Create(client *ticketmatic.Client, data *ticketmatic.CreateOrder) (*ticketmatic.Order, error) {
	r := client.NewRequest("POST", "/{accountname}/orders")
	r.Body(data)

	var obj *ticketmatic.Order
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Update an order
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.UpdateOrder) (*ticketmatic.Order, error) {
	r := client.NewRequest("PUT", "/{accountname}/orders/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.Order
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Confirm an order
//
// Marks the order as confirmed.
func Confirm(client *ticketmatic.Client, id int64) (*ticketmatic.Order, error) {
	r := client.NewRequest("POST", "/{accountname}/orders/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.Order
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Add tickets to order
func Addtickets(client *ticketmatic.Client, id int64, data *ticketmatic.AddTickets) (*ticketmatic.Order, error) {
	r := client.NewRequest("POST", "/{accountname}/orders/{id}/tickets")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.Order
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
