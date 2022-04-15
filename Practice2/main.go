package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomno int
}

type Result struct {
	job         Job
	sumofDigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func allocate(noofjobs int) {
	for i := 0; i < noofjobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
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
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}
func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	return sum
}
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofDigits)
	}
	done <- true
}
func main() {
	startTime := time.Now()

	noOfjobs := 100
	go allocate(noOfjobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time", diff.Microseconds())
}
