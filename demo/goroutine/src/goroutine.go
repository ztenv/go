/**
* @file goroutine.go
* @brief goroutine test demo
* @author shlian
* @version 1.0
* @date 2017-07-03
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

var wait_group *sync.WaitGroup = &sync.WaitGroup{}

func main() {
	var start int = 0
	const step int = 10

	for i := 0; i < 100; i++ {
		if i%10 == 0 {
			wait_group.Add(1)
			go sum(start, start+step)
			start += 10
		}
	}
	wait_group.Wait() //wait all goroutine done
}

/**
* @brief cal sum from start to stop
*
* @param start
* @param stop
*
* @return
 */
func sum(start, stop int) {
	var s int = 0
	for i := start; i < stop; i++ {
		s += i
		time.Sleep(1000 * time.Millisecond)
	}
	fmt.Printf("sum[%d,%d]=%d\n", start, stop, s)
	wait_group.Done()
}
