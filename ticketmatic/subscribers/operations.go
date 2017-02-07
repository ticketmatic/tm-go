package subscribers

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// Sync mailtool changes to Ticketmatic
func Sync(client *ticketmatic.Client, data []*ticketmatic.SubscriberSync) error {
	r := client.NewRequest("POST", "/{accountname}/subscribers/sync", "json")
	r.Body(data, "json")

	return r.Run(nil)
}

// Create a new communcation in Ticketmatic based on a list of subscriber
// emailaddresses
func Communications(client *ticketmatic.Client, data *ticketmatic.SubscriberCommunication) error {
	r := client.NewRequest("POST", "/{accountname}/subscribers/communications", "json")
	r.Body(data, "json")

	return r.Run(nil)
}
