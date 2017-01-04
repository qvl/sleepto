package flags

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type monthlist struct {
	list []time.Month
}

func (l *monthlist) String() string {
	s := make([]string, len(l.list))
	for i := range l.list {
		s[i] = string(l.list[i])
	}
	return strings.Join(s, ",")
}

func (l *monthlist) Set(s string) error {
	parts := strings.Split(s, ",")
	for i, p := range parts {
		i64, err := strconv.ParseInt(p, 10, 64)
		if err != nil {
			return fmt.Errorf("no integer at index %d: %s", i, p)
		}
		x := time.Month(i64)
		if x < 1 || x > 12 {
			return fmt.Errorf("invalid month: %d", x)
		}
		l.list = append(l.list, x)
	}
	return nil
}

// Monthlist defines a flag for a comma-separated list of months.
// Valid values are between 1 and 12.
func Monthlist(name, usage string) []time.Month {
	l := &monthlist{}
	flag.Var(l, name, usage)
	return l.list
}
