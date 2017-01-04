package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"qvl.io/runat/flags"
	"qvl.io/runat/match"
)

func main() {
	var (
		verbose = flag.Bool("verbose", false, "display next run time")
		month   = flags.Monthlist("month", "from 1, to 12")
		weekday = flags.Weekdaylist("weekday", "mo,tu,we,th,fr,sa,su")
		day     = flags.Intlist("day", "from 0 to 31", 0, 31)
		hour    = flags.Intlist("hour", "from 0 to 23", 0, 23)
		minute  = flags.Intlist("minute", "from 0 to 59", 0, 59)
		second  = flags.Intlist("second", "from 0 to 59", 0, 59)
	)

	flag.Parse()

	now := time.Now()

	next := match.Next(now, match.Condition{
		Month:   month(),
		Weekday: weekday(),
		Day:     day(),
		Hour:    hour(),
		Minute:  minute(),
		Second:  second(),
	})

	if *verbose {
		fmt.Fprintf(os.Stderr, "Running at %s\n", next.Format(time.RFC1123))
	}

	time.Sleep(next.Sub(now))
}
