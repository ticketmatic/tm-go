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
	Contacts map[string]*ticketmatic.Contact `json:"contacts"`

	// Delivery scenarios
	Deliveryscenarios map[string]*ticketmatic.DeliveryScenario `json:"deliveryscenarios"`

	// Events
	Events map[string]*ticketmatic.Event `json:"events"`

	// Payment methods
	Paymentmethods map[string]*ticketmatic.PaymentMethod `json:"paymentmethods"`

	// Payment scenarios
	Paymentscenarios map[string]*ticketmatic.PaymentScenario `json:"paymentscenarios"`

	// Price types
	Pricetypes map[string]*ticketmatic.PriceType `json:"pricetypes"`

	// Product categories
	Productcategories map[string]*ticketmatic.ProductCategory `json:"productcategories"`

	// Products
	Products map[string]*ticketmatic.Product `json:"products"`

	// Sales channels
	Saleschannels map[string]*ticketmatic.SalesChannel `json:"saleschannels"`

	// Service charges
	Servicecharges map[string]*ticketmatic.OrderFeeDefinition `json:"servicecharges"`

	// Ticket types
	Tickettypes map[string]*ticketmatic.OrderTickettype `json:"tickettypes"`

	// Voucher codes
	Vouchercodes map[string]string `json:"vouchercodes"`
}

// Get a list of orders
func Getlist(client *ticketmatic.Client, params *ticketmatic.OrderQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/orders")
	if params != nil {
		r.AddParameter("filter", params.Filter)
		r.AddParameter("includearchived", params.Includearchived)
		r.AddParameter("lastupdatesince", params.Lastupdatesince)
		r.AddParameter("limit", params.Limit)
		r.AddParameter("offset", params.Offset)
		r.AddParameter("orderby", params.Orderby)
		r.AddParameter("output", params.Output)
		r.AddParameter("searchterm", params.Searchterm)
		r.AddParameter("simplefilter", params.Simplefilter)
	}

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
//
// Note: This method may return a 429 Rate Limit Exceeded status when there is too
// much demand. See the article about rate limiting (/TODO) for more information on
// how to handle this.
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
//
// Note: This method may return a 429 Rate Limit Exceeded status when there is too
// much demand. See the article about rate limiting (/TODO) for more information on
// how to handle this.
func Addtickets(client *ticketmatic.Client, id int64, data *ticketmatic.AddTickets) (*ticketmatic.AddItemsResult, error) {
	r := client.NewRequest("POST", "/{accountname}/orders/{id}/tickets")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.AddItemsResult
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

// Add products to order
//
// Add products to order
func Addproducts(client *ticketmatic.Client, id int64, data *ticketmatic.AddProducts) (*ticketmatic.AddItemsResult, error) {
	r := client.NewRequest("POST", "/{accountname}/orders/{id}/products")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.AddItemsResult
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify products in order
//
// Individual products can be updated. Per call you can specify any number of
// product IDs and one operation.
//
// Each operation accepts different parameters, dependent on the operation type:
//
// * Set product holders: an array of ticket holder IDs (see Contact
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Contact)), one for each
// product (productholderids). *
func Updateproducts(client *ticketmatic.Client, id int64, data *ticketmatic.UpdateProducts) (*ticketmatic.Order, error) {
	r := client.NewRequest("PUT", "/{accountname}/orders/{id}/products")
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

// Remove products from order
func Deleteproducts(client *ticketmatic.Client, id int64, data *ticketmatic.DeleteProducts) (*ticketmatic.Order, error) {
	r := client.NewRequest("DELETE", "/{accountname}/orders/{id}/products")
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

// Add payments to order
func Addpayments(client *ticketmatic.Client, id int64, data *ticketmatic.AddPayments) (*ticketmatic.Order, error) {
	r := client.NewRequest("POST", "/{accountname}/orders/{id}/payments")
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

// Add refund for payment for order
func Addrefunds(client *ticketmatic.Client, id int64, data *ticketmatic.AddRefunds) (*ticketmatic.Order, error) {
	r := client.NewRequest("POST", "/{accountname}/orders/{id}/refunds")
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

// Get the log history for an order
func Getlogs(client *ticketmatic.Client, id int64) ([]*ticketmatic.LogItem, error) {
	r := client.NewRequest("GET", "/{accountname}/orders/{id}/logs")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj []*ticketmatic.LogItem
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get the PDF for (some or all) tickets in the order. DEPRECATED: Use /{id}/pdf
// instead.
func Postticketspdf(client *ticketmatic.Client, id int64, data *ticketmatic.TicketsPdfRequest) (*ticketmatic.Url, error) {
	r := client.NewRequest("POST", "/{accountname}/orders/{id}/tickets/pdf")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.Url
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get the PDF for (some or all) tickets and/or vouchercodes in the order
func Postpdf(client *ticketmatic.Client, id int64, data *ticketmatic.TicketsPdfRequest) (*ticketmatic.Url, error) {
	r := client.NewRequest("POST", "/{accountname}/orders/{id}/pdf")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.Url
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Send the delivery e-mail for the order
func Postticketsemaildelivery(client *ticketmatic.Client, id int64, data *ticketmatic.TicketsEmaildeliveryRequest) (*ticketmatic.Order, error) {
	r := client.NewRequest("POST", "/{accountname}/orders/{id}/tickets/emaildelivery")
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

// Send the payment instruction e-mail
//
// Send the payment instruction e-mail for the order that is linked to the payment
// scenario. Will only be sent if saldo <> 0 and paymentinstruction contains a
// valid payment instruction template.
func Postticketsemailpaymentinstruction(client *ticketmatic.Client, id int64) (*ticketmatic.Order, error) {
	r := client.NewRequest("POST", "/{accountname}/orders/{id}/tickets/emailpaymentinstruction")
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

// Create a payment request for an online payment for the order
func Postpaymentrequest(client *ticketmatic.Client, id int64, data *ticketmatic.PaymentRequest) (*ticketmatic.Url, error) {
	r := client.NewRequest("POST", "/{accountname}/orders/{id}/paymentrequest")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.Url
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get the PDF for a document for the order
func Getdocument(client *ticketmatic.Client, id int64, documentid string, language string) (*ticketmatic.Url, error) {
	r := client.NewRequest("GET", "/{accountname}/orders/{id}/documents/{documentid}/{language}")
	r.UrlParameters(map[string]interface{}{
		"id":         id,
		"documentid": documentid,
		"language":   language,
	})

	var obj *ticketmatic.Url
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
