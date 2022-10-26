package events

import (
	"os"
	"testing"

	"github.com/ticketmatic/tm-go/ticketmatic"
)

func TestBatch(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	event, err := Create(c, &ticketmatic.Event{
		Name: "Example",
		Contingents: []*ticketmatic.EventContingent{
			&ticketmatic.EventContingent{
				Amount: 100,
			},
		},
		Locationid: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	if event.Name != "Example" {
		t.Errorf("Unexpected event.Name, got %#v, expected %#v", event.Name, "Example")
	}

	if event.Contingents[0].Amount != 100 {
		t.Errorf("Unexpected event.Contingents[0].Amount, got %#v, expected %#v", event.Contingents[0].Amount, 100)
	}

	if event.Locationid != 1 {
		t.Errorf("Unexpected event.Locationid, got %#v, expected %#v", event.Locationid, 1)
	}

	event2, err := Create(c, &ticketmatic.Event{
		Name: "Example2",
		Contingents: []*ticketmatic.EventContingent{
			&ticketmatic.EventContingent{
				Amount: 100,
			},
		},
		Locationid: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	if event2.Name != "Example2" {
		t.Errorf("Unexpected event2.Name, got %#v, expected %#v", event2.Name, "Example2")
	}

	if event2.Contingents[0].Amount != 100 {
		t.Errorf("Unexpected event2.Contingents[0].Amount, got %#v, expected %#v", event2.Contingents[0].Amount, 100)
	}

	if event2.Locationid != 1 {
		t.Errorf("Unexpected event2.Locationid, got %#v, expected %#v", event2.Locationid, 1)
	}

	err = Batch(c, &ticketmatic.BatchEventOperation{
		Ids: []int64{
			event.Id,
			event2.Id,
		},
		Operation: "update",
		Parameters: &ticketmatic.BatchEventParameters{
			Updatefields: []*ticketmatic.BatchEventUpdateField{
				&ticketmatic.BatchEventUpdateField{
					Key:   "locationid",
					Value: 2,
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

}

func TestCreate(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	event, err := Create(c, &ticketmatic.Event{
		Name: "Example",
		Contingents: []*ticketmatic.EventContingent{
			&ticketmatic.EventContingent{
				Amount: 100,
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if event.Name != "Example" {
		t.Errorf("Unexpected event.Name, got %#v, expected %#v", event.Name, "Example")
	}

	if event.Contingents[0].Amount != 100 {
		t.Errorf("Unexpected event.Contingents[0].Amount, got %#v, expected %#v", event.Contingents[0].Amount, 100)
	}

}

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

func TestGetconditions(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	req, err := Get(c, 777717)
	if err != nil {
		t.Fatal(err)
	}

	if req.Id != 777717 {
		t.Errorf("Unexpected req.Id, got %#v, expected %#v", req.Id, 777717)
	}

	if req.Prices.Contingents[0].Pricetypes[0].Saleschannels[0].Conditions[0].Type != "orderticketlimit" {
		t.Errorf("Unexpected req.Prices.Contingents[0].Pricetypes[0].Saleschannels[0].Conditions[0].Type, got %#v, expected %#v", req.Prices.Contingents[0].Pricetypes[0].Saleschannels[0].Conditions[0].Type, "orderticketlimit")
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

	stream, err := Gettickets(c, list.Data[0].Id, &ticketmatic.EventTicketQuery{})
	if err != nil {
		t.Fatal(err)
	}

	tickets := make([]*ticketmatic.EventTicket, 0)
	for {
		n, err := stream.Next()
		if err != nil {
			t.Fatal(err)
		}
		if n == nil {
			break
		}
		tickets = append(tickets, n)
	}

}

func TestDeletefixedbundleevent(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	err = Delete(c, 777704)

	if err == nil {
		t.Fatal("Expected an error!")
	}

}

func TestLockunlocktickets(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	list, err := Getlist(c, &ticketmatic.EventQuery{
		Filter:  "select id from tm.event where seatingplanid is not null and id < 777800",
		Limit:   1,
		Orderby: "name",
		Output:  "ids",
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(list.Data) <= 0 {
		t.Errorf("Unexpected list.Data length, got %#v, expected greater than %#v", len(list.Data), 0)
	}

	stream, err := Gettickets(c, list.Data[0].Id, &ticketmatic.EventTicketQuery{})
	if err != nil {
		t.Fatal(err)
	}

	tickets := make([]*ticketmatic.EventTicket, 0)
	for {
		n, err := stream.Next()
		if err != nil {
			t.Fatal(err)
		}
		if n == nil {
			break
		}
		tickets = append(tickets, n)
	}

	if len(tickets) <= 0 {
		t.Errorf("Unexpected tickets length, got %#v, expected greater than %#v", len(tickets), 0)
	}

	err = Locktickets(c, list.Data[0].Id, &ticketmatic.EventLockTickets{
		Locktypeid: 1,
		Ticketids: []int64{
			tickets[0].Id,
			tickets[1].Id,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	err = Unlocktickets(c, list.Data[0].Id, &ticketmatic.EventUnlockTickets{
		Ticketids: []int64{
			tickets[0].Id,
			tickets[1].Id,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

}

func TestUpdateseatrankfortickets(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	list, err := Getlist(c, &ticketmatic.EventQuery{
		Filter:  "select id from tm.event where seatingplanid is not null and id < 777800",
		Limit:   1,
		Orderby: "name",
		Output:  "ids",
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(list.Data) <= 0 {
		t.Errorf("Unexpected list.Data length, got %#v, expected greater than %#v", len(list.Data), 0)
	}

	stream, err := Gettickets(c, list.Data[0].Id, &ticketmatic.EventTicketQuery{})
	if err != nil {
		t.Fatal(err)
	}

	tickets := make([]*ticketmatic.EventTicket, 0)
	for {
		n, err := stream.Next()
		if err != nil {
			t.Fatal(err)
		}
		if n == nil {
			break
		}
		tickets = append(tickets, n)
	}

	if len(tickets) <= 0 {
		t.Errorf("Unexpected tickets length, got %#v, expected greater than %#v", len(tickets), 0)
	}

	err = Updateseatrankfortickets(c, list.Data[0].Id, &ticketmatic.EventUpdateSeatRankForTickets{
		Seatrankid: 3,
		Ticketids: []int64{
			tickets[0].Id,
			tickets[1].Id,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

}
