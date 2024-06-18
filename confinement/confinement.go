package main

import (
	"fmt"
	"sync"
	"time"
)

func processSlowly(data int) int {
	time.Sleep(2 * time.Second)
	return data * 2
}

func processData(wg *sync.WaitGroup, resultDestination *int, data int) {
	defer wg.Done()
	processedData := processSlowly(data)
	*resultDestination = processedData
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	input := []int{1, 2, 3, 4, 5}
	result := make([]int, len(input))

	for i, data := range input {
		wg.Add(1)
		go processData(&wg, &result[i], data)
	}

	wg.Wait()

	fmt.Println(time.Since(start))
	fmt.Println(result)
}
