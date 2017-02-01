package vouchers

import (
	"os"
	"testing"

	"github.com/ticketmatic/tm-go/ticketmatic"
)

func TestCreatecodes(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	codes, err := Createcodes(c, 2, &ticketmatic.AddVoucherCodes{
		Amount: 10,
		Count:  3,
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(codes) != 3 {
		t.Errorf("Unexpected codes length, got %#v, expected %#v", len(codes), 3)
	}

	if codes[0].Code == "" {
		t.Errorf("Unexpected codes[0].Code, got %#v, expected different value", codes[0].Code)
	}

	err = Deactivatecodes(c, 2, []*ticketmatic.VoucherCode{
		&ticketmatic.VoucherCode{
			Code: codes[1].Code,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

}

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
