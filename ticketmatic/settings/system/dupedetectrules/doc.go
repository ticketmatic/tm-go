// Dupe detect rules are used for contact deduplication.
//
// Each rule has a set of criteria: a contact is considered a duplicate when all of
// the criteria match.
//
// # Fields and matchers
//
// The following field / matcher combinations are possible:
//
// | Field | Matchers allowed | |---------------|------------------| | email |
// exact | | firstname | exact, fuzzy | | lastname | exact, fuzzy | | address |
// normalized | | phonenumber | normalized |
//
// * exact: Identical match
//
// * fuzzy: Some mis-spelling is allowed
//
// * normalized: Normalization is applied before matching
//
// # Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/settings/system/dupedetectrules).
package dupedetectrules
