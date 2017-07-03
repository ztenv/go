/**
* @file map.go
* @brief map test demo
* @author shlian
* @version 1.0
* @date 2017-07-03
 */

package main

import (
	"fmt"
	"strconv"
)

type test_map map[int]string

func main() {

	//var kv map[int]string
	var kv test_map
	kv = make(map[int]string)

	for i := 0; i < 100; i++ {
		kv[i] = strconv.Itoa(i)
	}

	for k, v := range kv {
		fmt.Printf("map[%d]=%s\n", k, v)
	}
}
