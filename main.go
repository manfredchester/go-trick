package main

import (
	"fmt"
	"time"
)

func main() {

	makeBuff()
	deferCall()
	<-time.After(time.Second * 10)
}

func deferCall() {
	defer func() {
		fmt.Println("打印前")
	}()
	defer func() {
		fmt.Println("打印中")
	}()
	defer func() {
		fmt.Println("打印后")
	}()
	panic("触发panic")
}
