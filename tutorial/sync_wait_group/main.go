package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Go(func() { otherFunc(1) }) // it will auto handle wg.Add(1) and wg.Done()

	wg.Wait()
	fmt.Println("Worker task Completed")
}

func worker(i int, wg *sync.WaitGroup) { // calling this by manual methods of WaitGroup
	defer wg.Done()

	fmt.Printf("Worker %d started\n", i)
	// any task here
	fmt.Printf("Worker %d ended\n\n", i)
}

func otherFunc(i int) { // calling it by WaitGroup.Go()
	fmt.Printf("Num: %d\n\n", i)
}
