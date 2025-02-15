package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	 
	ctx = context.WithValue(ctx, "user", "Alice")

	// 3秒後にタイムアウトするコンテキスト
	ctx, timeoutCancel := context.WithTimeout(ctx, 3 * time.Second)
	defer timeoutCancel() // 確実にリソース開放

}