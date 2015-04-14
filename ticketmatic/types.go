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

type CreateOrderMailTemplate struct {
	Name         string            `json:"name,omitempty"`
	Typeid       int               `json:"typeid,omitempty"`
	Subject      string            `json:"subject,omitempty"`
	Body         string            `json:"body,omitempty"`
	Translations map[string]string `json:"translations,omitempty"`
}

type UpdateOrderMailTemplate struct {
	Name         string            `json:"name,omitempty"`
	Typeid       int               `json:"typeid,omitempty"`
	Subject      string            `json:"subject,omitempty"`
	Body         string            `json:"body,omitempty"`
	Translations map[string]string `json:"translations,omitempty"`
}

type WebSalesSkinParameters struct {
	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
}

type ListWebSalesSkin struct {
	// Unique ID
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`
}

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

type CreateWebSalesSkin struct {
	Name          string                `json:"name,omitempty"`
	Html          string                `json:"html,omitempty"`
	Css           string                `json:"css,omitempty"`
	Translations  map[string]string     `json:"translations,omitempty"`
	Configuration *WebskinConfiguration `json:"configuration,omitempty"`
}

type UpdateWebSalesSkin struct {
	Name          string                `json:"name,omitempty"`
	Html          string                `json:"html,omitempty"`
	Css           string                `json:"css,omitempty"`
	Translations  map[string]string     `json:"translations,omitempty"`
	Configuration *WebskinConfiguration `json:"configuration,omitempty"`
}

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

type CreatePriceAvailability struct {
	Name  string                  `json:"name,omitempty"`
	Rules *PriceAvailabilityRules `json:"rules,omitempty"`
}

type UpdatePriceAvailability struct {
	Name  string                  `json:"name,omitempty"`
	Rules *PriceAvailabilityRules `json:"rules,omitempty"`
}

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

type CreatePriceList struct {
	Name     string           `json:"name,omitempty"`
	Prices   *PricelistPrices `json:"prices,omitempty"`
	Hasranks bool             `json:"hasranks,omitempty"`
}

type UpdatePriceList struct {
	Name     string           `json:"name,omitempty"`
	Prices   *PricelistPrices `json:"prices,omitempty"`
	Hasranks bool             `json:"hasranks,omitempty"`
}

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

type ListPriceType struct {
	// Unique ID
	Id     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Typeid int    `json:"typeid,omitempty"`
	Remark string `json:"remark,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

type PriceType struct {
	// Unique ID
	Id     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Typeid int    `json:"typeid,omitempty"`
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

type CreatePriceType struct {
	Name   string `json:"name,omitempty"`
	Typeid int    `json:"typeid,omitempty"`
	Remark string `json:"remark,omitempty"`
}

type UpdatePriceType struct {
	Name   string `json:"name,omitempty"`
	Typeid int    `json:"typeid,omitempty"`
	Remark string `json:"remark,omitempty"`
}

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

type CreateRevenueSplitCategory struct {
	Name string `json:"name,omitempty"`
}

type UpdateRevenueSplitCategory struct {
	Name string `json:"name,omitempty"`
}

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

type CreateRevenueSplit struct {
	Name  string             `json:"name,omitempty"`
	Rules *RevenuesplitRules `json:"rules,omitempty"`
}

type UpdateRevenueSplit struct {
	Name  string             `json:"name,omitempty"`
	Rules *RevenuesplitRules `json:"rules,omitempty"`
}

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

type CreateTicketFee struct {
	Name  string          `json:"name,omitempty"`
	Rules *TicketfeeRules `json:"rules,omitempty"`
}

type UpdateTicketFee struct {
	Name  string          `json:"name,omitempty"`
	Rules *TicketfeeRules `json:"rules,omitempty"`
}

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

type CreateFilterDefinition struct {
	// Type ID
	Typeid         int    `json:"typeid,omitempty"`
	Description    string `json:"description,omitempty"`
	Sqlclause      string `json:"sqlclause,omitempty"`
	Filtertype     int    `json:"filtertype,omitempty"`
	Checklistquery string `json:"checklistquery,omitempty"`
}

type UpdateFilterDefinition struct {
	Description    string `json:"description,omitempty"`
	Sqlclause      string `json:"sqlclause,omitempty"`
	Filtertype     int    `json:"filtertype,omitempty"`
	Checklistquery string `json:"checklistquery,omitempty"`
}

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

type CreateLockType struct {
	Name       string `json:"name,omitempty"`
	Ishardlock bool   `json:"ishardlock,omitempty"`
}

type UpdateLockType struct {
	Name       string `json:"name,omitempty"`
	Ishardlock bool   `json:"ishardlock,omitempty"`
}

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

type CreateOrderFee struct {
	Name   string        `json:"name,omitempty"`
	Typeid int           `json:"typeid,omitempty"`
	Rule   *OrderfeeRule `json:"rule,omitempty"`
}

type UpdateOrderFee struct {
	Name   string        `json:"name,omitempty"`
	Typeid int           `json:"typeid,omitempty"`
	Rule   *OrderfeeRule `json:"rule,omitempty"`
}

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

type CreatePaymentMethod struct {
	Name                    string               `json:"name,omitempty"`
	Internalremark          string               `json:"internalremark,omitempty"`
	Paymentmethodtypeid     int                  `json:"paymentmethodtypeid,omitempty"`
	Paymentmethodreceiverid int                  `json:"paymentmethodreceiverid,omitempty"`
	Config                  *PaymentmethodConfig `json:"config,omitempty"`
}

type UpdatePaymentMethod struct {
	Name                    string               `json:"name,omitempty"`
	Internalremark          string               `json:"internalremark,omitempty"`
	Paymentmethodtypeid     int                  `json:"paymentmethodtypeid,omitempty"`
	Paymentmethodreceiverid int                  `json:"paymentmethodreceiverid,omitempty"`
	Config                  *PaymentmethodConfig `json:"config,omitempty"`
}

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
