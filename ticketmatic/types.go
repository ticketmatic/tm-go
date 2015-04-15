package ticketmatic

type DeliveryscenarioAvailability struct {
	// An array of sales channel IDs for which this delivery scenario can be used.
	Saleschannels []int `json:"saleschannels,omitempty"`

	// Use a script to refine the set of sales channels?
	Usescript bool `json:"usescript,omitempty"`

	// Script used to determine availability of the delivery scenario. More info on the
	// delivery scenario overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_deliveryscenarios)
	// page.
	Script string `json:"script,omitempty"`
}

type OrderfeeRule struct {
}

type PaymentmethodConfig struct {
}

type PaymentscenarioAvailability struct {
}

type PaymentscenarioExpiryParameters struct {
}

type PaymentscenarioOverdueParameters struct {
}

type PriceAvailabilityRules struct {
}

type PricelistPrices struct {
}

type RevenuesplitRules struct {
}

type TicketfeeRules struct {
}

type WebskinConfiguration struct {
}

// A timestamp returned by the diagnostic /time call
// (https://apps.ticketmatic.com/#/knowledgebase/api/diagnostics/time).
type Timestamp struct {
	// Current system time
	Systemtime Time `json:"systemtime,omitempty"`
}

type Event struct {
}

type ListEvent struct {
}

type EventParameters struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`
}

type Ticket struct {
}

type Order struct {
}

// Set of parameters used to filter order mail templates. More info: see the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ordermails/getlist).
type OrderMailTemplateParameters struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
}

// An item in a list of order mail templates. This differs from the normal
// OrderMailTemplate type: not all fields are present in the list. More info: see
// the getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ordermails/getlist).
type ListOrderMailTemplate struct {
	// Unique ID
	Id     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Typeid int    `json:"typeid,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// A single order mail template. More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ordermails/get).
type OrderMailTemplate struct {
	// Unique ID
	Id           int               `json:"id,omitempty"`
	Name         string            `json:"name,omitempty"`
	Typeid       int               `json:"typeid,omitempty"`
	Subject      string            `json:"subject,omitempty"`
	Body         string            `json:"body,omitempty"`
	Translations map[string]string `json:"translations,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// Convert OrderMailTemplate to UpdateOrderMailTemplate
func (o *OrderMailTemplate) Update() *UpdateOrderMailTemplate {
	return &UpdateOrderMailTemplate{
		Name:         o.Name,
		Typeid:       o.Typeid,
		Subject:      o.Subject,
		Body:         o.Body,
		Translations: o.Translations,
	}
}

// A set of fields to create a order mail template. More info: see the create
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ordermails/create).
type CreateOrderMailTemplate struct {
	Name         string            `json:"name,omitempty"`
	Typeid       int               `json:"typeid,omitempty"`
	Subject      string            `json:"subject,omitempty"`
	Body         string            `json:"body,omitempty"`
	Translations map[string]string `json:"translations,omitempty"`
}

// A set of fields to update a order mail template. More info: see the update
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ordermails/update).
type UpdateOrderMailTemplate struct {
	Name         string            `json:"name,omitempty"`
	Typeid       int               `json:"typeid,omitempty"`
	Subject      string            `json:"subject,omitempty"`
	Body         string            `json:"body,omitempty"`
	Translations map[string]string `json:"translations,omitempty"`
}

// Set of parameters used to filter web sales skins. More info: see the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_webskins/getlist).
type WebSalesSkinParameters struct {
	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
}

// An item in a list of web sales skins. This differs from the normal WebSalesSkin
// type: not all fields are present in the list. More info: see the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_webskins/getlist).
type ListWebSalesSkin struct {
	// Unique ID
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`
}

// A single web sales skin. More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_webskins/get).
type WebSalesSkin struct {
	// Unique ID
	Id            int                   `json:"id,omitempty"`
	Name          string                `json:"name,omitempty"`
	Html          string                `json:"html,omitempty"`
	Css           string                `json:"css,omitempty"`
	Translations  map[string]string     `json:"translations,omitempty"`
	Configuration *WebskinConfiguration `json:"configuration,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`
}

// Convert WebSalesSkin to UpdateWebSalesSkin
func (o *WebSalesSkin) Update() *UpdateWebSalesSkin {
	return &UpdateWebSalesSkin{
		Name:          o.Name,
		Html:          o.Html,
		Css:           o.Css,
		Translations:  o.Translations,
		Configuration: o.Configuration,
	}
}

// A set of fields to create a web sales skin. More info: see the create operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_webskins/create).
type CreateWebSalesSkin struct {
	Name          string                `json:"name,omitempty"`
	Html          string                `json:"html,omitempty"`
	Css           string                `json:"css,omitempty"`
	Translations  map[string]string     `json:"translations,omitempty"`
	Configuration *WebskinConfiguration `json:"configuration,omitempty"`
}

// A set of fields to update a web sales skin. More info: see the update operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_webskins/update).
type UpdateWebSalesSkin struct {
	Name          string                `json:"name,omitempty"`
	Html          string                `json:"html,omitempty"`
	Css           string                `json:"css,omitempty"`
	Translations  map[string]string     `json:"translations,omitempty"`
	Configuration *WebskinConfiguration `json:"configuration,omitempty"`
}

// Set of parameters used to filter event locations. More info: see the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_events_eventlocations/getlist).
type EventLocationParameters struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
}

// An item in a list of event locations. This differs from the normal EventLocation
// type: not all fields are present in the list. More info: see the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_events_eventlocations/getlist).
type ListEventLocation struct {
	// Unique ID
	Id int `json:"id,omitempty"`

	// Name of the location
	Name string `json:"name,omitempty"`

	// Street name
	Street1 string `json:"street1,omitempty"`

	// Nr. + Box
	Street2 string `json:"street2,omitempty"`

	// Zipcode
	Zip string `json:"zip,omitempty"`

	// City
	City string `json:"city,omitempty"`

	// Country code. Should be an ISO 3166-1 alpha-2
	// (http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) two-letter code.
	Countrycode string `json:"countrycode,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// A single event location. More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_events_eventlocations/get).
type EventLocation struct {
	// Unique ID
	Id int `json:"id,omitempty"`

	// Name of the location
	Name string `json:"name,omitempty"`

	// Street name
	Street1 string `json:"street1,omitempty"`

	// Nr. + Box
	Street2 string `json:"street2,omitempty"`

	// Zipcode
	Zip string `json:"zip,omitempty"`

	// City
	City string `json:"city,omitempty"`

	// Country code. Should be an ISO 3166-1 alpha-2
	// (http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) two-letter code.
	Countrycode string `json:"countrycode,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// Convert EventLocation to UpdateEventLocation
func (o *EventLocation) Update() *UpdateEventLocation {
	return &UpdateEventLocation{
		Name:        o.Name,
		Street1:     o.Street1,
		Street2:     o.Street2,
		Zip:         o.Zip,
		City:        o.City,
		Countrycode: o.Countrycode,
	}
}

// A set of fields to create a event location. More info: see the create operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_events_eventlocations/create).
type CreateEventLocation struct {
	// Name of the location
	Name string `json:"name,omitempty"`

	// Street name
	Street1 string `json:"street1,omitempty"`

	// Nr. + Box
	Street2 string `json:"street2,omitempty"`

	// Zipcode
	Zip string `json:"zip,omitempty"`

	// City
	City string `json:"city,omitempty"`

	// Country code. Should be an ISO 3166-1 alpha-2
	// (http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) two-letter code.
	Countrycode string `json:"countrycode,omitempty"`
}

// A set of fields to update a event location. More info: see the update operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_events_eventlocations/update).
type UpdateEventLocation struct {
	// Name of the location
	Name string `json:"name,omitempty"`

	// Street name
	Street1 string `json:"street1,omitempty"`

	// Nr. + Box
	Street2 string `json:"street2,omitempty"`

	// Zipcode
	Zip string `json:"zip,omitempty"`

	// City
	City string `json:"city,omitempty"`

	// Country code. Should be an ISO 3166-1 alpha-2
	// (http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) two-letter code.
	Countrycode string `json:"countrycode,omitempty"`
}

// Set of parameters used to filter price availabilities. More info: see the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_priceavailabilities/getlist).
type PriceAvailabilityParameters struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
}

// An item in a list of price availabilities. This differs from the normal
// PriceAvailability type: not all fields are present in the list. More info: see
// the getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_priceavailabilities/getlist).
type ListPriceAvailability struct {
	// Unique ID
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// A single price availability. More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_priceavailabilities/get).
type PriceAvailability struct {
	// Unique ID
	Id    int                     `json:"id,omitempty"`
	Name  string                  `json:"name,omitempty"`
	Rules *PriceAvailabilityRules `json:"rules,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// Convert PriceAvailability to UpdatePriceAvailability
func (o *PriceAvailability) Update() *UpdatePriceAvailability {
	return &UpdatePriceAvailability{
		Name:  o.Name,
		Rules: o.Rules,
	}
}

// A set of fields to create a price availability. More info: see the create
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_priceavailabilities/create).
type CreatePriceAvailability struct {
	Name  string                  `json:"name,omitempty"`
	Rules *PriceAvailabilityRules `json:"rules,omitempty"`
}

// A set of fields to update a price availability. More info: see the update
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_priceavailabilities/update).
type UpdatePriceAvailability struct {
	Name  string                  `json:"name,omitempty"`
	Rules *PriceAvailabilityRules `json:"rules,omitempty"`
}

// Set of parameters used to filter price lists. More info: see the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricelists/getlist).
type PriceListParameters struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
}

// An item in a list of price lists. This differs from the normal PriceList type:
// not all fields are present in the list. More info: see the getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricelists/getlist).
type ListPriceList struct {
	// Unique ID
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Hasranks bool   `json:"hasranks,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// A single price list. More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricelists/get).
type PriceList struct {
	// Unique ID
	Id       int              `json:"id,omitempty"`
	Name     string           `json:"name,omitempty"`
	Prices   *PricelistPrices `json:"prices,omitempty"`
	Hasranks bool             `json:"hasranks,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// Convert PriceList to UpdatePriceList
func (o *PriceList) Update() *UpdatePriceList {
	return &UpdatePriceList{
		Name:     o.Name,
		Prices:   o.Prices,
		Hasranks: o.Hasranks,
	}
}

// A set of fields to create a price list. More info: see the create operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricelists/create).
type CreatePriceList struct {
	Name     string           `json:"name,omitempty"`
	Prices   *PricelistPrices `json:"prices,omitempty"`
	Hasranks bool             `json:"hasranks,omitempty"`
}

// A set of fields to update a price list. More info: see the update operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricelists/update).
type UpdatePriceList struct {
	Name     string           `json:"name,omitempty"`
	Prices   *PricelistPrices `json:"prices,omitempty"`
	Hasranks bool             `json:"hasranks,omitempty"`
}

// Set of parameters used to filter price types. More info: see the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricetypes/getlist).
type PriceTypeParameters struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
}

// An item in a list of price types. This differs from the normal PriceType type:
// not all fields are present in the list. More info: see the getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricetypes/getlist).
type ListPriceType struct {
	// Unique ID
	Id int `json:"id,omitempty"`

	// Name of the price type
	Name string `json:"name,omitempty"`

	// The category of this price type, defines how the price is displayed. The
	// available values for this field can be found on the price type overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricetypes)
	// page.
	Typeid int `json:"typeid,omitempty"`

	// A remark that describes the price type. Will be shown to customers.
	Remark string `json:"remark,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// A single price type. More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricetypes/get).
type PriceType struct {
	// Unique ID
	Id int `json:"id,omitempty"`

	// Name of the price type
	Name string `json:"name,omitempty"`

	// The category of this price type, defines how the price is displayed. The
	// available values for this field can be found on the price type overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricetypes)
	// page.
	Typeid int `json:"typeid,omitempty"`

	// A remark that describes the price type. Will be shown to customers.
	Remark string `json:"remark,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// Convert PriceType to UpdatePriceType
func (o *PriceType) Update() *UpdatePriceType {
	return &UpdatePriceType{
		Name:   o.Name,
		Typeid: o.Typeid,
		Remark: o.Remark,
	}
}

// A set of fields to create a price type. More info: see the create operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricetypes/create).
type CreatePriceType struct {
	// Name of the price type
	Name string `json:"name,omitempty"`

	// The category of this price type, defines how the price is displayed. The
	// available values for this field can be found on the price type overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricetypes)
	// page.
	Typeid int `json:"typeid,omitempty"`

	// A remark that describes the price type. Will be shown to customers.
	Remark string `json:"remark,omitempty"`
}

// A set of fields to update a price type. More info: see the update operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricetypes/update).
type UpdatePriceType struct {
	// Name of the price type
	Name string `json:"name,omitempty"`

	// The category of this price type, defines how the price is displayed. The
	// available values for this field can be found on the price type overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricetypes)
	// page.
	Typeid int `json:"typeid,omitempty"`

	// A remark that describes the price type. Will be shown to customers.
	Remark string `json:"remark,omitempty"`
}

// Set of parameters used to filter revenue split categories. More info: see the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplitcategories/getlist).
type RevenueSplitCategoryParameters struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
}

// An item in a list of revenue split categories. This differs from the normal
// RevenueSplitCategory type: not all fields are present in the list. More info:
// see the getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplitcategories/getlist).
type ListRevenueSplitCategory struct {
	// Unique ID
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// A single revenue split category. More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplitcategories/get).
type RevenueSplitCategory struct {
	// Unique ID
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// Convert RevenueSplitCategory to UpdateRevenueSplitCategory
func (o *RevenueSplitCategory) Update() *UpdateRevenueSplitCategory {
	return &UpdateRevenueSplitCategory{
		Name: o.Name,
	}
}

// A set of fields to create a revenue split category. More info: see the create
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplitcategories/create).
type CreateRevenueSplitCategory struct {
	Name string `json:"name,omitempty"`
}

// A set of fields to update a revenue split category. More info: see the update
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplitcategories/update).
type UpdateRevenueSplitCategory struct {
	Name string `json:"name,omitempty"`
}

// Set of parameters used to filter revenue splits. More info: see the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplits/getlist).
type RevenueSplitParameters struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
}

// An item in a list of revenue splits. This differs from the normal RevenueSplit
// type: not all fields are present in the list. More info: see the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplits/getlist).
type ListRevenueSplit struct {
	// Unique ID
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// A single revenue split. More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplits/get).
type RevenueSplit struct {
	// Unique ID
	Id    int                `json:"id,omitempty"`
	Name  string             `json:"name,omitempty"`
	Rules *RevenuesplitRules `json:"rules,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// Convert RevenueSplit to UpdateRevenueSplit
func (o *RevenueSplit) Update() *UpdateRevenueSplit {
	return &UpdateRevenueSplit{
		Name:  o.Name,
		Rules: o.Rules,
	}
}

// A set of fields to create a revenue split. More info: see the create operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplits/create).
type CreateRevenueSplit struct {
	Name  string             `json:"name,omitempty"`
	Rules *RevenuesplitRules `json:"rules,omitempty"`
}

// A set of fields to update a revenue split. More info: see the update operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplits/update).
type UpdateRevenueSplit struct {
	Name  string             `json:"name,omitempty"`
	Rules *RevenuesplitRules `json:"rules,omitempty"`
}

// Set of parameters used to filter ticket fees. More info: see the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_ticketfees/getlist).
type TicketFeeParameters struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
}

// An item in a list of ticket fees. This differs from the normal TicketFee type:
// not all fields are present in the list. More info: see the getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_ticketfees/getlist).
type ListTicketFee struct {
	// Unique ID
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// A single ticket fee. More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_ticketfees/get).
type TicketFee struct {
	// Unique ID
	Id    int             `json:"id,omitempty"`
	Name  string          `json:"name,omitempty"`
	Rules *TicketfeeRules `json:"rules,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// Convert TicketFee to UpdateTicketFee
func (o *TicketFee) Update() *UpdateTicketFee {
	return &UpdateTicketFee{
		Name:  o.Name,
		Rules: o.Rules,
	}
}

// A set of fields to create a ticket fee. More info: see the create operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_ticketfees/create).
type CreateTicketFee struct {
	Name  string          `json:"name,omitempty"`
	Rules *TicketfeeRules `json:"rules,omitempty"`
}

// A set of fields to update a ticket fee. More info: see the update operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_ticketfees/update).
type UpdateTicketFee struct {
	Name  string          `json:"name,omitempty"`
	Rules *TicketfeeRules `json:"rules,omitempty"`
}

// Set of parameters used to filter filter definitions. More info: see the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_system_filterdefinitions/getlist).
type FilterDefinitionParameters struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// Only return items with the given typeid.
	Typeid int `json:"typeid,omitempty"`
}

// An item in a list of filter definitions. This differs from the normal
// FilterDefinition type: not all fields are present in the list. More info: see
// the getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_system_filterdefinitions/getlist).
type ListFilterDefinition struct {
	// Unique ID
	Id int `json:"id,omitempty"`

	// Type ID
	Typeid         int    `json:"typeid,omitempty"`
	Description    string `json:"description,omitempty"`
	Sqlclause      string `json:"sqlclause,omitempty"`
	Filtertype     int    `json:"filtertype,omitempty"`
	Checklistquery string `json:"checklistquery,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// A single filter definition. More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_system_filterdefinitions/get).
type FilterDefinition struct {
	// Unique ID
	Id int `json:"id,omitempty"`

	// Type ID
	Typeid         int    `json:"typeid,omitempty"`
	Description    string `json:"description,omitempty"`
	Sqlclause      string `json:"sqlclause,omitempty"`
	Filtertype     int    `json:"filtertype,omitempty"`
	Checklistquery string `json:"checklistquery,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// Convert FilterDefinition to UpdateFilterDefinition
func (o *FilterDefinition) Update() *UpdateFilterDefinition {
	return &UpdateFilterDefinition{
		Description:    o.Description,
		Sqlclause:      o.Sqlclause,
		Filtertype:     o.Filtertype,
		Checklistquery: o.Checklistquery,
	}
}

// A set of fields to create a filter definition. More info: see the create
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_system_filterdefinitions/create).
type CreateFilterDefinition struct {
	// Type ID
	Typeid         int    `json:"typeid,omitempty"`
	Description    string `json:"description,omitempty"`
	Sqlclause      string `json:"sqlclause,omitempty"`
	Filtertype     int    `json:"filtertype,omitempty"`
	Checklistquery string `json:"checklistquery,omitempty"`
}

// A set of fields to update a filter definition. More info: see the update
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_system_filterdefinitions/update).
type UpdateFilterDefinition struct {
	Description    string `json:"description,omitempty"`
	Sqlclause      string `json:"sqlclause,omitempty"`
	Filtertype     int    `json:"filtertype,omitempty"`
	Checklistquery string `json:"checklistquery,omitempty"`
}

// Set of parameters used to filter delivery scenarios. More info: see the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_deliveryscenarios/getlist).
type DeliveryScenarioParameters struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
}

// An item in a list of delivery scenarios. This differs from the normal
// DeliveryScenario type: not all fields are present in the list. More info: see
// the getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_deliveryscenarios/getlist).
type ListDeliveryScenario struct {
	// Unique ID
	Id int `json:"id,omitempty"`

	// Name of the delivery scenario
	Name string `json:"name,omitempty"`

	// A short description of the deilvery scenario. Will be shown to customers.
	Shortdescription string `json:"shortdescription,omitempty"`

	// An internal description field. Will not be shown to customers.
	Internalremark string `json:"internalremark,omitempty"`

	// The type of this delivery scenario, defines when this delivery scenario is
	// triggered. The available values for this field can be found on the delivery
	// scenario overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_deliveryscenarios)
	// page.
	Typeid int `json:"typeid,omitempty"`

	// A physical address is required
	Needsaddress bool `json:"needsaddress,omitempty"`

	// The ID of the order mail template that will be used for sending out this
	// delivery scenario. Can be 0 to indicate that no mail should be sent
	OrdermailtemplateidDelivery int `json:"ordermailtemplateid_delivery,omitempty"`

	// Are e-tickets allowed with this delivery scenario?
	Allowetickets int `json:"allowetickets,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// A single delivery scenario. More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_deliveryscenarios/get).
type DeliveryScenario struct {
	// Unique ID
	Id int `json:"id,omitempty"`

	// Name of the delivery scenario
	Name string `json:"name,omitempty"`

	// A short description of the deilvery scenario. Will be shown to customers.
	Shortdescription string `json:"shortdescription,omitempty"`

	// An internal description field. Will not be shown to customers.
	Internalremark string `json:"internalremark,omitempty"`

	// The type of this delivery scenario, defines when this delivery scenario is
	// triggered. The available values for this field can be found on the delivery
	// scenario overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_deliveryscenarios)
	// page.
	Typeid int `json:"typeid,omitempty"`

	// A physical address is required
	Needsaddress bool `json:"needsaddress,omitempty"`

	// The rules that define when this scenario is available. See the delivery scenario
	// overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_deliveryscenarios)
	// page for a description of this field
	Availability *DeliveryscenarioAvailability `json:"availability,omitempty"`

	// The ID of the order mail template that will be used for sending out this
	// delivery scenario. Can be 0 to indicate that no mail should be sent
	OrdermailtemplateidDelivery int `json:"ordermailtemplateid_delivery,omitempty"`

	// Are e-tickets allowed with this delivery scenario?
	Allowetickets int `json:"allowetickets,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// Convert DeliveryScenario to UpdateDeliveryScenario
func (o *DeliveryScenario) Update() *UpdateDeliveryScenario {
	return &UpdateDeliveryScenario{
		Name:                        o.Name,
		Shortdescription:            o.Shortdescription,
		Internalremark:              o.Internalremark,
		Typeid:                      o.Typeid,
		Needsaddress:                o.Needsaddress,
		Availability:                o.Availability,
		OrdermailtemplateidDelivery: o.OrdermailtemplateidDelivery,
		Allowetickets:               o.Allowetickets,
	}
}

// A set of fields to create a delivery scenario. More info: see the create
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_deliveryscenarios/create).
type CreateDeliveryScenario struct {
	// Name of the delivery scenario
	Name string `json:"name,omitempty"`

	// A short description of the deilvery scenario. Will be shown to customers.
	Shortdescription string `json:"shortdescription,omitempty"`

	// An internal description field. Will not be shown to customers.
	Internalremark string `json:"internalremark,omitempty"`

	// The type of this delivery scenario, defines when this delivery scenario is
	// triggered. The available values for this field can be found on the delivery
	// scenario overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_deliveryscenarios)
	// page.
	Typeid int `json:"typeid,omitempty"`

	// A physical address is required
	Needsaddress bool `json:"needsaddress,omitempty"`

	// The rules that define when this scenario is available. See the delivery scenario
	// overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_deliveryscenarios)
	// page for a description of this field
	Availability *DeliveryscenarioAvailability `json:"availability,omitempty"`

	// The ID of the order mail template that will be used for sending out this
	// delivery scenario. Can be 0 to indicate that no mail should be sent
	OrdermailtemplateidDelivery int `json:"ordermailtemplateid_delivery,omitempty"`

	// Are e-tickets allowed with this delivery scenario?
	Allowetickets int `json:"allowetickets,omitempty"`
}

// A set of fields to update a delivery scenario. More info: see the update
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_deliveryscenarios/update).
type UpdateDeliveryScenario struct {
	// Name of the delivery scenario
	Name string `json:"name,omitempty"`

	// A short description of the deilvery scenario. Will be shown to customers.
	Shortdescription string `json:"shortdescription,omitempty"`

	// An internal description field. Will not be shown to customers.
	Internalremark string `json:"internalremark,omitempty"`

	// The type of this delivery scenario, defines when this delivery scenario is
	// triggered. The available values for this field can be found on the delivery
	// scenario overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_deliveryscenarios)
	// page.
	Typeid int `json:"typeid,omitempty"`

	// A physical address is required
	Needsaddress bool `json:"needsaddress,omitempty"`

	// The rules that define when this scenario is available. See the delivery scenario
	// overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_deliveryscenarios)
	// page for a description of this field
	Availability *DeliveryscenarioAvailability `json:"availability,omitempty"`

	// The ID of the order mail template that will be used for sending out this
	// delivery scenario. Can be 0 to indicate that no mail should be sent
	OrdermailtemplateidDelivery int `json:"ordermailtemplateid_delivery,omitempty"`

	// Are e-tickets allowed with this delivery scenario?
	Allowetickets int `json:"allowetickets,omitempty"`
}

// Set of parameters used to filter lock types. More info: see the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_locktypes/getlist).
type LockTypeParameters struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
}

// An item in a list of lock types. This differs from the normal LockType type: not
// all fields are present in the list. More info: see the getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_locktypes/getlist).
type ListLockType struct {
	// Unique ID
	Id         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Ishardlock bool   `json:"ishardlock,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// A single lock type. More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_locktypes/get).
type LockType struct {
	// Unique ID
	Id         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Ishardlock bool   `json:"ishardlock,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// Convert LockType to UpdateLockType
func (o *LockType) Update() *UpdateLockType {
	return &UpdateLockType{
		Name:       o.Name,
		Ishardlock: o.Ishardlock,
	}
}

// A set of fields to create a lock type. More info: see the create operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_locktypes/create).
type CreateLockType struct {
	Name       string `json:"name,omitempty"`
	Ishardlock bool   `json:"ishardlock,omitempty"`
}

// A set of fields to update a lock type. More info: see the update operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_locktypes/update).
type UpdateLockType struct {
	Name       string `json:"name,omitempty"`
	Ishardlock bool   `json:"ishardlock,omitempty"`
}

// Set of parameters used to filter order fees. More info: see the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_orderfees/getlist).
type OrderFeeParameters struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
}

// An item in a list of order fees. This differs from the normal OrderFee type: not
// all fields are present in the list. More info: see the getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_orderfees/getlist).
type ListOrderFee struct {
	// Unique ID
	Id     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Typeid int    `json:"typeid,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// A single order fee. More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_orderfees/get).
type OrderFee struct {
	// Unique ID
	Id     int           `json:"id,omitempty"`
	Name   string        `json:"name,omitempty"`
	Typeid int           `json:"typeid,omitempty"`
	Rule   *OrderfeeRule `json:"rule,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// Convert OrderFee to UpdateOrderFee
func (o *OrderFee) Update() *UpdateOrderFee {
	return &UpdateOrderFee{
		Name:   o.Name,
		Typeid: o.Typeid,
		Rule:   o.Rule,
	}
}

// A set of fields to create a order fee. More info: see the create operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_orderfees/create).
type CreateOrderFee struct {
	Name   string        `json:"name,omitempty"`
	Typeid int           `json:"typeid,omitempty"`
	Rule   *OrderfeeRule `json:"rule,omitempty"`
}

// A set of fields to update a order fee. More info: see the update operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_orderfees/update).
type UpdateOrderFee struct {
	Name   string        `json:"name,omitempty"`
	Typeid int           `json:"typeid,omitempty"`
	Rule   *OrderfeeRule `json:"rule,omitempty"`
}

// Set of parameters used to filter payment methods. More info: see the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentmethods/getlist).
type PaymentMethodParameters struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
}

// An item in a list of payment methods. This differs from the normal PaymentMethod
// type: not all fields are present in the list. More info: see the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentmethods/getlist).
type ListPaymentMethod struct {
	// Unique ID
	Id                      int    `json:"id,omitempty"`
	Name                    string `json:"name,omitempty"`
	Internalremark          string `json:"internalremark,omitempty"`
	Paymentmethodtypeid     int    `json:"paymentmethodtypeid,omitempty"`
	Paymentmethodreceiverid int    `json:"paymentmethodreceiverid,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// A single payment method. More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentmethods/get).
type PaymentMethod struct {
	// Unique ID
	Id                      int                  `json:"id,omitempty"`
	Name                    string               `json:"name,omitempty"`
	Internalremark          string               `json:"internalremark,omitempty"`
	Paymentmethodtypeid     int                  `json:"paymentmethodtypeid,omitempty"`
	Paymentmethodreceiverid int                  `json:"paymentmethodreceiverid,omitempty"`
	Config                  *PaymentmethodConfig `json:"config,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// Convert PaymentMethod to UpdatePaymentMethod
func (o *PaymentMethod) Update() *UpdatePaymentMethod {
	return &UpdatePaymentMethod{
		Name:                    o.Name,
		Internalremark:          o.Internalremark,
		Paymentmethodtypeid:     o.Paymentmethodtypeid,
		Paymentmethodreceiverid: o.Paymentmethodreceiverid,
		Config:                  o.Config,
	}
}

// A set of fields to create a payment method. More info: see the create operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentmethods/create).
type CreatePaymentMethod struct {
	Name                    string               `json:"name,omitempty"`
	Internalremark          string               `json:"internalremark,omitempty"`
	Paymentmethodtypeid     int                  `json:"paymentmethodtypeid,omitempty"`
	Paymentmethodreceiverid int                  `json:"paymentmethodreceiverid,omitempty"`
	Config                  *PaymentmethodConfig `json:"config,omitempty"`
}

// A set of fields to update a payment method. More info: see the update operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentmethods/update).
type UpdatePaymentMethod struct {
	Name                    string               `json:"name,omitempty"`
	Internalremark          string               `json:"internalremark,omitempty"`
	Paymentmethodtypeid     int                  `json:"paymentmethodtypeid,omitempty"`
	Paymentmethodreceiverid int                  `json:"paymentmethodreceiverid,omitempty"`
	Config                  *PaymentmethodConfig `json:"config,omitempty"`
}

// Set of parameters used to filter payment scenarios. More info: see the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentscenarios/getlist).
type PaymentScenarioParameters struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
}

// An item in a list of payment scenarios. This differs from the normal
// PaymentScenario type: not all fields are present in the list. More info: see the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentscenarios/getlist).
type ListPaymentScenario struct {
	// Unique ID
	Id int `json:"id,omitempty"`

	// Name of the payment scenario
	Name string `json:"name,omitempty"`

	// Short description of the payment scenario, will be shown to customers
	Shortdescription string `json:"shortdescription,omitempty"`

	// An internal remark, which is never shown to customers. Can be used to
	// distinguish identically named payment scenarios.
	//
	// For example: You could have two VISA scenarios, one for the web sales and one
	// for the box office, each will have different fee configurations. Both will be
	// named VISA, this field can be used to distinguish them.
	Internalremark                        string `json:"internalremark,omitempty"`
	Typeid                                int    `json:"typeid,omitempty"`
	Paymentmethods                        []int  `json:"paymentmethods,omitempty"`
	OrdermailtemplateidPaymentinstruction int    `json:"ordermailtemplateid_paymentinstruction,omitempty"`
	OrdermailtemplateidOverdue            int    `json:"ordermailtemplateid_overdue,omitempty"`
	OrdermailtemplateidExpiry             int    `json:"ordermailtemplateid_expiry,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// A single payment scenario. More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentscenarios/get).
type PaymentScenario struct {
	// Unique ID
	Id int `json:"id,omitempty"`

	// Name of the payment scenario
	Name string `json:"name,omitempty"`

	// Short description of the payment scenario, will be shown to customers
	Shortdescription string `json:"shortdescription,omitempty"`

	// An internal remark, which is never shown to customers. Can be used to
	// distinguish identically named payment scenarios.
	//
	// For example: You could have two VISA scenarios, one for the web sales and one
	// for the box office, each will have different fee configurations. Both will be
	// named VISA, this field can be used to distinguish them.
	Internalremark                        string                            `json:"internalremark,omitempty"`
	Typeid                                int                               `json:"typeid,omitempty"`
	Overdueparameters                     *PaymentscenarioOverdueParameters `json:"overdueparameters,omitempty"`
	Expiryparameters                      *PaymentscenarioExpiryParameters  `json:"expiryparameters,omitempty"`
	Availability                          *PaymentscenarioAvailability      `json:"availability,omitempty"`
	Paymentmethods                        []int                             `json:"paymentmethods,omitempty"`
	OrdermailtemplateidPaymentinstruction int                               `json:"ordermailtemplateid_paymentinstruction,omitempty"`
	OrdermailtemplateidOverdue            int                               `json:"ordermailtemplateid_overdue,omitempty"`
	OrdermailtemplateidExpiry             int                               `json:"ordermailtemplateid_expiry,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// Convert PaymentScenario to UpdatePaymentScenario
func (o *PaymentScenario) Update() *UpdatePaymentScenario {
	return &UpdatePaymentScenario{
		Name:                                  o.Name,
		Shortdescription:                      o.Shortdescription,
		Internalremark:                        o.Internalremark,
		Typeid:                                o.Typeid,
		Overdueparameters:                     o.Overdueparameters,
		Expiryparameters:                      o.Expiryparameters,
		Availability:                          o.Availability,
		Paymentmethods:                        o.Paymentmethods,
		OrdermailtemplateidPaymentinstruction: o.OrdermailtemplateidPaymentinstruction,
		OrdermailtemplateidOverdue:            o.OrdermailtemplateidOverdue,
		OrdermailtemplateidExpiry:             o.OrdermailtemplateidExpiry,
	}
}

// A set of fields to create a payment scenario. More info: see the create
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentscenarios/create).
type CreatePaymentScenario struct {
	// Name of the payment scenario
	Name string `json:"name,omitempty"`

	// Short description of the payment scenario, will be shown to customers
	Shortdescription string `json:"shortdescription,omitempty"`

	// An internal remark, which is never shown to customers. Can be used to
	// distinguish identically named payment scenarios.
	//
	// For example: You could have two VISA scenarios, one for the web sales and one
	// for the box office, each will have different fee configurations. Both will be
	// named VISA, this field can be used to distinguish them.
	Internalremark                        string                            `json:"internalremark,omitempty"`
	Typeid                                int                               `json:"typeid,omitempty"`
	Overdueparameters                     *PaymentscenarioOverdueParameters `json:"overdueparameters,omitempty"`
	Expiryparameters                      *PaymentscenarioExpiryParameters  `json:"expiryparameters,omitempty"`
	Availability                          *PaymentscenarioAvailability      `json:"availability,omitempty"`
	Paymentmethods                        []int                             `json:"paymentmethods,omitempty"`
	OrdermailtemplateidPaymentinstruction int                               `json:"ordermailtemplateid_paymentinstruction,omitempty"`
	OrdermailtemplateidOverdue            int                               `json:"ordermailtemplateid_overdue,omitempty"`
	OrdermailtemplateidExpiry             int                               `json:"ordermailtemplateid_expiry,omitempty"`
}

// A set of fields to update a payment scenario. More info: see the update
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentscenarios/update).
type UpdatePaymentScenario struct {
	// Name of the payment scenario
	Name string `json:"name,omitempty"`

	// Short description of the payment scenario, will be shown to customers
	Shortdescription string `json:"shortdescription,omitempty"`

	// An internal remark, which is never shown to customers. Can be used to
	// distinguish identically named payment scenarios.
	//
	// For example: You could have two VISA scenarios, one for the web sales and one
	// for the box office, each will have different fee configurations. Both will be
	// named VISA, this field can be used to distinguish them.
	Internalremark                        string                            `json:"internalremark,omitempty"`
	Typeid                                int                               `json:"typeid,omitempty"`
	Overdueparameters                     *PaymentscenarioOverdueParameters `json:"overdueparameters,omitempty"`
	Expiryparameters                      *PaymentscenarioExpiryParameters  `json:"expiryparameters,omitempty"`
	Availability                          *PaymentscenarioAvailability      `json:"availability,omitempty"`
	Paymentmethods                        []int                             `json:"paymentmethods,omitempty"`
	OrdermailtemplateidPaymentinstruction int                               `json:"ordermailtemplateid_paymentinstruction,omitempty"`
	OrdermailtemplateidOverdue            int                               `json:"ordermailtemplateid_overdue,omitempty"`
	OrdermailtemplateidExpiry             int                               `json:"ordermailtemplateid_expiry,omitempty"`
}

// Set of parameters used to filter sales channels. More info: see the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_saleschannels/getlist).
type SalesChannelParameters struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
}

// An item in a list of sales channels. This differs from the normal SalesChannel
// type: not all fields are present in the list. More info: see the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_saleschannels/getlist).
type ListSalesChannel struct {
	// Unique ID
	Id int `json:"id,omitempty"`

	// Name of the sales channel
	Name string `json:"name,omitempty"`

	// The type of this sales channel, defines where this sales channel will be used.
	// The available values for this field can be found on the sales channel overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_saleschannels)
	// page.
	Typeid int `json:"typeid,omitempty"`

	// The ID of the order mail template to use for sending confirmations
	OrdermailtemplateidConfirmation int `json:"ordermailtemplateid_confirmation,omitempty"`

	// Always send the confirmation, regardless of the payment method configuration
	OrdermailtemplateidConfirmationSendalways bool `json:"ordermailtemplateid_confirmation_sendalways,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// A single sales channel. More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_saleschannels/get).
type SalesChannel struct {
	// Unique ID
	Id int `json:"id,omitempty"`

	// Name of the sales channel
	Name string `json:"name,omitempty"`

	// The type of this sales channel, defines where this sales channel will be used.
	// The available values for this field can be found on the sales channel overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_saleschannels)
	// page.
	Typeid int `json:"typeid,omitempty"`

	// The ID of the order mail template to use for sending confirmations
	OrdermailtemplateidConfirmation int `json:"ordermailtemplateid_confirmation,omitempty"`

	// Always send the confirmation, regardless of the payment method configuration
	OrdermailtemplateidConfirmationSendalways bool `json:"ordermailtemplateid_confirmation_sendalways,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// Convert SalesChannel to UpdateSalesChannel
func (o *SalesChannel) Update() *UpdateSalesChannel {
	return &UpdateSalesChannel{
		Name:   o.Name,
		Typeid: o.Typeid,
		OrdermailtemplateidConfirmation:           o.OrdermailtemplateidConfirmation,
		OrdermailtemplateidConfirmationSendalways: o.OrdermailtemplateidConfirmationSendalways,
	}
}

// A set of fields to create a sales channel. More info: see the create operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_saleschannels/create).
type CreateSalesChannel struct {
	// Name of the sales channel
	Name string `json:"name,omitempty"`

	// The type of this sales channel, defines where this sales channel will be used.
	// The available values for this field can be found on the sales channel overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_saleschannels)
	// page.
	Typeid int `json:"typeid,omitempty"`

	// The ID of the order mail template to use for sending confirmations
	OrdermailtemplateidConfirmation int `json:"ordermailtemplateid_confirmation,omitempty"`

	// Always send the confirmation, regardless of the payment method configuration
	OrdermailtemplateidConfirmationSendalways bool `json:"ordermailtemplateid_confirmation_sendalways,omitempty"`
}

// A set of fields to update a sales channel. More info: see the update operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_saleschannels/update).
type UpdateSalesChannel struct {
	// Name of the sales channel
	Name string `json:"name,omitempty"`

	// The type of this sales channel, defines where this sales channel will be used.
	// The available values for this field can be found on the sales channel overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_saleschannels)
	// page.
	Typeid int `json:"typeid,omitempty"`

	// The ID of the order mail template to use for sending confirmations
	OrdermailtemplateidConfirmation int `json:"ordermailtemplateid_confirmation,omitempty"`

	// Always send the confirmation, regardless of the payment method configuration
	OrdermailtemplateidConfirmationSendalways bool `json:"ordermailtemplateid_confirmation_sendalways,omitempty"`
}
