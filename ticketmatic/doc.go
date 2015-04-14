// The Ticketmatic API is a REST
// (http://en.wikipedia.org/wiki/Representational_state_transfer) API. The API is
// designed to have predictable, resource-oriented URLs and to use HTTP response
// codes to indicate API errors.
//
// You can make requests to Ticketmatic API from a terminal or browser, or from
// within any code. The examples in the documentation use curl, but you can use
// wget, libcurl or any other method to GET and POST in your preferred language.
// Your preferred coding language should include libraries to support making HTTP
// requests and parsing responses.
//
// URI
//
// The URIs for endpoints are served at
//
//
//    https://apps.ticketmatic.com/api/{accountname}/
//
//
//
// Note that Ticketmatic API is secure, so use HTTPS whenever you access any
// endpoints. Calls made over plain HTTP will fail.
//
// Getting started
//
// Be sure to read up on the core concepts first:
//
// * Authentication
// (https://apps.ticketmatic.com/#/knowledgebase/api/coreconcepts_authentication)
//
// * Custom fields
// (https://apps.ticketmatic.com/#/knowledgebase/api/coreconcepts_customfields)
//
// * Error handling
// (https://apps.ticketmatic.com/#/knowledgebase/api/coreconcepts_errors)
//
// * Translations
// (https://apps.ticketmatic.com/#/knowledgebase/api/coreconcepts_translations)
//
// Afterwards, check the pages of the individual endpoints you wish to use.
//
// Audit log
//
// All calls to the API are logged, and you can inspect these using the Audit Log
// tool in your account.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api).
package ticketmatic
