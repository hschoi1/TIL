package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <- chan int, results chan<- int) {
   for j := range jobs {
       fmt.Println("worker", id, "started job", j)
       time.Sleep(time.Second)
       fmt.Println("worker", id, "finished job", j)
       results <- j * 2
   }
}

func main() {
    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results) // run several concurrent instances. no jobs yet
    }
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }  // send jobs
    close(jobs) // close the channel to indicate that's all the work we have

    for a := 1; a <= numJobs; a++ {
        <- results  // collect all the results of the work. Ensures that the worker goroutines have finished
    }
}
//The program takes about 2 seconds because there are 3 workers operating concurrently
/*
worker 3 started job 1
worker 1 started job 2
worker 2 started job 3
worker 2 finished job 3
worker 2 started job 4
worker 1 finished job 2
worker 1 started job 5
worker 3 finished job 1
worker 1 finished job 5
worker 2 finished job 4

real	0m2.209s
*/
