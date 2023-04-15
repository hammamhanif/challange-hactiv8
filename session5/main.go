package main

import (
	"fmt"
	"strings"
	"sync"
)

func printResult(count int, data interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(data, count)
}

func main() {
	data1 := []interface{}{"coba1", "coba2", "coba3"}
	data2 := []interface{}{"bisa1", "bisa2", "bisa3"}
	wg := sync.WaitGroup{}

	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go printResult(i, data1, &wg)
		go printResult(i, data2, &wg)
	}

	wg.Wait()

	fmt.Println(strings.Repeat("#", 25))

	for i, j := 1, 1; i <= 4; i, j = i+1, j+1 {
		wg.Add(1)
		go printResult(i, data1, &wg)
		wg.Wait()

		wg.Add(1)
		go printResult(j, data2, &wg)
		wg.Wait()
	}
}
