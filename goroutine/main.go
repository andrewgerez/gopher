package main

import (
	"fmt"
	"time"
)

func counter(count int) {
	for i := 0; i < count; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}

func main() {
	go counter(10)
	go counter(10)
	go counter(10)
	counter(10)
}
