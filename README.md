# runat

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
- Waits forever (or til `SIGALRM`) when run without any conditions
- Logs next execution time (to stderr) before sleeping when `-verbose` is active
