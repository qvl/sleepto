package match_test

import (
	"testing"
	"time"

	"qvl.io/sleepto/match"
)

func TestNext(t *testing.T) {
	table := []struct {
		c   match.Condition
		in  time.Time
		out time.Time
	}{

		{
			match.Condition{},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
		},

		{
			match.Condition{
				Year:   []int{2020},
				Minute: []int{55, 13},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2020, 1, 1, 0, 13, 0, 0, time.UTC),
		},

		{
			match.Condition{
				Year:  []int{2014, 2016},
				Month: []time.Month{1, 6},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Time{},
		},

		{
			match.Condition{
				Month:  []time.Month{1, 6},
				Day:    []int{30, 15},
				Hour:   []int{13},
				Minute: []int{55, 13},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2017, 6, 15, 13, 13, 0, 0, time.UTC),
		},

		{
			match.Condition{
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
			match.Condition{
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
			match.Condition{
				Year:    []int{2019, 2020},
				Month:   []time.Month{1, 6},
				Day:     []int{30, 15},
				Weekday: []time.Weekday{time.Monday, time.Tuesday},
				Hour:    []int{13},
				Minute:  []int{55, 13},
				Second:  []int{44, 1, 13},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2019, 1, 15, 13, 13, 1, 0, time.UTC),
		},

		{
			match.Condition{
				Year: []int{2025},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		},

		{
			match.Condition{
				Year: []int{2017},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Time{},
		},

		{
			match.Condition{
				Year: []int{1990},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Time{},
		},

		{
			match.Condition{
				Month: []time.Month{time.February},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2018, 2, 1, 0, 0, 0, 0, time.UTC),
		},

		{
			match.Condition{
				Month: []time.Month{time.March},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2017, 3, 1, 0, 0, 0, 0, time.UTC),
		},

		{
			match.Condition{
				Month: []time.Month{time.February},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2018, 2, 1, 0, 0, 0, 0, time.UTC),
		},

		{
			match.Condition{
				Day: []int{4},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2017, 3, 4, 0, 0, 0, 0, time.UTC),
		},

		{
			match.Condition{
				Weekday: []time.Weekday{time.Friday},
			},
			time.Date(2017, 2, 10, 10, 7, 5, 8, time.UTC),
			time.Date(2017, 2, 17, 0, 0, 0, 0, time.UTC),
		},

		{
			match.Condition{
				Hour: []int{10},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2017, 2, 5, 10, 0, 0, 0, time.UTC),
		},

		{
			match.Condition{
				Hour: []int{0, 6, 18},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2017, 2, 4, 18, 0, 0, 0, time.UTC),
		},

		{
			match.Condition{
				Minute: []int{7},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2017, 2, 4, 11, 7, 0, 0, time.UTC),
		},

		{
			match.Condition{
				Minute: []int{55, 13},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2017, 2, 4, 10, 13, 0, 0, time.UTC),
		},

		{
			match.Condition{
				Second: []int{5},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2017, 2, 4, 10, 8, 5, 0, time.UTC),
		},

		{
			match.Condition{
				Second: []int{3, 4},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2017, 2, 4, 10, 8, 3, 0, time.UTC),
		},

		{
			match.Condition{
				Month: []time.Month{time.February},
				Day:   []int{10},
			},
			time.Date(2017, 2, 4, 0, 7, 5, 8, time.UTC),
			time.Date(2017, 2, 10, 0, 0, 0, 0, time.UTC),
		},

		{
			match.Condition{
				Month: []time.Month{time.February},
				Day:   []int{3},
			},
			time.Date(2017, 2, 4, 0, 7, 5, 8, time.UTC),
			time.Date(2018, 2, 3, 0, 0, 0, 0, time.UTC),
		},

		{
			match.Condition{
				Month:  []time.Month{time.February},
				Second: []int{30},
			},
			time.Date(2017, 2, 4, 0, 7, 5, 8, time.UTC),
			time.Date(2017, 2, 4, 0, 7, 30, 0, time.UTC),
		},

		{
			match.Condition{
				Month: []time.Month{time.February},
				Hour:  []int{10},
			},
			time.Date(2017, 2, 4, 0, 7, 5, 8, time.UTC),
			time.Date(2017, 2, 4, 10, 0, 0, 0, time.UTC),
		},

		{
			match.Condition{
				Day:    []int{4, 6},
				Minute: []int{2, 6},
			},
			time.Date(2017, 2, 4, 0, 7, 5, 8, time.UTC),
			time.Date(2017, 2, 4, 1, 2, 0, 0, time.UTC),
		},

		{
			match.Condition{
				Weekday: []time.Weekday{time.Friday, time.Tuesday},
				Minute:  []int{2, 6},
			},
			time.Date(2017, 2, 4, 0, 7, 5, 8, time.UTC),
			time.Date(2017, 2, 7, 0, 2, 0, 0, time.UTC),
		},

		{
			match.Condition{
				Hour:   []int{5, 15},
				Second: []int{0, 30},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2017, 2, 4, 15, 0, 0, 0, time.UTC),
		},

		{
			match.Condition{
				Hour:   []int{5, 15},
				Second: []int{0, 30},
			},
			time.Date(2017, 2, 4, 15, 7, 30, 0, time.UTC),
			time.Date(2017, 2, 4, 15, 8, 0, 0, time.UTC),
		},

		{
			match.Condition{
				Month: []time.Month{1, 10},
				Hour:  []int{2, 14},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2017, 10, 1, 2, 0, 0, 0, time.UTC),
		},

		{
			match.Condition{
				Day:  []int{1, 20},
				Hour: []int{2, 14},
			},
			time.Date(2017, 2, 4, 10, 7, 5, 8, time.UTC),
			time.Date(2017, 2, 20, 2, 0, 0, 0, time.UTC),
		},
	}

	for i, tt := range table {
		res := match.Next(tt.in, tt.c)
		if !res.Equal(tt.out) {
			t.Errorf(`
%d.
Expected: %v (%v)
Got:      %v (%v)
`, i, tt.out, tt.out.Weekday(), res, res.Weekday())
		}
	}

}
