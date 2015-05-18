package tools

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

	req, err := Queries(c, &ticketmatic.QueryRequest{
		Query: "SELECT * FROM tm.paymentscenario",
		Limit: 2,
	})
	if err != nil {
		t.Fatal(err)
	}

	if req.Nbrofresults <= 1 {
		t.Errorf("Unexpected req.Nbrofresults, got %#v, expected > %#v", req.Nbrofresults, 1)
	}

	if len(req.Results) != 2 {
		t.Errorf("Unexpected req.Results length, got %#v, expected %#v", len(req.Results), 2)
	}

}
