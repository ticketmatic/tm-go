package orders

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.Order `json:"data"`

	// The total number of results that are available without considering limit and offset, useful for paging.
	NbrOfResults int `json:"nbrofresults"`

	// Lookup data
	Lookups *Lookups `json:"lookup"`
}

type Lookups struct {
	// Contact details
	Contacts map[string]*ticketmatic.Contact `json:"contacts"`

	// Customfield values
	Customfieldvalues map[string]string `json:"customfieldvalues"`

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

	// Relation types
	Relationtypes map[string]string `json:"relationtypes"`

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
	r := client.NewRequest("GET", "/{accountname}/orders", "json")
	if params != nil {
		r.AddParameter("filter", params.Filter)
		r.AddParameter("includearchived", params.Includearchived)
		r.AddParameter("lastupdatesince", params.Lastupdatesince)
		r.AddParameter("limit", params.Limit)
		r.AddParameter("offset", params.Offset)
		r.AddParameter("orderby", params.Orderby)
		r.AddParameter("orderby_asc", params.OrderbyAsc)
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
	r := client.NewRequest("GET", "/{accountname}/orders/{id}", "json")
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
// (https://www.ticketmatic.com/docs/api/types/SalesChannel), which needs to be
// supplied when creating.
//
// Note: This method may return a 429 Rate Limit Exceeded status
// when there is too much demand. See the article about rate limiting
// (https://www.ticketmatic.com/docs/api/ratelimiting) for more information on how
// to handle this.
func Create(client *ticketmatic.Client, data *ticketmatic.CreateOrder) (*ticketmatic.Order, error) {
	r := client.NewRequest("POST", "/{accountname}/orders", "json")
	r.Body(data, "json")

	var obj *ticketmatic.Order
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Update an order
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.UpdateOrder) (*ticketmatic.Order, error) {
	r := client.NewRequest("PUT", "/{accountname}/orders/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	var obj *ticketmatic.Order
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Delete an order
//
// Delete an order.
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/orders/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}

// Batch operations
//
// Apply batch operations to a set of orders.
//
// The parameters required are specific to the type of operation.
//
// What will be affected?
//
// The operation will be applied to the orders with given IDs. The amount of IDs is
// limited to 1000 per call.
//
//	ids: [1, 2, 3]
//
// This will apply the operation to orders with ID 1, 2 and 3.
//
// # Batch operations
//
// The following operations are supported:
//
// * emaildelivery: Send the delivery email to a selection of orders.
//
// * pdf: Print a selection of orders.
//
// * update: Update a specific field for the
// selection of orders. See BatchOrderParameters
// (https://www.ticketmatic.com/docs/api/types/BatchOrderParameters) for more info.
func Batch(client *ticketmatic.Client, data *ticketmatic.BatchOrderOperation) error {
	r := client.NewRequest("POST", "/{accountname}/orders/batch", "json")
	r.Body(data, "json")

	return r.Run(nil)
}

// Delete orders
//
// Delete multiple orders.
func Deletebatch(client *ticketmatic.Client, data []int64) (*ticketmatic.BatchResult, error) {
	r := client.NewRequest("DELETE", "/{accountname}/orders", "json")
	r.Body(data, "json")

	var obj *ticketmatic.BatchResult
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
	r := client.NewRequest("POST", "/{accountname}/orders/{id}", "json")
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

// Split tickets and/or products from an order into a new one.
func Split(client *ticketmatic.Client, id int64, data *ticketmatic.SplitOrder) (*ticketmatic.Order, error) {
	r := client.NewRequest("POST", "/{accountname}/orders/{id}/split", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	var obj *ticketmatic.Order
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Add tickets to order
//
// When adding tickets, this is limited to 50 tickets per call.
// Note: This method may return a 429 Rate Limit Exceeded status
// when there is too much demand. See the article about rate limiting
// (https://www.ticketmatic.com/docs/api/ratelimiting) for more information on how
// to handle this.
func Addtickets(client *ticketmatic.Client, id int64, data *ticketmatic.AddTickets) (*ticketmatic.AddItemsResult, error) {
	r := client.NewRequest("POST", "/{accountname}/orders/{id}/tickets", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

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
// (https://www.ticketmatic.com/docs/api/types/Contact)), one for each ticket
// (ticketholderids).
//
// * Update price type: an array of ticket price type IDs (as can be found in the
// Event pricing (https://www.ticketmatic.com/docs/api/types/Event)), one for each
// ticket (tickettypepriceids)
//
// * Add to bundles: an array of bundle IDs, one for each ticket
//
// * Remove from bundles: none.
func Updatetickets(client *ticketmatic.Client, id int64, data *ticketmatic.UpdateTickets) (*ticketmatic.Order, error) {
	r := client.NewRequest("PUT", "/{accountname}/orders/{id}/tickets", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	var obj *ticketmatic.Order
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove tickets from order
func Deletetickets(client *ticketmatic.Client, id int64, data *ticketmatic.DeleteTickets) (*ticketmatic.Order, error) {
	r := client.NewRequest("DELETE", "/{accountname}/orders/{id}/tickets", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

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
	r := client.NewRequest("POST", "/{accountname}/orders/{id}/products", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

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
// (https://www.ticketmatic.com/docs/api/types/Contact)), one for each product
// (productholderids). *
func Updateproducts(client *ticketmatic.Client, id int64, data *ticketmatic.UpdateProducts) (*ticketmatic.Order, error) {
	r := client.NewRequest("PUT", "/{accountname}/orders/{id}/products", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	var obj *ticketmatic.Order
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove products from order
func Deleteproducts(client *ticketmatic.Client, id int64, data *ticketmatic.DeleteProducts) (*ticketmatic.Order, error) {
	r := client.NewRequest("DELETE", "/{accountname}/orders/{id}/products", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	var obj *ticketmatic.Order
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Add payments to order
func Addpayments(client *ticketmatic.Client, id int64, data *ticketmatic.AddPayments) (*ticketmatic.Order, error) {
	r := client.NewRequest("POST", "/{accountname}/orders/{id}/payments", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	var obj *ticketmatic.Order
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Add refund for payment for order
func Addrefunds(client *ticketmatic.Client, id int64, data *ticketmatic.AddRefunds) (*ticketmatic.Order, error) {
	r := client.NewRequest("POST", "/{accountname}/orders/{id}/refunds", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	var obj *ticketmatic.Order
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get the log history for an order
func Getlogs(client *ticketmatic.Client, id int64) ([]*ticketmatic.LogItem, error) {
	r := client.NewRequest("GET", "/{accountname}/orders/{id}/logs", "json")
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

// [DEPRECATED] Export tickets to PDF
//
// DEPRECATED: Use /{id}/pdf instead.
func Postticketspdf(client *ticketmatic.Client, id int64, data *ticketmatic.TicketsPdfRequest) (*ticketmatic.Url, error) {
	r := client.NewRequest("POST", "/{accountname}/orders/{id}/tickets/pdf", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	var obj *ticketmatic.Url
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Export tickets and/or vouchercodes to PDF
func Postpdf(client *ticketmatic.Client, id int64, data *ticketmatic.TicketsPdfRequest) (*ticketmatic.Url, error) {
	r := client.NewRequest("POST", "/{accountname}/orders/{id}/pdf", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	var obj *ticketmatic.Url
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Send the delivery e-mail for the order
func Postticketsemaildelivery(client *ticketmatic.Client, id int64, data *ticketmatic.TicketsEmaildeliveryRequest) (*ticketmatic.Order, error) {
	r := client.NewRequest("POST", "/{accountname}/orders/{id}/tickets/emaildelivery", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

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
	r := client.NewRequest("POST", "/{accountname}/orders/{id}/tickets/emailpaymentinstruction", "json")
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
	r := client.NewRequest("POST", "/{accountname}/orders/{id}/paymentrequest", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	var obj *ticketmatic.Url
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Cancel the outstanding payment request for the order
//
// A payment request can only be cancelled when its status is open.
func Cancelpaymentrequest(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/orders/{id}/paymentrequest", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}

// Get the PDF for a document for the order
func Getdocument(client *ticketmatic.Client, id int64, documentid string, language string) (*ticketmatic.Url, error) {
	r := client.NewRequest("GET", "/{accountname}/orders/{id}/documents/{documentid}/{language}", "json")
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

// Import historic orders
//
// Up to 100 orders can be sent per call.
//
// Many of the usual consistency checks are relaxed while importing orders.
// It is recommended that you only import orders that will not be changed anymore
// in the future.
func Import(client *ticketmatic.Client, data []*ticketmatic.ImportOrder) ([]*ticketmatic.OrderImportStatus, error) {
	r := client.NewRequest("POST", "/{accountname}/orders/import", "json")
	r.Body(data, "json")

	var obj []*ticketmatic.OrderImportStatus
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Reserve order IDs
//
// Importing orders with specified IDs is only possible when those IDs fall in the
// reserved ID range.
//
// Use this call to reserve a range of order IDs. Any unused ID lower than or equal
// to the specified ID will be reserved. New orders will receive IDs higher than
// the specified ID.
func Reserve(client *ticketmatic.Client, data *ticketmatic.OrderIdReservation) (*ticketmatic.OrderIdReservation, error) {
	r := client.NewRequest("POST", "/{accountname}/orders/import/reserve", "json")
	r.Body(data, "json")

	var obj *ticketmatic.OrderIdReservation
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Purge orders
//
// Purge all orders. This is only possible for test or staging accounts.
func Purge(client *ticketmatic.Client, params *ticketmatic.PurgeOrdersRequest) (string, error) {
	r := client.NewRequest("POST", "/{accountname}/orders/purge", "json")
	if params != nil {
		r.AddParameter("contacts", params.Contacts)
		r.AddParameter("createdsince", params.Createdsince)
		r.AddParameter("events", params.Events)
	}

	var obj string
	err := r.Run(&obj)
	if err != nil {
		return "", err
	}
	return obj, nil
}
