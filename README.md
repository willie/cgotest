# This demonstrates a breaking change in Go 1.15

Years ago, [I decided to statically compile the Go HTTP client so I could use it from C/C++.](https://twitter.com/bradfitz/status/880536642007711744). And since then, it’s become the best way for me to move some complex legacy C++ code forward while taking advantage of new libraries that are more mature in the Go ecosystem.

This has worked well until Go 1.15 and beyond. And that is what this repo is. A simple example of how it broke in Go 1.15. Hopefully, someone (cough, cough, Ian) can tell me what I’m doing wrong.

This sample is intended to be run on a Mac, but it fails on Linux as well (comment out the `-framework` directives to gcc).

## go 1.14
```
% cd test
% ./build.sh
go version go1.14 darwin/amd64
Hello Ian.
```

## go 1.15
```
% cd test
% ./build.sh
go version go1.15 darwin/amd64
# command-line-arguments
_cgo_export.c:29:3: error: unknown type name 'writeFunc'
...other errors follow...
```
