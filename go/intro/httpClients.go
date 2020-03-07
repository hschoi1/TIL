package main

import(
    "bufio"
    "fmt"
    "net/http"
)

func main() {

    resp, err := http.Get("http://gobyexample.com")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("Response status:", resp.Status)

    scanner := bufio.NewScanner(resp.Body)
    for i := 0; scanner.Scan() && i < 5; i++ { // print first 5 files of the response body
        fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
//Response status: 200 OK
//<!DOCTYPE html>
//<html>
//  <head>
//    <meta charset="utf-8">
//    <title>Go by Example</title>
