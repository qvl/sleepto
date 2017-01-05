#  :zzz: sleepto

[![GoDoc](https://godoc.org/qvl.io/sleepto?status.svg)](https://godoc.org/qvl.io/sleepto)
[![Build Status](https://travis-ci.org/qvl/sleepto.svg?branch=master)](https://travis-ci.org/qvl/sleepto)
[![Go Report Card](https://goreportcard.com/badge/qvl.io/sleepto)](https://goreportcard.com/report/qvl.io/sleepto)


`sleepto` is a simple alternative to task schedulers like [Cron](https://en.wikipedia.org/wiki/Cron).

It only handles the timing and doesn't run a daemon like other schedulers do.
Instead we encourage you to use your systems default init system (for example [Systemd](#systemd)) to control your jobs.
This allows you to:

- Use and watch scheduled jobs the way you use all other services running on your system (for example using `ps`).
- Start and pause jobs like any other service on your system.
- Use your systems default logging system.
- No conflicts - next task is only scheduled after previous one finished.
- Trigger immediate execution by sending a `SIGALRM` signal.
- Specify execution times with the precision of seconds (Cron only support minutes).
- Always know the time of the next execution.

Thanks to [runwhen](http://code.dogmap.org/runwhen/) for inspiration.


    Usage: sleepto [flags...] [command...]

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
      -day value
          0 to 31
      -hour value
          0 to 23
      -minute value
          0 to 59
      -month value
          1 to 12
      -second value
          0 to 59
      -silent
          Surpress all output
      -weekday value
          mo,tu,we,th,fr,sa,su

    For more visit: https://qvl.io/sleepto



## Install

- Via [Go](https://golang.org/) setup: `go get qvl.io/sleepto`

- Or download latest binary: https://github.com/qvl/sleepto/releases


## Setup

`sleepto` can be used in different scenarios but the most common one is probably to combine it with an init system.

## [Systemd](https://en.wikipedia.org/wiki/Systemd)

*TODO*


## Development

Make sure to use `gofmt` and create a [Pull Request](https://github.com/qvl/sleepto/pulls).


### Releasing

Run `./release.sh <version>` and upload the binaries on Github.


## License

[MIT](./license)
