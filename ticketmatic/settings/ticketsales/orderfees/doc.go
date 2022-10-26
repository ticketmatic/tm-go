// Order fees determine which extra costs or discounts can be added to an order.
//
// # Types
//
// There are two types:
//
// * Automatic (2401): Order fees of this type are added automatically to an order
// if the order matches the rules of the orderfee. Rules can be matched based on a
// saleschannel, deliveryscenario and/or paymentscenario.
//
// * Script (2402): Order fees of this type consist of a piece of custom
// javascript. When the order cost is evaluated, this javascript is executed and
// the return value is added to the order.
//
// * Manual (2403): Manually specified order fees.
//
// # Automatic
//
// An automatic order fee can have multiple rules. The general rules should be
// defined first and the exceptions later. Whenever the rules of an (automatic)
// order fee are checked it will execute the last rule that matches.
//
// A match will occur if the order has a saleschannel, delivery
// scenario and payment scenario that matches the OrderfeeAutoRule
// (https://www.ticketmatic.com/docs/api/types/OrderfeeAutoRule). If it's matched
// (and it is was the last rule that matched) the defined value will be added,
// based on the status (fixedfee or percentagefee).
//
// # Script
//
// An order fee of type script consists of a javascript. This
// javascript is always executed and the number that is returned
// is added to the order amount. This script has an order object
// available. You can find more info about writing order scripts here
// (https://www.ticketmatic.com/docs/tickets/configure_ticket_sales/orderfees).
//
// # Manual
//
// A manual order fee is never applied automatically. It can be used when importing
// historic orders, to manually set a price per order.
//
// # Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/orderfees).
package orderfees
