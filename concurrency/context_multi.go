package main

import (
	"time"
	"fmt"
	"context"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watchWithName(ctx, "【监控1】")
	go watchWithName(ctx, "【监控2】")
	go watchWithName(ctx, "【监控3】")

	time.Sleep(3 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(2 * time.Second)
}

func watchWithName(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了...")
			return
		default:
			fmt.Println(name, "goroutine监控中...")
		}
	}
}