// Package request HTTP client for golang
//  - Make http requests from Golang
//  - Intercept request and response
//  - Transform request and response data
//
// GET
//
//     client := request.Client{
//         URL:    "https://google.com",
//         Method: "GET",
//         Query: map[string]string{"hello": "world"},
//     }
//     resp := client.Send()
//
// POST
//
//     client := request.Client{
//         URL:    "https://google.com",
//         Method: "POST",
//         Query: map[string]string{"hello": "world"},
//         JSON:   []byte(`{"hello": "world"}`),
//     }
//     resp := client.Send()
//
// Content-Type
//
//     client := request.Client{
//         URL:          "https://google.com",
//         Method:       "POST",
//         ContentType: request.ApplicationXWwwFormURLEncoded, // default is "application/json"
//     }
//     resp := client.Send()
//
// Authorization
//
//     client := request.Client{
//         URL:       "https://google.com",
//         Method:    "POST",
//         BasicAuth:      request.BasicAuth{
//             Username:"user_xxx",
//             Password:"pwd_xxx",
//         }, // xxx:xxx
//     }
//
//     resp := client.Send()
//
// Cookies
//     client := request.Client{
//         URL:       "https://google.com",
//         Cookies:[]*http.Cookie{
//              {
//               Name:  "cookie_name",
//               Value: "cookie_value",
//              },
//         },
//     }
//
//     resp := client.Send()
package request
