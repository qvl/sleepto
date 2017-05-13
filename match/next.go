// Package match provides functionality to find matching times.
// This is the core logic of sleepto.
package match

import "time"

// Condition is used to match a time.
// All fields are optional and can be used in any combination.
// For each field one value of the list has
// to match to find a match for the condition.
type Condition struct {
	Year    []int
	Month   []time.Month
	Day     []int // 1 to 31
	Weekday []time.Weekday
	Hour    []int // 0 to 23
	Minute  []int // 0 to 59
	Second  []int // 0 to 59
}

// Next finds the next time the passed condition matches.
// Returns an empty time.Time when no possible match can be found.
// This can only happen when there is no future year condition.
// Use .IsZero() to test if result is empty.
func Next(start time.Time, c Condition) time.Time {
	if noMatch(start, c) {
		return time.Time{}
	}
	t := setBase(start, c)
	// Stop when when no condition
	if t.Equal(start) {
		return t
	}

	// Walk until all units match.
	// Adjust biggest unit first.
	for {
		switch {
		case wrong(c.Year, t.Year()):
			t = addYear(t)
		case wrongMonth(c.Month, t.Month()):
			t = addMonth(t)
		case wrong(c.Day, t.Day()) || wrongWeekday(c.Weekday, t.Weekday()):
			t = addDay(t)
		case wrong(c.Hour, t.Hour()):
			t = addHour(t)
		case wrong(c.Minute, t.Minute()):
			t = addMinute(t)
		case wrong(c.Second, t.Second()):
			t = addSecond(t)
		default:
			// Found matching time.
			return t
		}
	}
}

// Find smallest unit and start counting from there.
// At least have to increment by one.
func setBase(t time.Time, c Condition) time.Time {
	switch {
	case len(c.Second) > 0:
		return addSecond(t)
	case len(c.Minute) > 0:
		return addMinute(t)
	case len(c.Hour) > 0:
		return addHour(t)
	case len(c.Day) > 0 || len(c.Weekday) > 0:
		return addDay(t)
	case len(c.Month) > 0:
		return addMonth(t)
	case len(c.Year) > 0:
		return addYear(t)
	default:
		return t
	}
}

func noMatch(t time.Time, c Condition) bool {
	for _, y := range c.Year {
		if y <= t.Year() {
			return true
		}
	}
	return false
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

func addYear(t time.Time) time.Time {
	return t.AddDate(1, 1-int(t.Month()), 1-t.Day()).Truncate(time.Hour * 24)
}

func addMonth(t time.Time) time.Time {
	return t.AddDate(0, 1, 1-t.Day()).Truncate(time.Hour * 24)
}

func addDay(t time.Time) time.Time {
	return t.AddDate(0, 0, 1).Truncate(time.Hour * 24)
}

func addHour(t time.Time) time.Time {
	return t.Add(time.Hour).Truncate(time.Hour)
}

func addMinute(t time.Time) time.Time {
	return t.Add(time.Minute).Truncate(time.Minute)
}

func addSecond(t time.Time) time.Time {
	return t.Add(time.Second).Truncate(time.Second)
}
