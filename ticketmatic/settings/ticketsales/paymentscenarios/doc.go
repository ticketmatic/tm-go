// A payment scenario defines how a customer will pay for an order. This is not
// necessarily linked to a specific payment. When a customer selects a payment
// scenario for an order, a process is started. The actual process depends on the
// type of payment scenario:
//
// * Immediate payment: in this case, the payment must be executed immediately by
// using one of the payment methods that is linked to the payment scenario.
//
// * Deferred payment: in this case, the payment can be executed later and the
// order can be confirmed without payment. You can configure overdue and expiry
// parameters to define how long the order can stay unpaid before considering the
// order as overdue or expired and sending out reminder mails or even cancelling
// the order automatically
//
// An order fee can defined for a payment scenario: this fee will automatically be
// added to the order when the payment scenario is selected.false
//
// A payment scenario is available for one or more sales channels. A custom script
// can further refine when the payment scenario is available.
//
// Examples include: Credit card, Bank transfer, Cash
//
// # Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/paymentscenarios).
package paymentscenarios
