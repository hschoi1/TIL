package main

import (
    "fmt"
    "net/http"
)


func hello(w http.ResponseWriter, req *http.Request) {// The response writer is used to fill in the HTTP response

    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
    // read all the HTTP request headers and echo them into the response body
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {
    http.HandleFunc("/hello", hello) // convenience function. sets up the default router in the net/http pakcage
    http.HandleFunc("/headers", headers)

    http.ListenAndServe(":8090", nil)
}
//curl localhost:8090/hello -> hello
//curl localhost:8090/headers -> //User-Agent: curl/7.54.0 \n Accept: */*
