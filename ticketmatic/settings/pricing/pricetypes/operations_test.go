package pricetypes

import (
	"os"
	"testing"
	"time"

	"github.com/ticketmatic/tm-go/ticketmatic"
)

func TestGet(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	req, err := Getlist(c, &ticketmatic.PriceTypeQuery{})
	if err != nil {
		t.Fatal(err)
	}

	if len(req.Data) <= 0 {
		t.Errorf("Unexpected req.Data length, got %#v, expected greater than %#v", len(req.Data), 0)
	}

	req2, err := Getlist(c, &ticketmatic.PriceTypeQuery{
		Filter: "select id from conf.pricetype where typeid=2301",
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(req.Data) <= len(req2.Data) {
		t.Errorf("Unexpected req.Data length, got %#v, expected greater than %#v", len(req.Data), len(req2.Data))
	}

}

func TestCreatedelete(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	req, err := Getlist(c, &ticketmatic.PriceTypeQuery{})
	if err != nil {
		t.Fatal(err)
	}

	if len(req.Data) <= 0 {
		t.Errorf("Unexpected req.Data length, got %#v, expected greater than %#v", len(req.Data), 0)
	}

	req2, err := Create(c, &ticketmatic.PriceType{
		Name: "test",
	})
	if err != nil {
		t.Fatal(err)
	}

	if req2.Name != "test" {
		t.Errorf("Unexpected req2.Name, got %#v, expected %#v", req2.Name, "test")
	}

	if time.Since(req2.Createdts.Time()) > 24*time.Hour {
		t.Errorf("Unexpected req2.Createdts time, should be recent, got %s", req2.Createdts.Time())
	}

	req3, err := Getlist(c, &ticketmatic.PriceTypeQuery{})
	if err != nil {
		t.Fatal(err)
	}

	if len(req3.Data) <= len(req.Data) {
		t.Errorf("Unexpected req3.Data length, got %#v, expected greater than %#v", len(req3.Data), len(req.Data))
	}

	req4, err := Get(c, req2.Id)
	if err != nil {
		t.Fatal(err)
	}

	if req4.Name != "test" {
		t.Errorf("Unexpected req4.Name, got %#v, expected %#v", req4.Name, "test")
	}

	err = Delete(c, req2.Id)
	if err != nil {
		t.Fatal(err)
	}

	req6, err := Getlist(c, &ticketmatic.PriceTypeQuery{})
	if err != nil {
		t.Fatal(err)
	}

	if len(req.Data) != len(req6.Data) {
		t.Errorf("Unexpected req.Data length, got %#v, expected %#v", len(req.Data), len(req6.Data))
	}

}
