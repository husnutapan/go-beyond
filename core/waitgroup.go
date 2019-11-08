package main

import (
	"fmt"
	"sync"
	"time"
)

type Object struct {
}

func (object *Object) Update(wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	fmt.Println("Successfully updated...")
	wg.Done()
	return
}

func main() {
	var wg sync.WaitGroup
	list := make([]Object, 5)
	for _, object := range list {
		wg.Add(1)
		go object.Update(&wg)
	}
	wg.Wait()
	fmt.Println("Group Done...")
}
