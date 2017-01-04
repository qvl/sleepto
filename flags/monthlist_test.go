package flags

import (
	"testing"
	"time"
)

func TestMonthlist(t *testing.T) {
	tests := []struct {
		text    string
		parsed  []time.Month
		invalid bool
	}{
		{text: "1,2,3,4", parsed: []time.Month{1, 2, 3, 4}},
		{text: "1,12", parsed: []time.Month{1, 12}},
		{text: "1,2,2", parsed: []time.Month{1, 2, 2}},
		{text: "0,1,2", invalid: true},
		{text: "1,", invalid: true},
		{text: "0", invalid: true},
		{text: "", invalid: true},
		{text: "hello", invalid: true},
	}

	for i, tt := range tests {
		l := monthlist{}
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
		if !equalMonth(l.list, tt.parsed) {
			t.Errorf(`
%d.
Input:    %s
Expected: %v
Got       %v`, i, tt.text, tt.parsed, l.list)
		}
	}

}

func equalMonth(a []time.Month, b []time.Month) bool {
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
