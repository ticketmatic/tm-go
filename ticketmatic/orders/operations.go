package orders

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.Order `json:"data"`

	// Lookup data
	Lookups *Lookups `json:"lookup"`
}

type Lookups struct {
	// Contact details
	Contacts []*ticketmatic.Contact `json:"contacts"`

	// Delivery scenarios
	Deliveryscenarios []*ticketmatic.DeliveryScenario `json:"deliveryscenarios"`

	// events
	Events []*ticketmatic.Event `json:"events"`

	// Payment methods
	Paymentmethods []*ticketmatic.PaymentMethod `json:"paymentmethods"`

	// Payment scenarios
	Paymentscenarios []*ticketmatic.PaymentScenario `json:"paymentscenarios"`

	// Price types
	Pricetypes []*ticketmatic.PriceType `json:"pricetypes"`

	// Sales channels
	Saleschannels []*ticketmatic.SalesChannel `json:"saleschannels"`

	// Service charges
	Servicecharges []interface{} `json:"servicecharges"`

	// Ticket types
	Tickettypes []interface{} `json:"tickettypes"`
}

// Get a list of orders
func Getlist(client *ticketmatic.Client, params *ticketmatic.OrderQuery) (*List, error) {
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

	var obj *List
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
func Addtickets(client *ticketmatic.Client, id int64, data *ticketmatic.AddTickets) (*ticketmatic.AddTicketsResult, error) {
	r := client.NewRequest("POST", "/{accountname}/orders/{id}/tickets")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.AddTicketsResult
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify tickets in order
//
// Individual tickets can be updated. Per call you can specify any number of ticket
// IDs and one operation.
//
// Each operation accepts different parameters, dependent on the operation type:
//
// * Set ticket holders: an array of ticket holder IDs (see Contact
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Contact)), one for each
// ticket (ticketholderids).
//
// * Update price type: an array of ticket price type IDs (as can be found in the
// Event pricing (https://apps.ticketmatic.com/#/knowledgebase/api/types/Event)),
// one for each ticket (tickettypepriceids).
func Updatetickets(client *ticketmatic.Client, id int64, data *ticketmatic.UpdateTickets) (*ticketmatic.Order, error) {
	r := client.NewRequest("PUT", "/{accountname}/orders/{id}/tickets")
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

// Remove tickets from order
func Deletetickets(client *ticketmatic.Client, id int64, data *ticketmatic.DeleteTickets) (*ticketmatic.Order, error) {
	r := client.NewRequest("DELETE", "/{accountname}/orders/{id}/tickets")
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
