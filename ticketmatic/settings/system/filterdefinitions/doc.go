// Filter definitions define filters that can be used in the backoffice.
//
// The field typeid defines to which object a filter definition belongs:
//
// * 10001: order
//
// * 10002: customer
//
// * 10003: event
//
// * 10004: ticket
//
// * 10005: payment
//
// The field filtertype defines the kind of filter:
//
// * 0: Toggle
//
// * 1: Checklist
//
// * 2: Date range
//
// * 5: String
//
// * 6: Number range
//
// * 9: Tickets bought
//
// * 10: Option set
//
// # The field sqlclause defines the actual filter
//
// You can find more information about setting up filter definitions here
// (https://www.ticketmatic.com/docs/setup/filterdefinitions)
//
// # Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/settings/system/filterdefinitions).
package filterdefinitions
