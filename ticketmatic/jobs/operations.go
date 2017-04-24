package jobs

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// Get job
//
// Returns info on a job including the current status.
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.JobResult, error) {
	r := client.NewRequest("GET", "/{accountname}/jobs/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.JobResult
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
