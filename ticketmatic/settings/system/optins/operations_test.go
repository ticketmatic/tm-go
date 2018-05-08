package optins

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

	req, err := Getlist(c, &ticketmatic.OptInQuery{})
	if err != nil {
		t.Fatal(err)
	}

	if len(req.Data) <= 0 {
		t.Errorf("Unexpected req.Data length, got %#v, expected greater than %#v", len(req.Data), 0)
	}

}

func TestCreate(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	optin, err := Create(c, &ticketmatic.OptIn{
		Typeid: 40001,
		Name:   "Newsletter",
		Availability: []*ticketmatic.OptInAvailability{
			&ticketmatic.OptInAvailability{
				Saleschannelid: 4,
			},
		},
		Description: "Receive newsletters",
		Nocaption:   "No",
		Yescaption:  "Yes",
	})
	if err != nil {
		t.Fatal(err)
	}

	if optin.Id == 0 {
		t.Errorf("Unexpected optin.Id, got %#v, expected different value", optin.Id)
	}

	if optin.Typeid != 40001 {
		t.Errorf("Unexpected optin.Typeid, got %#v, expected %#v", optin.Typeid, 40001)
	}

	if optin.Name != "Newsletter" {
		t.Errorf("Unexpected optin.Name, got %#v, expected %#v", optin.Name, "Newsletter")
	}

}
