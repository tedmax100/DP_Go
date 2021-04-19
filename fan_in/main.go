package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	const workerCnt = 25
	wgWorkers := sync.WaitGroup{}
	wgWorkers.Add(workerCnt)

	genNumbers := func(min, max int) []int {
		numbers := make([]int, max-min+1)
		var idx = 0
		for i := min; i <= max; i++ {
			numbers[idx] = i
			idx++
		}
		return numbers
	}

	in := gen(genNumbers(1, 50)...)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	margedResult := merge(c1, c2)

	defer timeTrack(time.Now(), "")
	// Consume the merged output from c1 and c2.
	for w := 1; w <= workerCnt; w++ {
		go worker(w, margedResult, &wgWorkers)
	}
	wgWorkers.Wait()
}

// fan in : input from channels and return an output channel
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	// Start an output goroutine for each input channel in cs.
	//output copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	// 當job channel被關閉, 會觸發defer來結束wait group
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("worker %d started job with job %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker %d finished job for job %d with result %d\n", id, j, j*2)
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
