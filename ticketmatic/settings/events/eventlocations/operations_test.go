package eventlocations

import (
	"os"
	"testing"

	"github.com/ticketmatic/tm-go/ticketmatic"
)

func TestBadfilter(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	_, err = Getlist(c, &ticketmatic.EventLocationQuery{
		Filter: "SELECT * FROM tm.eventlocation",
	})

	if err == nil {
		t.Fatal("Expected an error!")
	}

}
