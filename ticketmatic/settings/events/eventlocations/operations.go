package eventlocations

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// Get a list of event locations
func Getlist(client *ticketmatic.Client, params *ticketmatic.EventLocationParameters) ([]*ticketmatic.ListEventLocation, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/events/eventlocations")
	r.AddParameter("includearchived", params.Includearchived)
	r.AddParameter("lastupdatesince", params.Lastupdatesince)
	r.AddParameter("filter", params.Filter)

	var obj []*ticketmatic.ListEventLocation
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get a single event location
func Get(client *ticketmatic.Client, id int) (*ticketmatic.EventLocation, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/events/eventlocations/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.EventLocation
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new event location
func Create(client *ticketmatic.Client, data *ticketmatic.CreateEventLocation) (*ticketmatic.EventLocation, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/events/eventlocations")
	r.Body(data)

	var obj *ticketmatic.EventLocation
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing event location
func Update(client *ticketmatic.Client, id int, data *ticketmatic.UpdateEventLocation) (*ticketmatic.EventLocation, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/events/eventlocations/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.EventLocation
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove an event location
//
// Event locations are archivable: this call won't actually delete the object from
// the database. Instead, it will mark the object as archived, which means it won't
// show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/events/eventlocations/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
