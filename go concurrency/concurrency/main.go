package main

import (
	"fmt"
	"sync"
)

// 共有する変数
var counter int

var mu sync.Mutex // ミューテックス

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock() // ロック (他のゴルーチンがアクセスできなくなる)
	counter++ // 競合する可能性のある変数を更新
	mu.Unlock() // アンロック (他のゴルーチンがアクセス可能になる)
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