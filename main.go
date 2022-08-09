package main

import (
	"context"
	"fmt"
	"time"
)

/*
	"WaitGroup" を使ってみる
*/
// func main() {
// 	wg := sync.WaitGroup{}
// 	wg.Add(2)
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 5; i++ {
// 			fmt.Printf("wg 1: %d / 5\n", i+1)
// 			time.Sleep(1 * time.Second)
// 		}
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 5; i++ {
// 			fmt.Printf("wg 2: %d / 5\n", i+1)
// 		}
// 	}()
// 	wg.Wait()
// 	fmt.Println("wg: done")
// }

/*
	非同期なループ処理
*/
// func main() {
// 	ctx, cannel := context.WithTimeout(context.Background(), 5*time.Second+500*time.Millisecond)
// 	defer cannel()

// 	i := 0
// L:
// 	for {
// 		fmt.Printf("loop %d\n", i)
// 		i++
// 		select {
// 		case <-ctx.Done():
// 			break L
// 		case <-time.After(1 * time.Second):

// 			// Sleep では sleep の途中に中断すると次のループに行ってしまうので、time.After で
// 			// default:
// 			// 	time.Sleep(1 * time.Second)
// 		}
// 	}
// }

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go LoopWithBefore(ctx)
	go LoopWithAfter(ctx)

	<-ctx.Done()
}

func LoopWithBefore(ctx context.Context) {
	beforeLoop := time.Now()

	for {
		loopTimer := time.After(3 * time.Second)

		HeavyProcess(ctx, "BEFORE")

		select {
		case <-ctx.Done():
			return
		case <-loopTimer:
			fmt.Printf("[BEFORE] loop duration: %.3fs\n", time.Now().Sub(beforeLoop).Seconds())
			beforeLoop = time.Now()
		}
	}
}

func LoopWithAfter(ctx context.Context) {
	beforeLoop := time.Now()
	for {
		HeavyProcess(ctx, "AFTER")

		select {
		case <-ctx.Done():
			return
		case <-time.After(3 * time.Second):
			fmt.Printf("[AFTER] loop duration: %.3fs\n", time.Now().Sub(beforeLoop).Seconds())
			beforeLoop = time.Now()
		}
	}
}

func HeavyProcess(ctx context.Context, pattern string) {
	fmt.Printf("[%s] heavy process\n", pattern)
	time.Sleep(1*time.Second + 500*time.Millisecond)
}
