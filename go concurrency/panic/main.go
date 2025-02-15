package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	go func() {
// 		defer func() {
// 			if r := recover(); r != nil {
// 				fmt.Println("ゴルーチンにリカバーをしました:", r)
// 			}
// 		}()
// 		panic("PANIC")
// 	}()

// 	time.Sleep(1 * time.Second) // ゴルーチンが終了する前に待機
// 	fmt.Println("Main function continues")
// }



// time.Sleepではなく、wait.Groupを使う場合
func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("ゴルーチンの処理開始")
	}()

	fmt.Println("メインゴルーチンが続く")

	wg.Wait() 
	fmt.Println("メインゴルーチン終了")
}
