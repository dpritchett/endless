# Endless Nameless
[![Gobuild Download](http://gobuild.io/badge/github.com/dpritchett/endless/download.png)](http://gobuild.io/github.com/dpritchett/endless)

Endless Nameless accepts any incoming HTTP requests, waits for a specified period of time, and then responds.

```sh
$ ./endless --port 1234 --delay 1.1 --response="*click*"
Endless nameless is now listening on port 1234 with a delay of 1.10 second(s)...
2014/06/04 11:18:50 [127.0.0.1]	GET /
2014/06/04 11:19:14 [127.0.0.1]	GET /
2014/06/04 11:19:17 [127.0.0.1]	POST /
2014/06/04 11:19:20 [127.0.0.1]	PUT /abcefg?q=things
```

```sh
$ curl localhost:1234
*click*
```

## Why?
I'd like to test some HTTP request timeout settings on an API client and I needed a way to simulate slow responses.

## Where
[Binary releases are available](http://gobuild.io/download/github.com/dpritchett/endless) for Linux, OSX, and Windows.  Note:  I do not know the people behind gobuild and cannot speak for its security.

## Golang users
```sh
go get -u  github.com/dpritchett/endless
go install github.com/dpritchett/endless
```
