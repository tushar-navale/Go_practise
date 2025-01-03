package main

import (
	"fmt"
)

// create go routine using the go keyword

func Display(s string) {
	for i := 0; i < 6; i++ {
		fmt.Println(s)
	}
}

func main() {
	go Display("go routine") // but this is not printed at all because the main function is executed first and the go routine is executed in the background
	Display("non go routine") // this block is printed first as the go routine is executed in the background
	
}
