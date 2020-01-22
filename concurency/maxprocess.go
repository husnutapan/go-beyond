package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	var processCounter sync.WaitGroup
	processCounter.Add(2)

	go func() {
		defer processCounter.Done()
		for i := 0; i < 10; i++ {
			time.Sleep(2 * time.Second)
			fmt.Println("first process is working...")
		}
	}()

	go func() {
		defer processCounter.Done()
		for i := 0; i < 10; i++ {
			time.Sleep(2 * time.Second)
			fmt.Println("second process is working...")

		}
	}()
	go func() {
		defer processCounter.Done()
		for i := 0; i < 10; i++ {
			time.Sleep(2 * time.Second)
			fmt.Println("third process one is working...")

		}
	}()
	processCounter.Wait()
}
