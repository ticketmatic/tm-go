// A delivery scenario defines how an order will be delivered.
//
// You can define an order fee for a delivery scenario and you can define in which
// conditions a delivery scenario is available. Typically, the customer will select
// the delivery scenario most appropriate for him.
//
// Examples include: Print at home, Mobile, Send by mail, Retrieve at desk
//
// # Types
//
// There are 3 types of delivery scenarios:
//
// * Manual (2501): delivery of the tickets will happen manually by the seller.
// This could be via a daily or weekly print.
//
// * Automatic after full payment (2502): Ticketmatic will deliver the tickets
// automatically once the outstanding balance for the order reaches 0 (full
// payment).
//
// * Automatic after confirmation (2503): Ticketmatic will deliver the tickets
// automatically as soon as the order is confirmed.
//
// # Availability
//
// The full rules for defining when a delivery scenario
// can be used are defined in DeliveryscenarioAvailability
// (https://www.ticketmatic.com/docs/api/types/DeliveryscenarioAvailability).
//
// # Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/deliveryscenarios).
package deliveryscenarios
