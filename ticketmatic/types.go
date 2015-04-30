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
	// An array of sales channel IDs for which this delivery scenario can be used
	Saleschannels []int64 `json:"saleschannels,omitempty"`

	// Use a script to refine the set of sales channels?
	Usescript bool `json:"usescript,omitempty"`

	// Script used to determine availability of the delivery scenario
	Script string `json:"script,omitempty"`
}

type Event struct {
	// Event ID
	Id int64 `json:"id,omitempty"`

	// Event name
	Name string `json:"name,omitempty"`

	// Event subtitle
	Subtitle string `json:"subtitle,omitempty"`

	// Event subtitle (2)
	Subtitle2 string `json:"subtitle2,omitempty"`

	// Small description that will be shown on the sales pages of this event
	Webremark string `json:"webremark,omitempty"`

	// Event start time
	Startts Time `json:"startts,omitempty"`

	// Time of start of sales.
	//
	// Used for all sales channels for which no specific sales period has been defined.
	Salestartts Time `json:"salestartts,omitempty"`

	// Time of end of sales.
	//
	// Used for all sales channels for which no specific sales period has been defined.
	Saleendts Time `json:"saleendts,omitempty"`

	// Event publish time
	Publishedts Time `json:"publishedts,omitempty"`

	// Event end time
	Endts Time `json:"endts,omitempty"`

	// Event code.
	//
	// Used as a unique identifier in web sales.
	Code string `json:"code,omitempty"`

	// External event code.
	//
	// This field is typically set when importing data from other systems.
	Externalcode string `json:"externalcode,omitempty"`

	// Production ID
	Productionid int64 `json:"productionid,omitempty"`

	// Event location ID
	//
	// See event locations
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_events_eventlocations)
	// for more info.
	Locationid int64 `json:"locationid,omitempty"`

	// Event location name
	//
	// Automatically derived using locationid.
	Locationname string `json:"locationname,omitempty"`

	// Seating plan ID
	//
	// Only set for events with fixed seats. See seating plans
	// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Seatingplan) for more
	// info.
	Seatingplanid int64 `json:"seatingplanid,omitempty"`

	// Price list ID for fixed seats.
	//
	// Only set for events with fixed seats. See price lists
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricelists)
	// for more info.
	Seatingplanpricelistid         int64       `json:"seatingplanpricelistid,omitempty"`
	Seatingplaneventspecificprices interface{} `json:"seatingplaneventspecificprices,omitempty"`
	Seatingplancontingents         interface{} `json:"seatingplancontingents,omitempty"`
	Contingents                    interface{} `json:"contingents,omitempty"`

	// Price availability ID
	//
	// Determines which price types are available for this event. See price
	// availabilities
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_priceavailabilities)
	// for more info.
	Priceavailabilityid int64 `json:"priceavailabilityid,omitempty"`

	// Ticket fee ID
	//
	// Determines which ticket fee rules are used for this event. See ticket fees
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_ticketfees)
	// for more info.
	Ticketfeeid int64 `json:"ticketfeeid,omitempty"`

	// Revenue split ID
	//
	// Determines which revenue split rules are used for this event. See revenue splits
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_revenuesplits)
	// for more info.
	Revenuesplitid int64 `json:"revenuesplitid,omitempty"`

	// Ticket layout ID
	//
	// See ticket layouts
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ticketlayouts)
	// for more info.
	Ticketlayoutid int64 `json:"ticketlayoutid,omitempty"`

	// Maximum number of tickets for this event that can be added to a basket
	Maxnbrofticketsperbasket int64 `json:"maxnbrofticketsperbasket,omitempty"`

	// Event status
	Currentstatus int64       `json:"currentstatus,omitempty"`
	Prices        interface{} `json:"prices,omitempty"`
	Saleschannels interface{} `json:"saleschannels,omitempty"`
	Tags          interface{} `json:"tags,omitempty"`
	Availability  interface{} `json:"availability,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`
}

type Order struct {
	// Order ID
	Orderid int64 `json:"orderid,omitempty"`

	// Order status
	Status int64 `json:"status,omitempty"`

	// Order code.
	//
	// Used as a unique identifier in web sales.
	Code string `json:"code,omitempty"`

	// Customer ID
	Customerid int64 `json:"customerid,omitempty"`

	// Has customer authenticated?
	Isauthenticatedcustomer bool `json:"isauthenticatedcustomer,omitempty"`

	// Total order amount.
	//
	// Includes all costs and fees.
	Totalamount float64 `json:"totalamount,omitempty"`

	// Total amount paid.
	Amountpaid                float64     `json:"amountpaid,omitempty"`
	Paymentstatus             int64       `json:"paymentstatus,omitempty"`
	Deliverystatus            int64       `json:"deliverystatus,omitempty"`
	Deliveryaddress           interface{} `json:"deliveryaddress,omitempty"`
	Deferredpaymentproperties interface{} `json:"deferredpaymentproperties,omitempty"`

	// Sales channel ID
	//
	// See sales channels
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_saleschannels)
	// for more info.
	Saleschannelid int64 `json:"saleschannelid,omitempty"`

	// Payment scenario ID
	//
	// See payment scenarios
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_paymentscenarios)
	// for more info.
	Paymentscenarioid int64 `json:"paymentscenarioid,omitempty"`

	// Delivery scenario ID
	//
	// See delivery scenarios
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_deliveryscenarios)
	// for more info.
	Deliveryscenarioid int64 `json:"deliveryscenarioid,omitempty"`

	// When the reminder mail will be sent
	Rappelts Time `json:"rappelts,omitempty"`

	// Whether the reminder mail has been sent
	Rappelsent bool `json:"rappelsent,omitempty"`

	// When the order will expire
	Expiryts   Time        `json:"expiryts,omitempty"`
	Tickets    interface{} `json:"tickets,omitempty"`
	Payments   interface{} `json:"payments,omitempty"`
	Lookup     interface{} `json:"lookup,omitempty"`
	Ordercosts interface{} `json:"ordercosts,omitempty"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`
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

// This type isn't documented yet
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/TODO).
type TODO struct {
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

type EventQuery struct {
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool   `json:"includearchived,omitempty"`
	Lastupdatesince Time   `json:"lastupdatesince,omitempty"`
	Limit           int64  `json:"limit,omitempty"`
	Offset          int64  `json:"offset,omitempty"`
	Orderby         string `json:"orderby,omitempty"`
	Output          string `json:"output,omitempty"`
	Searchterm      string `json:"searchterm,omitempty"`
	Simplefilter    string `json:"simplefilter,omitempty"`
	Context         string `json:"context,omitempty"`
}

type Ticket struct {
}

// Required data for creating an order
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Order).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/CreateOrder).
type CreateOrder struct {
	// Sales channel in which this order is created
	Saleschannelid int64 `json:"saleschannelid,omitempty"`
}

// Info for adding a ticket
// (https://apps.ticketmatic.com/#/knowledgebase/api/orders/addtickets) to an order
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Order).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/CreateTicket).
type CreateTicket struct {
	// The required ticket type price ID.
	Tickettypepriceid int64 `json:"tickettypepriceid,omitempty"`
}

// Request data used to add tickets
// (https://apps.ticketmatic.com/#/knowledgebase/api/orders/addtickets) to an order
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Order).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/AddTickets).
type AddTickets struct {
	// Ticket information
	Tickets []*CreateTicket `json:"tickets,omitempty"`
}

// Result when adding tickets
// (https://apps.ticketmatic.com/#/knowledgebase/api/orders/addtickets) to an order
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Order).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/AddTicketsResult).
type AddTicketsResult struct {
	// Number of tickets added
	Nbrofaddedtickets int64 `json:"nbrofaddedtickets,omitempty"`

	// The modified order
	Order *Order `json:"order,omitempty"`
}

type UpdateTickets struct {
}

type UpdateOrder struct {
	Deliveryscenarioid int64       `json:"deliveryscenarioid,omitempty"`
	Paymentscenarioid  int64       `json:"paymentscenarioid,omitempty"`
	Customerid         int64       `json:"customerid,omitempty"`
	Customfields       interface{} `json:"customfields,omitempty"`
}

type OrderQuery struct {
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool   `json:"includearchived,omitempty"`
	Lastupdatesince Time   `json:"lastupdatesince,omitempty"`
	Limit           int64  `json:"limit,omitempty"`
	Offset          int64  `json:"offset,omitempty"`
	Orderby         string `json:"orderby,omitempty"`
	Output          string `json:"output,omitempty"`
	Searchterm      string `json:"searchterm,omitempty"`
	Simplefilter    string `json:"simplefilter,omitempty"`
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
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/OrderMailTemplateQuery).
type OrderMailTemplateQuery struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
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
	//
	// Note: Ignored when creating a new order mail template.
	//
	// Note: Ignored when updating an existing order mail template.
	Id int64 `json:"id,omitempty"`

	// Name of the order mail template
	Name string `json:"name,omitempty"`

	// The type of this order mail template, defines where this template is used. The
	// available values for this field can be found on the order mail template overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ordermails)
	// page.
	Typeid int64 `json:"typeid,omitempty"`

	// Subject line for the order mail template
	//
	// Note: Not set when retrieving a list of order mail templates.
	Subject string `json:"subject,omitempty"`

	// Message body
	//
	// Note: Not set when retrieving a list of order mail templates.
	Body string `json:"body,omitempty"`

	// A map of language codes to gettext .po files
	// (http://en.wikipedia.org/wiki/Gettext). More info can be found on the order mail
	// template overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ordermails)
	// page.
	//
	// Note: Not set when retrieving a list of order mail templates.
	Translations map[string]string `json:"translations,omitempty"`

	// Created timestamp
	//
	// Note: Ignored when creating a new order mail template.
	//
	// Note: Ignored when updating an existing order mail template.
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new order mail template.
	//
	// Note: Ignored when updating an existing order mail template.
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new order mail template.
	//
	// Note: Ignored when updating an existing order mail template.
	Isarchived bool `json:"isarchived,omitempty"`
}

// Set of parameters used to filter ticket layouts.
//
// More info: see ticket layout
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/TicketLayout), the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ticketlayouts/getlist)
// and the ticket layouts endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ticketlayouts).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/TicketLayoutQuery).
type TicketLayoutQuery struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
}

// A single ticket layout.
//
// More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ticketlayouts/get)
// and the ticket layouts endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ticketlayouts).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/TicketLayout).
type TicketLayout struct {
	// Unique ID
	//
	// Note: Ignored when creating a new ticket layout.
	//
	// Note: Ignored when updating an existing ticket layout.
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`

	// Created timestamp
	//
	// Note: Ignored when creating a new ticket layout.
	//
	// Note: Ignored when updating an existing ticket layout.
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new ticket layout.
	//
	// Note: Ignored when updating an existing ticket layout.
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new ticket layout.
	//
	// Note: Ignored when updating an existing ticket layout.
	Isarchived bool `json:"isarchived,omitempty"`
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
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/WebSalesSkinQuery).
type WebSalesSkinQuery struct {
	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
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
	//
	// Note: Ignored when creating a new web sales skin.
	//
	// Note: Ignored when updating an existing web sales skin.
	Id int64 `json:"id,omitempty"`

	// Name of the web sales skin
	Name string `json:"name,omitempty"`

	// HTML template of the skin. See the web skin setup guide
	// (https://apps.ticketmatic.com/#/knowledgebase/designer_webskin) for more
	// information.
	//
	// Note: Not set when retrieving a list of web sales skins.
	Html string `json:"html,omitempty"`

	// CSS style rules. Should always include the style import.
	//
	// Note: Not set when retrieving a list of web sales skins.
	Css string `json:"css,omitempty"`

	// A map of language codes to gettext .po files
	// (http://en.wikipedia.org/wiki/Gettext). More info can be found on the web skin
	// overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_webskins)
	// page.
	//
	// Note: Not set when retrieving a list of web sales skins.
	Translations map[string]string `json:"translations,omitempty"`

	// Skin configuration.
	//
	// See the WebSalesSkinConfiguration reference
	// (https://apps.ticketmatic.com/#/knowledgebase/api/types/WebSalesSkinConfiguration)
	// for an overview of all possible options.
	//
	// Note: Not set when retrieving a list of web sales skins.
	Configuration *WebSalesSkinConfiguration `json:"configuration,omitempty"`

	// Created timestamp
	//
	// Note: Ignored when creating a new web sales skin.
	//
	// Note: Ignored when updating an existing web sales skin.
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new web sales skin.
	//
	// Note: Ignored when updating an existing web sales skin.
	Lastupdatets Time `json:"lastupdatets,omitempty"`
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
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/EventLocationQuery).
type EventLocationQuery struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
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
	//
	// Note: Ignored when creating a new event location.
	//
	// Note: Ignored when updating an existing event location.
	Id int64 `json:"id,omitempty"`

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
	//
	// Note: Ignored when creating a new event location.
	//
	// Note: Ignored when updating an existing event location.
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new event location.
	//
	// Note: Ignored when updating an existing event location.
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new event location.
	//
	// Note: Ignored when updating an existing event location.
	Isarchived bool `json:"isarchived,omitempty"`
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
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PriceAvailabilityQuery).
type PriceAvailabilityQuery struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
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
	//
	// Note: Ignored when creating a new price availability.
	//
	// Note: Ignored when updating an existing price availability.
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`

	// Note: Not set when retrieving a list of price availabilities.
	Rules *PriceAvailabilityRules `json:"rules,omitempty"`

	// Created timestamp
	//
	// Note: Ignored when creating a new price availability.
	//
	// Note: Ignored when updating an existing price availability.
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new price availability.
	//
	// Note: Ignored when updating an existing price availability.
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new price availability.
	//
	// Note: Ignored when updating an existing price availability.
	Isarchived bool `json:"isarchived,omitempty"`
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
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PriceListQuery).
type PriceListQuery struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
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
	//
	// Note: Ignored when creating a new price list.
	//
	// Note: Ignored when updating an existing price list.
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`

	// Note: Not set when retrieving a list of price lists.
	Prices   *PricelistPrices `json:"prices,omitempty"`
	Hasranks bool             `json:"hasranks,omitempty"`

	// Created timestamp
	//
	// Note: Ignored when creating a new price list.
	//
	// Note: Ignored when updating an existing price list.
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new price list.
	//
	// Note: Ignored when updating an existing price list.
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new price list.
	//
	// Note: Ignored when updating an existing price list.
	Isarchived bool `json:"isarchived,omitempty"`
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
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PriceTypeQuery).
type PriceTypeQuery struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
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
	//
	// Note: Ignored when creating a new price type.
	//
	// Note: Ignored when updating an existing price type.
	Id int64 `json:"id,omitempty"`

	// Name of the price type
	Name string `json:"name,omitempty"`

	// The category of this price type, defines how the price is displayed. The
	// available values for this field can be found on the price type overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricetypes)
	// page.
	Typeid int64 `json:"typeid,omitempty"`

	// A remark that describes the price type. Will be shown to customers.
	Remark string `json:"remark,omitempty"`

	// Created timestamp
	//
	// Note: Ignored when creating a new price type.
	//
	// Note: Ignored when updating an existing price type.
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new price type.
	//
	// Note: Ignored when updating an existing price type.
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new price type.
	//
	// Note: Ignored when updating an existing price type.
	Isarchived bool `json:"isarchived,omitempty"`
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
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/RevenueSplitCategoryQuery).
type RevenueSplitCategoryQuery struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
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
	//
	// Note: Ignored when creating a new revenue split category.
	//
	// Note: Ignored when updating an existing revenue split category.
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`

	// Created timestamp
	//
	// Note: Ignored when creating a new revenue split category.
	//
	// Note: Ignored when updating an existing revenue split category.
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new revenue split category.
	//
	// Note: Ignored when updating an existing revenue split category.
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new revenue split category.
	//
	// Note: Ignored when updating an existing revenue split category.
	Isarchived bool `json:"isarchived,omitempty"`
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
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/RevenueSplitQuery).
type RevenueSplitQuery struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
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
	//
	// Note: Ignored when creating a new revenue split.
	//
	// Note: Ignored when updating an existing revenue split.
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`

	// Note: Not set when retrieving a list of revenue splits.
	Rules *RevenuesplitRules `json:"rules,omitempty"`

	// Created timestamp
	//
	// Note: Ignored when creating a new revenue split.
	//
	// Note: Ignored when updating an existing revenue split.
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new revenue split.
	//
	// Note: Ignored when updating an existing revenue split.
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new revenue split.
	//
	// Note: Ignored when updating an existing revenue split.
	Isarchived bool `json:"isarchived,omitempty"`
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
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/TicketFeeQuery).
type TicketFeeQuery struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
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
	//
	// Note: Ignored when creating a new ticket fee.
	//
	// Note: Ignored when updating an existing ticket fee.
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`

	// Note: Not set when retrieving a list of ticket fees.
	Rules *TicketfeeRules `json:"rules,omitempty"`

	// Created timestamp
	//
	// Note: Ignored when creating a new ticket fee.
	//
	// Note: Ignored when updating an existing ticket fee.
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new ticket fee.
	//
	// Note: Ignored when updating an existing ticket fee.
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new ticket fee.
	//
	// Note: Ignored when updating an existing ticket fee.
	Isarchived bool `json:"isarchived,omitempty"`
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
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/FilterDefinitionQuery).
type FilterDefinitionQuery struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// Only return items with the given typeid.
	Typeid int64 `json:"typeid,omitempty"`
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
	//
	// Note: Ignored when creating a new filter definition.
	//
	// Note: Ignored when updating an existing filter definition.
	Id int64 `json:"id,omitempty"`

	// Type ID
	//
	// Note: Ignored when updating an existing filter definition.
	Typeid         int64  `json:"typeid,omitempty"`
	Description    string `json:"description,omitempty"`
	Sqlclause      string `json:"sqlclause,omitempty"`
	Filtertype     int64  `json:"filtertype,omitempty"`
	Checklistquery string `json:"checklistquery,omitempty"`

	// Created timestamp
	//
	// Note: Ignored when creating a new filter definition.
	//
	// Note: Ignored when updating an existing filter definition.
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new filter definition.
	//
	// Note: Ignored when updating an existing filter definition.
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new filter definition.
	//
	// Note: Ignored when updating an existing filter definition.
	Isarchived bool `json:"isarchived,omitempty"`
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
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/DeliveryScenarioQuery).
type DeliveryScenarioQuery struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
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
	//
	// Note: Ignored when creating a new delivery scenario.
	//
	// Note: Ignored when updating an existing delivery scenario.
	Id int64 `json:"id,omitempty"`

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
	Typeid int64 `json:"typeid,omitempty"`

	// A physical address is required
	Needsaddress bool `json:"needsaddress,omitempty"`

	// The rules that define when this scenario is available. See the delivery scenario
	// overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_deliveryscenarios)
	// page for a description of this field
	//
	// Note: Not set when retrieving a list of delivery scenarios.
	Availability *DeliveryscenarioAvailability `json:"availability,omitempty"`

	// The ID of the order mail template that will be used for sending out this
	// delivery scenario. Can be 0 to indicate that no mail should be sent
	OrdermailtemplateidDelivery int64 `json:"ordermailtemplateid_delivery,omitempty"`

	// Are e-tickets allowed with this delivery scenario?
	Allowetickets int64 `json:"allowetickets,omitempty"`

	// Created timestamp
	//
	// Note: Ignored when creating a new delivery scenario.
	//
	// Note: Ignored when updating an existing delivery scenario.
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new delivery scenario.
	//
	// Note: Ignored when updating an existing delivery scenario.
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new delivery scenario.
	//
	// Note: Ignored when updating an existing delivery scenario.
	Isarchived bool `json:"isarchived,omitempty"`
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
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/LockTypeQuery).
type LockTypeQuery struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
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
	//
	// Note: Ignored when creating a new lock type.
	//
	// Note: Ignored when updating an existing lock type.
	Id         int64  `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Ishardlock bool   `json:"ishardlock,omitempty"`

	// Created timestamp
	//
	// Note: Ignored when creating a new lock type.
	//
	// Note: Ignored when updating an existing lock type.
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new lock type.
	//
	// Note: Ignored when updating an existing lock type.
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new lock type.
	//
	// Note: Ignored when updating an existing lock type.
	Isarchived bool `json:"isarchived,omitempty"`
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
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/OrderFeeQuery).
type OrderFeeQuery struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
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
	//
	// Note: Ignored when creating a new order fee.
	//
	// Note: Ignored when updating an existing order fee.
	Id     int64  `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Typeid int64  `json:"typeid,omitempty"`

	// Note: Not set when retrieving a list of order fees.
	Rule *OrderFeeRule `json:"rule,omitempty"`

	// Created timestamp
	//
	// Note: Ignored when creating a new order fee.
	//
	// Note: Ignored when updating an existing order fee.
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new order fee.
	//
	// Note: Ignored when updating an existing order fee.
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new order fee.
	//
	// Note: Ignored when updating an existing order fee.
	Isarchived bool `json:"isarchived,omitempty"`
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
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PaymentMethodQuery).
type PaymentMethodQuery struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
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
	//
	// Note: Ignored when creating a new payment method.
	//
	// Note: Ignored when updating an existing payment method.
	Id                      int64  `json:"id,omitempty"`
	Name                    string `json:"name,omitempty"`
	Internalremark          string `json:"internalremark,omitempty"`
	Paymentmethodtypeid     int64  `json:"paymentmethodtypeid,omitempty"`
	Paymentmethodreceiverid int64  `json:"paymentmethodreceiverid,omitempty"`

	// Note: Not set when retrieving a list of payment methods.
	Config *PaymentmethodConfig `json:"config,omitempty"`

	// Created timestamp
	//
	// Note: Ignored when creating a new payment method.
	//
	// Note: Ignored when updating an existing payment method.
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new payment method.
	//
	// Note: Ignored when updating an existing payment method.
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new payment method.
	//
	// Note: Ignored when updating an existing payment method.
	Isarchived bool `json:"isarchived,omitempty"`
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
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PaymentScenarioQuery).
type PaymentScenarioQuery struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
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
	//
	// Note: Ignored when creating a new payment scenario.
	//
	// Note: Ignored when updating an existing payment scenario.
	Id int64 `json:"id,omitempty"`

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
	Internalremark string `json:"internalremark,omitempty"`
	Typeid         int64  `json:"typeid,omitempty"`

	// Note: Not set when retrieving a list of payment scenarios.
	Overdueparameters *PaymentscenarioOverdueParameters `json:"overdueparameters,omitempty"`

	// Note: Not set when retrieving a list of payment scenarios.
	Expiryparameters *PaymentscenarioExpiryParameters `json:"expiryparameters,omitempty"`

	// Note: Not set when retrieving a list of payment scenarios.
	Availability                          *PaymentscenarioAvailability `json:"availability,omitempty"`
	Paymentmethods                        []int64                      `json:"paymentmethods,omitempty"`
	OrdermailtemplateidPaymentinstruction int64                        `json:"ordermailtemplateid_paymentinstruction,omitempty"`
	OrdermailtemplateidOverdue            int64                        `json:"ordermailtemplateid_overdue,omitempty"`
	OrdermailtemplateidExpiry             int64                        `json:"ordermailtemplateid_expiry,omitempty"`

	// Created timestamp
	//
	// Note: Ignored when creating a new payment scenario.
	//
	// Note: Ignored when updating an existing payment scenario.
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new payment scenario.
	//
	// Note: Ignored when updating an existing payment scenario.
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new payment scenario.
	//
	// Note: Ignored when updating an existing payment scenario.
	Isarchived bool `json:"isarchived,omitempty"`
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
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/SalesChannelQuery).
type SalesChannelQuery struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
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
	//
	// Note: Ignored when creating a new sales channel.
	//
	// Note: Ignored when updating an existing sales channel.
	Id int64 `json:"id,omitempty"`

	// Name of the sales channel
	Name string `json:"name,omitempty"`

	// The type of this sales channel, defines where this sales channel will be used.
	// The available values for this field can be found on the sales channel overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_saleschannels)
	// page.
	Typeid int64 `json:"typeid,omitempty"`

	// The ID of the order mail template to use for sending confirmations
	OrdermailtemplateidConfirmation int64 `json:"ordermailtemplateid_confirmation,omitempty"`

	// Always send the confirmation, regardless of the payment method configuration
	OrdermailtemplateidConfirmationSendalways bool `json:"ordermailtemplateid_confirmation_sendalways,omitempty"`

	// Created timestamp
	//
	// Note: Ignored when creating a new sales channel.
	//
	// Note: Ignored when updating an existing sales channel.
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new sales channel.
	//
	// Note: Ignored when updating an existing sales channel.
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new sales channel.
	//
	// Note: Ignored when updating an existing sales channel.
	Isarchived bool `json:"isarchived,omitempty"`
}

type SubscriberSync struct {
	Email      string `json:"email,omitempty"`
	Firstname  string `json:"firstname,omitempty"`
	Lastname   string `json:"lastname,omitempty"`
	Subscribed bool   `json:"subscribed,omitempty"`
}
