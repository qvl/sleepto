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

type weekdaylist []time.Weekday

func (l *weekdaylist) String() string {
	s := make([]string, len(*l))
	for i := range *l {
		s[i] = dayToString((*l)[i])
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
		*l = weekdaylist(append(*l, x))
	}
	return nil
}

// Weekdaylist defines a flag for a comma-separated list of week days.
// Valid values are mo, tu, we, th, fr, sa, su.
func Weekdaylist(name, usage string) *[]time.Weekday {
	l := &[]time.Weekday{}
	flag.Var((*weekdaylist)(l), name, usage)
	return l
}
