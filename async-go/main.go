package main

import (
	"fmt"
	"sync"
	"time"
)

func p(v string, wg *sync.WaitGroup) {
	for i := 0; i < 20; i++ {
		fmt.Println(v)
		time.Sleep(time.Second)
	}
	defer wg.Done()
}

func main() {
	fmt.Println("vim-go")
	var wg sync.WaitGroup
	wg.Add(1)
	go p("first", &wg)
	wg.Add(1)
	go p("second", &wg)
	wg.Add(1)
	go p("third", &wg)
	wg.Wait()
}
