package tools

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// Execute a query on the public data model
//
// Use this method to execute random (read-only) queries on the public data model.
// Remark that this is not meant for long-running queries or for returning large
// resultsets. If the query executes too long or uses too much memory, an exception
// will be returned.
func Queries(client *ticketmatic.Client, data *ticketmatic.QueryRequest) (*ticketmatic.QueryResult, error) {
	r := client.NewRequest("POST", "/{accountname}/tools/queries")
	r.Body(data)

	var obj *ticketmatic.QueryResult
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
