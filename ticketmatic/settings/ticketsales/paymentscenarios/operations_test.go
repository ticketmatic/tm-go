package paymentscenarios

import (
	"os"
	"testing"

	"github.com/ticketmatic/tm-go/ticketmatic"
)

func TestCreate(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	paymentscenario, err := Create(c, &ticketmatic.PaymentScenario{
		Typeid: 2705,
		Name:   "Payment scenario test",
		Availability: &ticketmatic.PaymentscenarioAvailability{
			Saleschannels: []int64{
				1,
			},
			Usescript: false,
		},
		Expiryparameters: &ticketmatic.PaymentscenarioExpiryParameters{
			Daysaftercreation: 24,
		},
		Internalremark: "Testing",
		Overdueparameters: &ticketmatic.PaymentscenarioOverdueParameters{
			Daysaftercreation: 12,
		},
		Paymentmethods: []int64{
			1,
			2,
		},
		Shortdescription: "Short test",
	})
	if err != nil {
		t.Fatal(err)
	}

	if paymentscenario.Id == 0 {
		t.Errorf("Unexpected paymentscenario.Id, got %#v, expected different value", paymentscenario.Id)
	}

	if paymentscenario.Typeid != 2705 {
		t.Errorf("Unexpected paymentscenario.Typeid, got %#v, expected %#v", paymentscenario.Typeid, 2705)
	}

	if paymentscenario.Name != "Payment scenario test" {
		t.Errorf("Unexpected paymentscenario.Name, got %#v, expected %#v", paymentscenario.Name, "Payment scenario test")
	}

}

func TestGet(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	req, err := Get(c, 13)
	if err != nil {
		t.Fatal(err)
	}

	if req.Name != "PayPal" {
		t.Errorf("Unexpected req.Name, got %#v, expected %#v", req.Name, "PayPal")
	}

}
