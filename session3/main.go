package main

import "fmt"

type cacah = map[string]int

func main() {
	cacahKata := cacah{}
	kata := "selamat malam"
	for i := 0; i < len(kata); i++ {
		fmt.Println(string(kata[i]))
		cacahKata[string(kata[i])] += 1
	}
	fmt.Println(cacahKata)
}
