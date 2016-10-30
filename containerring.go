package main

import (
	"container/ring"
	"fmt"
	"time"
)

func main() {
	size := 3
	r := ring.New(size)
	for i := 0; i < size; i++ {
		r.Value = i * 2
		r = r.Next()
	}

	timer := time.NewTimer(10 * time.Second)
	for {
		select {
		case <-timer.C:
			return
		default:
			r = r.Next()
			fmt.Print(r.Value)
			fmt.Print(" ")
			time.Sleep(time.Second)
		}
	}
}