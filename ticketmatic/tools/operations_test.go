package tools

import (
	"os"
	"testing"

	"github.com/ticketmatic/tm-go/ticketmatic"
)

func TestInfo(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	info, err := Account(c)
	if err != nil {
		t.Fatal(err)
	}

	if info.Id != 998 {
		t.Errorf("Unexpected info.Id, got %#v, expected %#v", info.Id, 998)
	}

}

func TestGet(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	req, err := Queries(c, &ticketmatic.QueryRequest{
		Limit: 2,
		Query: "SELECT * FROM tm.paymentscenario",
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

func TestExport(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	req, err := Export(c, &ticketmatic.QueryRequest{
		Query: "SELECT * FROM tm.contact LIMIT 3",
	})
	if err != nil {
		t.Fatal(err)
	}

	stream := make([]map[string]interface{}, 0)
	for {
		n, err := req.Next()
		if err != nil {
			t.Fatal(err)
		}
		if n == nil {
			break
		}
		stream = append(stream, n)
	}

	if len(stream) != 3 {
		t.Errorf("Unexpected stream length, got %#v, expected %#v", len(stream), 3)
	}

}
