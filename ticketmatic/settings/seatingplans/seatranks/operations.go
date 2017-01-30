package seatranks

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.SeatRank `json:"data"`

	// The total number of results that are available without considering limit and offset, useful for paging.
	NbrOfResults int `json:"nbrofresults"`
}

// Get a list of seat ranks
func Getlist(client *ticketmatic.Client, params *ticketmatic.SeatRankQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/seatingplans/seatranks")
	if params != nil {
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

// Get a single seat rank
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.SeatRank, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/seatingplans/seatranks/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.SeatRank
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new seat rank
func Create(client *ticketmatic.Client, data *ticketmatic.SeatRank) (*ticketmatic.SeatRank, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/seatingplans/seatranks")
	r.Body(data)

	var obj *ticketmatic.SeatRank
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing seat rank
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.SeatRank) (*ticketmatic.SeatRank, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/seatingplans/seatranks/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.SeatRank
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a seat rank
//
// Seat ranks are archivable: this call won't actually delete the object from the
// database. Instead, it will mark the object as archived, which means it won't
// show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/seatingplans/seatranks/{id}")
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
	r := client.NewRequest("GET", "/{accountname}/settings/seatingplans/seatranks/{id}/translate")
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
	r := client.NewRequest("PUT", "/{accountname}/settings/seatingplans/seatranks/{id}/translate")
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
