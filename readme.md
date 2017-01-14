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
      -version
          Print binary version
      -weekday value
          mo,tu,we,th,fr,sa,su

    For more visit: https://qvl.io/sleepto



## Install

- With [Go](https://golang.org/):
```
go get qvl.io/sleepto
```

- With [Homebrew](http://brew.sh/):
```
brew install qvl/tap/sleepto
```

- Download from https://github.com/qvl/sleepto/releases


## Setup

`sleepto` can be used in different scenarios but the most common one is probably to combine it with an init system.

### [Systemd](https://en.wikipedia.org/wiki/Systemd)

[Systemd](https://en.wikipedia.org/wiki/Systemd) already runs on most Linux systems.

It even has its own [timer implementation](https://www.freedesktop.org/software/systemd/man/systemd.timer.html).
However, it's more complex than what you probably need and for many scenarios you might prefer to use the simple and portable `sleepto` instead.

See [ghbackup for an example](https://github.com/qvl/ghbackup#systemd-and-sleepto) on how to use `sleepto` in a service.

- See the logs for your service use:
```
sudo journalctl -u servicename
```
- List processes which should include your service:
```
ps fux
```
- Immediately finish sleeping:
```
kill -s ALRM YOUR_PID
```


## Development

Make sure to use `gofmt` and create a [Pull Request](https://github.com/qvl/sleepto/pulls).


### Releasing

Push a new Git tag and [GoReleaser](https://github.com/goreleaser/releaser) will automatically create a release.


## License

[MIT](./license)
