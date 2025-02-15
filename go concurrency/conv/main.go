package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "userID", 12345)

	processRequest(ctx) // ゴルーチンを開始

}

func processRequest(ctx context.Context) {
	// 値を取得
	userID := ctx.Value("userID")
	fmt.Println("処理中のユーザーID:", userID)
}