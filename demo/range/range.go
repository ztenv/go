package main

import "fmt"

func main(){
	var array [10] int 
	fmt.Println("starting main...")
	for i:=0;i<10;i++ {
		array[i]=i
	}

	for i:=0;i<10;i++{
		fmt.Printf("array[%d]=%d\n",i,array[i])
	}

	fmt.Println("stopped main")
}
