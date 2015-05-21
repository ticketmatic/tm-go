package subscribers

import (
	"os"
	"testing"

	"github.com/ticketmatic/tm-go/ticketmatic"
)

func TestSync(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	err = Sync(c, []*ticketmatic.SubscriberSync{
		&ticketmatic.SubscriberSync{
			Email:      "subscriber@ticketmatic.com",
			Subscribed: true,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

}

func TestCommunications(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	err = Communications(c, &ticketmatic.SubscriberCommunication{
		Ts: ticketmatic.NewTime(ticketmatic.MustParseTime("2015-02-02")),
		Addresses: []string{
			"test1@ticketmatic.com",
			"test2@ticketmatic.com",
		},
		Name: "test1",
	})
	if err != nil {
		t.Fatal(err)
	}

}
