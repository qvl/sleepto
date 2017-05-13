package flags

import "testing"

func TestYearlist(t *testing.T) {
	tests := []struct {
		text    string
		parsed  []int
		invalid bool
	}{
		{text: "0,2000", parsed: []int{0, 2000}},
		{text: "2017", parsed: []int{2017}},
		{text: "1999,57", parsed: []int{1999, 57}},
		{text: "-2000,2000", parsed: []int{-2000, 2000}},
		{text: "", invalid: true},
		{text: "-1, 2", invalid: true},
		{text: "hello", invalid: true},
	}

	for i, tt := range tests {
		l := yearlist{}
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
		if !equal(l.list, tt.parsed) {
			t.Errorf(`
%d.
Input:    %s
Expected: %v
Got       %v`, i, tt.text, tt.parsed, l.list)
		}
	}

	if (&yearlist{}).String() != "" {
		t.Errorf("Non empty String() output: %s", (&yearlist{}).String())
	}
}
