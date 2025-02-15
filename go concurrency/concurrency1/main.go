package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	data int
	rwMu sync.RWMutex // 書き込みロック
)

// データを読み取る (複数のゴルーチンが同時にできる)
func readData(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	rwMu.RLock() // 読み取りロック(複数同時可能)
	fmt.Printf("Reader %d: Data = %d\n", id, data)
	time.Sleep(1 * time.Second) // 疑似的な処理
	rwMu.RUnlock() // 読み取りロック解除
}

// データを書き込む (1つのゴルーチンしかできない)
func writeData(wg *sync.WaitGroup) {
	defer wg.Done()

	rwMu.Lock() // 書き込みロック(単独でしかできない)
	data++
	fmt.Println("Writer: Data updated to", data)
	time.Sleep(1 * time.Second) // 疑似的な処理
	rwMu.Unlock() // 書き込みロック解除
}

func main() {
	var wg sync.WaitGroup

	// リーダーを5つ起動
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go readData(i, &wg)
	}

	// ライターを1つ起動
	wg.Add(1)
	go writeData(&wg)

	wg.Wait()
}