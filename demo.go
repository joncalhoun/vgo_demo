package demo

import "fmt"

func Log(msg interface{}, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(msg)
	}
}
