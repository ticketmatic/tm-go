package subscribers

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// Sync mailtool changes to Ticketmatic
func Sync(client *ticketmatic.Client, data []*ticketmatic.SubscriberSync) error {
	r := client.NewRequest("POST", "/{accountname}/subscribers/sync")
	r.Body(data)

	return r.Run(nil)
}
