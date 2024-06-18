// we could also use a boolean channel to fix this race condition,
// however since we don't need communication between goroutines,
// it's better to use mutexes
package main

import (
	"fmt"
	"sync"
)

var total = 0

func count(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock() // critical section. without this mutex, the total will be incorrect
	total++
	m.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go count(&wg, &m)
	}

	wg.Wait()
	fmt.Println(total)
}
