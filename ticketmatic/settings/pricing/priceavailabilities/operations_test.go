package priceavailabilities

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

	priceavailability, err := Create(c, &ticketmatic.PriceAvailability{
		Name: "Price availability test",
		Rules: &ticketmatic.PriceAvailabilityRules{
			Defaultsaleschannelids: []int64{
				1,
				2,
			},
			Exceptions: []*ticketmatic.PriceAvailabilityRuleException{
				&ticketmatic.PriceAvailabilityRuleException{
					Pricetypeid:     1,
					Saleschannelids: []int64{},
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if priceavailability.Id == 0 {
		t.Errorf("Unexpected priceavailability.Id, got %#v, expected different value", priceavailability.Id)
	}

	if priceavailability.Name != "Price availability test" {
		t.Errorf("Unexpected priceavailability.Name, got %#v, expected %#v", priceavailability.Name, "Price availability test")
	}

	list, err := Getlist(c, &ticketmatic.PriceAvailabilityQuery{})
	if err != nil {
		t.Fatal(err)
	}

	if len(list.Data) <= 0 {
		t.Errorf("Unexpected list.Data length, got %#v, expected greater than %#v", len(list.Data), 0)
	}

	err = Delete(c, priceavailability.Id)
	if err != nil {
		t.Fatal(err)
	}

}
