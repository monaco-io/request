# Request

<img align="right" width="159px" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png">

[![Build Status](https://travis-ci.org/monaco-io/request.svg?branch=master)](https://travis-ci.org/monaco-io/request)
[![GoDoc](https://godoc.org/github.com/monaco-io/request?status.svg)](https://pkg.go.dev/github.com/monaco-io/request?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/monaco-io/request)](https://goreportcard.com/badge/github.com/monaco-io/request)
[![codecov](https://codecov.io/gh/monaco-io/request/branch/master/graph/badge.svg)](https://codecov.io/gh/monaco-io/request)
[![Release](https://img.shields.io/github/release/monaco-io/request.svg?style=flat-square)](https://github.com/monaco-io/request/releases)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/monaco-io/request)](https://www.tickgit.com/browse?repo=github.com/monaco-io/request)
[![Sourcegraph](https://sourcegraph.com/github.com/monaco-io/request/-/badge.svg)](https://sourcegraph.com/github.com/monaco-io/request?badge)
[![Open Source Helpers](https://www.codetriage.com/monaco-io/request/badges/users.svg)](https://www.codetriage.com/monaco-io/request)
[![Join the chat at https://gitter.im/monaco-io/request](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/monaco-io/request?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

HTTP client for golang, Inspired by [Javascript-axios](https://github.com/axios/axios) [Python-request](https://github.com/psf/requests).
If you have experience about axios or requests, you will love it.
No 3rd dependency.

## Features

- Make [http](https://golang.org) requests from Golang
- Intercept request and response
- Transform request and response data

## Installing

go mod:

```bash
go get github.com/monaco-io/request
```

## Methods

- OPTIONS
- GET
- HEAD
- POST
- PUT
- DELETE
- TRACE
- CONNECT

## Example

### GET

```go
package main

import (
    "log"

    "github.com/monaco-io/request"
)

func main() {
    client := request.Client{
        URL:    "https://baidu.com",
        Method: "GET",
        Params: map[string]string{"hello": "world"},
    }
    resp, err := client.Do()

    log.Println(string(resp), err)
}
```

### POST

```go
package main

import (
    "log"

    "github.com/monaco-io/request"
)

func main() {
    client := request.Client{
        URL:    "https://baidu.com",
        Method: "POST",
        Params: map[string]string{"hello": "world"},
        Body:   []byte(`{"hello": "world"}`),
    }
    resp, err := client.Do()

    log.Println(string(resp), err)
}
```

### Content-Type

```go
package main

import (
    "log"

    "github.com/monaco-io/request"
)

func main() {
    client := request.Client{
        URL:         "https://baidu.com",
        Method:      "POST",
        ContentType: request.ApplicationXWwwFormURLEncoded, // default is "application/json"
    }
    resp, err := client.Do()

    log.Println(string(resp), err)
}
```

### Authorization

```go
package main

import (
    "log"

    "github.com/monaco-io/request"
)

func main() {
    client := request.Client{
        URL:       "https://baidu.com",
        Method:    "POST",
        BasicAuth: request.BasicAuth{
            Username:"user_xxx",
            Password:"pwd_xxx",
        }, // xxx:xxx
    }

    resp, err := client.Do()

    log.Println(string(resp), err)
}
```

### Timeout

```go
package main

import (
    "log"

    "github.com/monaco-io/request"
)

func main() {
    client := request.Client{
        URL:       "https://baidu.com",
        Method:    "POST",
        Timeout:   10, // seconds
    }

    resp, err := client.Do()

    log.Println(string(resp), err)
}
```

## License

[MIT](LICENSE)