package main

import (
	"fmt"
	"sync"
	"time"
)

// mutexes is the naive way of solving this problem. 
// it bottlenecks performance as it only allows one goroutine to access the shared resource at a time.
var lock sync.Mutex

func processSlowly(data int) int {
	time.Sleep(2 * time.Second)
	return data * 2
}

func processData(wg *sync.WaitGroup, result *[]int, data int) {
	// lock.Lock()
	defer wg.Done()
	processedData := processSlowly(data)

	// this is the critical section, so we need to lock here instead of the whole function
	lock.Lock()
	*result = append(*result, processedData)
	lock.Unlock()
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	input := []int{1, 2, 3, 4, 5}
	result := []int{}

	for _, data := range input {
		wg.Add(1)
		go processData(&wg, &result, data)
	}

	wg.Wait()

	fmt.Println(time.Since(start))
	fmt.Println(result)
}
