# Go Future [![Last release](https://img.shields.io/github/release/euskadi31/go-future.svg)](https://github.com/euskadi31/go-future/releases/latest) [![Documentation](https://godoc.org/github.com/euskadi31/go-future?status.svg)](https://godoc.org/github.com/euskadi31/go-future)

[![Go Report Card](https://goreportcard.com/badge/github.com/euskadi31/go-future)](https://goreportcard.com/report/github.com/euskadi31/go-future)

| Branch | Status                                                                                                                                            | Coverage                                                                                                                                       |
| ------ | ------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| master | [![Go](https://github.com/euskadi31/go-einfo/actions/workflows/go.yml/badge.svg)](https://github.com/euskadi31/go-einfo/actions/workflows/go.yml) | [![Coveralls](https://img.shields.io/coveralls/euskadi31/go-future/master.svg)](https://coveralls.io/github/euskadi31/go-future?branch=master) |

go-future is an implementation of future in Go.

## Example

```go
package main

import (
    "fmt"

    "github.com/euskadi31/go-future"
)

func asyncFunc() *future.Future {
    f := future.New()

    go func(f *future.Future) {
        f.Value("my async value")
    }(f)

    return f
}

func main() {
    f := asyncFunc();

    v, err := f.Get()
    if err != nil {
        panic(err)
    }

    fmt.Println(v)

}
```

With `Fill`

```go
package main

import (
    "fmt"

    "github.com/euskadi31/go-future"
)

func asyncFunc() *future.Future {
    f := future.New()

    go func(f *future.Future) {
        f.Value("my async value")
    }(f)

    return f
}

func main() {
    f := asyncFunc();

    var v string
    if err := f.Fill(&v); err != nil {
        panic(err)
    }

    fmt.Println(v)
}
```

## License

go-future is licensed under [the MIT license](LICENSE.md).
