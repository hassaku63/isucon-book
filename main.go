package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// now := time.Now()  // 最後の ctx.Done() は何秒後に呼ばれるのか確認する実験
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// fmt.Println(cancel) // 1行上の defer cancel() をコメントアウトしても動くように追加した。 defer cancel() がなくてもこの関数は 10s で止まることがわかった

	// fmt.Printf("Elapsed[1]: %s\n", time.Since(now))  // 即時実行される
	go LoopWithBefore(ctx)
	// fmt.Printf("Elapsed[2]: %s\n", time.Since(now))  // 即時実行される
	go LoopWithAfter(ctx)
	// fmt.Printf("Elapsed[3]: %s\n", time.Since(now))  // 即時実行される

	// fmt.Printf("Elapsed[4]: %s\n", time.Since(now))  // 即時実行される
	<-ctx.Done()
	// fmt.Printf("Elapsed[5]: %s\n", time.Since(now))  // 10s 経過した後で実行される
}

func LoopWithBefore(ctx context.Context) {
	// 3s 経過したら debug print して
	beforeLoop := time.Now()

	for {
		// 時限式のチャネルを作って select で拾う
		loopTimer := time.After(3 * time.Second)

		// [疑問] これ、同期的な実行では？
		// この例では 1.5s かかるとわかっているが、それを超える場合はどうなるのか？
		// -> 予想通り 3.5s かかっていた
		HeavyProcess(ctx, "BEFORE")

		select {
		case <-ctx.Done():
			return
		case <-loopTimer:
			fmt.Printf("[BEFORE] loop duration: %.3fs %s\n", time.Since(beforeLoop).Seconds(), beforeLoop.GoString())
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
			fmt.Printf("[AFTER] loop duration: %.3fs %s\n", time.Since(beforeLoop).Seconds(), beforeLoop.GoString())
			beforeLoop = time.Now()
		}
	}
}

func HeavyProcess(ctx context.Context, pattern string) {
	fmt.Printf("[%s] heavy process\n", pattern)
	time.Sleep(1*time.Second + 500*time.Millisecond)
}
