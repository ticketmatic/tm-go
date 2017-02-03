package vouchers

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.Voucher `json:"data"`

	// The total number of results that are available without considering limit and offset, useful for paging.
	NbrOfResults int `json:"nbrofresults"`
}

// Get a list of vouchers
func Getlist(client *ticketmatic.Client, params *ticketmatic.VoucherQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/vouchers", "")
	if params != nil {
		r.AddParameter("typeid", params.Typeid)
		r.AddParameter("filter", params.Filter)
		r.AddParameter("includearchived", params.Includearchived)
		r.AddParameter("lastupdatesince", params.Lastupdatesince)
	}

	var obj *List
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get a single voucher
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.Voucher, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/vouchers/{id}", "")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.Voucher
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new voucher
func Create(client *ticketmatic.Client, data *ticketmatic.Voucher) (*ticketmatic.Voucher, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/vouchers", "")
	r.Body(data)

	var obj *ticketmatic.Voucher
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing voucher
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.Voucher) (*ticketmatic.Voucher, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/vouchers/{id}", "")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.Voucher
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a voucher
//
// Vouchers are archivable: this call won't actually delete the object from the
// database. Instead, it will mark the object as archived, which means it won't
// show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/vouchers/{id}", "")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}

// Fetch translatable fields
//
// Returns a dictionary with string values in all languages for each translatable
// field.
//
// See translations
// (https://www.ticketmatic.com/docs/api/coreconcepts/translations) for more
// information.
func Translations(client *ticketmatic.Client, id int64) (map[string]string, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/vouchers/{id}/translate", "")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj map[string]string
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Update translations
//
// Sets updated translation strings.
//
// See translations
// (https://www.ticketmatic.com/docs/api/coreconcepts/translations) for more
// information.
func Translate(client *ticketmatic.Client, id int64, data map[string]string) (map[string]string, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/vouchers/{id}/translate", "")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj map[string]string
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create voucher codes
//
// Creates individual voucher codes.
//
// Codes will be randomly generated unless supplied.
func Createcodes(client *ticketmatic.Client, id int64, data *ticketmatic.AddVoucherCodes) ([]*ticketmatic.VoucherCode, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/vouchers/{id}/codes", "")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj []*ticketmatic.VoucherCode
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Deactivate voucher codes
//
// Deactivates individual voucher codes.
func Deactivatecodes(client *ticketmatic.Client, id int64, data []*ticketmatic.VoucherCode) error {
	r := client.NewRequest("POST", "/{accountname}/settings/vouchers/{id}/deactivatecodes", "")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	return r.Run(nil)
}
