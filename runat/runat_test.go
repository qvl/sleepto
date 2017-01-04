package runat_test

import (
	"testing"
	"time"

	"qvl.io/runat/runat"
)

func TestRun(t *testing.T) {
	table := []struct {
		c   runat.Config
		in  time.Time
		out time.Time
	}{

		{
			runat.Config{
				Month:  []time.Month{1, 6},
				Day:    []int{30, 15},
				Hour:   []int{13},
				Minute: []int{55, 13},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2017, 6, 15, 13, 13, 0, 0, time.UTC),
		},

		{
			runat.Config{
				Month:  []time.Month{1, 6},
				Day:    []int{30, 15},
				Hour:   []int{13},
				Minute: []int{55, 13},
				Second: []int{44, 1, 13},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2017, 6, 15, 13, 13, 1, 0, time.UTC),
		},

		{
			runat.Config{
				Month:   []time.Month{1, 6},
				Day:     []int{30, 15},
				Weekday: []time.Weekday{time.Monday, time.Tuesday},
				Hour:    []int{13},
				Minute:  []int{55, 13},
				Second:  []int{44, 1, 13},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2018, 1, 15, 13, 13, 1, 0, time.UTC),
		},

		{
			runat.Config{
				Month: []time.Month{time.March},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2017, 3, 1, 0, 0, 0, 0, time.UTC),
		},

		{
			runat.Config{
				Month: []time.Month{time.February},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2018, 2, 1, 0, 0, 0, 0, time.UTC),
		},

		{
			runat.Config{
				Day: []int{3, 1},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2017, 3, 1, 0, 0, 0, 0, time.UTC),
		},

		{
			runat.Config{
				Weekday: []time.Weekday{time.Friday},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2017, 2, 10, 0, 0, 0, 0, time.UTC),
		},

		{
			runat.Config{
				Hour: []int{0, 6, 18},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2017, 2, 4, 18, 0, 0, 0, time.UTC),
		},

		{
			runat.Config{
				Minute: []int{55, 13},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2017, 2, 4, 10, 13, 0, 0, time.UTC),
		},

		{
			runat.Config{
				Second: []int{3, 4},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2017, 2, 4, 10, 8, 3, 0, time.UTC),
		},

		{
			runat.Config{
				Month: []time.Month{time.February},
				Day:   []int{10},
			},
			time.Date(2017, 2, 4, 0, 7, 5, 8, time.UTC),
			time.Date(2017, 2, 10, 0, 0, 0, 0, time.UTC),
		},

		{
			runat.Config{
				Month: []time.Month{time.February},
				Day:   []int{3},
			},
			time.Date(2017, 2, 4, 0, 7, 5, 8, time.UTC),
			time.Date(2018, 2, 3, 0, 0, 0, 0, time.UTC),
		},

		{
			runat.Config{
				Month:  []time.Month{time.February},
				Second: []int{30},
			},
			time.Date(2017, 2, 4, 0, 7, 5, 8, time.UTC),
			time.Date(2017, 2, 4, 0, 7, 30, 0, time.UTC),
		},

		{
			runat.Config{
				Month: []time.Month{time.February},
				Hour:  []int{10},
			},
			time.Date(2017, 2, 4, 0, 7, 5, 8, time.UTC),
			time.Date(2017, 2, 4, 10, 0, 0, 0, time.UTC),
		},

		{
			runat.Config{
				Day:    []int{4, 6},
				Minute: []int{2, 6},
			},
			time.Date(2017, 2, 4, 0, 7, 5, 8, time.UTC),
			time.Date(2017, 2, 4, 1, 2, 0, 0, time.UTC),
		},

		{
			runat.Config{
				Weekday: []time.Weekday{time.Friday, time.Tuesday},
				Minute:  []int{2, 6},
			},
			time.Date(2017, 2, 4, 0, 7, 5, 8, time.UTC),
			time.Date(2017, 2, 7, 0, 2, 0, 0, time.UTC),
		},

		{
			runat.Config{
				Hour:   []int{5, 15},
				Second: []int{0, 30},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2017, 2, 4, 15, 0, 0, 0, time.UTC),
		},

		{
			runat.Config{
				Hour:   []int{5, 15},
				Second: []int{0, 30},
			},
			time.Date(2017, 2, 4, 15, 7, 30, 0, time.UTC),
			time.Date(2017, 2, 4, 15, 8, 0, 0, time.UTC),
		},

		{
			runat.Config{
				Month: []time.Month{1, 10},
				Hour:  []int{2, 14},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2017, 10, 1, 2, 0, 0, 0, time.UTC),
		},

		{
			runat.Config{
				Day:  []int{1, 20},
				Hour: []int{2, 14},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2017, 2, 20, 2, 0, 0, 0, time.UTC),
		},
	}

	for i, tt := range table {
		res := runat.Run(tt.in, tt.c)
		if !res.Equal(tt.out) {
			t.Errorf(`
%d.
Expected: %v (%v)
Got:      %v (%v)
`, i, tt.out, tt.out.Weekday(), res, res.Weekday())
		}
	}

}
