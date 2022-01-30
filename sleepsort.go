package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func sleepsort(n int) {
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println(n)
}

func main() {

	var wg sync.WaitGroup

	for _, n := range os.Args[1:] {
		n, _ := strconv.Atoi(n)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sleepsort(i)
		}(n)
	}
	wg.Wait()
}
