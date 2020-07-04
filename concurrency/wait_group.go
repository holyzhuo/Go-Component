package main

import (
	"sync"
	"fmt"
)

func main() {
	var wg sync.WaitGroup

	funcs := []func() {
		func() {
			defer wg.Done()
			fmt.Println("第1个goroutine完成")

		},
		func() {
			defer wg.Done()
			fmt.Println("第2个goroutine完成")
		},
	}

	wg.Add(len(funcs))

	for _, f := range funcs {
		go f()
	}

	wg.Wait()
	fmt.Println("所有goroutine都完成")
}