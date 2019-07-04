


[![Build Status](https://travis-ci.org/mchirico/tlib.svg?branch=master)](https://travis-ci.org/mchirico/tlib)
[![codecov](https://codecov.io/gh/mchirico/tlib/branch/master/graph/badge.svg)](https://codecov.io/gh/mchirico/tlib)

[![Build Status](https://mchirico.visualstudio.com/tlib/_apis/build/status/mchirico.tlib?branchName=master)](https://mchirico.visualstudio.com/tlib/_build/latest?definitionId=6&branchName=master)
[Project](https://mchirico.visualstudio.com/tlib)
# tlib


## Testing Package

Trying to incorporate the following.

[Dave Cheney](https://www.youtube.com/watch?v=pN_lm6QqHcw)
[Globals](https://peter.bourgon.org/blog/2017/06/09/theory-of-modern-go.html)
[Testing Helpers](https://www.youtube.com/watch?v=yszygk1cpEc&feature=youtu.be&t=1609)

## Build with vendor
```
export GO111MODULE=on
go mod init
# Below will put all packages in a vendor folder
go mod vendor



go test -v -mod=vendor ./...

# Don't forget the "." in "./cmd/script" below
go build -v -mod=vendor ./...
```


## Don't forget golint

```

golint -set_exit_status $(go list ./... | grep -v /vendor/)

```


