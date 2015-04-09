package diagnostics

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// Get the backend time
//
// Returns the current system time, in UTC.
//
// The returned timestamp uses the ISO-8601 format.
func Time(client *ticketmatic.Client) (*ticketmatic.Timestamp, error) {
	r := client.NewRequest("GET", "/{accountname}/diagnostics/time")

	var obj *ticketmatic.Timestamp
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
