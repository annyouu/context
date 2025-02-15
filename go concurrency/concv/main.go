package main

import (
	"fmt"
	"sync"
	"time"
)

// 共有データ
var (
	data int
	rwMu sync.RWMutex // 読み書きロック
)

// 読み取りゴルーチン(複数実行可能)
func readData(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	rwMu.RLock() // 読み取りロック(他のリーダーは同時に実行可能)
	fmt.Printf("Reader %d: Reading Data = %d\n", id, data)
	time.Sleep(1 * time.Second)
	rwMu.RUnlock() //読み取り解除
	fmt.Printf("Reader %d: Unlock\n", id)
}

// 書き込みゴルーチン (1つだけ実行可能)
func writeData(wg *sync.WaitGroup) {
	defer wg.Done()

	rwMu.Lock() // 書き込みロック
	fmt.Println("Writer: Writing Data...")
	data++
	time.Sleep(2 * time.Second) 
	fmt.Println("Writer: Data Updated")
	rwMu.Unlock() // 書き込みロック解除
	fmt.Println("Writer: Unlock")
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go readData(i, &wg)
	}

	time.Sleep(500 * time.Millisecond)

	wg.Add(1)
	go writeData(&wg)

	//すべてのゴルーチンが終わるまで待機
	wg.Wait()
	fmt.Println("すべてのゴルーチン終わり！！")
}