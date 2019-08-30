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

	voucher, err := Get(c, 2)
	if err != nil {
		t.Fatal(err)
	}

	if voucher.Id != 2 {
		t.Errorf("Unexpected voucher.Id, got %#v, expected %#v", voucher.Id, 2)
	}

	if voucher.Nbrofcodes <= 0 {
		t.Errorf("Unexpected voucher.Nbrofcodes, got %#v, expected > %#v", voucher.Nbrofcodes, 0)
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

func TestCreatecodesupdate(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	codes, err := Createcodes(c, 2, &ticketmatic.AddVoucherCodes{
		Amount: 42,
		Codes: []*ticketmatic.VoucherCode{
			&ticketmatic.VoucherCode{
				Code:     "test12345",
				Expiryts: ticketmatic.NewTime(ticketmatic.MustParseTime("2019-05-01 14:00:00")),
			},
			&ticketmatic.VoucherCode{
				Code:     "foo12345",
				Expiryts: ticketmatic.NewTime(ticketmatic.MustParseTime("2019-05-01 14:00:00")),
			},
			&ticketmatic.VoucherCode{
				Code:     "bar12345",
				Expiryts: ticketmatic.NewTime(ticketmatic.MustParseTime("2019-05-01 14:00:00")),
			},
		},
		Count: 3,
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(codes) != 3 {
		t.Errorf("Unexpected codes length, got %#v, expected %#v", len(codes), 3)
	}

	if codes[0].Code != "test12345" {
		t.Errorf("Unexpected codes[0].Code, got %#v, expected %#v", codes[0].Code, "test12345")
	}

	if codes[0].Expiryts != "2019-05-01 14:00:00" {
		t.Errorf("Unexpected codes[0].Expiryts, got %#v, expected %#v", codes[0].Expiryts, "2019-05-01 14:00:00")
	}

	if codes[1].Code != "foo12345" {
		t.Errorf("Unexpected codes[1].Code, got %#v, expected %#v", codes[1].Code, "foo12345")
	}

	if codes[1].Expiryts != "2019-05-01 14:00:00" {
		t.Errorf("Unexpected codes[1].Expiryts, got %#v, expected %#v", codes[1].Expiryts, "2019-05-01 14:00:00")
	}

	if codes[2].Code != "bar12345" {
		t.Errorf("Unexpected codes[2].Code, got %#v, expected %#v", codes[2].Code, "bar12345")
	}

	if codes[2].Expiryts != "2019-05-01 14:00:00" {
		t.Errorf("Unexpected codes[2].Expiryts, got %#v, expected %#v", codes[2].Expiryts, "2019-05-01 14:00:00")
	}

	voucher, err := Get(c, 2)
	if err != nil {
		t.Fatal(err)
	}

	if voucher.Id != 2 {
		t.Errorf("Unexpected voucher.Id, got %#v, expected %#v", voucher.Id, 2)
	}

	if voucher.Nbrofcodes <= 0 {
		t.Errorf("Unexpected voucher.Nbrofcodes, got %#v, expected > %#v", voucher.Nbrofcodes, 0)
	}

	err = Deactivatecodes(c, 2, []*ticketmatic.VoucherCode{
		&ticketmatic.VoucherCode{
			Code: codes[1].Code,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	codes, err := Createcodes(c, 2, &ticketmatic.AddVoucherCodes{
		Amount: 42,
		Codes: []*ticketmatic.VoucherCode{
			&ticketmatic.VoucherCode{
				Code:     "foo12345",
				Expiryts: ticketmatic.NewTime(ticketmatic.MustParseTime("2019-05-01 14:00:00")),
			},
			&ticketmatic.VoucherCode{
				Code:     "bar12345",
				Expiryts: ticketmatic.NewTime(ticketmatic.MustParseTime("2019-05-01 14:00:00")),
			},
		},
		Count: 2,
	})

	if err == nil {
		t.Fatal("Expected an error!")
	}

	codes, err := Createcodes(c, 2, &ticketmatic.AddVoucherCodes{
		Amount: 42,
		Codes: []*ticketmatic.VoucherCode{
			&ticketmatic.VoucherCode{
				Code:     "foo12345",
				Expiryts: ticketmatic.NewTime(ticketmatic.MustParseTime("2019-05-02 14:00:00")),
			},
			&ticketmatic.VoucherCode{
				Code:     "bar12345",
				Expiryts: ticketmatic.NewTime(ticketmatic.MustParseTime("2019-05-02 14:00:00")),
			},
		},
		Count:  2,
		Update: true,
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(codes) != 2 {
		t.Errorf("Unexpected codes length, got %#v, expected %#v", len(codes), 2)
	}

	if codes[0].Code != "foo12345" {
		t.Errorf("Unexpected codes[0].Code, got %#v, expected %#v", codes[0].Code, "foo12345")
	}

	if codes[0].Expiryts != "2019-05-02 14:00:00" {
		t.Errorf("Unexpected codes[0].Expiryts, got %#v, expected %#v", codes[0].Expiryts, "2019-05-02 14:00:00")
	}

	if codes[1].Code != "bar12345" {
		t.Errorf("Unexpected codes[1].Code, got %#v, expected %#v", codes[1].Code, "bar12345")
	}

	if codes[1].Expiryts != "2019-05-02 14:00:00" {
		t.Errorf("Unexpected codes[1].Expiryts, got %#v, expected %#v", codes[1].Expiryts, "2019-05-02 14:00:00")
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
