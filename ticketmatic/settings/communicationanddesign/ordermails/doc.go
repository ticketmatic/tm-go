// Order mail templates are used to construct e-mails sent to customers.
//
// Before working with order mail templates, be sure to read the order mail design
// guide (https://apps.ticketmatic.com/#/knowledgebase/designer_ordermail).
//
// Types
//
// There are 5 types of order mail templates:
//
// * Confirmation (3101)
//
// * Delivery (3102)
//
// * Payment instructions (3103)
//
// * Overdue notices (3104)
//
// * Expiration notices (3105)
//
// Subject & Body
//
// Both the subject and body fields allow using Twig variables. These are described
// in the order mail design guide
// (https://apps.ticketmatic.com/#/knowledgebase/designer_ordermail).
//
// The body should contain a valid HTML document. Be sure to include a charset
// definition:
//
//
//    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
//
//
//
// Any CSS embedded into the document header will be inlined for proper display in
// mail clients.
//
// Translations
//
// Any HTML tag annotated with a translate attribute is considered translatable.
//
// These translations can be defined using the gettext
// (http://en.wikipedia.org/wiki/Gettext) format. The translations field contains a
// map with language codes as keys and .po files as their values.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_communicationanddesign_ordermails).
package ordermails
