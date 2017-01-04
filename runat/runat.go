package runat

import "time"

// Config ...
type Config struct {
	Month   []time.Month
	Day     []int
	Weekday []time.Weekday
	Hour    []int
	Minute  []int
	Second  []int
}

// Run ...
func Run(t time.Time, config Config) time.Time {
	noMonth := len(config.Month) == 0
	noDay := len(config.Day) == 0
	noHour := len(config.Hour) == 0
	noMinute := len(config.Minute) == 0
	noSecond := len(config.Second) == 0

	if noMonth && noDay && noHour && noMinute && noSecond {
		return t
	}
	if noDay && noHour && noMinute && noSecond {
		t = t.AddDate(0, 1, 1-t.Day()).Truncate(time.Hour * 24)
	} else if noHour && noMinute && noSecond {
		t = t.AddDate(0, 0, 1).Truncate(time.Hour * 24)
	} else if noMinute && noSecond {
		t = t.Add(time.Hour).Truncate(time.Hour)
	} else if noSecond {
		t = t.Add(time.Minute).Truncate(time.Minute)
	} else {
		t = t.Add(time.Second).Truncate(time.Second)
	}

	for {
		monthOk := noMonth
		for _, m := range config.Month {
			if t.Month() == m {
				monthOk = true
				break
			}
		}

		dayOk := noDay
		for _, d := range config.Day {
			if t.Day() == d {
				dayOk = true
				break
			}
		}

		hourOk := noHour
		for _, h := range config.Hour {
			if t.Hour() == h {
				hourOk = true
				break
			}
		}

		minuteOk := noMinute
		for _, m := range config.Minute {
			if t.Minute() == m {
				minuteOk = true
				break
			}
		}

		secondOk := noSecond
		for _, s := range config.Second {
			if t.Second() == s {
				secondOk = true
				break
			}
		}

		if !monthOk {
			t = t.AddDate(0, 1, 1-t.Day()).Truncate(time.Hour * 24)
		} else if !dayOk {
			t = t.AddDate(0, 0, 1).Truncate(time.Hour * 24)
		} else if !hourOk {
			t = t.Add(time.Hour).Truncate(time.Hour)
		} else if !minuteOk {
			t = t.Add(time.Minute).Truncate(time.Minute)
		} else if !secondOk {
			t = t.Add(time.Second).Truncate(time.Second)
		} else {
			return t
		}
	}
}
