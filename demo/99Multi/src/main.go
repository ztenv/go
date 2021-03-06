package main

import "fmt"
import "sub"

func main() {
	fmt.Println("starting main...")
	sub.NineMultiNine()
	fmt.Println("stopped main...")
}

func nineMultiNine() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%2d    ", j, i, j*i)
		}
		fmt.Println("")
	}
}
