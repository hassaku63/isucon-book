package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	// go rountine の中で Add() を実行すると、go routine が起動する前に Wait() が実行されてしまう可能性がある
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 5; i++ {
			fmt.Printf("wg 1: %d\n", i+1)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 5; i++ {
			fmt.Printf("wg 2: %d\n", i+1)
			time.Sleep(1 * time.Second)
		}
	}()

	wg.Wait()

	fmt.Println("wg: done")
}
