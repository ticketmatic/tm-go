// A delivery scenario defines how an order will be delivered.
//
// You can define an order fee for a delivery scenario and you can define in which
// conditions a delivery scenario is available. Typically, the customer will select
// the delivery scenario most appropriate for him.
//
// Examples include: Print at home, Mobile, Send by mail, Retrieve at desk
//
// Types
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
// Availability
//
// When a delivery scenario can be used can be defined in two ways:
//
// * By specifying a set of sales channels (required)
//
// * By writing a script (optional)
//
// In its simplest form, a DeliveryscenarioAvailability looks like this:
//
//
//    {
//        "saleschannels": [1, 2]
//    }
//
//
//
// This defines that the delivery scenario is available when used in the context of
// saleschannel 1 or 2.
//
// More complex logic can be specified by writing a small piece of JavaScript
// (http://en.wikipedia.org/wiki/JavaScript). To do so, you need to add a usescript
// and script field to the availability:
//
//
//    {
//        "saleschannels": [1, 2],
//        "usescript": true,
//        "script": "// script here"
//    }
//
//
//
// Note that the list of sales channel IDs is still required: the script can only
// restrict this set further.
//
// A simple example of a delivery scenario script:
//
//
//    return order.tickets.length < 3 && saleschannel.typeid == 3002;
//
//
//
// This script states that the current delivery scenario is only available if the
// amount of tickets in the order is less than 3 and the current sales channel is a
// web sales channel.
//
// With this script the DeliveryscenarioAvailability would look like this:
//
//
//    {
//        "saleschannels": [1, 2],
//        "usescript": true,
//        "script": "return order.tickets.length < 3 && saleschannel.typeid == 3002;"
//    }
//
//
//
// The following variables are available in the script:
//
// * order
//
// * saleschannel
//
// You can use any valid JavaScript syntax (including conditionals and loops). Note
// that each script has a strict time limit.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://apps.ticketmatic.com/#/knowledgebase/api/settings_ticketsales_deliveryscenarios).
package deliveryscenarios
