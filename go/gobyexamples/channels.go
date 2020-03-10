package main

import "fmt"

func main() {
   messages := make(chan string) // channels are the pipes that connect concurrent goroutines.

   go func() {messages <- "ping"}()

   msg := <-messages
   fmt.Println(msg) //ping
}
