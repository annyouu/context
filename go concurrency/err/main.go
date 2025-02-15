package main

import (
	"fmt"
	"context"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	//キャンセル可能なコンテキストを作成
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var eg errgroup.Group

	// ゴルーチン1
	eg.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			time.Sleep(1 * time.Second)
			return fmt.Errorf("ゴルーチン1でエラー発生")
		}
	})

	//ゴルーチン2
	eg.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			fmt.Println("ゴルーチン2は正常!!")
			return nil
		}
	})

	//ゴルーチン3
	eg.Go(func() error {
		select {
		case <- ctx.Done():
			return ctx.Err()
		default:
			time.Sleep(3 * time.Second)
			return fmt.Errorf("ゴルーチン3でエラー発生")
		}
	})

	if err := eg.Wait(); err != nil {
		fmt.Println("エラー発生:", err)
	} else {
		fmt.Println("全て正常に動いている")
	}
}


// var eg errgroup.Group

// eg.Go(func() error {
// 	return nil // 成功の場合
// })

// // エラーが起きればそのエラーを出す
// eg.Go(func() error {
// 	return fmt.Errorf("エラー発生")
// })

// // 全てのゴルーチンが終了するのを待ち、エラーがあれば取得
// if err := eg.Wait(); err != nil {
// 	fmt.Println("Error:", err)
// }





