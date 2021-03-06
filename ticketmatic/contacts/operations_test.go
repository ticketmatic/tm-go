package contacts

import (
	"os"
	"testing"

	"github.com/ticketmatic/tm-go/ticketmatic"
	"github.com/ticketmatic/tm-go/ticketmatic/settings/system/contactaddresstypes"
	"github.com/ticketmatic/tm-go/ticketmatic/settings/system/contacttitles"
	"github.com/ticketmatic/tm-go/ticketmatic/settings/system/phonenumbertypes"
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

	contact, err := Get(c, list.Data[0].Id, &ticketmatic.ContactGetQuery{})
	if err != nil {
		t.Fatal(err)
	}

	if contact.Id != list.Data[0].Id {
		t.Errorf("Unexpected contact.Id, got %#v, expected %#v", contact.Id, list.Data[0].Id)
	}

}

func TestBatch(t *testing.T) {
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

	contact2, err := Create(c, &ticketmatic.Contact{
		Firstname: "Bob",
	})
	if err != nil {
		t.Fatal(err)
	}

	if contact2.Id == 0 {
		t.Errorf("Unexpected contact2.Id, got %#v, expected different value", contact2.Id)
	}

	if contact2.Firstname != "Bob" {
		t.Errorf("Unexpected contact2.Firstname, got %#v, expected %#v", contact2.Firstname, "Bob")
	}

	err = Batch(c, &ticketmatic.BatchContactOperation{
		Ids: []int64{
			contact.Id,
		},
		Operation: "update",
		Parameters: &ticketmatic.BatchContactParameters{
			Fields: &ticketmatic.ContactBatchUpdate{
				Languagecode: "EN",
			},
		},
	})
	if err != nil {
		t.Fatal(err)
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
				Typeid:      addrtypes.Data[0].Id,
				City:        "Nieuwerkerk Aan Den Ijssel",
				Countrycode: "NL",
				Street1:     "Kerkstraat",
				Street2:     "1",
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
				Typeid: ptypes.Data[0].Id,
				Number: "+31222222222",
			},
			&ticketmatic.Phonenumber{
				Typeid: ptypes.Data[1].Id,
				Number: "+31222222222",
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

	if contact.Addresses[0].Countrycode != "NL" {
		t.Errorf("Unexpected contact.Addresses[0].Countrycode, got %#v, expected %#v", contact.Addresses[0].Countrycode, "NL")
	}

	if contact.Addresses[0].Country != "Netherlands" {
		t.Errorf("Unexpected contact.Addresses[0].Country, got %#v, expected %#v", contact.Addresses[0].Country, "Netherlands")
	}

	err = Delete(c, contact.Id)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCreateunicode(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	contact, err := Create(c, &ticketmatic.Contact{
		Email:     "john@test.com",
		Firstname: "JØhñ",
		Lastname:  "ポテト 👌 ไก่",
	})
	if err != nil {
		t.Fatal(err)
	}

	if contact.Id == 0 {
		t.Errorf("Unexpected contact.Id, got %#v, expected different value", contact.Id)
	}

	if contact.Firstname != "JØhñ" {
		t.Errorf("Unexpected contact.Firstname, got %#v, expected %#v", contact.Firstname, "JØhñ")
	}

	if contact.Lastname != "ポテト 👌 ไก่" {
		t.Errorf("Unexpected contact.Lastname, got %#v, expected %#v", contact.Lastname, "ポテト 👌 ไก่")
	}

	contact2, err := Get(c, contact.Id, &ticketmatic.ContactGetQuery{})
	if err != nil {
		t.Fatal(err)
	}

	if contact2.Id == 0 {
		t.Errorf("Unexpected contact2.Id, got %#v, expected different value", contact2.Id)
	}

	if contact2.Firstname != "JØhñ" {
		t.Errorf("Unexpected contact2.Firstname, got %#v, expected %#v", contact2.Firstname, "JØhñ")
	}

	if contact2.Lastname != "ポテト 👌 ไก่" {
		t.Errorf("Unexpected contact2.Lastname, got %#v, expected %#v", contact2.Lastname, "ポテト 👌 ไก่")
	}

	contact3, err := Get(c, 0, &ticketmatic.ContactGetQuery{
		Email: "john@test.com",
	})
	if err != nil {
		t.Fatal(err)
	}

	if contact3.Id == 0 {
		t.Errorf("Unexpected contact3.Id, got %#v, expected different value", contact3.Id)
	}

	err = Delete(c, contact.Id)
	if err != nil {
		t.Fatal(err)
	}

}

func TestArchived(t *testing.T) {
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

	req, err := Getlist(c, &ticketmatic.ContactQuery{
		Includearchived: true,
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(req.Data) <= 0 {
		t.Errorf("Unexpected req.Data length, got %#v, expected greater than %#v", len(req.Data), 0)
	}

	err = Delete(c, contact.Id)
	if err != nil {
		t.Fatal(err)
	}

	req2, err := Getlist(c, &ticketmatic.ContactQuery{})
	if err != nil {
		t.Fatal(err)
	}

	if len(req.Data) <= len(req2.Data) {
		t.Errorf("Unexpected req.Data length, got %#v, expected greater than %#v", len(req.Data), len(req2.Data))
	}

	req3, err := Getlist(c, &ticketmatic.ContactQuery{
		Includearchived: true,
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(req.Data) != len(req3.Data) {
		t.Errorf("Unexpected req.Data length, got %#v, expected %#v", len(req.Data), len(req3.Data))
	}

}

func TestImport(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	contacts, err := Import(c, []*ticketmatic.Contact{
		&ticketmatic.Contact{
			Firstname: "Test",
			Lastname:  "Mc Cheer",
		},
		&ticketmatic.Contact{
			Email:     "invalid",
			Firstname: "Last",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if contacts[0].Ok != true {
		t.Errorf("Unexpected contacts[0].Ok, got %#v, expected %#v", contacts[0].Ok, true)
	}

	if contacts[1].Ok != false {
		t.Errorf("Unexpected contacts[1].Ok, got %#v, expected %#v", contacts[1].Ok, false)
	}

	if contacts[0].Id <= 0 {
		t.Errorf("Unexpected contacts[0].Id, got %#v, expected > %#v", contacts[0].Id, 0)
	}

	if contacts[1].Error != "Invalid email" {
		t.Errorf("Unexpected contacts[1].Error, got %#v, expected %#v", contacts[1].Error, "Invalid email")
	}

	err = Delete(c, contacts[0].Id)
	if err != nil {
		t.Fatal(err)
	}

}

func TestUpdatewithoptins(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	contact, err := Create(c, &ticketmatic.Contact{
		Email:     "john34@test.com",
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

	if contact.Email != "john34@test.com" {
		t.Errorf("Unexpected contact.Email, got %#v, expected %#v", contact.Email, "john34@test.com")
	}

	if len(contact.Optins) != 0 {
		t.Errorf("Unexpected contact.Optins length, got %#v, expected %#v", len(contact.Optins), 0)
	}

	updated, err := Update(c, contact.Id, &ticketmatic.Contact{
		Optins: []*ticketmatic.ContactOptIn{
			&ticketmatic.ContactOptIn{
				Info: &ticketmatic.ContactOptInInfo{
					Method:  "api",
					Remarks: "remarks",
				},
				Optinid: 1,
				Status:  7602,
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if updated.Id != contact.Id {
		t.Errorf("Unexpected updated.Id, got %#v, expected %#v", updated.Id, contact.Id)
	}

	if len(updated.Optins) != 1 {
		t.Errorf("Unexpected updated.Optins length, got %#v, expected %#v", len(updated.Optins), 1)
	}

	if updated.Optins[0].Optinid != 1 {
		t.Errorf("Unexpected updated.Optins[0].Optinid, got %#v, expected %#v", updated.Optins[0].Optinid, 1)
	}

	if updated.Optins[0].Status != 7602 {
		t.Errorf("Unexpected updated.Optins[0].Status, got %#v, expected %#v", updated.Optins[0].Status, 7602)
	}

	if updated.Optins[0].Info.Method != "api" {
		t.Errorf("Unexpected updated.Optins[0].Info.Method, got %#v, expected %#v", updated.Optins[0].Info.Method, "api")
	}

	if updated.Optins[0].Info.Remarks != "remarks" {
		t.Errorf("Unexpected updated.Optins[0].Info.Remarks, got %#v, expected %#v", updated.Optins[0].Info.Remarks, "remarks")
	}

}

func TestRemarks(t *testing.T) {
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

	remark, err := Createremark(c, contact.Id, &ticketmatic.ContactRemark{
		Content: "Hello World",
	})
	if err != nil {
		t.Fatal(err)
	}

	if remark.Id == 0 {
		t.Errorf("Unexpected remark.Id, got %#v, expected different value", remark.Id)
	}

	if remark.Content != "Hello World" {
		t.Errorf("Unexpected remark.Content, got %#v, expected %#v", remark.Content, "Hello World")
	}

	if remark.Pinned != false {
		t.Errorf("Unexpected remark.Pinned, got %#v, expected %#v", remark.Pinned, false)
	}

	updated, err := Updateremark(c, contact.Id, remark.Id, &ticketmatic.ContactRemark{
		Content: "Hello World 2",
		Pinned:  true,
	})
	if err != nil {
		t.Fatal(err)
	}

	if updated.Content != "Hello World 2" {
		t.Errorf("Unexpected updated.Content, got %#v, expected %#v", updated.Content, "Hello World 2")
	}

	if updated.Pinned != true {
		t.Errorf("Unexpected updated.Pinned, got %#v, expected %#v", updated.Pinned, true)
	}

	get, err := Getremark(c, contact.Id, remark.Id)
	if err != nil {
		t.Fatal(err)
	}

	if get.Id != remark.Id {
		t.Errorf("Unexpected get.Id, got %#v, expected %#v", get.Id, remark.Id)
	}

	if get.Content != "Hello World 2" {
		t.Errorf("Unexpected get.Content, got %#v, expected %#v", get.Content, "Hello World 2")
	}

	if get.Pinned != true {
		t.Errorf("Unexpected get.Pinned, got %#v, expected %#v", get.Pinned, true)
	}

	err = Deleteremark(c, contact.Id, remark.Id)
	if err != nil {
		t.Fatal(err)
	}

}
