package diagnostics

import (
	"os"
	"testing"
	"time"

	"github.com/ticketmatic/tm-go/ticketmatic"
)

func TestGetTime(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	req, err := Time(c)
	if err != nil {
		t.Fatal(err)
	}

	if time.Since(req.Systemtime.Time()) > 24*time.Hour {
		t.Errorf("Unexpected req.Systemtime time, should be recent, got %s", req.Systemtime.Time())
	}

}
