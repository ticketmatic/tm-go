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

}

func TestCreatecodesfixed(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	codes, err := Createcodes(c, 2, &ticketmatic.AddVoucherCodes{
		Amount: 12,
		Codes: []string{
			"ABC123DEF456",
		},
		Count: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(codes) != 1 {
		t.Errorf("Unexpected codes length, got %#v, expected %#v", len(codes), 1)
	}

	if codes[0].Code != "ABC123DEF456" {
		t.Errorf("Unexpected codes[0].Code, got %#v, expected %#v", codes[0].Code, "ABC123DEF456")
	}

}
