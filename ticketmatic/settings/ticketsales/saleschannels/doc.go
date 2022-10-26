// In Ticketmatic, each order is created in the context of a sales channel
// (https://www.ticketmatic.com/docs/tickets/configure_ticket_sales/saleschannels).
//
// # Types
//
// There are 3 types of sales channels:
//
// * Desk (3001): for all orders in the Orders application
//
// * Web (3002): for orders from websites
//
// * External (3003): for orders sold by partners
//
// There is always exactly one sales channel of type Desk. Additionally you can
// define multiple sales channels of types Web and External.
//
// # Order mails
//
// Each sales channel can be configured with an order mail template. This template
// is used to send order confirmations to customers. When (and if) this mail is
// sent is defined by the payment method, but can be overridden per saleschannel.
//
// # Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/saleschannels).
package saleschannels
