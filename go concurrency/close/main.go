package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct {})

	go func() {
		<-done //ここでブロック
		fmt.Println("ゴルーチン1:終了通知を受け取った!")
	}()

	go func() {
		<-done // ここでブロック
		fmt.Println("ゴルーチン2:終了通知を受け取った!")
	}()

	time.Sleep(1 * time.Second) //ゴルーチンが待機する時間を作る
	close(done) // 全ゴルーチンに「終了」を通知！

	time.Sleep(1 * time.Second) // メッセージを表示する時間を確保
}