package ticketmatic

import (
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	amsLocation, err := time.LoadLocation("Europe/Amsterdam")
	if err != nil {
		t.Fatalf("failed to load Europe/Amsterdam")
	}
	time.Local = amsLocation
	testcases := []struct {
		input    string
		expected string
	}{
		{"2019-04-29 20:00:00", "2019-04-29T20:00:00+02:00"},
		{"2019-04-29T20:00:00", "2019-04-29T20:00:00+02:00"},
		{"2019-04-29", "2019-04-29T00:00:00+02:00"},
	}
	for _, tc := range testcases {
		got, err := ParseTime(tc.input)
		if err != nil {
			t.Fatalf("parse '%s' failed: %+v", tc.input, err)
		}
		gotFormatted := got.Format(time.RFC3339Nano)
		if gotFormatted != tc.expected {
			t.Fatalf("expected '%s' but got '%s'", tc.expected, gotFormatted)
		}
	}
}
