package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
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
        go worker(w, jobs, results)
    }

		// Send 5 jobs
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
		// Close the channel to indicate that's all the work we have.
    close(jobs)

		// Collect all the results of the work. This also ensures that the worker goroutines have finished
    for a := 1; a <= numJobs; a++ {
        <-results
    }
}
