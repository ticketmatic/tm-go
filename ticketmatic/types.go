package ticketmatic

import (
	"encoding/json"
	"strings"
)

// Address, used for physical deliveries and contact details.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Address).
type Address struct {
	// Address ID
	//
	// Note: Only available when used for a contact
	// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Contact) address.
	Id int64 `json:"id,omitempty"`

	// Contact this address belongs to
	//
	// Note: Only available when used for a contact
	// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Contact) address.
	Customerid int64 `json:"customerid,omitempty"`

	// Addressee
	//
	// Note: Only available when used as an order
	// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Order) deliver address.
	Addressee string `json:"addressee,omitempty"`

	// ISO 3166-1 alpha-2 country code
	// (http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
	Countrycode string `json:"countrycode,omitempty"`

	// Country name (based on typeid, returned as a convenience).
	//
	// Note: Only available when used for a contact
	// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Contact) address.
	Country string `json:"country,omitempty"`

	// Zip code
	Zip string `json:"zip,omitempty"`

	// City
	City string `json:"city,omitempty"`

	// Street field 1 (typically the street name)
	Street1 string `json:"street1,omitempty"`

	// Street field 2 (typically the number)
	Street2 string `json:"street2,omitempty"`

	// Street field 3 (sometimes used for box numbers or suffixes)
	Street3 string `json:"street3,omitempty"`

	// Street field 4 (rarely used)
	Street4 string `json:"street4,omitempty"`

	// Address type ID
	//
	// Note: Only available when used for a contact
	// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Contact) address.
	Typeid int64 `json:"typeid,omitempty"`

	// Address type (based on typeid, returned as a convenience).
	//
	// Note: Only available when used for a contact
	// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Contact) address.
	Type string `json:"type,omitempty"`
}

// A single contact.
//
// More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/contacts/get) and the contacts
// endpoint (https://apps.ticketmatic.com/#/knowledgebase/api/contacts).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Contact).
type Contact struct {
	// Contact ID
	//
	// Note: Ignored when creating a new contact.
	//
	// Note: Ignored when updating an existing contact.
	Id int64 `json:"id,omitempty"`

	// E-mail address
	Email string `json:"email,omitempty"`

	// Customer title ID (also determines the gender of the contact)
	Customertitleid int64 `json:"customertitleid,omitempty"`

	// First name
	Firstname string `json:"firstname,omitempty"`

	// Middle name
	Middlename string `json:"middlename,omitempty"`

	// Last name
	Lastname string `json:"lastname,omitempty"`

	// Language (ISO 639-1 code (http://en.wikipedia.org/wiki/List_of_ISO_639-1_codes))
	Languagecode string `json:"languagecode,omitempty"`

	// Sex
	Sex string `json:"sex,omitempty"`

	// Birth date
	Birthdate Time `json:"birthdate,omitempty"`

	// Company
	Company string `json:"company,omitempty"`

	// Job function
	Organizationfunction string `json:"organizationfunction,omitempty"`

	// Addresses
	Addresses []*Address `json:"addresses"`

	// VAT Number (for organizations)
	Vatnumber string `json:"vatnumber,omitempty"`

	// Phone numbers
	Phonenumbers []*Phonenumber `json:"phonenumbers"`

	// Relation type IDs
	Relationtypes []int64 `json:"relationtypes"`

	// Whether or not this contact is subscribed in the e-mail marketing integration
	//
	// Note: Ignored when creating a new contact.
	//
	// Note: Ignored when updating an existing contact.
	Subscribed bool `json:"subscribed,omitempty"`

	// Contact status
	//
	// Possible values:
	//
	// * deleted: Contact has been deleted
	//
	// * incomplete: Contact misses crucial account information
	//
	// * (blank): Normal contact
	//
	// Note: Ignored when creating a new contact.
	//
	// Note: Ignored when updating an existing contact.
	Status string `json:"status,omitempty"`

	// Account type.
	//
	// Indicates the authentication type supported for this contact (used when
	// authentication is enabled in web sales).
	//
	// Possible values:
	//
	// * 0: No authentication
	//
	// * 1901: Password authentication
	//
	// * 1902: Facebook
	//
	// * 1903: Google
	//
	// * 1904: Twitter
	//
	// Note: Ignored when creating a new contact.
	//
	// Note: Ignored when updating an existing contact.
	AccountType int64 `json:"account_type,omitempty"`

	// Whether or not this contact has been deleted
	//
	// Note: Ignored when creating a new contact.
	//
	// Note: Ignored when updating an existing contact.
	Isdeleted bool `json:"isdeleted,omitempty"`

	// Created timestamp
	//
	// Note: Ignored when creating a new contact.
	//
	// Note: Ignored when updating an existing contact.
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new contact.
	//
	// Note: Ignored when updating an existing contact.
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Custom fields
	CustomFields map[string]interface{} `json:"-"`
}

// Custom unmarshaller with support for custom fields
func (o *Contact) UnmarshalJSON(data []byte) error {
	// Alias the type, to avoid calling UnmarshalJSON. Unpack it.
	type tmp Contact
	var obj tmp
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}

	*o = Contact(obj)

	// Unpack it again, this time to a map, so we can pull out the custom fields.
	var raw map[string]interface{}
	err = json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	o.CustomFields = make(map[string]interface{})
	for key, val := range raw {
		if strings.HasPrefix(key, "c_") {
			o.CustomFields[key[2:]] = val
		}
	}

	// Note: We're doing a double JSON decode here, I'd love to get rid of it
	// but I'm not sure how we can do this easily. Suggestions welcome:
	// developers@ticketmatic.com!

	return nil
}

// Custom marshaller with support for custom fields
func (o *Contact) MarshalJSON() ([]byte, error) {
	// Use a custom type to avoid the custom marshaller, marshal the data.
	type tmp struct {
		Id                   int64          `json:"id,omitempty"`
		Email                string         `json:"email,omitempty"`
		Customertitleid      int64          `json:"customertitleid,omitempty"`
		Firstname            string         `json:"firstname,omitempty"`
		Middlename           string         `json:"middlename,omitempty"`
		Lastname             string         `json:"lastname,omitempty"`
		Languagecode         string         `json:"languagecode,omitempty"`
		Sex                  string         `json:"sex,omitempty"`
		Birthdate            Time           `json:"birthdate,omitempty"`
		Company              string         `json:"company,omitempty"`
		Organizationfunction string         `json:"organizationfunction,omitempty"`
		Addresses            []*Address     `json:"addresses,omitempty"`
		Vatnumber            string         `json:"vatnumber,omitempty"`
		Phonenumbers         []*Phonenumber `json:"phonenumbers,omitempty"`
		Relationtypes        []int64        `json:"relationtypes,omitempty"`
		Subscribed           bool           `json:"subscribed,omitempty"`
		Status               string         `json:"status,omitempty"`
		AccountType          int64          `json:"account_type,omitempty"`
		Isdeleted            bool           `json:"isdeleted,omitempty"`
		Createdts            Time           `json:"createdts,omitempty"`
		Lastupdatets         Time           `json:"lastupdatets,omitempty"`
	}

	obj := tmp{
		Id:                   o.Id,
		Email:                o.Email,
		Customertitleid:      o.Customertitleid,
		Firstname:            o.Firstname,
		Middlename:           o.Middlename,
		Lastname:             o.Lastname,
		Languagecode:         o.Languagecode,
		Sex:                  o.Sex,
		Birthdate:            o.Birthdate,
		Company:              o.Company,
		Organizationfunction: o.Organizationfunction,
		Addresses:            o.Addresses,
		Vatnumber:            o.Vatnumber,
		Phonenumbers:         o.Phonenumbers,
		Relationtypes:        o.Relationtypes,
		Subscribed:           o.Subscribed,
		Status:               o.Status,
		AccountType:          o.AccountType,
		Isdeleted:            o.Isdeleted,
		Createdts:            o.Createdts,
		Lastupdatets:         o.Lastupdatets,
	}
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	// Unpack it again, to get the wire representation
	var raw map[string]interface{}
	err = json.Unmarshal(data, &raw)
	if err != nil {
		return nil, err
	}

	for key, val := range o.CustomFields {
		raw["c_"+key] = val
	}

	// Pack it again
	return json.Marshal(raw)

	// Note: Like UnmarshalJSON, this is quite crazy. But it works beautifully.
	// Know a way to do this better? Get in touch!
}

// A CustomfieldAvailability configures in what saleschannels a custom field is
// available (during the checkout).
//
// It can also further refine the availability based on a script written in
// JavaScript.
//
// More information about writing order scripts can be found here
// (https://apps.ticketmatic.com/#/knowledgebase/developer_writingorderscripts).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/CustomfieldAvailability).
type CustomfieldAvailability struct {
	// The custom field will be available for these saleschannels. It this is empty the
	// custom field will not be available.
	Saleschannels []int64 `json:"saleschannels"`

	// Indicates if the script will be used.
	Usescript bool `json:"usescript,omitempty"`

	// A Javascript that needs to return a boolean. It has the current order available.
	Script string `json:"script,omitempty"`
}

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
	Saleschannels []int64 `json:"saleschannels"`

	// Use a script to refine the set of sales channels?
	Usescript bool `json:"usescript,omitempty"`

	// Script used to determine availability of the delivery scenario
	Script string `json:"script,omitempty"`
}

// A single Event.
//
// Status
//
// The currentstatus field of an event can have any of the following values:
//
// * Draft (19001)
//
// * Active (19002)
//
// * Closed (19003)
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Event).
type Event struct {
	// Event ID
	//
	// Note: Ignored when creating a new event.
	//
	// Note: Ignored when updating an existing event.
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
	//
	// Note: Ignored when creating a new event.
	//
	// Note: Ignored when updating an existing event.
	Publishedts Time `json:"publishedts,omitempty"`

	// Event end time
	Endts Time `json:"endts,omitempty"`

	// Event code.
	//
	// Used as a unique identifier in web sales.
	//
	// Note: Ignored when creating a new event.
	//
	// Note: Ignored when updating an existing event.
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
	//
	// Note: Ignored when creating a new event.
	//
	// Note: Ignored when updating an existing event.
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
	Seatingplanpricelistid int64 `json:"seatingplanpricelistid,omitempty"`

	// Event specific prices in addition to the prices defined in the field
	// seatingplanpricelistid. Prices from the pricelist and the event specific prices
	// are combined in one pricelist for the event. The optional position attribute
	// defines where the event specific prices will be positioned in the resulting
	// pricelist
	Seatingplaneventspecificprices *PricelistPrices `json:"seatingplaneventspecificprices,omitempty"`

	// Information about the contingents defined in the seatingplan. Read-only.
	//
	// Note: Ignored when creating a new event.
	//
	// Note: Ignored when updating an existing event.
	Seatingplancontingents []*EventSeatingplanContingent `json:"seatingplancontingents"`

	// Information about the contingents in the Event that are not in the seatingplan
	Contingents []*EventContingent `json:"contingents"`

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

	// Queue ID
	//
	// See rate limiting
	// (https://apps.ticketmatic.com/#/knowledgebase/api/ratelimiting) for more info.
	Queuetoken int64 `json:"queuetoken,omitempty"`

	// Maximum number of tickets for this event that can be added to a basket
	Maxnbrofticketsperbasket int64 `json:"maxnbrofticketsperbasket,omitempty"`

	// Event status
	//
	// The available values for this field can be found on the price type overview
	// (https://apps.ticketmatic.com/#/knowledgebase/api/events) page.
	Currentstatus int64 `json:"currentstatus,omitempty"`

	// Information on the available prices for the event
	//
	// Note: Ignored when creating a new event.
	//
	// Note: Ignored when updating an existing event.
	Prices *EventPrices `json:"prices,omitempty"`

	// Per-sales channel information about when this event is for sale.
	//
	// Note: Ignored when creating a new event.
	//
	// Note: Ignored when updating an existing event.
	Saleschannels []*EventSalesChannel `json:"saleschannels"`

	// Information on the availability of tickets per contingent. Read-only.
	//
	// Note: Ignored when creating a new event.
	//
	// Note: Ignored when updating an existing event.
	Availability []*EventContingentAvailability `json:"availability"`

	// Created timestamp
	//
	// Note: Ignored when creating a new event.
	//
	// Note: Ignored when updating an existing event.
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new event.
	//
	// Note: Ignored when updating an existing event.
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Custom fields
	CustomFields map[string]interface{} `json:"-"`
}

// Custom unmarshaller with support for custom fields
func (o *Event) UnmarshalJSON(data []byte) error {
	// Alias the type, to avoid calling UnmarshalJSON. Unpack it.
	type tmp Event
	var obj tmp
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}

	*o = Event(obj)

	// Unpack it again, this time to a map, so we can pull out the custom fields.
	var raw map[string]interface{}
	err = json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	o.CustomFields = make(map[string]interface{})
	for key, val := range raw {
		if strings.HasPrefix(key, "c_") {
			o.CustomFields[key[2:]] = val
		}
	}

	// Note: We're doing a double JSON decode here, I'd love to get rid of it
	// but I'm not sure how we can do this easily. Suggestions welcome:
	// developers@ticketmatic.com!

	return nil
}

// Custom marshaller with support for custom fields
func (o *Event) MarshalJSON() ([]byte, error) {
	// Use a custom type to avoid the custom marshaller, marshal the data.
	type tmp struct {
		Id                             int64                          `json:"id,omitempty"`
		Name                           string                         `json:"name,omitempty"`
		Subtitle                       string                         `json:"subtitle,omitempty"`
		Subtitle2                      string                         `json:"subtitle2,omitempty"`
		Webremark                      string                         `json:"webremark,omitempty"`
		Startts                        Time                           `json:"startts,omitempty"`
		Salestartts                    Time                           `json:"salestartts,omitempty"`
		Saleendts                      Time                           `json:"saleendts,omitempty"`
		Publishedts                    Time                           `json:"publishedts,omitempty"`
		Endts                          Time                           `json:"endts,omitempty"`
		Code                           string                         `json:"code,omitempty"`
		Externalcode                   string                         `json:"externalcode,omitempty"`
		Productionid                   int64                          `json:"productionid,omitempty"`
		Locationid                     int64                          `json:"locationid,omitempty"`
		Locationname                   string                         `json:"locationname,omitempty"`
		Seatingplanid                  int64                          `json:"seatingplanid,omitempty"`
		Seatingplanpricelistid         int64                          `json:"seatingplanpricelistid,omitempty"`
		Seatingplaneventspecificprices *PricelistPrices               `json:"seatingplaneventspecificprices,omitempty"`
		Seatingplancontingents         []*EventSeatingplanContingent  `json:"seatingplancontingents,omitempty"`
		Contingents                    []*EventContingent             `json:"contingents,omitempty"`
		Ticketfeeid                    int64                          `json:"ticketfeeid,omitempty"`
		Revenuesplitid                 int64                          `json:"revenuesplitid,omitempty"`
		Ticketlayoutid                 int64                          `json:"ticketlayoutid,omitempty"`
		Queuetoken                     int64                          `json:"queuetoken,omitempty"`
		Maxnbrofticketsperbasket       int64                          `json:"maxnbrofticketsperbasket,omitempty"`
		Currentstatus                  int64                          `json:"currentstatus,omitempty"`
		Prices                         *EventPrices                   `json:"prices,omitempty"`
		Saleschannels                  []*EventSalesChannel           `json:"saleschannels,omitempty"`
		Availability                   []*EventContingentAvailability `json:"availability,omitempty"`
		Createdts                      Time                           `json:"createdts,omitempty"`
		Lastupdatets                   Time                           `json:"lastupdatets,omitempty"`
	}

	obj := tmp{
		Id:                             o.Id,
		Name:                           o.Name,
		Subtitle:                       o.Subtitle,
		Subtitle2:                      o.Subtitle2,
		Webremark:                      o.Webremark,
		Startts:                        o.Startts,
		Salestartts:                    o.Salestartts,
		Saleendts:                      o.Saleendts,
		Publishedts:                    o.Publishedts,
		Endts:                          o.Endts,
		Code:                           o.Code,
		Externalcode:                   o.Externalcode,
		Productionid:                   o.Productionid,
		Locationid:                     o.Locationid,
		Locationname:                   o.Locationname,
		Seatingplanid:                  o.Seatingplanid,
		Seatingplanpricelistid:         o.Seatingplanpricelistid,
		Seatingplaneventspecificprices: o.Seatingplaneventspecificprices,
		Seatingplancontingents:         o.Seatingplancontingents,
		Contingents:                    o.Contingents,
		Ticketfeeid:                    o.Ticketfeeid,
		Revenuesplitid:                 o.Revenuesplitid,
		Ticketlayoutid:                 o.Ticketlayoutid,
		Queuetoken:                     o.Queuetoken,
		Maxnbrofticketsperbasket:       o.Maxnbrofticketsperbasket,
		Currentstatus:                  o.Currentstatus,
		Prices:                         o.Prices,
		Saleschannels:                  o.Saleschannels,
		Availability:                   o.Availability,
		Createdts:                      o.Createdts,
		Lastupdatets:                   o.Lastupdatets,
	}
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	// Unpack it again, to get the wire representation
	var raw map[string]interface{}
	err = json.Unmarshal(data, &raw)
	if err != nil {
		return nil, err
	}

	for key, val := range o.CustomFields {
		raw["c_"+key] = val
	}

	// Pack it again
	return json.Marshal(raw)

	// Note: Like UnmarshalJSON, this is quite crazy. But it works beautifully.
	// Know a way to do this better? Get in touch!
}

// Information about a contingent for an event
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Event).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/EventContingent).
type EventContingent struct {
	// Contingent ID
	Id int64 `json:"id,omitempty"`

	// Event ID
	Eventid int64 `json:"eventid,omitempty"`

	// Name of the contingent
	Name string `json:"name,omitempty"`

	// Number of tickets in the contingent
	Amount int64 `json:"amount,omitempty"`

	// Whether the barcodes for the tickets in this contingent were imported (true), or
	// were generated internally (false)
	Withimportedbarcodes bool `json:"withimportedbarcodes,omitempty"`

	// Price list ID for this contingent
	Pricelistid int64 `json:"pricelistid,omitempty"`

	// Event specific prices in addition to the prices defined in the field
	// pricelistid. Prices from the pricelist and the event specific prices are
	// combined in one pricelist. The optional position attribute defines where the
	// event specific prices will be positioned in the resulting pricelist
	Eventspecificprices *PricelistPrices `json:"eventspecificprices,omitempty"`

	// Locked tickets for the contingent
	Locks []*EventContingentLock `json:"locks"`
}

// Information about the availability of tickets for a contingent
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/EventContingentAvailability).
type EventContingentAvailability struct {
	// Contingent ID
	Tickettypeid int64 `json:"tickettypeid,omitempty"`

	// Total number of tickets in the contingent
	Total int64 `json:"total,omitempty"`

	// Number of locked tickets with a hard lock type
	LockedHard int64 `json:"locked_hard,omitempty"`

	// Number of locked tickets with a soft lock type
	LockedSoft int64 `json:"locked_soft,omitempty"`

	// Number of tickets reserved in unconfirmed orders
	Reserved int64 `json:"reserved,omitempty"`

	// Number of tickets in confirmed orders that are completely paid
	SoldPaid int64 `json:"sold_paid,omitempty"`

	// Number of tickets in confirmed orders that are not completely paid
	SoldUnpaid int64 `json:"sold_unpaid,omitempty"`

	// Number of complimentary tickets
	Complimentary int64 `json:"complimentary,omitempty"`

	// Number of available tickets
	Free int64 `json:"free,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`
}

// Information about locked tickets in a Contingent.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/EventContingentLock).
type EventContingentLock struct {
	// Contingent ID
	Tickettypeid int64 `json:"tickettypeid,omitempty"`

	// Lock type ID
	Locktypeid int64 `json:"locktypeid,omitempty"`

	// Number of tickets in the contingent
	Amount int64 `json:"amount,omitempty"`
}

// Information about the prices for an event.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/EventPrices).
type EventPrices struct {
	// Price information for the contingents
	Contingents []*EventPricesContingent `json:"contingents"`
}

// Information about the prices for a contingent for an event.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/EventPricesContingent).
type EventPricesContingent struct {
	// Contingent ID
	Contingentid int64 `json:"contingentid,omitempty"`

	// Price information for the pricetypes
	Pricetypes []*EventPricesPricetype `json:"pricetypes"`
}

// Information about costs for a price for an event.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/EventPricesCost).
type EventPricesCost struct {
	// Cost ID
	Costid int64 `json:"costid,omitempty"`

	// The actual cost
	Cost float64 `json:"cost,omitempty"`
}

// Information about the price for a pricetype for the specific sales channel for
// an event.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/EventPricesPricetype).
type EventPricesPricetype struct {
	// Pricetype ID
	Pricetypeid int64 `json:"pricetypeid,omitempty"`

	// Ticket type price ID, used to add tickets to an order
	Tickettypepriceid int64 `json:"tickettypepriceid,omitempty"`

	// Price information for this pricetype for the different sales channels
	Saleschannels []*EventPricesSaleschannel `json:"saleschannels"`
}

// Information about the price for a pricetype for the specific sales channel for
// an event.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/EventPricesSaleschannel).
type EventPricesSaleschannel struct {
	// Saleschannel ID
	Saleschannelid int64 `json:"saleschannelid,omitempty"`

	// Tickettypeprice ID
	Tickettypepriceid int64 `json:"tickettypepriceid,omitempty"`

	// The actual price
	Price float64 `json:"price,omitempty"`

	// The actual servicecharge
	Servicecharge float64 `json:"servicecharge,omitempty"`

	// The costs associated with this price
	Costs []*EventPricesCost `json:"costs"`
}

// Information about the sales period for a specific sales channel
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/SalesChannel) in an
// event (https://apps.ticketmatic.com/#/knowledgebase/api/types/Event).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/EventSalesChannel).
type EventSalesChannel struct {
	// Event ID
	Eventid int64 `json:"eventid,omitempty"`

	// Sales channel ID
	Saleschannelid int64 `json:"saleschannelid,omitempty"`

	// When the sales start
	Salestartts Time `json:"salestartts,omitempty"`

	// When the sales end
	Saleendts Time `json:"saleendts,omitempty"`

	// Whether or not this sales channel is active for this event
	Isactive bool `json:"isactive,omitempty"`
}

// Information about a contingent in the seating plan for an event
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Event).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/EventSeatingplanContingent).
type EventSeatingplanContingent struct {
	// Contingent ID
	Id int64 `json:"id,omitempty"`

	// Event ID
	Eventid int64 `json:"eventid,omitempty"`

	// Seat rank ID
	Seatrankid int64 `json:"seatrankid,omitempty"`

	// Name of the contingent
	Name string `json:"name,omitempty"`

	// Number of tickets in the contingent
	Amount int64 `json:"amount,omitempty"`
}

// Key-value item
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/KeyValueItem).
type KeyValueItem struct {
	// Key
	Key string `json:"key,omitempty"`

	// Value
	Value string `json:"value,omitempty"`
}

// A single Order.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Order).
type Order struct {
	// Order ID
	Orderid int64 `json:"orderid,omitempty"`

	// Order status
	//
	// Possible values:
	//
	// * 21001: Unconfirmed
	//
	// * 21002: Confirmed
	//
	// * 21003: Archived
	Status int64 `json:"status,omitempty"`

	// Order code
	//
	// Used as a unique identifier in web sales.
	Code string `json:"code,omitempty"`

	// Customer ID
	Customerid int64 `json:"customerid,omitempty"`

	// Has customer authenticated?
	Isauthenticatedcustomer bool `json:"isauthenticatedcustomer,omitempty"`

	// Total order amount
	//
	// Includes all costs and fees.
	Totalamount float64 `json:"totalamount,omitempty"`

	// Total amount paid
	Amountpaid float64 `json:"amountpaid,omitempty"`

	// Payment status
	//
	// Possible values:
	//
	// * 0: Incomplete
	//
	// * 1: Fully paid
	//
	// * 2: Overpaid
	Paymentstatus int64 `json:"paymentstatus,omitempty"`

	// Number of tickets in the order. Read-only
	Nbroftickets int64 `json:"nbroftickets,omitempty"`

	// Delivery status
	//
	// Possible values:
	//
	// * 2601: Not delivered
	//
	// * 2602: Delivered
	//
	// * 2603: Changed after delivery
	Deliverystatus int64 `json:"deliverystatus,omitempty"`

	// Address used when delivering physically
	Deliveryaddress *Address `json:"deliveryaddress,omitempty"`

	// Information on the deferred payment scenario. Structure depends on payment
	// method
	Deferredpaymentproperties map[string]interface{} `json:"deferredpaymentproperties,omitempty"`

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

	// Whether the overdue order has been handled (and optionally reminder mail has
	// been sent)
	Rappelhandled bool `json:"rappelhandled,omitempty"`

	// When the order will expire
	Expiryts Time `json:"expiryts,omitempty"`

	// Whether the expired order has been handled (and optionally expiry mail has been
	// sent)
	Expiryhandled bool `json:"expiryhandled,omitempty"`

	// Tickets in the order
	Tickets []*OrderTicket `json:"tickets"`

	// Products in the order
	Products []*OrderProduct `json:"products"`

	// Payments for the order
	Payments []*Payment `json:"payments"`

	// Queue tokens for rate limiting
	Queuetokens []int64 `json:"queuetokens"`

	// Promocodes active for the Order
	Promocodes []string `json:"promocodes"`

	// Related objects
	//
	// See the lookup fields on the getlist operation
	// (https://apps.ticketmatic.com/#/knowledgebase/api/orders/getlist) for a full
	// description.
	Lookup map[string]interface{} `json:"lookup,omitempty"`

	// Order fees for the order
	Ordercosts []*Ordercost `json:"ordercosts"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Custom fields
	CustomFields map[string]interface{} `json:"-"`
}

// Custom unmarshaller with support for custom fields
func (o *Order) UnmarshalJSON(data []byte) error {
	// Alias the type, to avoid calling UnmarshalJSON. Unpack it.
	type tmp Order
	var obj tmp
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}

	*o = Order(obj)

	// Unpack it again, this time to a map, so we can pull out the custom fields.
	var raw map[string]interface{}
	err = json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	o.CustomFields = make(map[string]interface{})
	for key, val := range raw {
		if strings.HasPrefix(key, "c_") {
			o.CustomFields[key[2:]] = val
		}
	}

	// Note: We're doing a double JSON decode here, I'd love to get rid of it
	// but I'm not sure how we can do this easily. Suggestions welcome:
	// developers@ticketmatic.com!

	return nil
}

// Custom marshaller with support for custom fields
func (o *Order) MarshalJSON() ([]byte, error) {
	// Use a custom type to avoid the custom marshaller, marshal the data.
	type tmp struct {
		Orderid                   int64                  `json:"orderid,omitempty"`
		Status                    int64                  `json:"status,omitempty"`
		Code                      string                 `json:"code,omitempty"`
		Customerid                int64                  `json:"customerid,omitempty"`
		Isauthenticatedcustomer   bool                   `json:"isauthenticatedcustomer,omitempty"`
		Totalamount               float64                `json:"totalamount,omitempty"`
		Amountpaid                float64                `json:"amountpaid,omitempty"`
		Paymentstatus             int64                  `json:"paymentstatus,omitempty"`
		Nbroftickets              int64                  `json:"nbroftickets,omitempty"`
		Deliverystatus            int64                  `json:"deliverystatus,omitempty"`
		Deliveryaddress           *Address               `json:"deliveryaddress,omitempty"`
		Deferredpaymentproperties map[string]interface{} `json:"deferredpaymentproperties,omitempty"`
		Saleschannelid            int64                  `json:"saleschannelid,omitempty"`
		Paymentscenarioid         int64                  `json:"paymentscenarioid,omitempty"`
		Deliveryscenarioid        int64                  `json:"deliveryscenarioid,omitempty"`
		Rappelts                  Time                   `json:"rappelts,omitempty"`
		Rappelhandled             bool                   `json:"rappelhandled,omitempty"`
		Expiryts                  Time                   `json:"expiryts,omitempty"`
		Expiryhandled             bool                   `json:"expiryhandled,omitempty"`
		Tickets                   []*OrderTicket         `json:"tickets,omitempty"`
		Products                  []*OrderProduct        `json:"products,omitempty"`
		Payments                  []*Payment             `json:"payments,omitempty"`
		Queuetokens               []int64                `json:"queuetokens,omitempty"`
		Promocodes                []string               `json:"promocodes,omitempty"`
		Lookup                    map[string]interface{} `json:"lookup,omitempty"`
		Ordercosts                []*Ordercost           `json:"ordercosts,omitempty"`
		Createdts                 Time                   `json:"createdts,omitempty"`
		Lastupdatets              Time                   `json:"lastupdatets,omitempty"`
	}

	obj := tmp{
		Orderid:                   o.Orderid,
		Status:                    o.Status,
		Code:                      o.Code,
		Customerid:                o.Customerid,
		Isauthenticatedcustomer:   o.Isauthenticatedcustomer,
		Totalamount:               o.Totalamount,
		Amountpaid:                o.Amountpaid,
		Paymentstatus:             o.Paymentstatus,
		Nbroftickets:              o.Nbroftickets,
		Deliverystatus:            o.Deliverystatus,
		Deliveryaddress:           o.Deliveryaddress,
		Deferredpaymentproperties: o.Deferredpaymentproperties,
		Saleschannelid:            o.Saleschannelid,
		Paymentscenarioid:         o.Paymentscenarioid,
		Deliveryscenarioid:        o.Deliveryscenarioid,
		Rappelts:                  o.Rappelts,
		Rappelhandled:             o.Rappelhandled,
		Expiryts:                  o.Expiryts,
		Expiryhandled:             o.Expiryhandled,
		Tickets:                   o.Tickets,
		Products:                  o.Products,
		Payments:                  o.Payments,
		Queuetokens:               o.Queuetokens,
		Promocodes:                o.Promocodes,
		Lookup:                    o.Lookup,
		Ordercosts:                o.Ordercosts,
		Createdts:                 o.Createdts,
		Lastupdatets:              o.Lastupdatets,
	}
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	// Unpack it again, to get the wire representation
	var raw map[string]interface{}
	err = json.Unmarshal(data, &raw)
	if err != nil {
		return nil, err
	}

	for key, val := range o.CustomFields {
		raw["c_"+key] = val
	}

	// Pack it again
	return json.Marshal(raw)

	// Note: Like UnmarshalJSON, this is quite crazy. But it works beautifully.
	// Know a way to do this better? Get in touch!
}

// A single order fee for an order.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Ordercost).
type Ordercost struct {
	// Order ID
	Orderid int64 `json:"orderid,omitempty"`

	// Order fee ID
	Servicechargedefinitionid int64 `json:"servicechargedefinitionid,omitempty"`

	// Payment amount
	Amount float64 `json:"amount,omitempty"`
}

// More info about order fees can be found here
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_orderfees).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/OrderfeeRule).
type OrderfeeRule struct {
	// This is required if the order fee type is set to automatic. It is a set of rules
	// that define the order fee.
	Auto []*OrderfeeAutoRule `json:"auto"`

	// This is required if the order fee type is set to script. The javascript needs to
	// return a value.
	Script string `json:"script,omitempty"`

	// This can be set if the order fee type is set to script. It allows adding extra
	// information to the script environment.
	Context []*OrderfeeScriptContext `json:"context"`
}

// More info about order fees can be found here
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_orderfees).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/OrderfeeScriptContext).
type OrderfeeScriptContext struct {
	// The name of the variable that will be added to the script environment.
	Key string `json:"key,omitempty"`

	// The query that will be executed on the public data model. The result will be
	// available in the script environment.
	Query string `json:"query,omitempty"`

	// If set to true the query will be cached for 60 seconds. If not set the query
	// will be executed again every time a script is executed.
	Cacheable bool `json:"cacheable,omitempty"`
}

// More info about order fees can be found here
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_orderfees).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/OrderfeeAutoRule).
type OrderfeeAutoRule struct {
	// The sales channels that this order fee is applicable for. If not set it defaults
	// to 'all'. This is only needed if the order fee type is set to automatic.
	Saleschannelids []int64 `json:"saleschannelids"`

	// The delivery scenarios that this order fee is applicable for. If not set it
	// defaults to 'all'. This is only needed if the order fee type is set to
	// automatic.
	Deliveryscenarioids []int64 `json:"deliveryscenarioids"`

	// The payment scenarios that this order fee is applicable for. If not set it
	// default to 'all'. This is only needed if the order fee type is set to automatic.
	Paymentscenarioids []int64 `json:"paymentscenarioids"`

	// Can be fixedfee or percentagefee. Defauls to fixedfee. This is only needed if
	// the order fee type is set to automatic.
	Status string `json:"status,omitempty"`

	// The value (amount) that will be added to the order. Is required if the order fee
	// type is set to automatic.
	Value float64 `json:"value,omitempty"`
}

// A single product in an order.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/OrderProduct).
type OrderProduct struct {
	// Orderproduct ID
	Id int64 `json:"id,omitempty"`

	// Unique code for this orderproduct
	Code string `json:"code,omitempty"`

	// Order ID
	Orderid int64 `json:"orderid,omitempty"`

	// Product ID
	Productid int64 `json:"productid,omitempty"`

	// Property values for this product
	Properties map[string]string `json:"properties,omitempty"`

	// Ticket price
	Price float64 `json:"price,omitempty"`

	// Vouchercode ID for the voucher that is linked to this orderproduct
	Vouchercodeid int64 `json:"vouchercodeid,omitempty"`
}

// A single ticket in an order.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/OrderTicket).
type OrderTicket struct {
	// Ticket ID
	Id int64 `json:"id,omitempty"`

	// Order ID
	Orderid int64 `json:"orderid,omitempty"`

	// Contingent ID
	Tickettypeid int64 `json:"tickettypeid,omitempty"`

	// Id for the tickettypeprice of this ticket for the order
	Baskettickettypepriceid int64 `json:"baskettickettypepriceid,omitempty"`

	// Ticket price
	Price float64 `json:"price,omitempty"`

	// Service charge
	Servicecharge float64 `json:"servicecharge,omitempty"`

	// Ticket holder ID
	Ticketholderid int64 `json:"ticketholderid,omitempty"`

	// Name for the ticket holder
	Ticketname string `json:"ticketname,omitempty"`

	// The abo ticketid for tickets that belong to an abo
	Aboparentid int64 `json:"aboparentid,omitempty"`

	// Event id
	Eventid int64 `json:"eventid,omitempty"`

	// Pricetype id
	Pricetypeid int64 `json:"pricetypeid,omitempty"`

	// Description of the ticket
	Seatdescription string `json:"seatdescription,omitempty"`

	// Name of the seat
	Seatname string `json:"seatname,omitempty"`

	// Contingent name
	Tickettypename string `json:"tickettypename,omitempty"`

	// The voucher code that was linked to this ticket
	Vouchercodeid int64 `json:"vouchercodeid,omitempty"`

	// The id of the product this ticket is linked to
	Bundleid int64 `json:"bundleid,omitempty"`
}

// A single payment.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Payment).
type Payment struct {
	// Payment ID
	Id int64 `json:"id,omitempty"`

	// Order ID
	Orderid int64 `json:"orderid,omitempty"`

	// Payment method ID
	Paymentmethodid int64 `json:"paymentmethodid,omitempty"`

	// Id for the original payment if this payment is a refund
	Refundpaymentid int64 `json:"refundpaymentid,omitempty"`

	// Timestamp of payment
	Paidts Time `json:"paidts,omitempty"`

	// Payment amount
	Amount float64 `json:"amount,omitempty"`

	// Additional properties for the payment. Structure depends on the payment method
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// A PaymentscenarioAvailability configures in what saleschannels a payment
// scenario is available.
//
// It can also further refine the availability based on a script written in
// JavaScript.
//
// More information about writing order scripts can be found here
// (https://apps.ticketmatic.com/#/knowledgebase/developer_writingorderscripts).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PaymentscenarioAvailability).
type PaymentscenarioAvailability struct {
	// The payment scenario will be available for these saleschannels. It this is empty
	// the payment scenario will not be available.
	Saleschannels []int64 `json:"saleschannels"`

	// Indicates if the script will be used.
	Usescript bool `json:"usescript,omitempty"`

	// A Javascript that needs to return a boolean. It has the current order and
	// saleschannel available. More info
	// (https://apps.ticketmatic.com/#/knowledgebase/developer_writingorderscripts)
	Script string `json:"script,omitempty"`
}

// The PaymentscenarioExpiryParameters can only be set when the Paymentscenario is
// of a type deferred payment.
//
// It determines the moment in time when an order expires. It's calculated as
// MIN(<order creation date> + daysafterordercreation, <date of first event in
// order> - daysbeforeevent). If deleteonexpiry is set to true, the order will be
// deleted.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PaymentscenarioExpiryParameters).
type PaymentscenarioExpiryParameters struct {
	// The amount of days after an order has been created that the order becomes
	// overdue.
	Daysafterordercreation int64 `json:"daysafterordercreation,omitempty"`

	// The number of days before an event that an order becomes overdue.
	Daysbeforeevent int64 `json:"daysbeforeevent,omitempty"`

	// Indicates is the order will be deleted when it's expired.
	Deleteonexpiry bool `json:"deleteonexpiry,omitempty"`
}

// The PaymentscenarioOverdueParameters can only be set when the Paymentscenario is
// of type deferred payment.
//
// It determines the moment in time when an order becomes overdue. It's calculated
// as MIN(<order creation date> + daysafterordercreation, <date of first event in
// order> - daysbeforeevent).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PaymentscenarioOverdueParameters).
type PaymentscenarioOverdueParameters struct {
	// The amount of days after an order has been created that the order becomes
	// overdue.
	Daysafterordercreation int64 `json:"daysafterordercreation,omitempty"`

	// The number of days before an event that an order becomes overdue.
	Daysbeforeevent int64 `json:"daysbeforeevent,omitempty"`
}

// See contact (https://apps.ticketmatic.com/#/knowledgebase/api/types/Contact) for
// more information.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Phonenumber).
type Phonenumber struct {
	// Address ID
	//
	// Note: Only available when used for a contact
	// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Contact) address.
	Id int64 `json:"id,omitempty"`

	// Contact this address belongs to
	//
	// Note: Only available when used for a contact
	// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Contact) address.
	Customerid int64 `json:"customerid,omitempty"`

	// Phone number
	Number string `json:"number,omitempty"`

	// Phone number type ID
	Typeid int64 `json:"typeid,omitempty"`

	// Phone number type (based on typeid, returned as a convenience)
	Type string `json:"type,omitempty"`
}

// You can find more information about prices in the endpoint documentation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricelists).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PricelistPrices).
type PricelistPrices struct {
	// The seatranks for which this pricelist lists prices.
	Seatrankids []int64 `json:"seatrankids"`

	// The set of prices for this pricelist.
	Prices []*PricelistPrice `json:"prices"`
}

// You can find more information about price in the endpoint documentation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_pricelists).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PricelistPrice).
type PricelistPrice struct {
	// The pricetype for this price.
	Pricetypeid int64 `json:"pricetypeid,omitempty"`

	// The (decimal) prices for this PricelistPrice. If no seatrankids has been set,
	// this should consist of 1 price. If seatrankids are set this should an equal
	// number of prices as the number of seatranks.
	Prices []float64 `json:"prices"`

	// Array of booleans indicating if the corresponding price is available for this
	// PricelistPrice. Should contain the same number of booleans as prices.
	Availabilities []bool `json:"availabilities"`

	// The list of saleschannels for which this PricelistPrice is active.
	Saleschannels []int64 `json:"saleschannels"`

	// Extra conditions for this price. This can be a promocode, a ticketlimit per
	// order, ... .
	Conditions []*PricelistPriceCondition `json:"conditions"`

	// Optional, and only used for eventspecificprices. Indicates the position of this
	// price in the pricelist.
	Position int64 `json:"position,omitempty"`
}

// These are the possible condition and example values:
//
// Ticketlimit
//
// There is a limited amount of tickets available for the selected pricetype.
//
//
//    {
//        "type": "ticketlimit",
//        "value": 10
//    }
//
//
//
// Date validity
//
// The price type is only available in this period.
//
// Absolute
//
//
//    {
//        "type": "date",
//        "value": {
//            "datetype": "absolute",
//            "absoluteStart": "2015-05-20",
//            "absoluteEnd": "2015-05-27"
//        }
//    }
//
//
//
// Relative
//
//
//    {
//        "type": "date",
//        "value": {
//            "datetype": "relative_eventdate",
//            "relativeStart": 10,
//            "relativeEnd": 5
//        }
//    }
//
//
//
// Promocode
//
// The price type is only available if the customer provides a promocode.
//
//
//    {
//        "type": "promocode",
//        "value": "TM"
//    }
//
//
//
// Max number of tickets per customer
//
// Limit the maximum number of tickets a customer can buy of this specific price
// type.
//
//
//    {
//        "type": "orderticketlimit",
//        "value": 2
//    }
//
//
//
// Voucherids
//
// When buying a ticket of this pricetype, a valid vouchercode with voucherid one
// of the values should be attached to the ticket.
//
//
//    {
//        "type": "voucherids",
//        "value": [1,2,3]
//    }
//
//
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PricelistPriceCondition).
type PricelistPriceCondition struct {
	// The type of condition. Possible values:
	//
	// * ticketlimit
	//
	// * date
	//
	// * promocode
	//
	// * orderticketlimit
	//
	// * voucherids
	Type string `json:"type,omitempty"`

	// The value of this condition. See type for info about what should be filled in.
	Value interface{} `json:"value,omitempty"`
}

// Product instancevalue exception
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/ProductInstanceException).
type ProductInstanceException struct {
	// Properties for which this exception is valid
	Properties map[string][]string `json:"properties,omitempty"`

	// Value for this exception
	Value *ProductInstanceValue `json:"value,omitempty"`
}

// Product Instance Pricetype Value
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/ProductInstancePricetypeValue).
type ProductInstancePricetypeValue struct {
	// Min amount from which the pricetype will be applied
	From int64 `json:"from,omitempty"`

	// Pricetype id
	Id int64 `json:"id,omitempty"`
}

// Product Instance Value
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/ProductInstanceValue).
type ProductInstanceValue struct {
	// Price
	Price float64 `json:"price,omitempty"`

	// Voucher
	Voucher *ProductVoucherValue `json:"voucher,omitempty"`

	// Set of tickettypeprices (used in fixedbundle products)
	Tickettypeprices []int64 `json:"tickettypeprices"`

	// Set of pricetype values (used in optionbundle products)
	Pricetypes []*ProductInstancePricetypeValue `json:"pricetypes"`

	// Set of tickettypes (used in optionbundle products)
	Tickettypes []int64 `json:"tickettypes"`
}

// Product instancevalues
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/ProductInstancevalues).
type ProductInstancevalues struct {
	// Default values
	Default *ProductInstanceValue `json:"default,omitempty"`

	// Exceptions on the default values
	Exceptions []*ProductInstanceException `json:"exceptions"`
}

// Product property
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/ProductProperty).
type ProductProperty struct {
	// Key
	Key string `json:"key,omitempty"`

	// Name
	Name string `json:"name,omitempty"`

	// Values
	Values []*KeyValueItem `json:"values"`

	// Description
	Description string `json:"description,omitempty"`
}

// Product Voucher Value
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/ProductVoucherValue).
type ProductVoucherValue struct {
	// Voucher id
	Voucherid int64 `json:"voucherid,omitempty"`

	// Amount (only used for vouchers of type Payment)
	Amount float64 `json:"amount,omitempty"`
}

// Defines which fees are active for specific price types and sales channels. It's
// possible to define a fixed fee and a percentage based fee. The default rule (if
// none is specified for a specific sales channel) is always a fixed fee of 0.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/TicketfeeRules).
type TicketfeeRules struct {
	// The default ticket fee rule, one rule for each saleschannel.
	Default []*TicketfeeSaleschannelRule `json:"default"`

	// An array of exception rules for specific pricetypes.
	Exceptions []*TicketfeeException `json:"exceptions"`
}

// An exception to the default rule for a specific pricetype and a set of
// saleschannels.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/TicketfeeException).
type TicketfeeException struct {
	// The pricetype for which this exception is active.
	Pricetypeid int64 `json:"pricetypeid,omitempty"`

	// The set of rules (one for each saleschannel).
	Saleschannels []*TicketfeeSaleschannelRule `json:"saleschannels"`
}

// This is a rule for a specific saleschannel that indicates the fee based on a
// fixed amount or a percentage.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/TicketfeeSaleschannelRule).
type TicketfeeSaleschannelRule struct {
	// The saleschannel for which this rule is active.
	Saleschannelid int64 `json:"saleschannelid,omitempty"`

	// The status sets the type of rule. Possible values:
	//
	// * fixedfee: A fixed ticket fee.
	//
	// * percentagefee: A fee thats a percentage of the ticket.
	Status string `json:"status,omitempty"`

	// The value of this ticket fee. Can be an absolute amount (fixedfee) or a
	// percentage (percentagefee). In both cases only provide a decimal.
	Value float64 `json:"value,omitempty"`
}

// The definition of the validity of a voucher.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/VoucherValidity).
type VoucherValidity struct {
	// The max number of times the vouchercode can be used
	Maxusages int64 `json:"maxusages,omitempty"`

	// The max number of times the vouchercode can be used for a single event
	Maxusagesperevent int64 `json:"maxusagesperevent,omitempty"`

	// The fixed expiry date for a voucher
	ExpiryFixeddate Time `json:"expiry_fixeddate,omitempty"`

	// The relative expiry date for a voucher: voucher code expires this number of
	// months after creation
	ExpiryMonthsaftercreation int64 `json:"expiry_monthsaftercreation,omitempty"`
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

// Filter parameters to fetch a list of contacts
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/ContactQuery).
type ContactQuery struct {
	// A SQL query that returns contact IDs
	//
	// Can be used to do arbitrary filtering. See the database documentation for
	// contact (/db/contact) for more information.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// Only include contacts that have been updated since the given timestamp.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Limit results to at most the given amount of contacts.
	Limit int64 `json:"limit,omitempty"`

	// Skip the first X contacts.
	Offset int64 `json:"offset,omitempty"`

	// Order by the given field.
	//
	// Supported values: name, lastupdatets, createdts.
	Orderby string `json:"orderby,omitempty"`

	// Output format.
	//
	// Possible values:
	//
	// * ids: Only fill the ID field
	//
	// * minimal: A minimal set of order fields
	//
	// * default: Return all order fields (also used when the output parameter is
	// omitted)
	Output string `json:"output,omitempty"`

	// A text filter string.
	//
	// Matches against the contact name and contact details.
	Searchterm string `json:"searchterm,omitempty"`
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

// Filter parameters to fetch a list of events
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/EventQuery).
type EventQuery struct {
	// A SQL query that returns event IDs
	//
	// Can be used to do arbitrary filtering. See the database documentation for event
	// (/db/event) for more information.
	Filter string `json:"filter,omitempty"`

	// Only include events that have been updated since the given timestamp.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Limit results to at most the given amount of events.
	Limit int64 `json:"limit,omitempty"`

	// Skip the first X events.
	Offset int64 `json:"offset,omitempty"`

	// Order by the given field.
	//
	// Supported values: name, startts.
	Orderby string `json:"orderby,omitempty"`

	// Output format.
	//
	// Possible values:
	//
	// * ids: Only fill the ID field
	//
	// * default: Return all event fields (also used when the output parameter is
	// omitted)
	//
	// * withlookup: Returns all event fields and an additional lookup field which
	// contains all dependent objects
	Output string `json:"output,omitempty"`

	// A text filter string.
	//
	// Matches against the start of the event name, the production name or the
	// subtitle.
	Searchterm string `json:"searchterm,omitempty"`

	// Filters the events based on a given set of fields. Currently supports:
	// productionid.
	Simplefilter *EventFilter `json:"simplefilter,omitempty"`

	// Restrict the event information to a specific context.
	//
	// Currently allows you to filter the event information (both the events and the
	// pricing information within each event) to a specific saleschannel. This makes it
	// very easy to show the correct information on a website.
	Context *EventContext `json:"context,omitempty"`
}

// Used when requesting events, to filter events.
//
// Currently allows you to filter based on the production ID.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/EventFilter).
type EventFilter struct {
	// The ID of the production
	Productionid int64 `json:"productionid,omitempty"`

	// The event status. By default, events with status Active or Closed will be
	// returned
	Status int64 `json:"status,omitempty"`
}

// Used when requesting events, to restrict the event information to a specific
// context.
//
// Currently allows you to filter the event information (both the events and the
// pricing information within each event) to a specific saleschannel. This makes it
// very easy to show the correct information on a website.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/EventContext).
type EventContext struct {
	// The ID of the saleschannel used to restrict the event information
	Saleschannelid int64 `json:"saleschannelid,omitempty"`
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

	// Event IDs that might end up in this order, used to hint the rate limiter
	// (https://apps.ticketmatic.com/#/knowledgebase/api/ratelimiting) of what might
	// come.
	Events []int64 `json:"events"`

	// Product IDs that might end up in this order, used to hint the rate limiter
	// (https://apps.ticketmatic.com/#/knowledgebase/api/ratelimiting) of what might
	// come.
	Products []int64 `json:"products"`
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

// Info for requesting a PDF ticket for one or more tickets or vouchercodes in an
// order (https://apps.ticketmatic.com/#/knowledgebase/api/types/Order).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/TicketsPdfRequest).
type TicketsPdfRequest struct {
	// Ticketids
	Tickets []int64 `json:"tickets"`

	// Vouchercodeids
	Vouchercodes []int64 `json:"vouchercodes"`
}

// Info for requesting a e-mail delivery for an order
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Order).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/TicketsEmaildeliveryRequest).
type TicketsEmaildeliveryRequest struct {
	// Template id
	Templateid int64 `json:"templateid,omitempty"`
}

// Url.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Url).
type Url struct {
	// Url.
	Url string `json:"url,omitempty"`
}

// Order tickettype
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/OrderTickettype).
type OrderTickettype struct {
	// Tickettype id
	Id int64 `json:"id,omitempty"`

	// Tickettype name
	Name string `json:"name,omitempty"`

	// Tickettype full name
	Fulltypename string `json:"fulltypename,omitempty"`
}

// Info for requesting an immediate payment in an order
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Order).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/PaymentRequest).
type PaymentRequest struct {
	// The returnurl that will be called after the payment request was done.
	Returnurl string `json:"returnurl,omitempty"`

	// The language to be used during the payment processing
	Language string `json:"language,omitempty"`
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
	Tickets []*CreateTicket `json:"tickets"`
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

// Request data used to delete tickets
// (https://apps.ticketmatic.com/#/knowledgebase/api/orders/deletetickets) from an
// order (https://apps.ticketmatic.com/#/knowledgebase/api/types/Order).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/DeleteTickets).
type DeleteTickets struct {
	// Ticket IDs
	Tickets []int64 `json:"tickets"`
}

// Used when requesting orders, to filter orders.
//
// Specify any of the supported fields to filter the list of orders.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/OrderFilter).
type OrderFilter struct {
	// Only include orders older than the given timestamp
	Createdsince Time `json:"createdsince,omitempty"`

	// Filter orders based on saleschannel
	Saleschannelid int64 `json:"saleschannelid,omitempty"`

	// Filter orders based on customer
	Customerid int64 `json:"customerid,omitempty"`

	// Only include orders with a given status
	//
	// Possible values:
	//
	// * 21001: Unconfirmed orders
	//
	// * 21002: Confirmed orders
	//
	// * 21003: Archived orders
	Status int64 `json:"status,omitempty"`
}

// Filter parameters to fetch a list of orders
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/OrderQuery).
type OrderQuery struct {
	// A SQL query that returns order IDs
	//
	// Can be used to do arbitrary filtering. See the database documentation for order
	// (/db/order) for more information.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// Only include orders that have been updated since the given timestamp.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Limit results to at most the given amount of orders.
	Limit int64 `json:"limit,omitempty"`

	// Skip the first X orders.
	Offset int64 `json:"offset,omitempty"`

	// Order by the given field.
	//
	// Supported values: createdts, lastupdatets.
	Orderby string `json:"orderby,omitempty"`

	// Output format.
	//
	// Possible values:
	//
	// * ids: Only fill the ID field
	//
	// * minimal: A minimal set of order fields
	//
	// * default: Return all order fields (also used when the output parameter is
	// omitted)
	//
	// * withlookup: Returns all order fields and an additional lookup field which
	// contains all dependent objects
	Output string `json:"output,omitempty"`

	// A text filter string.
	//
	// Matches against the order ID or the customer details..
	Searchterm string `json:"searchterm,omitempty"`

	// Filters the orders based on a given set of fields. Currently supports:
	// createdsince, saleschannelid, customerid, status.
	Simplefilter *OrderFilter `json:"simplefilter,omitempty"`
}

// Used to update an order. Each of the fields is optional. Omitting a field will
// leave it unchanged.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/UpdateOrder).
type UpdateOrder struct {
	// New delivery scenario ID
	Deliveryscenarioid int64 `json:"deliveryscenarioid,omitempty"`

	// Delivery address
	Deliveryaddress *Address `json:"deliveryaddress,omitempty"`

	// New payment scenario ID
	Paymentscenarioid int64 `json:"paymentscenarioid,omitempty"`

	// New customer ID
	Customerid int64 `json:"customerid,omitempty"`

	// Change custom field values
	Customfields map[string]interface{} `json:"customfields,omitempty"`

	// Rappel timestamp, as string in ISO 8601 format. Cannot be in the past.
	Rappelts string `json:"rappelts,omitempty"`

	// Expiry timestamp, as string in ISO 8601 format. Cannot be in the past.
	Expiryts string `json:"expiryts,omitempty"`
}

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
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/UpdateTickets).
type UpdateTickets struct {
	// Ticket IDs
	Tickets []int64 `json:"tickets"`

	// Operation to execute.
	//
	// Supported values:
	//
	// * setticketholders
	//
	// * updatepricetype
	Operation string `json:"operation,omitempty"`

	// Operation parameters
	Params map[string]interface{} `json:"params,omitempty"`
}

// Request data used to add a payment
// (https://apps.ticketmatic.com/#/knowledgebase/api/orders/addpayments) to an
// order (https://apps.ticketmatic.com/#/knowledgebase/api/types/Order).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/AddPayments).
type AddPayments struct {
	// Id of the payment method to be used for the payment
	Paymentmethodid int64 `json:"paymentmethodid,omitempty"`

	// Amount for the payment
	Amount float64 `json:"amount,omitempty"`
}

// Request data used to refund a payment
// (https://apps.ticketmatic.com/#/knowledgebase/api/orders/addrefunds) for an
// order (https://apps.ticketmatic.com/#/knowledgebase/api/types/Order).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/AddRefunds).
type AddRefunds struct {
	// Id of the payment that needs to be refunded
	Paymentid int64 `json:"paymentid,omitempty"`

	// Amount that needs to be refunded
	Amount float64 `json:"amount,omitempty"`
}

// Log item returned when requesting the log history of an order
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Order).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/LogItem).
type LogItem struct {
	// Id of the log item
	Id int64 `json:"id,omitempty"`

	// Log item timestamp
	Ts Time `json:"ts,omitempty"`

	// User id
	Userid int64 `json:"userid,omitempty"`

	// Log item type
	Typeid int64 `json:"typeid,omitempty"`

	// Order id
	Orderid int64 `json:"orderid,omitempty"`

	// User name
	Username string `json:"username,omitempty"`

	// Info
	Info map[string]interface{} `json:"info,omitempty"`

	// Model
	Model map[string]interface{} `json:"model,omitempty"`

	// Lookup info
	Lookupinfo map[string]interface{} `json:"lookupinfo,omitempty"`
}

// Rate limiting status. See rate limiting
// (https://apps.ticketmatic.com/#/knowledgebase/api/ratelimiting) for more details
// on rate limiting.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/QueueStatus).
type QueueStatus struct {
	// Queueing ID, used to request status updates
	Id string `json:"id,omitempty"`

	// Status code: 1: wait more, 2: ready to proceed
	Progress int64 `json:"progress,omitempty"`

	// Number of people waiting ahead. Might not be returned when the queue hasn't
	// started yet.
	Ahead int64 `json:"ahead,omitempty"`

	// Number of milliseconds to wait before requesting a new status update
	Backoff int64 `json:"backoff,omitempty"`

	// Whether the queue has started
	Started bool `json:"started,omitempty"`

	// When the queue will start
	Starttime Time `json:"starttime,omitempty"`

	// The ID of the newly created order. Only returned when a throttled "create order"
	// call has finished queueing.
	Orderid int64 `json:"orderid,omitempty"`
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

	// Only return items with the given typeid.
	Typeid int64 `json:"typeid,omitempty"`
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
	Id int64 `json:"id,omitempty"`

	// Type ID
	//
	// Note: Ignored when updating an existing ticket layout.
	Typeid int64 `json:"typeid,omitempty"`

	// Name for the ticket layout
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

// Set of parameters used to filter ticket layout templates.
//
// More info: see ticket layout template
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/TicketLayoutTemplate),
// the getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ticketlayouttemplates/getlist)
// and the ticket layout templates endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ticketlayouttemplates).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/TicketLayoutTemplateQuery).
type TicketLayoutTemplateQuery struct {
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

// A single ticket layout template.
//
// More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ticketlayouttemplates/get)
// and the ticket layout templates endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ticketlayouttemplates).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/TicketLayoutTemplate).
type TicketLayoutTemplate struct {
	// Unique ID
	//
	// Note: Ignored when creating a new ticket layout template.
	//
	// Note: Ignored when updating an existing ticket layout template.
	Id int64 `json:"id,omitempty"`

	// Type ID
	//
	// Note: Ignored when updating an existing ticket layout template.
	Typeid int64 `json:"typeid,omitempty"`

	// Name for the ticket layout template
	Name string `json:"name,omitempty"`

	// Html template containing the definition for the ticket layout template
	//
	// Note: Not set when retrieving a list of ticket layout templates.
	Htmltemplate string `json:"htmltemplate,omitempty"`

	// Css classes for the ticket layout template
	//
	// Note: Not set when retrieving a list of ticket layout templates.
	Css string `json:"css,omitempty"`

	// Translations for the ticket layout template
	//
	// Note: Not set when retrieving a list of ticket layout templates.
	Translations map[string]string `json:"translations,omitempty"`

	// Deliveryscenario's for which this ticket layout template will be used
	Deliveryscenarios []int64 `json:"deliveryscenarios"`

	// Created timestamp
	//
	// Note: Ignored when creating a new ticket layout template.
	//
	// Note: Ignored when updating an existing ticket layout template.
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new ticket layout template.
	//
	// Note: Ignored when updating an existing ticket layout template.
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new ticket layout template.
	//
	// Note: Ignored when updating an existing ticket layout template.
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

// Set of parameters used to filter order fee definitions.
//
// More info: see order fee definition
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/OrderFeeDefinition), the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_orderfeedefinitions/getlist)
// and the order fee definitions endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_orderfeedefinitions).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/OrderFeeDefinitionQuery).
type OrderFeeDefinitionQuery struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
}

// A single order fee definition.
//
// More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_orderfeedefinitions/get)
// and the order fee definitions endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_pricing_orderfeedefinitions).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/OrderFeeDefinition).
type OrderFeeDefinition struct {
	// Unique ID
	//
	// Note: Ignored when creating a new order fee definition.
	//
	// Note: Ignored when creating a new order fee.
	Id int64 `json:"id,omitempty"`

	// Name for the order fee
	Name string `json:"name,omitempty"`

	// Type of the order fee. Can be Automatic (2401) or Script (2402)
	Typeid int64 `json:"typeid,omitempty"`

	// Definition of the rule that defines when the order fee will be applied
	//
	// Note: Not set when retrieving a list of order fee definitions.
	//
	// Note: Not set when retrieving a list of order fees.
	Rule *OrderfeeRule `json:"rule,omitempty"`

	// Created timestamp
	//
	// Note: Ignored when creating a new order fee definition.
	//
	// Note: Ignored when creating a new order fee.
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new order fee definition.
	//
	// Note: Ignored when creating a new order fee.
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new order fee definition.
	//
	// Note: Ignored when creating a new order fee.
	Isarchived bool `json:"isarchived,omitempty"`

	// Archived timestamp
	//
	// Note: Ignored when creating a new order fee definition.
	//
	// Note: Ignored when creating a new order fee.
	Archivedts Time `json:"archivedts,omitempty"`
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
	Id int64 `json:"id,omitempty"`

	// Name for the pricelist
	Name string `json:"name,omitempty"`

	// Definition of the actual prices and conditions for the pricelist
	//
	// Note: Not set when retrieving a list of price lists.
	Prices *PricelistPrices `json:"prices,omitempty"`

	// Boolean indicating whether this pricelist has ranks or not
	Hasranks bool `json:"hasranks,omitempty"`

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
	Id int64 `json:"id,omitempty"`

	// Name for the ticket fee scheme
	Name string `json:"name,omitempty"`

	// Definition of the rules that define when the ticket fee will be applied
	//
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

// Set of parameters used to filter productcategories.
//
// More info: see productcategory
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Productcategory), the
// getlist operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_productcategories/getlist)
// and the productcategories endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_productcategories).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/ProductcategoryQuery).
type ProductcategoryQuery struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
}

// A single productcategory.
//
// More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_productcategories/get)
// and the productcategories endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_productcategories).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Productcategory).
type Productcategory struct {
	// Unique ID
	//
	// Note: Ignored when creating a new productcategory.
	//
	// Note: Ignored when updating an existing productcategory.
	Id int64 `json:"id,omitempty"`

	// Name for the product category
	Name string `json:"name,omitempty"`

	// Name for the product category in plural
	Nameplural string `json:"nameplural,omitempty"`

	// Name for the holder/owner of this product
	Contactname string `json:"contactname,omitempty"`

	// Name for the holder/owner of this product in plural
	Contactnameplural string `json:"contactnameplural,omitempty"`

	// Created timestamp
	//
	// Note: Ignored when creating a new productcategory.
	//
	// Note: Ignored when updating an existing productcategory.
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new productcategory.
	//
	// Note: Ignored when updating an existing productcategory.
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new productcategory.
	//
	// Note: Ignored when updating an existing productcategory.
	Isarchived bool `json:"isarchived,omitempty"`
}

// Set of parameters used to filter products.
//
// More info: see product
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Product), the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_products/getlist) and
// the products endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_products).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/ProductQuery).
type ProductQuery struct {
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

// A single product.
//
// More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_products/get) and the
// products endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_products).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/Product).
type Product struct {
	// Unique ID
	//
	// Note: Ignored when creating a new product.
	//
	// Note: Ignored when updating an existing product.
	Id int64 `json:"id,omitempty"`

	// Type ID
	//
	// Note: Ignored when updating an existing product.
	Typeid int64 `json:"typeid,omitempty"`

	// Unique 12-digit for the product
	Code string `json:"code,omitempty"`

	// Name for the product
	Name string `json:"name,omitempty"`

	// Description for the product
	Description string `json:"description,omitempty"`

	// Category for the product
	Categoryid int64 `json:"categoryid,omitempty"`

	// Optional layout for the product. If not specified, there will be no ticket
	// generated for the product
	Layoutid int64 `json:"layoutid,omitempty"`

	// Definition of possible properties for the product
	Properties []*ProductProperty `json:"properties"`

	// Definition of the values for an instance of the product. These depend on the
	// properties
	//
	// Note: Not set when retrieving a list of products.
	Instancevalues *ProductInstancevalues `json:"instancevalues,omitempty"`

	// Start of sales
	Salestartts Time `json:"salestartts,omitempty"`

	// End of sales
	Saleendts Time `json:"saleendts,omitempty"`

	// Sales is active for these saleschannels
	//
	// Note: Not set when retrieving a list of products.
	Saleschannels []int64 `json:"saleschannels"`

	// Translations for the product properties
	//
	// Note: Not set when retrieving a list of products.
	Translations map[string]string `json:"translations,omitempty"`

	// Created timestamp
	//
	// Note: Ignored when creating a new product.
	//
	// Note: Ignored when updating an existing product.
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new product.
	//
	// Note: Ignored when updating an existing product.
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new product.
	//
	// Note: Ignored when updating an existing product.
	Isarchived bool `json:"isarchived,omitempty"`
}

// Set of parameters used to filter seat ranks.
//
// More info: see seat rank
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/SeatRank), the getlist
// operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_seatranks/getlist)
// and the seat ranks endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_seatranks).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/SeatRankQuery).
type SeatRankQuery struct {
	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`
}

// A single seat rank.
//
// More info: see the get operation
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_seatranks/get) and
// the seat ranks endpoint
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_seatranks).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/SeatRank).
type SeatRank struct {
	// Unique ID
	//
	// Note: Ignored when creating a new seat rank.
	//
	// Note: Ignored when updating an existing seat rank.
	Id int64 `json:"id,omitempty"`

	// Name for the seat rank
	Name string `json:"name,omitempty"`

	// Created timestamp
	//
	// Note: Ignored when creating a new seat rank.
	//
	// Note: Ignored when updating an existing seat rank.
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new seat rank.
	//
	// Note: Ignored when updating an existing seat rank.
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new seat rank.
	//
	// Note: Ignored when updating an existing seat rank.
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
	Typeid int64 `json:"typeid,omitempty"`

	// Name for the filter
	Description string `json:"description,omitempty"`

	// The sql clause that defines how the filter will work
	Sqlclause string `json:"sqlclause,omitempty"`

	// The type of filter definition defines the UI and resulting parameters that will
	// be used when a user selects the filter. The possible values can be found here
	// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_system_filterdefinitions).
	Filtertype int64 `json:"filtertype,omitempty"`

	// For certain filter types, the user must select a value from a list. The
	// checklistquery contains the sql clause to retrieve the list of available values.
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
	Id int64 `json:"id,omitempty"`

	// Name for the lock type
	Name string `json:"name,omitempty"`

	// Indicates whether this lock is a hard lock (meaning that it normally never will
	// be released and does not count for the inventory) or a soft lock
	Ishardlock bool `json:"ishardlock,omitempty"`

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
	// Note: Ignored when creating a new order fee definition.
	//
	// Note: Ignored when creating a new order fee.
	Id int64 `json:"id,omitempty"`

	// Name for the order fee
	Name string `json:"name,omitempty"`

	// Type of the order fee. Can be Automatic (2401) or Script (2402)
	Typeid int64 `json:"typeid,omitempty"`

	// Definition of the rule that defines when the order fee will be applied
	//
	// Note: Not set when retrieving a list of order fee definitions.
	//
	// Note: Not set when retrieving a list of order fees.
	Rule *OrderfeeRule `json:"rule,omitempty"`

	// Created timestamp
	//
	// Note: Ignored when creating a new order fee definition.
	//
	// Note: Ignored when creating a new order fee.
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new order fee definition.
	//
	// Note: Ignored when creating a new order fee.
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new order fee definition.
	//
	// Note: Ignored when creating a new order fee.
	Isarchived bool `json:"isarchived,omitempty"`

	// Archived timestamp
	//
	// Note: Ignored when creating a new order fee definition.
	//
	// Note: Ignored when creating a new order fee.
	Archivedts Time `json:"archivedts,omitempty"`
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
	Id int64 `json:"id,omitempty"`

	// Name of the payment method
	Name string `json:"name,omitempty"`

	// Internal remark, will not be shown to customers
	Internalremark string `json:"internalremark,omitempty"`

	// Type of the paymentmethod.
	Paymentmethodtypeid int64 `json:"paymentmethodtypeid,omitempty"`

	// Specific configuration for the payment method, content depends on the payment
	// method type.
	//
	// Note: Not set when retrieving a list of payment methods.
	Config map[string]interface{} `json:"config,omitempty"`

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

	// Type for the payment scenario. Can be 'Immediate payment' (2701), 'Mollie bank
	// transfer' (2702), 'Regular bank transfer' (2703), 'Deferred online payment'
	// (2704), 'Deferred other' (2705).
	Typeid int64 `json:"typeid,omitempty"`

	// Rules that define when an order becomes overdue. Not used for type 2701.
	//
	// Note: Not set when retrieving a list of payment scenarios.
	Overdueparameters *PaymentscenarioOverdueParameters `json:"overdueparameters,omitempty"`

	// Rules that define when an order becomes expired. Not used for type 2701.
	//
	// Note: Not set when retrieving a list of payment scenarios.
	Expiryparameters *PaymentscenarioExpiryParameters `json:"expiryparameters,omitempty"`

	// Rules that define in what conditions this payment scenario is available
	//
	// Note: Not set when retrieving a list of payment scenarios.
	Availability *PaymentscenarioAvailability `json:"availability,omitempty"`

	// Set of payment methods that are linked to this payment scenario. Depending on
	// the type, this field has different usage.
	Paymentmethods []int64 `json:"paymentmethods"`

	// Link to the order mail template that will be sent as payment instruction. Can be
	// 0 to indicate that no mail should be sent. Not used for type 2701.
	OrdermailtemplateidPaymentinstruction int64 `json:"ordermailtemplateid_paymentinstruction,omitempty"`

	// Link to the order mail template that will be sent when the order is overdue. Can
	// be 0 to indicate that no mail should be sent. Not used for type 2701.
	OrdermailtemplateidOverdue int64 `json:"ordermailtemplateid_overdue,omitempty"`

	// Link to the order mail template that will be sent when the order is expired. Can
	// be 0 to indicate that no mail should be sent. Not used for type 2701.
	OrdermailtemplateidExpiry int64 `json:"ordermailtemplateid_expiry,omitempty"`

	// Bank account number to be used. Only used for type 2703 (Regular bank transfer)
	Bankaccountnumber string `json:"bankaccountnumber,omitempty"`

	// BIC code for the bank account number. Only used for type 2703 (Regular bank
	// transfer)
	Bankaccountbic string `json:"bankaccountbic,omitempty"`

	// Beneficiary for the bank account number. Only used for type 2703 (Regular bank
	// transfer)
	Bankaccountbeneficiary string `json:"bankaccountbeneficiary,omitempty"`

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

	// The ID of the order mail template to use for sending confirmations. Can be 0 to
	// indicate that no mail should be sent
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

// A subscriber record to sync state back to Ticketmatic
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/SubscriberSync).
type SubscriberSync struct {
	// Subscriber e-mail
	Email string `json:"email,omitempty"`

	// Previous value of the email field, to indicate a changed e-mail address.
	//
	// Used to find the correct contact. The normal email field will be used when this
	// field is ommitted or empty.
	Oldemail string `json:"oldemail,omitempty"`

	// Subscriber first name
	Firstname string `json:"firstname,omitempty"`

	// Subscriber last name
	Lastname string `json:"lastname,omitempty"`

	// Whether or not the subscriber is still subscribed
	Subscribed bool `json:"subscribed,omitempty"`
}

// A newly created communication
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/SubscriberCommunication).
type SubscriberCommunication struct {
	// Name of the communication
	Name string `json:"name,omitempty"`

	// Optional description of the communication
	Remark string `json:"remark,omitempty"`

	// Timestamp for the communication
	Ts Time `json:"ts,omitempty"`

	// E-mail addresses to which the communication has been sent
	Addresses []string `json:"addresses"`
}

// Required data for creating a query on the public data model.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/QueryRequest).
type QueryRequest struct {
	// Actual query to execute
	Query string `json:"query,omitempty"`

	// Optional offset for the result. Default 0
	Offset int64 `json:"offset,omitempty"`

	// Optional limit for the result. Default 100
	Limit int64 `json:"limit,omitempty"`
}

// Result of a query on the public data model.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/QueryResult).
type QueryResult struct {
	// The number of rows in the result
	Nbrofresults int64 `json:"nbrofresults,omitempty"`

	// The actual resulting rows
	Results []map[string]interface{} `json:"results"`
}

// Required data for requesting the ticketsprocessedstatistics.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/TicketsprocessedRequest).
type TicketsprocessedRequest struct {
	// Start date of the period
	Startts string `json:"startts,omitempty"`

	// End date of the period
	Endts string `json:"endts,omitempty"`

	// How the results are grouped. Values can be 'day' or 'month'
	Groupby string `json:"groupby,omitempty"`
}

// Statistics on the number of tickets processed in a certain period.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/types/TicketsprocessedStatistics).
type TicketsprocessedStatistics struct {
	// Start of the period
	Ts Time `json:"ts,omitempty"`

	// The number of tickets processed
	Processed int64 `json:"processed,omitempty"`

	// The number of tickets sold online
	Soldonline int64 `json:"soldonline,omitempty"`
}
