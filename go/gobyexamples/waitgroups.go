package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) { // a WaitGroup must be passed to functions by pointer

    defer wg.Done()  //Block until the WaitGroup counter goes back to 0

    fmt.Printf("Worker %d starting\n", id)

    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {

   var wg sync.WaitGroup // The WaitGroup is used to wait for all the go routines launched here to finish

   for i := 1; i <= 5; i++ {
       wg.Add(1)
       go worker(i, &wg)
   }

   wg.Wait()
}
// Worker 1 starting
// Worker 4 starting
// Worker 5 starting
// Worker 3 starting
// Worker 2 starting
// Worker 2 done
// Worker 3 done
// Worker 1 done
// Worker 4 done
// Worker 5 done
