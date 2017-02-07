package ratelimiting

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// Request a status update
//
// Request a new rate limiting status update. See rate limiting
// (https://www.ticketmatic.com/docs/api/ratelimiting) for more details on rate
// limiting.
func Status(client *ticketmatic.Client, id int64) (*ticketmatic.QueueStatus, error) {
	r := client.NewRequest("POST", "/{accountname}/ratelimiting/status/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.QueueStatus
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
