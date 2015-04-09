package ticketmatic

import (
	"fmt"
	"time"
)

// Time is a wrapper that takes care of correct time serialization. You can
// easily convert from/to a normal time.Time using the NewTime() and Time()
// functions.
type Time struct {
	ts time.Time
}

// Wrap a time.Time into a ticketmatic.Time
func NewTime(t time.Time) Time {
	return Time{ts: t}
}

// Convert a ticketmatic.Time into a regular time.Time
func (t Time) Time() time.Time {
	return t.ts
}

// Custom unmarshalling to handle the different time formats that may be
// returned by Ticketmatic.
func (t *Time) UnmarshalJSON(data []byte) error {
	s := string(data)
	s = s[1 : len(s)-1]
	if len(s) >= 19 {
		if s[10] == ' ' {
			// We're not always fully consistent in date outputs (for
			// historical reasons). In this case got a string like
			// "2015-04-09 14:19:49". Correcting.
			s = fmt.Sprintf("%sT%s", s[:10], s[11:])
		}

		// "2015-04-09T14:19:49"
		if len(s) == 19 {
			ts, err := time.Parse("2006-01-02T15:04:05", s)
			if err != nil {
				return err
			}
			t.ts = ts
			return nil
		} else {
			ts, err := time.Parse("2006-01-02T15:04:05.999999", s)
			if err != nil {
				return err
			}
			t.ts = ts
			return nil
		}
	} else if len(s) == 10 {
		// "2015-04-09"
		ts, err := time.Parse("2006-01-02", s)
		if err != nil {
			return err
		}
		t.ts = ts
		return nil
	}

	return fmt.Errorf("Unknown date format: %s", s)
}

// Marshal Time to JSON.
func (t *Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ts.Format("2006-01-02T15:04:05.999999"))), nil
}
