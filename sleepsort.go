package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)
var size int
var wg sync.WaitGroup

func sleepsort(n int) int {
	defer wg.Done()
	time.Sleep(time.Duration(n) * time.Second)
	return n
}

func printer(in <-chan int) {
	for i := 0; i < size; i++ {
		fmt.Print(<-in," ")
	}
}

type Chan chan int

func main() {
	size, _  = strconv.Atoi(os.Args[1])
	c := make(Chan, size)

	for _, n := range os.Args[2:] {
		n, _ := strconv.Atoi(n)
		wg.Add(1)
		go func(c Chan, i int) {
			c <- sleepsort(i)
		}(c, n)
	}
	wg.Wait()
	printer(c)
	close(c)
}
