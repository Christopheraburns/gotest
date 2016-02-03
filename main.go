package main

import "fmt"

func main() {
	//fmt.Println(42)
	//fmt.Printf("%d -%b - %x", 42, 42, 42)
	//fmt.Printf("%d -%#X - %#x", 42, 42, 42)
	for i := 0; i < 128; i++ {
		fmt.Printf("%d \t %b \t %#x %q \n", i, i, i, i)
	}
}
