package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex

	mu.Lock()

	go func() {
		<-time.After(1 * time.Second)
		mu.Unlock()
	}()

	// LockWithValue(mu) // copylocks
	LockWithRef(&mu)

	fmt.Println("mutex unlocked.")
}

func LockWithValue(mu sync.Mutex) {
	mu.Lock()
	mu.Unlock()
}

func LockWithRef(mu *sync.Mutex) {
	mu.Lock()
	mu.Unlock()
}
