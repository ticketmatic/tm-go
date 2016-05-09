package events

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

	list, err := Getlist(c, &ticketmatic.EventQuery{
		Output: "withlookup",
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(list.Data) <= 0 {
		t.Errorf("Unexpected list.Data length, got %#v, expected greater than %#v", len(list.Data), 0)
	}

	event, err := Get(c, list.Data[0].Id)
	if err != nil {
		t.Fatal(err)
	}

	if event.Id != list.Data[0].Id {
		t.Errorf("Unexpected event.Id, got %#v, expected %#v", event.Id, list.Data[0].Id)
	}

}

func TestGetdraft(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	event, err := Create(c, &ticketmatic.Event{
		Name: "Draft",
	})
	if err != nil {
		t.Fatal(err)
	}

	if event.Name != "Draft" {
		t.Errorf("Unexpected event.Name, got %#v, expected %#v", event.Name, "Draft")
	}

	list, err := Getlist(c, &ticketmatic.EventQuery{
		Filter: "select id from tm.event where nameen='Draft'",
		Simplefilter: &ticketmatic.EventFilter{
			Status: []int64{
				19001,
				19002,
				19003,
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(list.Data) <= 0 {
		t.Errorf("Unexpected list.Data length, got %#v, expected greater than %#v", len(list.Data), 0)
	}

	err = Delete(c, event.Id)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGettickets(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	list, err := Getlist(c, &ticketmatic.EventQuery{})
	if err != nil {
		t.Fatal(err)
	}

	if len(list.Data) <= 0 {
		t.Errorf("Unexpected list.Data length, got %#v, expected greater than %#v", len(list.Data), 0)
	}

	tickets, err := Gettickets(c, list.Data[0].Id, &ticketmatic.EventTicketQuery{})
	if err != nil {
		t.Fatal(err)
	}

	if len(tickets.Data) <= 0 {
		t.Errorf("Unexpected tickets.Data length, got %#v, expected greater than %#v", len(tickets.Data), 0)
	}

}
