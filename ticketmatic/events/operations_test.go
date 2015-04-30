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

	list, err := Getlist(c, &ticketmatic.EventQuery{})
	if err != nil {
		t.Fatal(err)
	}

	if len(list) <= 0 {
		t.Errorf("Unexpected list length, got %#v, expected greater than %#v", len(list), 0)
	}

	event, err := Get(c, list[0].Id)
	if err != nil {
		t.Fatal(err)
	}

	if event.Id != list[0].Id {
		t.Errorf("Unexpected event.Id, got %#v, expected %#v", event.Id, list[0].Id)
	}

}
