// Price lists are used to define the actual prices that will be available for one
// or more events. You can create a price list for a selection of seat ranks or for
// a simple contingent without seatranks.
//
// In each price list prices are defined for a selection of price types.
// Additionally, conditions for each price type can be defined.
//
// The possible conditions are listed in PricelistPriceCondition
// (https://www.ticketmatic.com/docs/api/types/PricelistPriceCondition).
//
// The prices for an event are defined by linking a price list to the event.
// The same price list can be linked to multiple events.
//
// Changing the price in a price list will automatically change the price in all
// events that have linked this price list. (Remark: the new prices will only
// be applied for new orders, prices for tickets that are already sold will not
// change)
//
// # Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/settings/pricing/pricelists).
package pricelists
