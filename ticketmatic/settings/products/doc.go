// Product types
//
// There are 4 product types:
//
// * Simple (26001)
//
// * Voucher (26002)
//
// * Fixed bundle (26003)
//
// * Option bundle (26004)
//
// You can find more information about these types here
// (https://www.ticketmatic.com/docs/events/products).
//
// Each product can have one or more variants (t-shirt sizes for example). These
// variants are controlled by properties and each property can have one or more
// values. For instance if you want to create a t-shirt product that has multiple
// sizes, you should create a single product, with one property (size) which has
// multiple values (Small, Medium, Large).
//
// To create a price for each specific variant, the instancevalues property should
// be populated. A product should always have a default price configured and if
// different prices for specific variants of a product are desired, these can be
// controlled with instance value exceptions.
//
// Each instance value exception contains one or more property values to match.
// If all the property values match, this instance value exception is selected to
// control the price for that specific product variant.
//
// For non simple products (like vouchers and bundle) the instancevalues also
// contain other content, such as which voucherid should be used when generating a
// new voucher or which tickettypeprices should be used for the fixed bundle.
//
// Take a look at the examples
// (https://www.ticketmatic.com/docs/api/settings/products/create/) for creating a
// new product to see how properties and instancevalues work.
//
// # Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/settings/products).
package products
