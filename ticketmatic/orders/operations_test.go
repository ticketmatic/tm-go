package orders

import (
	"os"
	"testing"

	"github.com/ticketmatic/tm-go/ticketmatic"
	"github.com/ticketmatic/tm-go/ticketmatic/events"
)

func TestBatch(t *testing.T) {
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

	order2, err := Create(c, &ticketmatic.CreateOrder{
		Saleschannelid: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	if order2.Orderid == 0 {
		t.Errorf("Unexpected order2.Orderid, got %#v, expected different value", order2.Orderid)
	}

	if order2.Saleschannelid != 1 {
		t.Errorf("Unexpected order2.Saleschannelid, got %#v, expected %#v", order2.Saleschannelid, 1)
	}

	err = Batch(c, &ticketmatic.BatchOrderOperation{
		Ids: []int64{
			order.Orderid,
			order2.Orderid,
		},
		Operation: "update",
		Parameters: &ticketmatic.BatchOrderParameters{
			Updatefields: []*ticketmatic.BatchOrderUpdateField{
				&ticketmatic.BatchOrderUpdateField{
					Key:   "deliveryscenarioid",
					Value: 1,
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

}

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

	list2, err := Getlist(c, &ticketmatic.OrderQuery{
		Limit: 100,
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(list2.Data) != 100 {
		t.Errorf("Unexpected list2.Data length, got %#v, expected %#v", len(list2.Data), 100)
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

	ttps, err := events.Get(c, 777701)
	if err != nil {
		t.Fatal(err)
	}

	if ttps.Id == 0 {
		t.Errorf("Unexpected ttps.Id, got %#v, expected different value", ttps.Id)
	}

	ticketsadded, err := Addtickets(c, order.Orderid, &ticketmatic.AddTickets{
		Tickets: []*ticketmatic.CreateTicket{
			&ticketmatic.CreateTicket{
				Tickettypepriceid: ttps.Prices.Contingents[0].Pricetypes[0].Tickettypepriceid,
			},
			&ticketmatic.CreateTicket{
				Tickettypepriceid: ttps.Prices.Contingents[0].Pricetypes[0].Tickettypepriceid,
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

func TestSplit(t *testing.T) {
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

	ttps, err := events.Get(c, 777701)
	if err != nil {
		t.Fatal(err)
	}

	if ttps.Id == 0 {
		t.Errorf("Unexpected ttps.Id, got %#v, expected different value", ttps.Id)
	}

	ticketsadded, err := Addtickets(c, order.Orderid, &ticketmatic.AddTickets{
		Tickets: []*ticketmatic.CreateTicket{
			&ticketmatic.CreateTicket{
				Tickettypepriceid: ttps.Prices.Contingents[0].Pricetypes[0].Tickettypepriceid,
			},
			&ticketmatic.CreateTicket{
				Tickettypepriceid: ttps.Prices.Contingents[0].Pricetypes[0].Tickettypepriceid,
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(ticketsadded.Order.Tickets) != 2 {
		t.Errorf("Unexpected ticketsadded.Order.Tickets length, got %#v, expected %#v", len(ticketsadded.Order.Tickets), 2)
	}

	ticketids := []int64{
		ticketsadded.Order.Tickets[0].Id,
	}

	_, err = Confirm(c, order.Orderid)
	if err != nil {
		t.Fatal(err)
	}

	split, err := Split(c, order.Orderid, &ticketmatic.SplitOrder{
		Deliveryscenarioid: 3,
		Tickets:            ticketids,
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(split.Tickets) != 1 {
		t.Errorf("Unexpected split.Tickets length, got %#v, expected %#v", len(split.Tickets), 1)
	}

	if split.Deliveryscenarioid != 3 {
		t.Errorf("Unexpected split.Deliveryscenarioid, got %#v, expected %#v", split.Deliveryscenarioid, 3)
	}

	if split.Paymentscenarioid != 3 {
		t.Errorf("Unexpected split.Paymentscenarioid, got %#v, expected %#v", split.Paymentscenarioid, 3)
	}

}
