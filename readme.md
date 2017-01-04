#  :zzz: sleepto

[![GoDoc](https://godoc.org/qvl.io/sleepto?status.svg)](https://godoc.org/qvl.io/sleepto)
[![Build Status](https://travis-ci.org/qvl/sleepto.svg?branch=master)](https://travis-ci.org/qvl/sleepto)
[![Go Report Card](https://goreportcard.com/badge/qvl.io/sleepto)](https://goreportcard.com/report/qvl.io/sleepto)


Like [runwhen](http://code.dogmap.org/runwhen/) but simpler:

    Usage: sleepto [conditions]

    Sleep until next time the specified conditions match.

    Conditions are specified with flags.
    All flags are optional and can be used in any combination.
    The condition flags take one or more value each.
    Values are separated by comma.

    Examples:
      # Every day at 1am
      sleepto -hour 1 && dbbackup.sh
      # Every 10th of month at 3pm
      sleepto -day 10 -hour 15 && send-report
      # Every 15 minutes
      sleepto -minute 0,15,30,45 && say "Hello"

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
      -verbose
            display next run time
      -weekday value
            mo,tu,we,th,fr,sa,su

    For more visit: https://qvl.io/sleepto



## Install

- Via [Go](https://golang.org/) setup: `go get qvl.io/sleepto`

- Or download latest binary: https://github.com/qvl/sleepto/releases


## Setup

*TODO*


## Development

Make sure to use `gofmt` and create a [Pull Request](https://github.com/qvl/sleepto/pulls).


### Releasing

Run `./release.sh <version>` and upload the binaries on Github.


## License

[MIT](./license)
