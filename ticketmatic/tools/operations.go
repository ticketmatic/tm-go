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

// Get statistics on the tickets processed during a certain period
//
// Use this method to retrieve the statistics on the number of tickets processed
// and sold online during a certain period. The results can be grouped by day or
// month. These statistics are often used as basis for invoicing or reporting.
func Ticketsprocessedstatistics(client *ticketmatic.Client, params *ticketmatic.TicketsprocessedRequest) (*ticketmatic.TicketsprocessedStatistics, error) {
	r := client.NewRequest("GET", "/{accountname}/tools/ticketsprocessedstatistics")
	if params != nil {
		r.AddParameter("startts", params.Startts)
		r.AddParameter("endts", params.Endts)
		r.AddParameter("groupby", params.Groupby)
	}

	var obj *ticketmatic.TicketsprocessedStatistics
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
