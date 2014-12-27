
# go-trace

 Golang tracing utility.

 View the [docs](http://godoc.org/github.com/tj/go-trace).

## Installation

```
$ go get github.com/tj/go-trace
```

## About

 This package allows you to instrument a program with trace probes which
 emit arbitrary data. Zero or more plugins can act on the events emitted,
 otherwise they are simply discarded.

 The primary use-case for this package is to provide instrumentation for
 top-like utilities, as well as debugging when logging does not suffice.

# License

MIT
