#  :floppy_disk: runat

[![GoDoc](https://godoc.org/qvl.io/runat?status.svg)](https://godoc.org/qvl.io/runat)
[![Build Status](https://travis-ci.org/qvl/runat.svg?branch=master)](https://travis-ci.org/qvl/runat)
[![Go Report Card](https://goreportcard.com/badge/github.com/qvl/runat)](https://goreportcard.com/report/github.com/qvl/runat)


Like [runwhen](http://code.dogmap.org/runwhen/) but simpler:

## Syntax
```
runat -a b ...
```

with:

```
month   in 1,...,12
day     in 1,...,31
weekday in mo,tu,we,th,fr,sa,su
hour    in 0,...,24
minute  in 0,...,60
second  in 0,...,60
```

(24/60 alias to 0)


## Examples

``` sh
# Every day at 1am
runat -hour 1 && ~/bin/dbbackup.sh
# Every 1st of month at 1am
runat -day 1 -hour 1 && ~/bin/letsencrypt-renew.sh
# Every 15 minutes
runat -minute 0,15,30,45 ~/bin/dbbackup.sh
```


- Trigger right away on `SIGALRM`
- Logs next execution time (to stderr) before sleeping when `-verbose` is active


## Install

- Via [Go](https://golang.org/) setup: `go get qvl.io/runat`

- Or download latest binary: https://github.com/qvl/runat/releases


## Setup

*TODO*


## Use as Go package

From another Go program you can directly use the `runat` sub-package.
Have a look at the [GoDoc](https://godoc.org/qvl.io/runat/runat).


## Development

Make sure to use `gofmt` and create a [Pull Request](https://github.com/qvl/runat/pulls).

### Releasing

Run `./release.sh <version>` and upload the binaries on Github.


## License

[MIT](./license)
