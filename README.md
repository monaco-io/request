# Request

HTTP client for golang, Inspired by [axios](https://github.com/axios/axios) [request](https://github.com/psf/requests)

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
        URL:          "https://baidu.com",
        Method:       "POST",
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
        Auth:      request.Auth{
            Username:"user_xxx",
            Password:"pwd_xxx",
        },
    }

    resp, err := client.Do()

    log.Println(string(resp), err)
}
```

## License

[MIT](LICENSE)
