package contacts

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

	list, err := Getlist(c, &ticketmatic.ContactQuery{})
	if err != nil {
		t.Fatal(err)
	}

	if len(list.Data) <= 0 {
		t.Errorf("Unexpected list.Data length, got %#v, expected greater than %#v", len(list.Data), 0)
	}

	contact, err := Get(c, list.Data[0].Id)
	if err != nil {
		t.Fatal(err)
	}

	if contact.Id != list.Data[0].Id {
		t.Errorf("Unexpected contact.Id, got %#v, expected %#v", contact.Id, list.Data[0].Id)
	}

}

func TestCreate(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	contact, err := Create(c, &ticketmatic.Contact{
		Firstname: "John",
	})
	if err != nil {
		t.Fatal(err)
	}

	if contact.Id == 0 {
		t.Errorf("Unexpected contact.Id, got %#v, expected different value", contact.Id)
	}

	if contact.Firstname != "John" {
		t.Errorf("Unexpected contact.Firstname, got %#v, expected %#v", contact.Firstname, "John")
	}

	updated, err := Update(c, contact.Id, &ticketmatic.Contact{
		Lastname: "Doe",
	})
	if err != nil {
		t.Fatal(err)
	}

	if updated.Id != contact.Id {
		t.Errorf("Unexpected updated.Id, got %#v, expected %#v", updated.Id, contact.Id)
	}

	if updated.Firstname != "John" {
		t.Errorf("Unexpected updated.Firstname, got %#v, expected %#v", updated.Firstname, "John")
	}

	if updated.Lastname != "Doe" {
		t.Errorf("Unexpected updated.Lastname, got %#v, expected %#v", updated.Lastname, "Doe")
	}

	err = Delete(c, contact.Id)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCreatecustom(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	titles, err := contacttitles.Getlist(c, &ticketmatic.ContactTitleQuery{})
	if err != nil {
		t.Fatal(err)
	}

	if len(titles.Data) <= 0 {
		t.Errorf("Unexpected titles.Data length, got %#v, expected greater than %#v", len(titles.Data), 0)
	}

	addrtypes, err := contactaddresstypes.Getlist(c, &ticketmatic.ContactAddressTypeQuery{})
	if err != nil {
		t.Fatal(err)
	}

	if len(addrtypes.Data) <= 0 {
		t.Errorf("Unexpected addrtypes.Data length, got %#v, expected greater than %#v", len(addrtypes.Data), 0)
	}

	ptypes, err := phonenumbertypes.Getlist(c, &ticketmatic.PhoneNumberTypeQuery{})
	if err != nil {
		t.Fatal(err)
	}

	if len(ptypes.Data) <= 1 {
		t.Errorf("Unexpected ptypes.Data length, got %#v, expected greater than %#v", len(ptypes.Data), 1)
	}

	contact, err := Create(c, &ticketmatic.Contact{
		Addresses: []*ticketmatic.Address{
			&ticketmatic.Address{
				City:        "Nieuwerkerk Aan Den Ijssel",
				Countrycode: "NL",
				Street1:     "Kerkstraat",
				Street2:     "1",
				Typeid:      addrtypes.Data[0].Id,
				Zip:         "2914 AH",
			},
		},
		Birthdate:       ticketmatic.NewTime(ticketmatic.MustParseTime("1959-09-21")),
		Customertitleid: titles.Data[0].Id,
		Email:           "john@worldonline.nl",
		Firstname:       "John",
		Lastname:        "Johns",
		Middlename:      "J",
		Phonenumbers: []*ticketmatic.Phonenumber{
			&ticketmatic.Phonenumber{
				Number: "+31222222222",
				Typeid: ptypes.Data[0].Id,
			},
			&ticketmatic.Phonenumber{
				Number: "+31222222222",
				Typeid: ptypes.Data[1].Id,
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if contact.Id == 0 {
		t.Errorf("Unexpected contact.Id, got %#v, expected different value", contact.Id)
	}

	if contact.Firstname != "John" {
		t.Errorf("Unexpected contact.Firstname, got %#v, expected %#v", contact.Firstname, "John")
	}

	err = Delete(c, contact.Id)
	if err != nil {
		t.Fatal(err)
	}

}
