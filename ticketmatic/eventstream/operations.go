package eventstream

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// Poll eventstream
//
// Poll the eventstream.
func Eventstream(client *ticketmatic.Client, params *ticketmatic.EventstreamRequest) (*ticketmatic.EventstreamResult, error) {
	r := client.NewRequest("GET", "/{accountname}/eventstream", "json")
	if params != nil {
		r.AddParameter("id", params.Id)
		r.AddParameter("eventtypes", params.Eventtypes)
		r.AddParameter("ts", params.Ts)
	}

	var obj *ticketmatic.EventstreamResult
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
