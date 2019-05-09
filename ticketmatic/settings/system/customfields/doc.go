// Custom fields allow you to extend the Ticketmatic data model with your own data
// fields. See setting up customfields
// (https://www.ticketmatic.com/docs/setup/customfields) for more information.
//
// The list of custom field types are:
//
// * Order (13001)
//
// * Contact (13002)
//
// * Event (13003)
//
// * Ticket (13004)
//
// * Product (13005)
//
// * Pricetype (13006)
//
// * Payment method (13007)
//
// * Payment scenario (13008)
//
// * Delivery scenario (13009)
//
// * Event location (13010)
//
// * User (13011)
//
// * Waiting list request (13012)
//
// * Waiting list request item (13013)
//
// The list of possible custom field field types are:
//
// * String (12001)
//
// * Integer (12002)
//
// * Date (12003)
//
// * Boolean (12004)
//
// * Text (12005)
//
// * Multi-language string (12006)
//
// * Multi-language text (12007)
//
// * Decimal (12008)
//
// * Select (single dropdown) (12009)
//
// * Select (multi dropdown) (12010)
//
// * Select (single optionset) (12011)
//
// * Select (multi checklist) (12012)
//
// * Text with formatting (12013)
//
// Possible edit types:
//
// * Hidden (22001): only via the API
//
// * CRUD (22002): in the backoffice
//
// * Checkout (22003): by the customer, during ticket sales
//
// Possible required types: * None (30001): not required * Everywhere (30002):
// required everywhere * Checkout (30003): Only required during online checkout
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/settings/system/customfields).
package customfields
