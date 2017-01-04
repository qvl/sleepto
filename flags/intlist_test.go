package flags

import "testing"

func TestIntlist(t *testing.T) {
	tests := []struct {
		text    string
		parsed  []int
		invalid bool
	}{
		{text: "0,1,2,3", parsed: []int{0, 1, 2, 3}},
		{text: "0", parsed: []int{0}},
		{text: "0,10", parsed: []int{0, 10}},
		{text: "0,0,2,2,3", parsed: []int{0, 0, 2, 2, 3}},
		{text: "", invalid: true},
		{text: "-1,2,3", invalid: true},
		{text: "11", invalid: true},
		{text: "100", invalid: true},
		{text: "hello", invalid: true},
	}

	for i, tt := range tests {
		l := intlist{min: 0, max: 10}
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

}

func equal(a []int, b []int) bool {
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
