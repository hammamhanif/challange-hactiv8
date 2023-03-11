package main

import "fmt"

func main() {
	var i, j, n = 0, 0, 0
	var c = []rune{'\u0421', '\u0410', '\u0428', '\u0410', '\u0420', '\u0412', '\u041E'}
	for i < 5 {
		fmt.Println("Nilai i =", i)
		i++
		continue
	}
	for j <= 10 {
		if j == 5 {
			for _, l := range c {
				fmt.Printf("character %U '%c' starts at byte position %d\n", l, l, n)
				n += 2
				continue
			}
		}
		fmt.Printf("Nilai j = %d\n", j)
		j++
	}
}
