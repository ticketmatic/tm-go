package contacts

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.Contact `json:"data"`

	// The total number of results that are available without considering limit and offset, useful for paging.
	NbrOfResults int `json:"nbrofresults"`
}

// Get a list of contacts
func Getlist(client *ticketmatic.Client, params *ticketmatic.ContactQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/contacts", "json")
	if params != nil {
		r.AddParameter("filter", params.Filter)
		r.AddParameter("includearchived", params.Includearchived)
		r.AddParameter("lastupdatesince", params.Lastupdatesince)
		r.AddParameter("limit", params.Limit)
		r.AddParameter("offset", params.Offset)
		r.AddParameter("orderby", params.Orderby)
		r.AddParameter("output", params.Output)
		r.AddParameter("searchterm", params.Searchterm)
	}

	var obj *List
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get a single contact
//
// To retrieve a contact based on the e-mail address, pass 0 as the id and supply
// an email parameter.
func Get(client *ticketmatic.Client, id int64, params *ticketmatic.ContactGetQuery) (*ticketmatic.Contact, error) {
	r := client.NewRequest("GET", "/{accountname}/contacts/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	if params != nil {
		r.AddParameter("email", params.Email)
	}

	var obj *ticketmatic.Contact
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new contact
//
// Creates a new contact
func Create(client *ticketmatic.Client, data *ticketmatic.Contact) (*ticketmatic.Contact, error) {
	r := client.NewRequest("POST", "/{accountname}/contacts", "json")
	r.Body(data, "json")

	var obj *ticketmatic.Contact
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Update a contact
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.Contact) (*ticketmatic.Contact, error) {
	r := client.NewRequest("PUT", "/{accountname}/contacts/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	var obj *ticketmatic.Contact
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a contact
//
// Contacts are archivable: this call won't actually delete the object from the
// database. Instead, it will mark the contact as deleted, which means it won't
// show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/contacts/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}

// Batch operations
//
// Apply batch operations to a set of contacts.
//
// The parameters required are specific to the type of operation.
//
// What will be affected?
//
// If you don't specify anything, the batch operation will be applied to all
// contacts.
//
// To restrict the operation to a strict set of contacts, pass in the IDs:
//
//
//    ids: [1, 2, 3]
//
//
//
// This will only apply the operation to contacts with ID 1, 2 and 3.
//
// You can also apply the operation to all contacts except for a set of IDs, using
// excludeids:
//
//
//    excludeids: [4, 5]
//
//
//
// This will apply the operation to all contacts, except for contacts with ID 4 and
// 5.
//
// Batch operations
//
// The following operations are supported:
//
// * addrelationtypes: Adds the specified relation types to the selection of
// contacts. The parameters object should contain an ids field with a set of
// relation type IDs.
//
// * removerelationtypes: Remove the specified relation types from the selection of
// contacts. The parameters object should contain an ids field with a set of
// relation type IDs.
//
// * delete: Deletes the selection of contacts.
//
// * subscribe: Subscribes the selected contacts using the mailing tool.
//
// * unsubscribe: Unsubscribes the selected contacts using the mailing tool.
//
// * sendselection: Send a selection of contacts to the mailing tool. These
// contacts can then be used to send out a mailing. The parameters object can
// optionally contain a name field that will be used to identify the selection.
//
// * update: Update a specific field for the selection of contacts. See
// BatchContactParameters
// (https://www.ticketmatic.com/docs/api/types/BatchContactParameters) for more
// info.
//
// * merge: Merges the selected contacts into a specified contact. The primary
// parameter should be supplied to indicate which contact gets preference.
func Batch(client *ticketmatic.Client, data *ticketmatic.BatchContactOperation) error {
	r := client.NewRequest("POST", "/{accountname}/contacts/batch", "json")
	r.Body(data, "json")

	return r.Run(nil)
}

// Import contact
//
// Up to 1000 contacts can be sent per call.
func Import(client *ticketmatic.Client, data []*ticketmatic.Contact) ([]*ticketmatic.ContactImportStatus, error) {
	r := client.NewRequest("POST", "/{accountname}/contacts/import", "json")
	r.Body(data, "json")

	var obj []*ticketmatic.ContactImportStatus
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Reserve contact IDs
//
// Importing contacts with the specified IDs is only possible when those IDs fall
// in the reserved ID range. Use this call to reserve a range of contact IDs. Any
// unused ID lower than or equal to the specified ID will be reserved. New contacts
// will receive IDs higher than the specified ID.
func Reserve(client *ticketmatic.Client, data *ticketmatic.ContactIdReservation) (*ticketmatic.ContactIdReservation, error) {
	r := client.NewRequest("POST", "/{accountname}/contacts/import/reserve", "json")
	r.Body(data, "json")

	var obj *ticketmatic.ContactIdReservation
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Send a reset password request for this contact.
//
// An e-mail will be sent with reset instructions. An error will be thrown when
// this contact does not have a linked password account.
func Resetpassword(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("POST", "/{accountname}/contacts/{id}/resetpassword", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}

// Create a password account
//
// Link this contact with a password account to be used in authenticated widgets.
// An e-mail will be sent with password instructions.
func Createaccount(client *ticketmatic.Client, id int64) (*ticketmatic.Contact, error) {
	r := client.NewRequest("POST", "/{accountname}/contacts/{id}/account", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.Contact
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Delete password account
//
// Remove the linked account from this contact. The contact will not be able to
// login anymore in authenticated widgets.
func Deleteaccount(client *ticketmatic.Client, id int64) (*ticketmatic.Contact, error) {
	r := client.NewRequest("DELETE", "/{accountname}/contacts/{id}/account", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.Contact
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
