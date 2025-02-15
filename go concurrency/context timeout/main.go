package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)

	go worker(ctx)

	time.Sleep(5 * time.Second) // 5秒待機(ゴルーチンが終了するのを確認)

}

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("タイムアウト!")
			return
		default:
			fmt.Println("実行中...")
			time.Sleep(1 * time.Second)
		}
	}
}