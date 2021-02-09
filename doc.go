// Package request HTTP client for golang
//  - Make http requests from Golang
//  - Intercept request and response
//  - Transform request and response data
//
// package main
//
// import (
//     "github.com/monaco-io/request"
// )
//
// func main() {
//     var body = struct {
//          A string
//          B int
//         }{A: "A", B: 001}
//     var result interface{}
//
//     client := request.Client{
//         URL:    "https://google.com",
//         Method: "POST",
//         Query: map[string]string{"hello": "world"},
//         JSON:   body,
//     }
//     if err := client.Send().Scan(&result).Error(); err != nil{
//         // handle error
//     }
//
//     // str := client.Send().String()
//     // bytes := client.Send().Bytes()
// ```
package request
