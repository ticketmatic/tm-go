package orders

import (
	"os"
	"testing"

	"github.com/ticketmatic/tm-go/ticketmatic"
	"github.com/ticketmatic/tm-go/ticketmatic/events"
)

func TestGet(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	list, err := Getlist(c, &ticketmatic.OrderQuery{
		Output: "withlookup",
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(list.Data) <= 0 {
		t.Errorf("Unexpected list.Data length, got %#v, expected greater than %#v", len(list.Data), 0)
	}

	order, err := Get(c, list.Data[0].Orderid)
	if err != nil {
		t.Fatal(err)
	}

	if order.Orderid != list.Data[0].Orderid {
		t.Errorf("Unexpected order.Orderid, got %#v, expected %#v", order.Orderid, list.Data[0].Orderid)
	}

}

func TestCreate(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	order, err := Create(c, &ticketmatic.CreateOrder{
		Saleschannelid: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	if order.Orderid == 0 {
		t.Errorf("Unexpected order.Orderid, got %#v, expected different value", order.Orderid)
	}

	if order.Saleschannelid != 1 {
		t.Errorf("Unexpected order.Saleschannelid, got %#v, expected %#v", order.Saleschannelid, 1)
	}

	updated, err := Update(c, order.Orderid, &ticketmatic.UpdateOrder{
		Customerid:         777701,
		Deliveryscenarioid: 2,
		Paymentscenarioid:  3,
	})
	if err != nil {
		t.Fatal(err)
	}

	if updated.Orderid != order.Orderid {
		t.Errorf("Unexpected updated.Orderid, got %#v, expected %#v", updated.Orderid, order.Orderid)
	}

	if updated.Deliveryscenarioid != 2 {
		t.Errorf("Unexpected updated.Deliveryscenarioid, got %#v, expected %#v", updated.Deliveryscenarioid, 2)
	}

	if updated.Paymentscenarioid != 3 {
		t.Errorf("Unexpected updated.Paymentscenarioid, got %#v, expected %#v", updated.Paymentscenarioid, 3)
	}

	if updated.Customerid != 777701 {
		t.Errorf("Unexpected updated.Customerid, got %#v, expected %#v", updated.Customerid, 777701)
	}

	ticketsadded, err := Addtickets(c, order.Orderid, &ticketmatic.AddTickets{
		Tickets: []*ticketmatic.CreateTicket{
			&ticketmatic.CreateTicket{
				Tickettypepriceid: 584,
			},
			&ticketmatic.CreateTicket{
				Tickettypepriceid: 584,
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(ticketsadded.Order.Tickets) != 2 {
		t.Errorf("Unexpected ticketsadded.Order.Tickets length, got %#v, expected %#v", len(ticketsadded.Order.Tickets), 2)
	}

	_, err = Confirm(c, order.Orderid)
	if err != nil {
		t.Fatal(err)
	}

	ticketids := []int64{
		ticketsadded.Order.Tickets[0].Id,
	}

	updated2, err := Updatetickets(c, order.Orderid, &ticketmatic.UpdateTickets{
		Operation: "setticketholders",
		Params: map[string]interface{}{
			"ticketholderids": []interface{}{777701},
		},
		Tickets: ticketids,
	})
	if err != nil {
		t.Fatal(err)
	}

	if updated2.Tickets[0].Ticketholderid != 777701 {
		t.Errorf("Unexpected updated2.Tickets[0].Ticketholderid, got %#v, expected %#v", updated2.Tickets[0].Ticketholderid, 777701)
	}

	deleted, err := Deletetickets(c, order.Orderid, &ticketmatic.DeleteTickets{
		Tickets: ticketids,
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(deleted.Tickets) != 1 {
		t.Errorf("Unexpected deleted.Tickets length, got %#v, expected %#v", len(deleted.Tickets), 1)
	}

}

func TestCreatequeued(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	_, err = Create(c, &ticketmatic.CreateOrder{
		Events: []int64{
			777714,
		},
		Saleschannelid: 1,
	})
	var exc *ticketmatic.QueueStatus
	if err != nil {
		if e, ok := err.(*ticketmatic.RateLimitError); ok {
			exc = e.Status
		} else {
			t.Fatal(err)
		}
	}

	if exc == nil {
		t.Fatal("Expected a rate limiting error")
	}

	if exc.Id == "" {
		t.Errorf("Unexpected exc.Id, got %#v, expected different value", exc.Id)
	}

}

func TestAddticketsqueued(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	order, err := Create(c, &ticketmatic.CreateOrder{
		Saleschannelid: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	if order.Orderid == 0 {
		t.Errorf("Unexpected order.Orderid, got %#v, expected different value", order.Orderid)
	}

	if order.Saleschannelid != 1 {
		t.Errorf("Unexpected order.Saleschannelid, got %#v, expected %#v", order.Saleschannelid, 1)
	}

	ttps, err := events.Get(c, 777713)
	if err != nil {
		t.Fatal(err)
	}

	if ttps.Id == 0 {
		t.Errorf("Unexpected ttps.Id, got %#v, expected different value", ttps.Id)
	}

	_, err = Addtickets(c, order.Orderid, &ticketmatic.AddTickets{
		Tickets: []*ticketmatic.CreateTicket{
			&ticketmatic.CreateTicket{
				Tickettypepriceid: ttps.Prices.Contingents[0].Pricetypes[0].Tickettypepriceid,
			},
		},
	})
	var exc *ticketmatic.QueueStatus
	if err != nil {
		if e, ok := err.(*ticketmatic.RateLimitError); ok {
			exc = e.Status
		} else {
			t.Fatal(err)
		}
	}

	if exc == nil {
		t.Fatal("Expected a rate limiting error")
	}

	if exc.Id == "" {
		t.Errorf("Unexpected exc.Id, got %#v, expected different value", exc.Id)
	}

}
