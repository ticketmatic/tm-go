package seatingplans

import (
	"os"
	"testing"

	"github.com/ticketmatic/tm-go/ticketmatic"
)

func TestCreatesinglezone(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	seatingplan, err := Create(c, &ticketmatic.SeatingPlan{
		Name:      "The Ballroom",
		Status:    "draft",
		Useszones: false,
	})
	if err != nil {
		t.Fatal(err)
	}

	if seatingplan.Name != "The Ballroom" {
		t.Errorf("Unexpected seatingplan.Name, got %#v, expected %#v", seatingplan.Name, "The Ballroom")
	}

	if seatingplan.Status != "draft" {
		t.Errorf("Unexpected seatingplan.Status, got %#v, expected %#v", seatingplan.Status, "draft")
	}

	if seatingplan.Useszones != false {
		t.Errorf("Unexpected seatingplan.Useszones, got %#v, expected %#v", seatingplan.Useszones, false)
	}

}

func TestCreatemultizone(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	seatingplan, err := Create(c, &ticketmatic.SeatingPlan{
		Name:      "The Opera House",
		Status:    "draft",
		Useszones: true,
		Zones: []int64{
			1,
			2,
			3,
			4,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if seatingplan.Name != "The Opera House" {
		t.Errorf("Unexpected seatingplan.Name, got %#v, expected %#v", seatingplan.Name, "The Opera House")
	}

	if seatingplan.Status != "draft" {
		t.Errorf("Unexpected seatingplan.Status, got %#v, expected %#v", seatingplan.Status, "draft")
	}

	if seatingplan.Useszones != false {
		t.Errorf("Unexpected seatingplan.Useszones, got %#v, expected %#v", seatingplan.Useszones, false)
	}

	if seatingplan.Zones != []interface{}{1, 2, 3, 4} {
		t.Errorf("Unexpected seatingplan.Zones, got %#v, expected %#v", seatingplan.Zones, []interface{}{1, 2, 3, 4})
	}

}

func TestGet(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	list, err := Getlist(c, &ticketmatic.SeatingPlanQuery{})
	if err != nil {
		t.Fatal(err)
	}

	if len(list.Data) <= 0 {
		t.Errorf("Unexpected list.Data length, got %#v, expected greater than %#v", len(list.Data), 0)
	}

	seatingplan, err := Get(c, list.Data[0].Id)
	if err != nil {
		t.Fatal(err)
	}

	if seatingplan.Id != list.Data[0].Id {
		t.Errorf("Unexpected seatingplan.Id, got %#v, expected %#v", seatingplan.Id, list.Data[0].Id)
	}

}
