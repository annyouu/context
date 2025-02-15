package main

import (
	"fmt"
	"sync"
)

var m sync.Mutex
var counter int

func increment() {
	m.Lock() //ロック
	defer m.Unlock() // 処理が終わったらロックを解除
	counter++
}

func main() {
	var wg sync.WaitGroup

	// ゴルーチンを並行実行
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment() // counterを増やす
		}()
	}

	wg.Wait()
	fmt.Println("Final counter:", counter)
}