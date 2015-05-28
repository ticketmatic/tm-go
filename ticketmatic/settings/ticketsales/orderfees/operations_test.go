package orderfees

import (
	"os"
	"testing"

	"github.com/ticketmatic/tm-go/ticketmatic"
)

func TestCreateAndDelete(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	orderfee, err := Create(c, &ticketmatic.OrderFee{
		Name: "Order fee test",
		Rule: &ticketmatic.OrderfeeRule{
			Auto: []*ticketmatic.OrderfeeAutoRule{
				&ticketmatic.OrderfeeAutoRule{
					Deliveryscenarioids: []int64{
						1,
					},
					Paymentscenarioids: []int64{
						1,
					},
					Saleschannelids: []int64{
						1,
					},
					Status: "fixedfee",
					Value:  1,
				},
			},
		},
		Typeid: 2401,
	})
	if err != nil {
		t.Fatal(err)
	}

	if orderfee.Id == 0 {
		t.Errorf("Unexpected orderfee.Id, got %#v, expected different value", orderfee.Id)
	}

	if orderfee.Typeid != 2401 {
		t.Errorf("Unexpected orderfee.Typeid, got %#v, expected %#v", orderfee.Typeid, 2401)
	}

	if orderfee.Name != "Order fee test" {
		t.Errorf("Unexpected orderfee.Name, got %#v, expected %#v", orderfee.Name, "Order fee test")
	}

	orderfeescript, err := Create(c, &ticketmatic.OrderFee{
		Name: "Order fee script test",
		Rule: &ticketmatic.OrderfeeRule{
			Script: "return 2;",
		},
		Typeid: 2402,
	})
	if err != nil {
		t.Fatal(err)
	}

	if orderfeescript.Id == 0 {
		t.Errorf("Unexpected orderfeescript.Id, got %#v, expected different value", orderfeescript.Id)
	}

	if orderfeescript.Typeid != 2402 {
		t.Errorf("Unexpected orderfeescript.Typeid, got %#v, expected %#v", orderfeescript.Typeid, 2402)
	}

	if orderfeescript.Name != "Order fee script test" {
		t.Errorf("Unexpected orderfeescript.Name, got %#v, expected %#v", orderfeescript.Name, "Order fee script test")
	}

	if orderfeescript.Rule.Script != "return 2;" {
		t.Errorf("Unexpected orderfeescript.Rule.Script, got %#v, expected %#v", orderfeescript.Rule.Script, "return 2;")
	}

	list, err := Getlist(c, &ticketmatic.OrderFeeQuery{})
	if err != nil {
		t.Fatal(err)
	}

	if len(list.Data) <= 1 {
		t.Errorf("Unexpected list.Data length, got %#v, expected greater than %#v", len(list.Data), 1)
	}

	err = Delete(c, orderfee.Id)
	if err != nil {
		t.Fatal(err)
	}

	err = Delete(c, orderfeescript.Id)
	if err != nil {
		t.Fatal(err)
	}

}
