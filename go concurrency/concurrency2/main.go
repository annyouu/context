package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64 // int64型の変数を使う

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	atomic.AddInt64(&counter, 1)
}

func main() {
	var wg sync.WaitGroup

	// 10個のゴルーチンを並行実行
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Final counter:", counter)
}