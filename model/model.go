package model

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func Printer(data string) {
	for _, v := range data {
		fmt.Printf("Printer:%c\n", v)
		time.Sleep(time.Second)
	}
}

func Person1() {
	Printer("1234567890")
	ch <- 123
}

func Person2() {
	<-ch
	Printer("qwertyuiop")
}

// func Person3() {
// 	Printer("!@#$%^&*()")
// }
