package main


import (
    "fmt"
    "net/http"
    "time"
)
func hello(w http.ResponseWriter, req *http.Request) {
    ctx := req.Context() // a Context carries deadlines, cancellation signals, and other request-scoped values
    // across API boundaries and goroutines
    fmt.Println("server: hello handler started")
    defer fmt.Println("server: hello handler ended")

    select {
    case <- time.After(10 * time.Second):
      fmt.Fprintf(w, "hello\n")
    case <- ctx.Done(): // keeping an eye on the context's Done() channel for a signal that we should cancel the work and return ASAP

       err := ctx.Err()
       fmt.Println("server:", err)
       internalError := http.StatusInternalServerError
       http.Error(w, err.Error(), internalError)
    }
}

func main() {
    http.HandleFunc("/hello", hello)
    http.ListenAndServe(":8090", nil)
}
/*
server: hello handler started
server: hello handler ended
hello
*/
/*
server: hello handler started
^Cserver: context canceled
server: hello handler ended
*/
