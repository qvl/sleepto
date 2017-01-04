package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"qvl.io/runat/flags"
	"qvl.io/runat/match"
)

const usage = `Usage: %s [conditions]

Sleep until next time the specified conditions match.

Conditions are specified with flags.
All flags are optional and can be used in any combination.
The condition flags take one or more value each.
Values are separated by comma.

Examples:
  # Every day at 1am
  runat -hour 1 && dbbackup.sh
  # Every 10th of month at 3pm
  runat -day 10 -hour 15 && send-report
  # Every 15 minutes
  runat -minute 0,15,30,45 && say "Hello"

Flags:
`

func main() {
	var (
		verbose = flag.Bool("verbose", false, "display next run time")
		month   = flags.Monthlist("month", "1 to 12")
		weekday = flags.Weekdaylist("weekday", "mo,tu,we,th,fr,sa,su")
		day     = flags.Intlist("day", "0 to 31", 0, 31)
		hour    = flags.Intlist("hour", "0 to 23", 0, 23)
		minute  = flags.Intlist("minute", "0 to 59", 0, 59)
		second  = flags.Intlist("second", "0 to 59", 0, 59)
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
		flag.PrintDefaults()
	}
	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(1)
	}
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
