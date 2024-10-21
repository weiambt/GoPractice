package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test_1(t *testing.T) {

}

func test1() {
	ctx := context.Background()

	// 1. 传递数据
	ctx = context.WithValue(ctx, "key", "value")
	fmt.Println(ctx.Value("key"))

	// 2. 传递超时时间
	ctx, cancel := context.WithTimeout(ctx, 1000)
	defer cancel()

	// 3. 传递取消函数
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()
}

func test2() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(mctx context.Context) {
		for {
			select {
			case <-mctx.Done():
				fmt.Println("done")
				return
			default:
				fmt.Println("default")
				time.Sleep(time.Second)
			}
		}
	}(ctx)

	time.Sleep(3 * time.Second)
	cancel()

}

func main() {
	//test1()

	test2()
}
