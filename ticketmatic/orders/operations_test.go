package orders

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

	list, err := Getlist(c, &ticketmatic.OrderQuery{})
	if err != nil {
		t.Fatal(err)
	}

	if len(list) <= 0 {
		t.Errorf("Unexpected list length, got %#v, expected greater than %#v", len(list), 0)
	}

	order, err := Get(c, list[0].Orderid)
	if err != nil {
		t.Fatal(err)
	}

	if order.Orderid != list[0].Orderid {
		t.Errorf("Unexpected order.Orderid, got %#v, expected %#v", order.Orderid, list[0].Orderid)
	}

}
