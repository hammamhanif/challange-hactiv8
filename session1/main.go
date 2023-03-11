package main

import "fmt"

func main() {
	i := 1
	j := true
	rusia := 'Я'
	k := 21
	f := 15
	f2 := 19
	var m float64 = 123.456

	fmt.Printf("Nilai i : %v\n", i)
	fmt.Printf("Tipe Data I : %T\n", i)
	fmt.Println("%")
	fmt.Println(j)
	fmt.Printf("Karakter rusia : %c \n", rusia)
	fmt.Printf("base 10 : %d \n", k)
	fmt.Printf("base 8 : %o \n", k)
	fmt.Printf("base 16 : %x \n", f)
	fmt.Printf("base 16 : %X %X\n", f, f2)
	fmt.Printf("unicode  Я : %U \n", rusia)
	fmt.Printf("Float 64 : %f\n", m)
	fmt.Printf("Scientific : %E\n", m)

}
