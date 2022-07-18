package main

import (
	"fmt"
	"time"

	"github.com/wwqdrh/logger/pyroscope"
)

func main() {
	pyroscope.Start("simple.golang.app", "http://127.0.0.1:4040")

	for i := 0; i < 100; i++ {
		_ = make([]int, 10)
		time.Sleep(5 * time.Second)
	}
	fmt.Println("done")
}
