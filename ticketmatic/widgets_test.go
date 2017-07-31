package ticketmatic

import (
	"net/url"
	"testing"
)

func TestWidgets(t *testing.T) {
	widgets := NewWidgets("club", "142dda885ec6024f934a40c1", "abd2e5893bd447dc7331af1db8df42fdc62fc5c8f9f04784")

	u := widgets.GenerateUrl("addtickets", map[string]string{
		"event":     "123",
		"skinid":    "25",
		"returnurl": "http://www.ticketmatic.com",
		"l":         "fr",
	})

	url, err := url.Parse(u)
	if err != nil {
		t.Fatal(err)
	}

	if url.Path != "/widgets/club/addtickets" {
		t.Errorf("Unexpected path: %s", url.Path)
	}

	q := url.Query()

	expected := map[string]string{
		"event":     "123",
		"skinid":    "25",
		"returnurl": "http://www.ticketmatic.com",
		"l":         "fr",
		"accesskey": "142dda885ec6024f934a40c1",
		"signature": "ae727e02cea8c27322a24af487af950b4d8d26978e57151bfed7e356dd593c00",
	}

	if len(q) != len(expected) {
		t.Errorf("Bad number of parameters: %d", len(q))
	}
	for k, v := range expected {
		if q.Get(k) != v {
			t.Errorf("Unexpected value for %s: %s", k, q.Get(k))
		}
	}
}
