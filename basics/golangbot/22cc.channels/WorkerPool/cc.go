package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(1 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

//The allocate function above takes the number of jobs to be created as input parameter, generates pseudo random numbers with a maximum value of 998, creates Job struct using the random number and the for loop counter i as the id and then writes them to the jobs channel. It closes the jobs channel after writing all jobs.

type Job struct {
	id       int
	randomno int
}

var jobs = make(chan Job, 10)

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

//The result function reads the results channel and prints the job id, input random no and the sum of digits of the random no. The result function also takes a done channel as parameter to which it writes to once it has printed all the results.

type Result struct {
	job         Job
	sumofdigits int
}

var results = make(chan Result, 10)

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func main() {
	fmt.Printf("%T\n", results)
	startTime := time.Now()

	noOfJobs := 100
	go allocate(noOfJobs)

	done := make(chan bool)
	go result(done)

	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
