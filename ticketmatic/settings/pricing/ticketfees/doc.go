// Ticket fees are schemes that define what fee will be added to the ticket price,
// depending on the price type and the sales channel used when the ticket is sold.
// The fee can either be a fixed cost or a percentage cost.
//
// By linking a ticketfee scheme to an event, the fees that will be applied
// to tickets for that event are defined. The same ticket fee can be linked to
// multiple events. Changing a ticket fee scheme will automatically update this for
// all linked events (the new fees will only be applied for new orders, fees for
// tickets that are already sold will not change).
//
// # Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/settings/pricing/ticketfees).
package ticketfees
