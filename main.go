// Package main starts the sleepto binary.
// Argument parsing, usage information and the actual execution can be found here.
// See package match for independent usage of the timing functionality.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"qvl.io/sleepto/flags"
	"qvl.io/sleepto/match"
)

// Can be set in build step using -ldflags
var version string

const (
	usage = `Usage: %s [flags...] [command...]

Sleep until next time the specified conditions match.

Conditions are specified with flags.
All flags are optional and can be used in any combination.
The condition flags take one or more value each.
Values are separated by comma.

A command can be specified optionally.
All arguments following the command are passed to it.

When the process receives a SIGALRM signal it finishes immediately.

Examples:
  # Every 10th of month at 3pm
  sleepto -day 10 -hour 15 /bin/send-report
  # Every 15 minutes
  sleepto -minute 0,15,30,45 say "Hello human"
  # Every day at 1am
  sleepto -hour 1 && ~/dbbackup.sh

Flags:
`
	more = "\nFor more visit: https://qvl.io/sleepto"
)

func main() {
	var (
		silent      = flag.Bool("silent", false, "Suppress all output")
		versionFlag = flag.Bool("version", false, "Print binary version")
		month       = flags.Monthlist("month", "1 to 12")
		weekday     = flags.Weekdaylist("weekday", "mo,tu,we,th,fr,sa,su")
		day         = flags.Intlist("day", "0 to 31", 0, 31)
		hour        = flags.Intlist("hour", "0 to 23", 0, 23)
		minute      = flags.Intlist("minute", "0 to 59", 0, 59)
		second      = flags.Intlist("second", "0 to 59", 0, 59)
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, more)
	}
	flag.Parse()

	if *versionFlag {
		fmt.Printf("sleepto %s %s %s\n", version, runtime.GOOS, runtime.GOARCH)
		os.Exit(0)
	}

	now := time.Now()

	next := match.Next(now, match.Condition{
		Month:   *month,
		Weekday: *weekday,
		Day:     *day,
		Hour:    *hour,
		Minute:  *minute,
		Second:  *second,
	})

	// No conditions specified
	if next.Equal(now) {
		flag.Usage()
		os.Exit(1)
	}

	if !*silent {
		fmt.Fprintf(os.Stderr, "sleeping until: %s\n", next.Format(time.RFC1123))
	}

	// Wait for timer or SIGALRM
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGALRM)
	select {
	case <-sigChan:
	case <-time.After(next.Sub(now)):
	}

	// Replace current process if command is specified
	args := flag.Args()
	if len(args) == 0 {
		return
	}
	cmd, err := exec.LookPath(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	err = syscall.Exec(cmd, args, os.Environ())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
