package main

import (
  "fmt"
  "time"
)

func f(from string) {
  for i := 0; i < 3; i++ {
     fmt.Println(from, ":", i)
  }
}

func main() {
   f("direct") // usual way, synchronously

   go f("goroutine")

   go func(msg string) {
       fmt.Println(msg)
   }("going") // goroutine for an anonymous function call

   time.Sleep(time.Second)
   fmt.Println("done")
}
// direct : 0
// direct : 1
// direct : 2
// going
// goroutine : 0
// goroutine : 1
// goroutine :2
// done
