package main

import (
	"fmt"
	"sync"
)

func main() {
	userIds := []int{}
	userIdsLock := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			userIdsLock.Lock()
			userIds = append(userIds, id)
			userIdsLock.Unlock()
		}(i)

		wg.Wait()
	}
	fmt.Printf("userIds: %v\n", userIds)
}
