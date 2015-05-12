package orders

import (
	"os"
	"testing"

	"github.com/ticketmatic/tm-go/ticketmatic"
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
		Deliveryscenarioid: 2,
		Paymentscenarioid:  3,
		Customerid:         208,
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

	if updated.Customerid != 208 {
		t.Errorf("Unexpected updated.Customerid, got %#v, expected %#v", updated.Customerid, 208)
	}

	_, err = Addtickets(c, order.Orderid, &ticketmatic.AddTickets{
		Tickets: []*ticketmatic.CreateTicket{
			&ticketmatic.CreateTicket{
				Tickettypepriceid: 584,
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = Confirm(c, order.Orderid)
	if err != nil {
		t.Fatal(err)
	}

}
