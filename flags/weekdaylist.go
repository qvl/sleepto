package flags

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

var strDays = map[string]time.Weekday{
	"mo": time.Monday,
	"tu": time.Tuesday,
	"we": time.Wednesday,
	"th": time.Thursday,
	"fr": time.Friday,
	"sa": time.Saturday,
	"su": time.Sunday,
}

type weekdaylist struct {
	list []time.Weekday
}

func (l *weekdaylist) String() string {
	s := make([]string, len(l.list))
	for i := range l.list {
		s[i] = dayToString(l.list[i])
	}
	return strings.Join(s, ",")
}

func dayToString(d time.Weekday) string {
	for k, v := range strDays {
		if v == d {
			return k
		}
	}
	return ""
}

func (l *weekdaylist) Set(s string) error {
	days := strings.Split(s, ",")
	for i, d := range days {
		x, ok := strDays[d]
		if !ok {
			return fmt.Errorf("invalid day at index %d: %s", i, d)
		}
		l.list = append(l.list, x)
	}
	return nil
}

// Weekdaylist defines a flag for a comma-separated list of week days.
// Valid values are mo, tu, we, th, fr, sa, su.
// Call the returned function after flag.Parse to get the value.
func Weekdaylist(name string) func() []time.Weekday {
	l := &weekdaylist{}
	flag.Var(l, name, "mo,tu,we,th,fr,sa,su")
	return func() []time.Weekday {
		return l.list
	}
}
