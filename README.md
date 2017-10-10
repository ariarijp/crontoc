# crontoc

Show crontab lines with next execution time.

## Installation

```bash
$ go get -u github.com/ariarijp/crontoc
```

## Usage

```bash
$ crontab -l
# m h  dom mon dow   command
* * * * * vmstat >> /tmp/vmstat.log
10 0 * * * df -h >> /tmp/df.log
*/5 0 * * * free >> /tmp/free.log
0 10 1 * * date >> /tmp/date.log
$ crontoc --help
Usage of crontoc:
  -from string
    	Show TOC from its time. (default "2017-10-10T21:58:34")
  -sort
    	Sort by next execution time.
$ crontab -l | crontoc
Next: 2017-09-19 23:33:00 +0900 JST # * * * * * vmstat >> /tmp/vmstat.log
Next: 2017-09-20 00:10:00 +0900 JST # 10 0 * * * df -h >> /tmp/df.log
Next: 2017-09-20 00:00:00 +0900 JST # */5 0 * * * free >> /tmp/free.log
Next: 2017-10-01 10:00:00 +0900 JST # 0 10 1 * * date >> /tmp/date.log
$ crontab -l | crontoc -sort
2017-09-19 23:35:00 +0900 JST # * * * * * vmstat >> /tmp/vmstat.log
2017-09-20 00:00:00 +0900 JST # */5 0 * * * free >> /tmp/free.log
2017-09-20 00:10:00 +0900 JST # 10 0 * * * df -h >> /tmp/df.log
2017-10-01 10:00:00 +0900 JST # 0 10 1 * * date >> /tmp/date.log
```

## License

MIT

## Author

[ariarijp](https://github.com/ariarijp)
