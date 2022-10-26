package vouchers

import (
	"os"
	"testing"

	"github.com/ticketmatic/tm-go/ticketmatic"
)

func TestValidity(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	voucher, err := Create(c, &ticketmatic.Voucher{
		Typeid:       24001,
		Name:         "test",
		Codeformatid: 27001,
		Validity: &ticketmatic.VoucherValidity{
			ExpiryMonthsaftercreation: 12,
			Maxusages:                 5,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if voucher.Validity.ExpiryMonthsaftercreation != 12 {
		t.Errorf("Unexpected voucher.Validity.ExpiryMonthsaftercreation, got %#v, expected %#v", voucher.Validity.ExpiryMonthsaftercreation, 12)
	}

	if voucher.Validity.Maxusages != 5 {
		t.Errorf("Unexpected voucher.Validity.Maxusages, got %#v, expected %#v", voucher.Validity.Maxusages, 5)
	}

}
