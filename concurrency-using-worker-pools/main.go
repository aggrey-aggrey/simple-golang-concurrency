package main

import (
	"fmt"
)


func main (){
	/*use case : when you have queues of jobs to be done 
	and you have multiple concurrent workers pulling from the queue .
	This is a good example to test that we are taking advantage of multi-co processor**/

	// create a buffer channels
	jobs := make(chan int,100)
	results := make(chan int,100)

	//create worker as concurrent go routine and pass it the two channels
	go worker (jobs, results)
	go worker (jobs, results)
	go worker (jobs, results)

	//fill jobs channels with 100 #s
	for  i :=0; i < 100; i ++ {
		jobs <- i
	}
	close(jobs)
	
	//receive the  100 #s in the results channel and output to the terminal
	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

	//jobs channel for sending jobs and results channel receives jobs results
	func worker (jobs <-chan int, results chan<- int){
		//calculate the nth fib and  send it to results channel
		//the jobs are queues of #s
		for n := range jobs {
			results <- fib(n)
		}
	}

	//calculate the n-th fibannaci  #
	func fib(n int ) int {
		if n<= 1 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
