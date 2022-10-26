package diagnostics

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// Get the backend time
//
// Returns the current system time, in UTC.
//
// The returned timestamp uses the ISO-8601 format.
//
// This call does not require an Authorization header to be
// set (it's the only call that allows this) and can be used to
// investigate timestamp issues when trying to sign API requests
// (https://www.ticketmatic.com/docs/api/coreconcepts/authentication).
func Time(client *ticketmatic.Client) (*ticketmatic.Timestamp, error) {
	r := client.NewRequest("GET", "/{accountname}/diagnostics/time", "json")

	var obj *ticketmatic.Timestamp
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
