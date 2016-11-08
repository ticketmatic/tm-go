package contactfields

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

	req, err := Getlist(c)
	if err != nil {
		t.Fatal(err)
	}

	if len(req.Data) <= 0 {
		t.Errorf("Unexpected req.Data length, got %#v, expected greater than %#v", len(req.Data), 0)
	}

	reqget, err := Get(c, req.Data[0].Id)
	if err != nil {
		t.Fatal(err)
	}

	if reqget.Id <= 0 {
		t.Errorf("Unexpected reqget.Id, got %#v, expected > %#v", reqget.Id, 0)
	}

	_, err = Translations(c, reqget.Id)
	if err != nil {
		t.Fatal(err)
	}

}
