package flags

import (
	"testing"
	"time"
)

func TestWeekdaylist(t *testing.T) {
	tests := []struct {
		text    string
		parsed  []time.Weekday
		invalid bool
	}{
		{text: "tu,th,fr", parsed: []time.Weekday{time.Tuesday, time.Thursday, time.Friday}},
		{text: "sa,su", parsed: []time.Weekday{time.Saturday, time.Sunday}},
		{text: "mo", parsed: []time.Weekday{time.Monday}},
		{text: "mo,tu,", invalid: true},
		{text: "1,2", invalid: true},
		{text: "0", invalid: true},
		{text: "", invalid: true},
		{text: "hello", invalid: true},
	}

	for i, tt := range tests {
		l := weekdaylist{}
		if err := l.Set(tt.text); err != nil {
			if !tt.invalid {
				t.Errorf("parsing %s failed unexpectedly: %v", tt.text, err)
			}
			continue
		}
		if tt.invalid {
			t.Errorf("parsing %s should have failed", tt.text)
			continue
		}
		if !equalWeekday(l.list, tt.parsed) {
			t.Errorf(`
%d.
Input:    %s
Expected: %v
Got       %v`, i, tt.text, tt.parsed, l.list)
		}
	}

	if (&weekdaylist{}).String() != "" {
		t.Errorf("Non empty String() output: %s", (&weekdaylist{}).String())
	}
}

func equalWeekday(a []time.Weekday, b []time.Weekday) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
