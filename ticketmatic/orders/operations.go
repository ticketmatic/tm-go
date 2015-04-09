package orders

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// Add tickets to order
func Addtickets(client *ticketmatic.Client, id int, data []*ticketmatic.Ticket) (*ticketmatic.Order, error) {
	r := client.NewRequest("POST", "/{accountname}/orders/{id}/tickets")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.Order
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
