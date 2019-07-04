


[![Build Status](https://travis-ci.org/mchirico/tlib.svg?branch=master)](https://travis-ci.org/mchirico/tlib)
[![codecov](https://codecov.io/gh/mchirico/tlib/branch/master/graph/badge.svg)](https://codecov.io/gh/mchirico/tlib)
[![Build Status](https://mchirico.visualstudio.com/tlib/_apis/build/status/mchirico.tlib?branchName=master)](https://mchirico.visualstudio.com/tlib/_build/latest?definitionId=6&branchName=master)

[![Board Status](https://mchirico.visualstudio.com/3942235e-81df-4d3e-b690-764f0e8da2b6/6e3787fa-7f1b-4ada-90e5-9b0822963cdd/_apis/work/boardbadge/8ce57aad-6547-4e5e-861d-17c9cf9237df)](https://mchirico.visualstudio.com/3942235e-81df-4d3e-b690-764f0e8da2b6/_boards/board/t/6e3787fa-7f1b-4ada-90e5-9b0822963cdd/Microsoft.RequirementCategory/)

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


