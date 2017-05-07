package main

import "fmt"

func main() {
	fmt.Println("Hello!")
	fmt.Printf("%d \t %b \t %#x\n", 42, 42, 42)

	for i := 0; i < 200; i++ {
		fmt.Printf("%d\t%b\t%#x\t%q\n", i, i, i, i)
	}
}
