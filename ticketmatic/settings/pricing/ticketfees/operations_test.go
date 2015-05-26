package ticketfees

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

	fee, err := Create(c, &ticketmatic.TicketFee{
		Name: "Fee",
		Rules: &ticketmatic.TicketfeeRules{
			Default: []*ticketmatic.TicketfeeSaleschannelRule{
				&ticketmatic.TicketfeeSaleschannelRule{
					Saleschannelid: 1,
					Status:         "fixedfee",
					Value:          1,
				},
			},
			Exceptions: []*ticketmatic.TicketfeeException{
				&ticketmatic.TicketfeeException{
					Pricetypeid: 2,
					Saleschannels: []*ticketmatic.TicketfeeSaleschannelRule{
						&ticketmatic.TicketfeeSaleschannelRule{
							Saleschannelid: 1,
							Status:         "percentagefee",
							Value:          3.5,
						},
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if fee.Name != "Fee" {
		t.Errorf("Unexpected fee.Name, got %#v, expected %#v", fee.Name, "Fee")
	}

	if fee.Id == 0 {
		t.Errorf("Unexpected fee.Id, got %#v, expected different value", fee.Id)
	}

}
