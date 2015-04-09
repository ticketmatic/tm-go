package ticketmatic

type DeliveryscenarioAvailability struct {
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

	// All items that were updated since this timestamp will be returned. Timestamp should be passed
	// in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that returns the ids.
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
	// All items that were updated since this timestamp will be returned. Timestamp should be passed
	// in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that returns the ids.
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

	// All items that were updated since this timestamp will be returned. Timestamp should be passed
	// in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that returns the ids.
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

	// Country code. Should be an ISO 3166-1 alpha-2 (http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
	// two-letter code.
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

	// Country code. Should be an ISO 3166-1 alpha-2 (http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
	// two-letter code.
	Countrycode string `json:"countrycode,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
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

	// Country code. Should be an ISO 3166-1 alpha-2 (http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
	// two-letter code.
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

	// Country code. Should be an ISO 3166-1 alpha-2 (http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
	// two-letter code.
	Countrycode string `json:"countrycode,omitempty"`
}

type PriceAvailabilityParameters struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp should be passed
	// in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that returns the ids.
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

	// All items that were updated since this timestamp will be returned. Timestamp should be passed
	// in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that returns the ids.
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

	// All items that were updated since this timestamp will be returned. Timestamp should be passed
	// in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that returns the ids.
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

	// All items that were updated since this timestamp will be returned. Timestamp should be passed
	// in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that returns the ids.
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

type CreateRevenueSplitCategory struct {
	Name string `json:"name,omitempty"`
}

type UpdateRevenueSplitCategory struct {
	Name string `json:"name,omitempty"`
}

type RevenueSplitParameters struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp should be passed
	// in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that returns the ids.
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

	// All items that were updated since this timestamp will be returned. Timestamp should be passed
	// in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that returns the ids.
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

	// All items that were updated since this timestamp will be returned. Timestamp should be passed
	// in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that returns the ids.
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

	// All items that were updated since this timestamp will be returned. Timestamp should be passed
	// in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that returns the ids.
	Filter string `json:"filter,omitempty"`
}

type ListDeliveryScenario struct {
	// Unique ID
	Id                          int    `json:"id,omitempty"`
	Name                        string `json:"name,omitempty"`
	Shortdescription            string `json:"shortdescription,omitempty"`
	Internalremark              string `json:"internalremark,omitempty"`
	Typeid                      int    `json:"typeid,omitempty"`
	Needsaddress                bool   `json:"needsaddress,omitempty"`
	OrdermailtemplateidDelivery int    `json:"ordermailtemplateid_delivery,omitempty"`
	Allowetickets               int    `json:"allowetickets,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

type DeliveryScenario struct {
	// Unique ID
	Id                          int                           `json:"id,omitempty"`
	Name                        string                        `json:"name,omitempty"`
	Shortdescription            string                        `json:"shortdescription,omitempty"`
	Internalremark              string                        `json:"internalremark,omitempty"`
	Typeid                      int                           `json:"typeid,omitempty"`
	Needsaddress                bool                          `json:"needsaddress,omitempty"`
	Availability                *DeliveryscenarioAvailability `json:"availability,omitempty"`
	OrdermailtemplateidDelivery int                           `json:"ordermailtemplateid_delivery,omitempty"`
	Allowetickets               int                           `json:"allowetickets,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

type CreateDeliveryScenario struct {
	Name                        string                        `json:"name,omitempty"`
	Shortdescription            string                        `json:"shortdescription,omitempty"`
	Internalremark              string                        `json:"internalremark,omitempty"`
	Typeid                      int                           `json:"typeid,omitempty"`
	Needsaddress                bool                          `json:"needsaddress,omitempty"`
	Availability                *DeliveryscenarioAvailability `json:"availability,omitempty"`
	OrdermailtemplateidDelivery int                           `json:"ordermailtemplateid_delivery,omitempty"`
	Allowetickets               int                           `json:"allowetickets,omitempty"`
}

type UpdateDeliveryScenario struct {
	Name                        string                        `json:"name,omitempty"`
	Shortdescription            string                        `json:"shortdescription,omitempty"`
	Internalremark              string                        `json:"internalremark,omitempty"`
	Typeid                      int                           `json:"typeid,omitempty"`
	Needsaddress                bool                          `json:"needsaddress,omitempty"`
	Availability                *DeliveryscenarioAvailability `json:"availability,omitempty"`
	OrdermailtemplateidDelivery int                           `json:"ordermailtemplateid_delivery,omitempty"`
	Allowetickets               int                           `json:"allowetickets,omitempty"`
}

type LockTypeParameters struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp should be passed
	// in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that returns the ids.
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

	// All items that were updated since this timestamp will be returned. Timestamp should be passed
	// in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that returns the ids.
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

	// All items that were updated since this timestamp will be returned. Timestamp should be passed
	// in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that returns the ids.
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

	// All items that were updated since this timestamp will be returned. Timestamp should be passed
	// in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that returns the ids.
	Filter string `json:"filter,omitempty"`
}

type ListPaymentScenario struct {
	// Unique ID
	Id int `json:"id,omitempty"`

	// Name of the payment scenario
	Name string `json:"name,omitempty"`

	// Short description of the payment scenario, will be shown to customers
	Shortdescription string `json:"shortdescription,omitempty"`

	// An internal remark, which is never shown to customers. Can be used to distinguish identically
	// named payment scenarios.
	//
	// For example: You could have two VISA scenarios, one for the web sales and one for the box
	// office, each will have different fee configurations. Both will be named VISA, this field can be
	// used to distinguish them.
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

	// An internal remark, which is never shown to customers. Can be used to distinguish identically
	// named payment scenarios.
	//
	// For example: You could have two VISA scenarios, one for the web sales and one for the box
	// office, each will have different fee configurations. Both will be named VISA, this field can be
	// used to distinguish them.
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

type CreatePaymentScenario struct {
	// Name of the payment scenario
	Name string `json:"name,omitempty"`

	// Short description of the payment scenario, will be shown to customers
	Shortdescription string `json:"shortdescription,omitempty"`

	// An internal remark, which is never shown to customers. Can be used to distinguish identically
	// named payment scenarios.
	//
	// For example: You could have two VISA scenarios, one for the web sales and one for the box
	// office, each will have different fee configurations. Both will be named VISA, this field can be
	// used to distinguish them.
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

	// An internal remark, which is never shown to customers. Can be used to distinguish identically
	// named payment scenarios.
	//
	// For example: You could have two VISA scenarios, one for the web sales and one for the box
	// office, each will have different fee configurations. Both will be named VISA, this field can be
	// used to distinguish them.
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

	// All items that were updated since this timestamp will be returned. Timestamp should be passed
	// in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that returns the ids.
	Filter string `json:"filter,omitempty"`
}

type ListSalesChannel struct {
	// Unique ID
	Id int `json:"id,omitempty"`

	// Name of the sales channel
	Name string `json:"name,omitempty"`

	// The type of this sales channel, defines where this sales channel will be used.
	// The available values for this field can be found on the sales channel overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_saleschannels) page.
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
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_saleschannels) page.
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

type CreateSalesChannel struct {
	// Name of the sales channel
	Name string `json:"name,omitempty"`

	// The type of this sales channel, defines where this sales channel will be used.
	// The available values for this field can be found on the sales channel overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_saleschannels) page.
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
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_saleschannels) page.
	Typeid int `json:"typeid,omitempty"`

	// The ID of the order mail template to use for sending confirmations
	OrdermailtemplateidConfirmation int `json:"ordermailtemplateid_confirmation,omitempty"`

	// Always send the confirmation, regardless of the payment method configuration
	OrdermailtemplateidConfirmationSendalways bool `json:"ordermailtemplateid_confirmation_sendalways,omitempty"`
}
