package main

import (
	"context"
	"fmt"
	"time"

	"github.com/wwqdrh/logger/pprofx"
)

func main() {
	pprofx.Start(context.Background(), pprofx.NewPprofOption(
		"http://127.0.0.1:4040",
		"simple.golang.app",
		pprofx.WithPprofType(pprofx.AllTypeOptions...),
	))

	for i := 0; i < 100; i++ {
		_ = make([]int, 10)
		time.Sleep(5 * time.Second)
	}
	fmt.Println("done")
}
