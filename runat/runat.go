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
	// Find smallest unit and start counting from there.
	// At least have to increment by one.
	switch {
	case len(config.Second) > 0:
		t = t.Add(time.Second).Truncate(time.Second)
	case len(config.Minute) > 0:
		t = t.Add(time.Minute).Truncate(time.Minute)
	case len(config.Hour) > 0:
		t = t.Add(time.Hour).Truncate(time.Hour)
	case len(config.Day) > 0:
		t = t.AddDate(0, 0, 1).Truncate(time.Hour * 24)
	case len(config.Month) > 0:
		t = t.AddDate(0, 1, 1-t.Day()).Truncate(time.Hour * 24)
	default:
		// Empty config
		return t
	}

	// Walk until all units match.
	// Adjust biggest unit first.
	for {
		switch {
		case wrongMonth(config.Month, t.Month()):
			t = t.AddDate(0, 1, 1-t.Day()).Truncate(time.Hour * 24)
		case wrong(config.Day, t.Day()):
			t = t.AddDate(0, 0, 1).Truncate(time.Hour * 24)
		case wrong(config.Hour, t.Hour()):
			t = t.Add(time.Hour).Truncate(time.Hour)
		case wrong(config.Minute, t.Minute()):
			t = t.Add(time.Minute).Truncate(time.Minute)
		case wrong(config.Second, t.Second()):
			t = t.Add(time.Second).Truncate(time.Second)
		default:
			// Found matching time.
			return t
		}
	}
}

func wrong(xs []int, x int) bool {
	if len(xs) == 0 {
		return false
	}
	for _, y := range xs {
		if x == y {
			return false
		}
	}
	return true
}

func wrongMonth(ms []time.Month, m time.Month) bool {
	xs := make([]int, len(ms))
	for i := range ms {
		xs[i] = int(ms[i])
	}
	return wrong(xs, int(m))
}

func wrongWeekday(ds []time.Weekday, d time.Weekday) bool {
	xs := make([]int, len(ds))
	for i := range ds {
		xs[i] = int(ds[i])
	}
	return wrong(xs, int(d))
}
