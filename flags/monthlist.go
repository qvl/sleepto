package flags

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type monthlist []time.Month

func (l *monthlist) String() string {
	s := make([]string, len(*l))
	for i := range *l {
		s[i] = strconv.Itoa(int((*l)[i]))
	}
	return strings.Join(s, ",")
}

func (l *monthlist) Set(s string) error {
	parts := strings.Split(s, ",")
	for i, p := range parts {
		xint, err := strconv.Atoi(p)
		if err != nil {
			return fmt.Errorf("no integer at index %d: %s", i, p)
		}
		x := time.Month(xint)
		if x < 1 || x > 12 {
			return fmt.Errorf("invalid month: %d", x)
		}
		*l = monthlist(append(*l, x))
	}
	return nil
}

// Monthlist defines a flag for a comma-separated list of months.
// Valid values are between 1 and 12.
func Monthlist(name, usage string) *[]time.Month {
	l := &[]time.Month{}
	flag.Var((*monthlist)(l), name, usage)
	return l
}
