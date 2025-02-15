package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx) // ゴルーチン開始

	time.Sleep(2 * time.Second) // 少し待つ
	fmt.Println("メイン関数:キャンセルを発行!")
	cancel() // ここでキャンセル

	time.Sleep(1 * time.Second) // ゴルーチンが終了するのを待つ
}

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("キャンセルされた!")
			return
		default:
			fmt.Println("実行中...")
			time.Sleep(500 * time.Millisecond) // 0.5秒ごとに実行
		}
	}
}