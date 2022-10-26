// A payment method defines the method of an actual payment. Each payment is done
// by a specific payment method.
//
// In the payment method, the type and if appropriate the technical parameters
// for performing the payment are defined. These parameters depend on the payment
// gateway used.
//
// Not all payment methods can be used in all contexts.
//
// This list of possible paymentmethod types are:
//
// * Internal (1001)
//
// * Stripe (1005)
//
// * Mollie (1006)
//
// * Voucher (1007)
//
// * PayPal (1008)
//
// # Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/paymentmethods).
package paymentmethods
