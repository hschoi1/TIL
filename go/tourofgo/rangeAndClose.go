//https://tour.golang.org/concurrency/4
package main

import (
    "fmt"
)

func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    close(c) // indicate that no more values will be sent. Only the sender should close a channel
} // v, ok := <-ch

func main() {
    c := make(chan int, 10)
    go fibonacci(cap(c), c)
    for i := range c {
        fmt.Println(i)
    }
}
/*
0
1
1
2
3
5
8
13
21
34
*/
