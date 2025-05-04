package math

import (
	"fmt"
	"sync"
	"time"
)

func Worker(i int, wg *sync.WaitGroup) {
	fmt.Printf(" worker %d started ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf(" worker %d ended ", i)
	wg.Done()
}
