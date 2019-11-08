package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func consoleNumber() {
	for i := 1; i <= 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(i)
	}
	defer wg.Done()
}

func consoleAlphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(2000 * time.Millisecond)
		fmt.Printf("%c ", i)
		fmt.Println()
	}
	defer wg.Done()
}

func main() {
	wg.Add(2)
	go consoleNumber()
	go consoleAlphabets()
	wg.Wait()
	fmt.Println("Main function completed...")

}
