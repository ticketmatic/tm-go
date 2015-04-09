package events

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// Get a list of events
func List(client *ticketmatic.Client, params *ticketmatic.EventParameters) ([]*ticketmatic.ListEvent, error) {
	r := client.NewRequest("GET", "/{accountname}/events")
	r.AddParameter("includearchived", params.Includearchived)

	var obj []*ticketmatic.ListEvent
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get an event
func Single(client *ticketmatic.Client, id int) (*ticketmatic.Event, error) {
	r := client.NewRequest("GET", "/{accountname}/events/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.Event
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
