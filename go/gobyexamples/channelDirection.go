package main

import "fmt"

func ping(pings chan<- string, msg string) { // only accepts a channel for sending. Otherwise, a compile error
    pings <- msg
}

func pong(pings <- chan string, pongs chan <- string) { // accepts one channel for receives and a second for sends
    msg := <- pings
    pongs <- msg
}


func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs) // passed message
}
