#  :floppy_disk: sleepto

[![GoDoc](https://godoc.org/qvl.io/sleepto?status.svg)](https://godoc.org/qvl.io/sleepto)
[![Build Status](https://travis-ci.org/qvl/sleepto.svg?branch=master)](https://travis-ci.org/qvl/sleepto)
[![Go Report Card](https://goreportcard.com/badge/github.com/qvl/sleepto)](https://goreportcard.com/report/github.com/qvl/sleepto)


Like [runwhen](http://code.dogmap.org/runwhen/) but simpler:

## Syntax
```
sleepto -a b ...
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
sleepto -hour 1 && ~/bin/dbbackup.sh
# Every 1st of month at 1am
sleepto -day 1 -hour 1 && ~/bin/letsencrypt-renew.sh
# Every 15 minutes
sleepto -minute 0,15,30,45 ~/bin/dbbackup.sh
```


- Trigger right away on `SIGALRM`
- Logs next execution time (to stderr) before sleeping when `-verbose` is active


## Install

- Via [Go](https://golang.org/) setup: `go get qvl.io/sleepto`

- Or download latest binary: https://github.com/qvl/sleepto/releases


## Setup

*TODO*


## Use as Go package

From another Go program you can directly use the `sleepto` sub-package.
Have a look at the [GoDoc](https://godoc.org/qvl.io/sleepto/sleepto).


## Development

Make sure to use `gofmt` and create a [Pull Request](https://github.com/qvl/sleepto/pulls).

### Releasing

Run `./release.sh <version>` and upload the binaries on Github.


## License

[MIT](./license)
