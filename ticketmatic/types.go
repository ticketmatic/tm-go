package ticketmatic

// A DeliveryscenarioAvailability defines when a delivery scenario
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/DeliveryScenario) is
// available.
//
// This can be done in two ways:
//
// * By specifying a set of sales channels (required)
//
// * By writing a script (optional)
//
// In its simplest form, a DeliveryscenarioAvailability looks like this:
//
//
//    {
//        "saleschannels": [1, 2]
//    }
//
//
//
// This defines that the delivery scenario is available when used in the context of
// saleschannel 1 or 2.
//
// More complex logic can be specified by writing a small piece of JavaScript
// (http://en.wikipedia.org/wiki/JavaScript). To do so, you need to add a usescript
// and script field to the availability:
//
//
//    {
//        "saleschannels": [1, 2],
//        "usescript": true,
//        "script": "// script here"
//    }
//
//
//
// Note that the list of sales channel IDs is still required: the script can only
// restrict this set further.
//
// A simple example of a delivery scenario script:
//
//
//    return order.tickets.length < 3 && saleschannel.typeid == 3002;
//
//
//
// This script states that the current delivery scenario is only available if the
// amount of tickets in the order is less than 3 and the current sales channel is a
// web sales channel.
//
// With this script the DeliveryscenarioAvailability would look like this:
//
//
//    {
//        "saleschannels": [1, 2],
//        "usescript": true,
//        "script": "return order.tickets.length < 3 && saleschannel.typeid == 3002;"
//    }
//
//
//
// The following variables are available in the script:
//
// * order
//
// * saleschannel
//
// You can use any valid JavaScript syntax (including conditionals and loops). Note
// that each script has a strict time limit.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/DeliveryscenarioAvailability).
type DeliveryscenarioAvailability struct {
	// An array of sales channel IDs for which this delivery scenario can be used.
	Saleschannels []int `json:"saleschannels,omitempty"`

	// Use a script to refine the set of sales channels?
	Usescript bool `json:"usescript,omitempty"`

	// Script used to determine availability of the delivery scenario.
	Script string `json:"script,omitempty"`
}

type OrderFeeRule struct {
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

// Configuration settings and parameters for a web sales skin
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/WebSalesSkin).
//
// Page titles
//
// The title field contains a template for the page title. The same variables as in
// the HTML of the skin itself can be used.
//
// Check the web sales skin setup guide
// (https://apps.ticketmatic.com/#/knowledgebase/designer_webskin) for more
// information.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/WebSalesSkinConfiguration).
type WebSalesSkinConfiguration struct {
	// Page title
	Title string `json:"title,omitempty"`

	// Asset path to favicon image.
	Favicon string `json:"favicon,omitempty"`

	// Facebook app ID to use for Facebook authentication.
	//
	// The default Ticketmatic Facebook app will be used if you leave this field blank
	Fbappid string `json:"fbappid,omitempty"`

	// Google Analytics tracking ID. Can be left blank.
	Googleanalyticsid string `json:"googleanalyticsid,omitempty"`
}

// A timestamp returned by the diagnostic /time call
// (https://apps.ticketmatic.com/#/knowledgebase/api/diagnostics/time).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Timestamp).
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

// Set of parameters used to filter order mail templates.
//
// More info: see order mail template
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/OrderMailTemplate), the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ordermails/getlist)
// and the order mail templates endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ordermails).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/OrderMailTemplateParameters).
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

// An item in a list of order mail templates.
//
// This differs from the normal OrderMailTemplate type: not all fields are present
// in the list.
//
// More info: see order mail template
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/OrderMailTemplate), the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ordermails/getlist)
// and the order mail templates endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ordermails).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/ListOrderMailTemplate).
type ListOrderMailTemplate struct {
	// Unique ID
	Id int `json:"id,omitempty"`

	// Name of the order mail template
	Name string `json:"name,omitempty"`

	// The type of this order mail template, defines where this template is used. The
	// available values for this field can be found on the order mail template overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ordermails)
	// page.
	Typeid int `json:"typeid,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	Isarchived bool `json:"isarchived,omitempty"`
}

// A single order mail template.
//
// More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ordermails/get)
// and the order mail templates endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ordermails).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/OrderMailTemplate).
type OrderMailTemplate struct {
	// Unique ID
	Id int `json:"id,omitempty"`

	// Name of the order mail template
	Name string `json:"name,omitempty"`

	// The type of this order mail template, defines where this template is used. The
	// available values for this field can be found on the order mail template overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ordermails)
	// page.
	Typeid int `json:"typeid,omitempty"`

	// Subject line for the order mail template
	Subject string `json:"subject,omitempty"`

	// Message body
	Body string `json:"body,omitempty"`

	// A map of language codes to gettext .po files
	// (http://en.wikipedia.org/wiki/Gettext). More info can be found on the order mail
	// template overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ordermails)
	// page.
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

// A set of fields to create a order mail template.
//
// More info: see order mail template
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/OrderMailTemplate), the
// create operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ordermails/create)
// and the order mail templates endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ordermails).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/CreateOrderMailTemplate).
type CreateOrderMailTemplate struct {
	// Name of the order mail template
	Name string `json:"name,omitempty"`

	// The type of this order mail template, defines where this template is used. The
	// available values for this field can be found on the order mail template overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ordermails)
	// page.
	Typeid int `json:"typeid,omitempty"`

	// Subject line for the order mail template
	Subject string `json:"subject,omitempty"`

	// Message body
	Body string `json:"body,omitempty"`

	// A map of language codes to gettext .po files
	// (http://en.wikipedia.org/wiki/Gettext). More info can be found on the order mail
	// template overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ordermails)
	// page.
	Translations map[string]string `json:"translations,omitempty"`
}

// A set of fields to update a order mail template.
//
// More info: see order mail template
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/OrderMailTemplate), the
// update operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ordermails/update)
// and the order mail templates endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ordermails).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/UpdateOrderMailTemplate).
type UpdateOrderMailTemplate struct {
	// Name of the order mail template
	Name string `json:"name,omitempty"`

	// The type of this order mail template, defines where this template is used. The
	// available values for this field can be found on the order mail template overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ordermails)
	// page.
	Typeid int `json:"typeid,omitempty"`

	// Subject line for the order mail template
	Subject string `json:"subject,omitempty"`

	// Message body
	Body string `json:"body,omitempty"`

	// A map of language codes to gettext .po files
	// (http://en.wikipedia.org/wiki/Gettext). More info can be found on the order mail
	// template overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ordermails)
	// page.
	Translations map[string]string `json:"translations,omitempty"`
}

// Set of parameters used to filter web sales skins.
//
// More info: see web sales skin
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/WebSalesSkin), the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_webskins/getlist)
// and the web sales skins endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_webskins).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/WebSalesSkinParameters).
type WebSalesSkinParameters struct {
	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
}

// An item in a list of web sales skins.
//
// This differs from the normal WebSalesSkin type: not all fields are present in
// the list.
//
// More info: see web sales skin
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/WebSalesSkin), the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_webskins/getlist)
// and the web sales skins endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_webskins).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/ListWebSalesSkin).
type ListWebSalesSkin struct {
	// Unique ID
	Id int `json:"id,omitempty"`

	// Name of the web sales skin
	Name string `json:"name,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`
}

// A single web sales skin.
//
// More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_webskins/get)
// and the web sales skins endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_webskins).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/WebSalesSkin).
type WebSalesSkin struct {
	// Unique ID
	Id int `json:"id,omitempty"`

	// Name of the web sales skin
	Name string `json:"name,omitempty"`

	// HTML template of the skin. See the web skin setup guide
	// (https://apps.ticketmatic.com/#/knowledgebase/designer_webskin) for more
	// information.
	Html string `json:"html,omitempty"`

	// CSS style rules. Should always include the style import.
	Css string `json:"css,omitempty"`

	// A map of language codes to gettext .po files
	// (http://en.wikipedia.org/wiki/Gettext). More info can be found on the web skin
	// overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_webskins)
	// page.
	Translations map[string]string `json:"translations,omitempty"`

	// Skin configuration.
	//
	// See the WebSalesSkinConfiguration reference
	// (https://apps.ticketmatic.com/#/knowledgebase/api/types/WebSalesSkinConfiguration)
	// for an overview of all possible options.
	Configuration *WebSalesSkinConfiguration `json:"configuration,omitempty"`

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

// A set of fields to create a web sales skin.
//
// More info: see web sales skin
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/WebSalesSkin), the
// create operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_webskins/create)
// and the web sales skins endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_webskins).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/CreateWebSalesSkin).
type CreateWebSalesSkin struct {
	// Name of the web sales skin
	Name string `json:"name,omitempty"`

	// HTML template of the skin. See the web skin setup guide
	// (https://apps.ticketmatic.com/#/knowledgebase/designer_webskin) for more
	// information.
	Html string `json:"html,omitempty"`

	// CSS style rules. Should always include the style import.
	Css string `json:"css,omitempty"`

	// A map of language codes to gettext .po files
	// (http://en.wikipedia.org/wiki/Gettext). More info can be found on the web skin
	// overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_webskins)
	// page.
	Translations map[string]string `json:"translations,omitempty"`

	// Skin configuration.
	//
	// See the WebSalesSkinConfiguration reference
	// (https://apps.ticketmatic.com/#/knowledgebase/api/types/WebSalesSkinConfiguration)
	// for an overview of all possible options.
	Configuration *WebSalesSkinConfiguration `json:"configuration,omitempty"`
}

// A set of fields to update a web sales skin.
//
// More info: see web sales skin
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/WebSalesSkin), the
// update operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_webskins/update)
// and the web sales skins endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_webskins).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/UpdateWebSalesSkin).
type UpdateWebSalesSkin struct {
	// Name of the web sales skin
	Name string `json:"name,omitempty"`

	// HTML template of the skin. See the web skin setup guide
	// (https://apps.ticketmatic.com/#/knowledgebase/designer_webskin) for more
	// information.
	Html string `json:"html,omitempty"`

	// CSS style rules. Should always include the style import.
	Css string `json:"css,omitempty"`

	// A map of language codes to gettext .po files
	// (http://en.wikipedia.org/wiki/Gettext). More info can be found on the web skin
	// overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_webskins)
	// page.
	Translations map[string]string `json:"translations,omitempty"`

	// Skin configuration.
	//
	// See the WebSalesSkinConfiguration reference
	// (https://apps.ticketmatic.com/#/knowledgebase/api/types/WebSalesSkinConfiguration)
	// for an overview of all possible options.
	Configuration *WebSalesSkinConfiguration `json:"configuration,omitempty"`
}

// Set of parameters used to filter event locations.
//
// More info: see event location
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/EventLocation), the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_events_eventlocations/getlist)
// and the event locations endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_events_eventlocations).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/EventLocationParameters).
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

// An item in a list of event locations.
//
// This differs from the normal EventLocation type: not all fields are present in
// the list.
//
// More info: see event location
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/EventLocation), the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_events_eventlocations/getlist)
// and the event locations endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_events_eventlocations).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/ListEventLocation).
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

// A single event location.
//
// More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_events_eventlocations/get)
// and the event locations endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_events_eventlocations).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/EventLocation).
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

// A set of fields to create a event location.
//
// More info: see event location
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/EventLocation), the
// create operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_events_eventlocations/create)
// and the event locations endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_events_eventlocations).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/CreateEventLocation).
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

// A set of fields to update a event location.
//
// More info: see event location
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/EventLocation), the
// update operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_events_eventlocations/update)
// and the event locations endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_events_eventlocations).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/UpdateEventLocation).
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

// Set of parameters used to filter price availabilities.
//
// More info: see price availability
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PriceAvailability), the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_priceavailabilities/getlist)
// and the price availabilities endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_priceavailabilities).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PriceAvailabilityParameters).
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

// An item in a list of price availabilities.
//
// This differs from the normal PriceAvailability type: not all fields are present
// in the list.
//
// More info: see price availability
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PriceAvailability), the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_priceavailabilities/getlist)
// and the price availabilities endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_priceavailabilities).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/ListPriceAvailability).
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

// A single price availability.
//
// More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_priceavailabilities/get)
// and the price availabilities endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_priceavailabilities).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PriceAvailability).
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

// A set of fields to create a price availability.
//
// More info: see price availability
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PriceAvailability), the
// create operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_priceavailabilities/create)
// and the price availabilities endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_priceavailabilities).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/CreatePriceAvailability).
type CreatePriceAvailability struct {
	Name  string                  `json:"name,omitempty"`
	Rules *PriceAvailabilityRules `json:"rules,omitempty"`
}

// A set of fields to update a price availability.
//
// More info: see price availability
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PriceAvailability), the
// update operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_priceavailabilities/update)
// and the price availabilities endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_priceavailabilities).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/UpdatePriceAvailability).
type UpdatePriceAvailability struct {
	Name  string                  `json:"name,omitempty"`
	Rules *PriceAvailabilityRules `json:"rules,omitempty"`
}

// Set of parameters used to filter price lists.
//
// More info: see price list
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PriceList), the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricelists/getlist)
// and the price lists endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricelists).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PriceListParameters).
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

// An item in a list of price lists.
//
// This differs from the normal PriceList type: not all fields are present in the
// list.
//
// More info: see price list
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PriceList), the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricelists/getlist)
// and the price lists endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricelists).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/ListPriceList).
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

// A single price list.
//
// More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricelists/get)
// and the price lists endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricelists).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PriceList).
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

// A set of fields to create a price list.
//
// More info: see price list
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PriceList), the create
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricelists/create)
// and the price lists endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricelists).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/CreatePriceList).
type CreatePriceList struct {
	Name     string           `json:"name,omitempty"`
	Prices   *PricelistPrices `json:"prices,omitempty"`
	Hasranks bool             `json:"hasranks,omitempty"`
}

// A set of fields to update a price list.
//
// More info: see price list
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PriceList), the update
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricelists/update)
// and the price lists endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricelists).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/UpdatePriceList).
type UpdatePriceList struct {
	Name     string           `json:"name,omitempty"`
	Prices   *PricelistPrices `json:"prices,omitempty"`
	Hasranks bool             `json:"hasranks,omitempty"`
}

// Set of parameters used to filter price types.
//
// More info: see price type
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PriceType), the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricetypes/getlist)
// and the price types endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricetypes).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PriceTypeParameters).
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

// An item in a list of price types.
//
// This differs from the normal PriceType type: not all fields are present in the
// list.
//
// More info: see price type
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PriceType), the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricetypes/getlist)
// and the price types endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricetypes).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/ListPriceType).
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

// A single price type.
//
// More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricetypes/get)
// and the price types endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricetypes).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PriceType).
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

// A set of fields to create a price type.
//
// More info: see price type
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PriceType), the create
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricetypes/create)
// and the price types endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricetypes).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/CreatePriceType).
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

// A set of fields to update a price type.
//
// More info: see price type
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PriceType), the update
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricetypes/update)
// and the price types endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricetypes).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/UpdatePriceType).
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

// Set of parameters used to filter revenue split categories.
//
// More info: see revenue split category
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/RevenueSplitCategory),
// the getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplitcategories/getlist)
// and the revenue split categories endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplitcategories).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/RevenueSplitCategoryParameters).
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

// An item in a list of revenue split categories.
//
// This differs from the normal RevenueSplitCategory type: not all fields are
// present in the list.
//
// More info: see revenue split category
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/RevenueSplitCategory),
// the getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplitcategories/getlist)
// and the revenue split categories endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplitcategories).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/ListRevenueSplitCategory).
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

// A single revenue split category.
//
// More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplitcategories/get)
// and the revenue split categories endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplitcategories).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/RevenueSplitCategory).
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

// A set of fields to create a revenue split category.
//
// More info: see revenue split category
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/RevenueSplitCategory),
// the create operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplitcategories/create)
// and the revenue split categories endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplitcategories).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/CreateRevenueSplitCategory).
type CreateRevenueSplitCategory struct {
	Name string `json:"name,omitempty"`
}

// A set of fields to update a revenue split category.
//
// More info: see revenue split category
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/RevenueSplitCategory),
// the update operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplitcategories/update)
// and the revenue split categories endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplitcategories).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/UpdateRevenueSplitCategory).
type UpdateRevenueSplitCategory struct {
	Name string `json:"name,omitempty"`
}

// Set of parameters used to filter revenue splits.
//
// More info: see revenue split
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/RevenueSplit), the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplits/getlist)
// and the revenue splits endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplits).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/RevenueSplitParameters).
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

// An item in a list of revenue splits.
//
// This differs from the normal RevenueSplit type: not all fields are present in
// the list.
//
// More info: see revenue split
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/RevenueSplit), the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplits/getlist)
// and the revenue splits endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplits).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/ListRevenueSplit).
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

// A single revenue split.
//
// More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplits/get)
// and the revenue splits endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplits).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/RevenueSplit).
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

// A set of fields to create a revenue split.
//
// More info: see revenue split
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/RevenueSplit), the
// create operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplits/create)
// and the revenue splits endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplits).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/CreateRevenueSplit).
type CreateRevenueSplit struct {
	Name  string             `json:"name,omitempty"`
	Rules *RevenuesplitRules `json:"rules,omitempty"`
}

// A set of fields to update a revenue split.
//
// More info: see revenue split
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/RevenueSplit), the
// update operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplits/update)
// and the revenue splits endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplits).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/UpdateRevenueSplit).
type UpdateRevenueSplit struct {
	Name  string             `json:"name,omitempty"`
	Rules *RevenuesplitRules `json:"rules,omitempty"`
}

// Set of parameters used to filter ticket fees.
//
// More info: see ticket fee
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/TicketFee), the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_ticketfees/getlist)
// and the ticket fees endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_ticketfees).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/TicketFeeParameters).
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

// An item in a list of ticket fees.
//
// This differs from the normal TicketFee type: not all fields are present in the
// list.
//
// More info: see ticket fee
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/TicketFee), the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_ticketfees/getlist)
// and the ticket fees endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_ticketfees).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/ListTicketFee).
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

// A single ticket fee.
//
// More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_ticketfees/get)
// and the ticket fees endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_ticketfees).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/TicketFee).
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

// A set of fields to create a ticket fee.
//
// More info: see ticket fee
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/TicketFee), the create
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_ticketfees/create)
// and the ticket fees endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_ticketfees).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/CreateTicketFee).
type CreateTicketFee struct {
	Name  string          `json:"name,omitempty"`
	Rules *TicketfeeRules `json:"rules,omitempty"`
}

// A set of fields to update a ticket fee.
//
// More info: see ticket fee
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/TicketFee), the update
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_ticketfees/update)
// and the ticket fees endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_ticketfees).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/UpdateTicketFee).
type UpdateTicketFee struct {
	Name  string          `json:"name,omitempty"`
	Rules *TicketfeeRules `json:"rules,omitempty"`
}

// Set of parameters used to filter filter definitions.
//
// More info: see filter definition
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/FilterDefinition), the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_system_filterdefinitions/getlist)
// and the filter definitions endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_system_filterdefinitions).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/FilterDefinitionParameters).
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

// An item in a list of filter definitions.
//
// This differs from the normal FilterDefinition type: not all fields are present
// in the list.
//
// More info: see filter definition
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/FilterDefinition), the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_system_filterdefinitions/getlist)
// and the filter definitions endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_system_filterdefinitions).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/ListFilterDefinition).
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

// A single filter definition.
//
// More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_system_filterdefinitions/get)
// and the filter definitions endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_system_filterdefinitions).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/FilterDefinition).
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

// A set of fields to create a filter definition.
//
// More info: see filter definition
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/FilterDefinition), the
// create operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_system_filterdefinitions/create)
// and the filter definitions endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_system_filterdefinitions).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/CreateFilterDefinition).
type CreateFilterDefinition struct {
	// Type ID
	Typeid         int    `json:"typeid,omitempty"`
	Description    string `json:"description,omitempty"`
	Sqlclause      string `json:"sqlclause,omitempty"`
	Filtertype     int    `json:"filtertype,omitempty"`
	Checklistquery string `json:"checklistquery,omitempty"`
}

// A set of fields to update a filter definition.
//
// More info: see filter definition
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/FilterDefinition), the
// update operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_system_filterdefinitions/update)
// and the filter definitions endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_system_filterdefinitions).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/UpdateFilterDefinition).
type UpdateFilterDefinition struct {
	Description    string `json:"description,omitempty"`
	Sqlclause      string `json:"sqlclause,omitempty"`
	Filtertype     int    `json:"filtertype,omitempty"`
	Checklistquery string `json:"checklistquery,omitempty"`
}

// Set of parameters used to filter delivery scenarios.
//
// More info: see delivery scenario
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/DeliveryScenario), the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_deliveryscenarios/getlist)
// and the delivery scenarios endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_deliveryscenarios).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/DeliveryScenarioParameters).
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

// An item in a list of delivery scenarios.
//
// This differs from the normal DeliveryScenario type: not all fields are present
// in the list.
//
// More info: see delivery scenario
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/DeliveryScenario), the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_deliveryscenarios/getlist)
// and the delivery scenarios endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_deliveryscenarios).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/ListDeliveryScenario).
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

// A single delivery scenario.
//
// More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_deliveryscenarios/get)
// and the delivery scenarios endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_deliveryscenarios).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/DeliveryScenario).
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

// A set of fields to create a delivery scenario.
//
// More info: see delivery scenario
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/DeliveryScenario), the
// create operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_deliveryscenarios/create)
// and the delivery scenarios endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_deliveryscenarios).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/CreateDeliveryScenario).
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

// A set of fields to update a delivery scenario.
//
// More info: see delivery scenario
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/DeliveryScenario), the
// update operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_deliveryscenarios/update)
// and the delivery scenarios endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_deliveryscenarios).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/UpdateDeliveryScenario).
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

// Set of parameters used to filter lock types.
//
// More info: see lock type
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/LockType), the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_locktypes/getlist)
// and the lock types endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_locktypes).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/LockTypeParameters).
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

// An item in a list of lock types.
//
// This differs from the normal LockType type: not all fields are present in the
// list.
//
// More info: see lock type
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/LockType), the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_locktypes/getlist)
// and the lock types endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_locktypes).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/ListLockType).
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

// A single lock type.
//
// More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_locktypes/get)
// and the lock types endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_locktypes).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/LockType).
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

// A set of fields to create a lock type.
//
// More info: see lock type
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/LockType), the create
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_locktypes/create)
// and the lock types endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_locktypes).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/CreateLockType).
type CreateLockType struct {
	Name       string `json:"name,omitempty"`
	Ishardlock bool   `json:"ishardlock,omitempty"`
}

// A set of fields to update a lock type.
//
// More info: see lock type
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/LockType), the update
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_locktypes/update)
// and the lock types endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_locktypes).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/UpdateLockType).
type UpdateLockType struct {
	Name       string `json:"name,omitempty"`
	Ishardlock bool   `json:"ishardlock,omitempty"`
}

// Set of parameters used to filter order fees.
//
// More info: see order fee
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/OrderFee), the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_orderfees/getlist)
// and the order fees endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_orderfees).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/OrderFeeParameters).
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

// An item in a list of order fees.
//
// This differs from the normal OrderFee type: not all fields are present in the
// list.
//
// More info: see order fee
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/OrderFee), the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_orderfees/getlist)
// and the order fees endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_orderfees).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/ListOrderFee).
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

// A single order fee.
//
// More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_orderfees/get)
// and the order fees endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_orderfees).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/OrderFee).
type OrderFee struct {
	// Unique ID
	Id     int           `json:"id,omitempty"`
	Name   string        `json:"name,omitempty"`
	Typeid int           `json:"typeid,omitempty"`
	Rule   *OrderFeeRule `json:"rule,omitempty"`

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

// A set of fields to create a order fee.
//
// More info: see order fee
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/OrderFee), the create
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_orderfees/create)
// and the order fees endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_orderfees).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/CreateOrderFee).
type CreateOrderFee struct {
	Name   string        `json:"name,omitempty"`
	Typeid int           `json:"typeid,omitempty"`
	Rule   *OrderFeeRule `json:"rule,omitempty"`
}

// A set of fields to update a order fee.
//
// More info: see order fee
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/OrderFee), the update
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_orderfees/update)
// and the order fees endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_orderfees).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/UpdateOrderFee).
type UpdateOrderFee struct {
	Name   string        `json:"name,omitempty"`
	Typeid int           `json:"typeid,omitempty"`
	Rule   *OrderFeeRule `json:"rule,omitempty"`
}

// Set of parameters used to filter payment methods.
//
// More info: see payment method
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PaymentMethod), the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentmethods/getlist)
// and the payment methods endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentmethods).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PaymentMethodParameters).
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

// An item in a list of payment methods.
//
// This differs from the normal PaymentMethod type: not all fields are present in
// the list.
//
// More info: see payment method
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PaymentMethod), the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentmethods/getlist)
// and the payment methods endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentmethods).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/ListPaymentMethod).
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

// A single payment method.
//
// More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentmethods/get)
// and the payment methods endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentmethods).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PaymentMethod).
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

// A set of fields to create a payment method.
//
// More info: see payment method
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PaymentMethod), the
// create operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentmethods/create)
// and the payment methods endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentmethods).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/CreatePaymentMethod).
type CreatePaymentMethod struct {
	Name                    string               `json:"name,omitempty"`
	Internalremark          string               `json:"internalremark,omitempty"`
	Paymentmethodtypeid     int                  `json:"paymentmethodtypeid,omitempty"`
	Paymentmethodreceiverid int                  `json:"paymentmethodreceiverid,omitempty"`
	Config                  *PaymentmethodConfig `json:"config,omitempty"`
}

// A set of fields to update a payment method.
//
// More info: see payment method
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PaymentMethod), the
// update operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentmethods/update)
// and the payment methods endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentmethods).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/UpdatePaymentMethod).
type UpdatePaymentMethod struct {
	Name                    string               `json:"name,omitempty"`
	Internalremark          string               `json:"internalremark,omitempty"`
	Paymentmethodtypeid     int                  `json:"paymentmethodtypeid,omitempty"`
	Paymentmethodreceiverid int                  `json:"paymentmethodreceiverid,omitempty"`
	Config                  *PaymentmethodConfig `json:"config,omitempty"`
}

// Set of parameters used to filter payment scenarios.
//
// More info: see payment scenario
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PaymentScenario), the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentscenarios/getlist)
// and the payment scenarios endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentscenarios).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PaymentScenarioParameters).
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

// An item in a list of payment scenarios.
//
// This differs from the normal PaymentScenario type: not all fields are present in
// the list.
//
// More info: see payment scenario
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PaymentScenario), the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentscenarios/getlist)
// and the payment scenarios endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentscenarios).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/ListPaymentScenario).
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

// A single payment scenario.
//
// More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentscenarios/get)
// and the payment scenarios endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentscenarios).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PaymentScenario).
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

// A set of fields to create a payment scenario.
//
// More info: see payment scenario
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PaymentScenario), the
// create operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentscenarios/create)
// and the payment scenarios endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentscenarios).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/CreatePaymentScenario).
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

// A set of fields to update a payment scenario.
//
// More info: see payment scenario
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PaymentScenario), the
// update operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentscenarios/update)
// and the payment scenarios endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentscenarios).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/UpdatePaymentScenario).
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

// Set of parameters used to filter sales channels.
//
// More info: see sales channel
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/SalesChannel), the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_saleschannels/getlist)
// and the sales channels endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_saleschannels).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/SalesChannelParameters).
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

// An item in a list of sales channels.
//
// This differs from the normal SalesChannel type: not all fields are present in
// the list.
//
// More info: see sales channel
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/SalesChannel), the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_saleschannels/getlist)
// and the sales channels endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_saleschannels).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/ListSalesChannel).
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

// A single sales channel.
//
// More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_saleschannels/get)
// and the sales channels endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_saleschannels).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/SalesChannel).
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

// A set of fields to create a sales channel.
//
// More info: see sales channel
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/SalesChannel), the
// create operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_saleschannels/create)
// and the sales channels endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_saleschannels).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/CreateSalesChannel).
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

// A set of fields to update a sales channel.
//
// More info: see sales channel
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/SalesChannel), the
// update operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_saleschannels/update)
// and the sales channels endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_saleschannels).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/UpdateSalesChannel).
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
