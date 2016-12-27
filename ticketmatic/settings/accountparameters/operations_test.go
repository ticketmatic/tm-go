package accountparameters

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

	list, err := Getlist(c)
	if err != nil {
		t.Fatal(err)
	}

	if len(list) <= 0 {
		t.Errorf("Unexpected list length, got %#v, expected greater than %#v", len(list), 0)
	}

	param, err := Get(c, "accountname")
	if err != nil {
		t.Fatal(err)
	}

	if param.Key != "accountname" {
		t.Errorf("Unexpected param.Key, got %#v, expected %#v", param.Key, "accountname")
	}

	if param.Value != "qa" {
		t.Errorf("Unexpected param.Value, got %#v, expected %#v", param.Value, "qa")
	}

	newparam, err := Set(c, &ticketmatic.AccountParameter{
		Key:   "testparam",
		Value: 123,
	})
	if err != nil {
		t.Fatal(err)
	}

	if newparam.Key != "testparam" {
		t.Errorf("Unexpected newparam.Key, got %#v, expected %#v", newparam.Key, "testparam")
	}

	if newparam.Value != 123 {
		t.Errorf("Unexpected newparam.Value, got %#v, expected %#v", newparam.Value, 123)
	}

}
